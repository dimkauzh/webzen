//go:build js && wasm
// +build js,wasm

package draw

import (
	"github.com/dimkauzh/webzen/src/backend/colors"
	"github.com/dimkauzh/webzen/src/backend/document"
)

func DrawRect(width, height, x, y float64, color [4]int) {
	canvas := document.GetElementById("webzen")
	context := canvas.GetContext("2d")

	rgba := colors.GetRGBA(color)
	context.Set("fillStyle", rgba)
	context.FillRect(x, y, width, height)
}

func FillBackground(color [4]int) {
	canvas := document.GetElementById("webzen")
	context := canvas.GetContext("2d")

	rgba := colors.GetRGBA(color)
	context.Set("fillStyle", rgba)
	context.FillRect(0, 0, canvas.Get("width").Float(), canvas.Get("height").Float())
}
