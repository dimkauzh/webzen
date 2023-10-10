//go:build js && wasm
// +build js,wasm

package entity

import (
  "github.com/dimkauzh/webzen/src/webzen/image"
  rect "github.com/dimkauzh/webzen/src/webzen/shape"
)

type Entity struct {
  x      int
  y      int
  width  int
  height int
  image  image.Image
  rect   rect.Rect
}

func NewEntity(x int, y int, width int, height int) Entity {
  return Entity{x, y, width, height, nil, nil}
}

func (e *Entity) SetImage(image image.Image) {
  e.image = image
}

func (e *Entity) SetRect(rect rect.Rect) {
  e.rect = rect
}
