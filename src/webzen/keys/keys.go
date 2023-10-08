//go:build js && wasm
// +build js,wasm

package keys

import (
	"syscall/js"

	"github.com/dimkauzh/webzen/src/backend/document"
)

var keyPressed = make(map[string]bool)

func SetupEventListeners() {
	// Initialize event listeners once
	document.AddEventListener("keydown", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		event := p[0]
		key := event.Get("key").String()

		// Set the key's state to pressed.
		keyPressed[key] = true

		return nil
	}))

	document.AddEventListener("keyup", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		event := p[0]
		key := event.Get("key").String()

		// Set the key's state to released.
		keyPressed[key] = false

		return nil
	}))
}

func KeyPressed(key string) bool {
	return keyPressed[key]
}
