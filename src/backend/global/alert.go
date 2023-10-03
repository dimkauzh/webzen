//go:build js && wasm
// +build js,wasm

package global

import "syscall/js"

func Alert(text string) {
	js.Global().Call("alert", text)
}
