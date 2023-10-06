//go:build js && wasm
// +build js,wasm

package draw

import (
	"github.com/dimkauzh/webzen/src/backend/colors"
	"github.com/dimkauzh/webzen/src/backend/document"
)

func SetPixel(x, y float64, color [4]int) {
	canvas := document.GetElementById("canvas")
	// Get the canvas context
	context := canvas.GetContext("2d")

	// Draw the rectangle
	rgba := colors.GetRGBA(color)
	context.Set("fillStyle", rgba)
	context.FillRect(x, y, 1, 1) // Set the rectangle dimensions as needed
}
