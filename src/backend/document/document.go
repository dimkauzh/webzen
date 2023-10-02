package document

import "syscall/js"

type _body struct{}

// Document is a wrapper for the JavaScript Document object.

func CreateElement(element string) js.Value {
	return js.Global().Get("document").Call("createElement", element)
}

func (b *_body) AppendChild(appendChild js.Value) {
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

var Body = _body{}
