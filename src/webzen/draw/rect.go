package draw

import (
	"strconv"
	"syscall/js"
)

func DrawRect(width, height, x, y float64, color [4]int) {
	// Create a canvas element if it doesn't exist
	canvas := js.Global().Get("document").Call("getElementById", "canvas")

	// Get the canvas context
	context := canvas.Call("getContext", "2d")

	// Set the canvas size
	canvas.Set("width", width)
	canvas.Set("height", height)

	// Draw the rectangle
	rgba := "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
	context.Set("fillStyle", rgba)
	context.Call("fillRect", x, y, width, height)
}
