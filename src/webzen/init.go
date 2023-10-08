//go:build js && wasm
// +build js,wasm

package webzen

import (
	"syscall/js"
  "time"

	"github.com/dimkauzh/webzen/src/backend/document"
	"github.com/dimkauzh/webzen/src/backend/window"
  "github.com/dimkauzh/webzen/src/webzen/keys"
)

func Init() {
	canvas := document.GetElementById("canvas")
	if canvas.IsNull() {
		canvas = document.CreateCanvasElement()
		canvas.Set("id", "canvas")
		document.Body.AppendCanvasChild(canvas)
	}

	canvas.Set("width", document.DocumentElement.ClientWidth())
	canvas.Set("height", document.DocumentElement.ClientHeight())

	document.DocumentElement.Style.Set("overflow", "hidden")
	document.Body.Style.Set("overflow", "hidden")

	window.AddEventListener("resize", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		canvas.Set("width", document.DocumentElement.ClientWidth())
		canvas.Set("height", document.DocumentElement.ClientHeight())
		return nil
	}))

  keys.SetupEventListeners()
}

func Update() {
  time.Sleep(time.Millisecond * 16)
 
}
