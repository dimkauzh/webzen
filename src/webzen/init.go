package webzen

import "syscall/js"

func Init() {
	var canvas js.Value
	// Create a canvas element if it doesn't exist
	canvas = js.Global().Get("document").Call("getElementById", "canvas")
	if canvas.IsNull() {
		canvas = js.Global().Get("document").Call("createElement", "canvas")
		canvas.Set("id", "canvas")
		js.Global().Get("document").Get("body").Call("appendChild", canvas)
	}

	// Set the canvas size to match the viewport size
	canvas.Set("width", js.Global().Get("document").Get("documentElement").Get("clientWidth"))
	canvas.Set("height", js.Global().Get("document").Get("documentElement").Get("clientHeight"))

	// Prevent scrolling on the HTML and body elements
	js.Global().Get("document").Get("documentElement").Get("style").Set("overflow", "hidden")
	js.Global().Get("document").Get("body").Get("style").Set("overflow", "hidden")

	// Listen for window resize events to adjust the canvas size
	js.Global().Get("window").Call("addEventListener", "resize", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		canvas.Set("width", js.Global().Get("document").Get("documentElement").Get("clientWidth"))
		canvas.Set("height", js.Global().Get("document").Get("documentElement").Get("clientHeight"))
		return nil
	}))
}
