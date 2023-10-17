//go:build js && wasm
// +build js,wasm

package entity

import (
    "github.com/dimkauzh/webzen/src/image"
    "github.com/dimkauzh/webzen/src/shape"
    "github.com/dimkauzh/webzen/src/vector"
)

type Entity struct {
    pos    vector.Vector2D
    width  int
    height int
    image  image.Image
    rect   shape.Rect
}

func NewEntity(x int, y int, width int, height int) Entity {
    return Entity{vector.NewVector2D(x, y), width, height, nil, nil}
}

func (e *Entity) SetImage(image image.Image) {
    e.image = image
}

func (e *Entity) SetRect(rect shape.Rect) {
    e.rect = rect
}
