package draw

import (
	"strconv"
	"syscall/js"
)

func DrawRect(width, height int, x, y float64, color [4]int) {
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	// Get the canvas context
	context := canvas.Call("getContext", "2d")

	// Draw the rectangle
	rgba := "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
	context.Set("fillStyle", rgba)
	context.Call("fillRect", x, y, width, height) // Set the rectangle dimensions as needed
}

func FillBackground(color [4]int) {
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	// Get the canvas context
	context := canvas.Call("getContext", "2d")

	// Fill the canvas with the specified color
	rgba := "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
	context.Set("fillStyle", rgba)
	context.Call("fillRect", 0, 0, canvas.Get("width").Float(), canvas.Get("height").Float())
}
