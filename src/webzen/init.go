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

var fpsInterval = time.Millisecond * 16

func SetFps(fps int) {
	fpsInterval = time.Second / time.Duration(fps)
}

func clearCanvas() {
	canvas := document.GetElementById("webzen")
	context := canvas.GetContext("2d")
	context.ClearRect(0, 0, canvas.Get("width").Float(), canvas.Get("height").Float())
}

func Init() {
	canvas := document.GetElementById("webzen")
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

func CustomInit(width, height int) {
	canvas := document.GetElementById("webzen")
	if canvas.IsNull() {
		canvas = document.CreateCanvasElement()
		canvas.Set("id", "canvas")
		document.Body.AppendCanvasChild(canvas)
	}

	canvas.Set("width", width)
	canvas.Set("height", height)

	document.DocumentElement.Style.Set("overflow", "hidden")
	document.Body.Style.Set("overflow", "hidden")

	keys.SetupEventListeners()
}

func Update() {
	time.Sleep(time.Millisecond * 16)
	clearCanvas()
}

func GetWidth() int {
	return document.DocumentElement.ClientIntWidth()
}

func GetHeight() int {
	return document.DocumentElement.ClientIntHeight()
}
