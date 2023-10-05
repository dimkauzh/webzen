//go:build js && wasm
// +build js,wasm

package draw

import (
	"strconv"

	"github.com/dimkauzh/webzen/src/backend/document"
)

func SetPixel(x, y float64, color [4]int) {
	canvas := document.GetElementById("canvas")
	// Get the canvas context
	context := canvas.GetContext("2d")

	// Draw the rectangle
	rgba := "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
	context.Set("fillStyle", rgba)
	context.FillRect(x, y, 1, 1) // Set the rectangle dimensions as needed
}
