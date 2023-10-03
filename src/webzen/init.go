package webzen

import "syscall/js"

func Init() {
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	canvas.Set("id", "canvas")
	js.Global().Get("document").Get("body").Call("appendChild", canvas)
}
