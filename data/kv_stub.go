//go:build !(js && wasm)

package data

import (
	"os"
	"path/filepath"
)

// LoadFromKV retrieves data from local disk in native Go mode
func LoadFromKV(key string) ([]byte, error) {
	path := filepath.Join("data", key)
	if _, err := os.Stat(path); err == nil {
		return os.ReadFile(path)
	}
	return nil, nil
}

// SaveToKV stores data to local disk in native Go mode
func SaveToKV(key string, data []byte) error {
	path := filepath.Join("data", key)
	_ = os.MkdirAll(filepath.Dir(path), 0755)
	return os.WriteFile(path, data, 0644)
}

// DeleteFromKV removes data from local disk in native Go mode
func DeleteFromKV(key string) error {
	path := filepath.Join("data", key)
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	}
	return nil
}
