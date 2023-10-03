package window

import "syscall/js"

func AddEventListener(ev string, fn js.Func) {
	js.Global().Get("window").Call("addEventListener", ev, fn)
}
