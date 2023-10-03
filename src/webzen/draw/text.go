package draw

import (
	"strconv"
	"syscall/js"
)

func DrawText(text string, size, x, y int) {
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	// Get the canvas context
	context := canvas.Call("getContext", "2d")

	// Set the text size and font
	context.Set("font", strconv.Itoa(size)+"px Arial")

	// Set the text color
	context.Set("fillStyle", "black") // You can change the text color as needed

	// Draw the text at the specified position
	context.Call("fillText", text, float64(x), float64(y))
}
