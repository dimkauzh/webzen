//go:build js && wasm
// +build js,wasm

package vector

type Vector2D struct {
    x int
    y int
}

func NewVector2D(x, y int) Vector2D {
    return Vector2D{x, y}
}
