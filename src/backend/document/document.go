//go:build js && wasm
// +build js,wasm

package document

import "syscall/js"

var Body = body{}
var DocumentElement = documentElement{}

type body struct {
	Style styleBody
}

type element struct {
	el    js.Value
	Style js.Value
}

type documentElement struct {
	Style styleDocument
}
type styleDocument struct{}
type styleBody struct{}

func CreateElement(element_str string) element {
	cd := js.Global().Get("document").Call("createElement", element_str)
	style := cd.Get("style")
	return element{cd, style}
}

func CreateCanvasElement() canvas {
	cd := js.Global().Get("document").Call("createElement", "canvas")
	return canvas{cd}
}

func AddEventListener(event string, f js.Func) {
	js.Global().Get("document").Call("addEventListener", event, f)
}

func (e *element) Set(key string, value interface{}) {
	e.el.Set(key, value)
}

func (e *element) Call(type1, type2 string, f js.Func) {
	e.el.Call(type1, type2, f)
}

func (e *element) AddEventListener(event string, f js.Func) {
	e.el.Call("addEventListener", event, f)
}

func (b *body) AppendChild(el element) {
	appendChild := el.el
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

func (b *body) AppendCanvasChild(ca canvas) {
	appendChild := ca.canvas
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

func GetElementById(id string) canvas {
	return canvas{js.Global().Get("document").Call("getElementById", "canvas")}
}

func (d *documentElement) ClientWidth() js.Value {
	return js.Global().Get("document").Get("documentElement").Get("clientWidth")
}

func (d *documentElement) ClientHeight() js.Value {
	return js.Global().Get("document").Get("documentElement").Get("clientHeight")
}

func (d *documentElement) ClientIntWidth() int {
	return js.Global().Get("document").Get("documentElement").Get("clientWidth").Int()
}

func (d *documentElement) ClientIntHeight() int {
	return js.Global().Get("document").Get("documentElement").Get("clientHeight").Int()
}

func (d *styleDocument) Set(key string, value interface{}) {
	js.Global().Get("document").Get("documentElement").Get("style").Set(key, value)
}

func (b *styleBody) Set(key string, value interface{}) {
	js.Global().Get("document").Get("body").Get("style").Set(key, value)
}
