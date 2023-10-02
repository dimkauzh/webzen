package document

import "syscall/js"

var Body = body{}

type body struct{}

type element struct {
	el    js.Value
	Style js.Value
}

func CreateElement(element_str string) element {
	cd := js.Global().Get("document").Call("createElement", element_str)
	style := cd.Get("style")
	return element{cd, style}
}

func (e *element) Set(key string, value interface{}) {
	e.el.Set(key, value)
}

func (b *body) AppendChild(el element) {
	appendChild := el.el
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}
