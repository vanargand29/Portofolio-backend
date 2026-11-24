//go:build js && wasm

package data

import (
	"fmt"
	"syscall/js"

	"github.com/syumai/workers-go/cloudflare"
)

func awaitPromise(promiseVal js.Value) (js.Value, error) {
	resultCh := make(chan js.Value)
	errCh := make(chan error)
	var then, catch js.Func
	then = js.FuncOf(func(_ js.Value, args []js.Value) any {
		defer then.Release()
		if len(args) > 0 {
			resultCh <- args[0]
		} else {
			resultCh <- js.Undefined()
		}
		return js.Undefined()
	})
	catch = js.FuncOf(func(_ js.Value, args []js.Value) any {
		defer catch.Release()
		msg := "unknown JS error"
		if len(args) > 0 {
			msg = args[0].Call("toString").String()
		}
		errCh <- fmt.Errorf("%s", msg)
		return js.Undefined()
	})
	promiseVal.Call("then", then).Call("catch", catch)
	select {
	case res := <-resultCh:
		return res, nil
	case err := <-errCh:
		return js.Undefined(), err
	}
}

func getKVInstance() (val js.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("KV runtime context not ready: %v", r)
		}
	}()
	inst := cloudflare.GetBinding("STORE")
	if inst.IsUndefined() || inst.IsNull() {
		return js.Undefined(), fmt.Errorf("STORE KV binding not found")
	}
	return inst, nil
}

// LoadFromKV retrieves data by key from Cloudflare KV
func LoadFromKV(key string) (result []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("KV load panic: %v", r)
			result = nil
		}
	}()

	inst, err := getKVInstance()
	if err != nil || inst.IsUndefined() || inst.IsNull() {
		return nil, err
	}

	opt := js.Global().Get("Object").New()
	opt.Set("type", "arrayBuffer")
	p := inst.Call("get", key, opt)
	v, err := awaitPromise(p)
	if err != nil {
		return nil, err
	}
	if v.IsNull() || v.IsUndefined() {
		// Key does not exist in KV
		return nil, nil
	}

	byteLength := v.Get("byteLength").Int()
	if byteLength == 0 {
		return []byte{}, nil
	}

	u8 := js.Global().Get("Uint8Array").New(v)
	buf := make([]byte, byteLength)
	js.CopyBytesToGo(buf, u8)
	return buf, nil
}

// SaveToKV stores data by key into Cloudflare KV
func SaveToKV(key string, data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("KV save panic: %v", r)
		}
	}()

	inst, err := getKVInstance()
	if err != nil || inst.IsUndefined() || inst.IsNull() {
		return err
	}

	u8 := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(u8, data)

	p := inst.Call("put", key, u8.Get("buffer"))
	_, err = awaitPromise(p)
	return err
}

// DeleteFromKV removes data by key from Cloudflare KV
func DeleteFromKV(key string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("KV delete panic: %v", r)
		}
	}()

	inst, err := getKVInstance()
	if err != nil || inst.IsUndefined() || inst.IsNull() {
		return err
	}

	p := inst.Call("delete", key)
	_, err = awaitPromise(p)
	return err
}
