//go:build js && wasm
// +build js,wasm

package webzen

import (
	"syscall/js"

	"github.com/dimkauzh/webzen/src/backend/document"
	"github.com/dimkauzh/webzen/src/backend/window"
)

func Init() {
	canvas := document.GetElementById("canvas")
	if canvas.IsNull() {
		canvas = document.CreateCanvasElement()
		canvas.Set("id", "canvas")
		document.Body.AppendCanvasChild(canvas)
	}

	// Set the canvas size to match the viewport size
	canvas.Set("width", document.DocumentElement.ClientWidth())
	canvas.Set("height", document.DocumentElement.ClientHeight())

	// Prevent scrolling on the HTML and body elements
	document.DocumentElement.Style.Set("overflow", "hidden")
	document.Body.Style.Set("overflow", "hidden")

	// Listen for window resize events to adjust the canvas size
	window.AddEventListener("resize", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		canvas.Set("width", document.DocumentElement.ClientWidth())
		canvas.Set("height", document.DocumentElement.ClientHeight())
		return nil
	}))
}
