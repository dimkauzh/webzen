package image

import "syscall/js"

func DrawImage(imagePath string, width, height, x, y float64) {
	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	context := canvas.Call("getContext", "2d")

	// Load and draw the image
	img := js.Global().Get("document").Call("createElement", "img")
	img.Set("src", imagePath)
	img.Call("addEventListener", "load", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		context.Call("drawImage", img, x, y, width, height)
		return nil
	}))

	// Handle image loading errors
	img.Call("addEventListener", "error", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		js.Global().Call("alert", "Failed to load image")
		return nil
	}))
}
