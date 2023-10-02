package document

import "syscall/js"

var Body = _body{}

type _body struct{}

type _element struct {
	el    js.Value
	Style js.Value
}

func CreateElement(element string) _element {
	cd := js.Global().Get("document").Call("createElement", element)
	style := cd.Get("style")
	return _element{cd, style}
}

func (e *_element) Set(key string, value interface{}) {
	e.el.Set(key, value)
}

func (b *_body) AppendChild(element _element) {
	appendChild := element.el
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}
