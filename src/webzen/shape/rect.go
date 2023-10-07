//go:build js && wasm
// +build js,wasm

package shape

import (
	"github.com/dimkauzh/webzen/src/backend/colors"
	"github.com/dimkauzh/webzen/src/backend/document"
)

type Rect struct {
	x      float64
	y      float64
	width  float64
	height float64
	color  [4]int
}

func NewRect(x, y, width, height float64, color [4]int) Rect {
	return Rect{x, y, width, height, color}
}

func (r *Rect) Draw() {
	canvas := document.GetElementById("canvas")
	context := canvas.GetContext("2d")

	rgba := colors.GetRGBA(r.color)
	context.Set("fillStyle", rgba)
	context.FillRect(r.x, r.y, r.width, r.height)
}
