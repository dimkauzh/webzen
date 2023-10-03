package global

import "syscall/js"

func Alert(text string) {
	js.Global().Call("alert", text)
}
