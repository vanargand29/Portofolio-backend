package admin

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	_ "embed"

	"github.com/syumai/workers-go/cloudflare"
	"golang.org/x/crypto/bcrypt"
)

//go:embed login.html
var loginHTML string

const (
	sessionCookieName = "admin_session"

	maxFailedAttempts = 5
	lockoutDuration   = 15 * time.Minute
	sessionMaxAge     = 60 * 60 * 8 // 8 hours
)

// getSecret safely retrieves a secret from Cloudflare Workers context, falling back to os.Getenv
func getSecret(name string) string {
	defer func() {
		_ = recover()
	}()
	if v := cloudflare.Getenv(name); v != "" {
		return strings.TrimSpace(v)
	}
	if v := os.Getenv(name); v != "" {
		return strings.TrimSpace(v)
	}
	return ""
}

// Credential helpers - loaded securely from Cloudflare Worker secrets
func getAdminUsername() string {
	if u := getSecret("ADMIN_USERNAME"); u != "" {
		return u
	}
	return "admin"
}

func getAdminPasswordHash() string {
	return getSecret("ADMIN_PASSWORD_HASH")
}

func getSessionSecret() string {
	if s := getSecret("SESSION_SECRET"); s != "" {
		return s
	}
	return "local-dev-only-secret-not-for-production"
}

// ----- Rate limiter -----

type attemptInfo struct {
	count       int
	lastAttempt time.Time
	lockedUntil time.Time
}

var (
	loginAttempts   = map[string]*attemptInfo{}
	loginAttemptsMu sync.Mutex
)

func clientIP(r *http.Request) string {
	if cf := r.Header.Get("CF-Connecting-IP"); cf != "" {
		return cf
	}
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.TrimSpace(strings.SplitN(xff, ",", 2)[0])
	}
	return r.RemoteAddr
}

func isRateLimited(ip string) bool {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()
	a, ok := loginAttempts[ip]
	if !ok {
		return false
	}
	if time.Now().Before(a.lockedUntil) {
		return true
	}
	// Reset after lockout expires
	if a.count >= maxFailedAttempts && time.Now().After(a.lockedUntil) {
		delete(loginAttempts, ip)
	}
	return false
}

func recordFailedAttempt(ip string) {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()
	a, ok := loginAttempts[ip]
	if !ok {
		a = &attemptInfo{}
		loginAttempts[ip] = a
	}
	a.count++
	a.lastAttempt = time.Now()
	if a.count >= maxFailedAttempts {
		a.lockedUntil = time.Now().Add(lockoutDuration)
	}
}

func clearAttempts(ip string) {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()
	delete(loginAttempts, ip)
}

// ----- Stateless Cryptographic HMAC Session Tokens -----
// Serverless-ready: Verified anywhere without shared RAM or database lookups!

func createSessionToken(username string) string {
	exp := time.Now().Add(time.Duration(sessionMaxAge) * time.Second).Unix()
	payload := fmt.Sprintf("%s:%d", username, exp)
	mac := hmac.New(sha256.New, []byte(getSessionSecret()))
	mac.Write([]byte(payload))
	sig := hex.EncodeToString(mac.Sum(nil))
	rawToken := fmt.Sprintf("%s:%s", payload, sig)
	return base64.RawURLEncoding.EncodeToString([]byte(rawToken))
}

func isTokenValid(token string) bool {
	if token == "" {
		return false
	}
	rawBytes, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return false
	}
	parts := strings.Split(string(rawBytes), ":")
	if len(parts) != 3 {
		return false
	}
	u := parts[0]
	expStr := parts[1]
	givenSig := parts[2]

	expTime, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil || time.Now().Unix() > expTime {
		return false
	}

	mac := hmac.New(sha256.New, []byte(getSessionSecret()))
	mac.Write([]byte(fmt.Sprintf("%s:%s", u, expStr)))
	expectedSig := hex.EncodeToString(mac.Sum(nil))

	if subtle.ConstantTimeCompare([]byte(givenSig), []byte(expectedSig)) != 1 {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(u), []byte(getAdminUsername())) == 1
}

// ----- Security headers -----

func setSecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
}

// ----- Auth middleware -----

// IsAuthenticated checks if the request carries a valid signed session cookie.
func IsAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil || cookie.Value == "" {
		return false
	}
	return isTokenValid(cookie.Value)
}

// RequireAuth wraps a handler: unauthenticated requests are redirected to /admin/login.
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		if !IsAuthenticated(r) {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}

// LoginHandler handles GET /admin/login and POST /admin/login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)

	switch r.Method {
	case http.MethodGet:
		// If already authenticated, redirect to /admin
		if IsAuthenticated(r) {
			http.Redirect(w, r, "/admin", http.StatusFound)
			return
		}

		errMsg := r.URL.Query().Get("error")
		page := loginHTML
		if errMsg == "rate" {
			page = strings.ReplaceAll(page, "<!--ERROR_PLACEHOLDER-->",
				`<div class="error-box">Terlalu banyak percobaan gagal. Coba lagi dalam 15 menit.</div>`)
		} else if errMsg != "" {
			page = strings.ReplaceAll(page, "<!--ERROR_PLACEHOLDER-->",
				`<div class="error-box">Username atau password salah. Coba lagi.</div>`)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, page)

	case http.MethodPost:
		ip := clientIP(r)

		if isRateLimited(ip) {
			http.Redirect(w, r, "/admin/login?error=rate", http.StatusFound)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Redirect(w, r, "/admin/login?error=1", http.StatusFound)
			return
		}
		username := strings.TrimSpace(r.FormValue("username"))
		password := r.FormValue("password")

		expectedUsername := getAdminUsername()
		expectedHash := getAdminPasswordHash()

		usernameOK := subtle.ConstantTimeCompare([]byte(username), []byte(expectedUsername)) == 1
		var passwordOK bool
		if expectedHash != "" {
			passwordOK = bcrypt.CompareHashAndPassword([]byte(expectedHash), []byte(password)) == nil
		}

		if usernameOK && passwordOK {
			clearAttempts(ip)
			token := createSessionToken(expectedUsername)

			isHTTPS := r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" || r.Header.Get("CF-Visitor") != ""

			http.SetCookie(w, &http.Cookie{
				Name:     sessionCookieName,
				Value:    token,
				Path:     "/",
				HttpOnly: true,
				Secure:   isHTTPS,
				SameSite: http.SameSiteLaxMode, // Lax ensures cookie is sent on 302 navigation
				MaxAge:   sessionMaxAge,
			})
			http.Redirect(w, r, "/admin", http.StatusFound)
			return
		}

		recordFailedAttempt(ip)
		http.Redirect(w, r, "/admin/login?error=1", http.StatusFound)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// LogoutHandler clears the session cookie and redirects to /admin/login.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}
