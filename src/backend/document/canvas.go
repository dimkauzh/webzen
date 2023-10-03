//go:build js && wasm
// +build js,wasm

package document

import "syscall/js"

type canvas struct {
	canvas js.Value
}

type context struct {
	context js.Value
}

func (c *canvas) GetContext(context_str string) context {
	return context{c.canvas.Call("getContext", context_str)}
}

func (c *canvas) Get(key string) js.Value {
	return c.canvas.Get(key)
}

func (c *canvas) Set(key string, value interface{}) {
	c.canvas.Set(key, value)
}

func (c *canvas) IsNull() bool {
	return c.canvas.IsNull()
}

func (c *context) Set(key string, value interface{}) {
	c.context.Set(key, value)
}

func (c *context) FillText(text string, x, y float64) {
	c.context.Call("fillText", text, x, y)
}

func (c *context) FillRect(x, y, width, height float64) {
	c.context.Call("fillRect", x, y, width, height)
}

func (c *context) DrawImage(img js.Value, x, y, width, height float64) {
	c.context.Call("drawImage", img, x, y, width, height)
}
