//go:build js && wasm
// +build js,wasm

package image

import (
	"syscall/js"

	"github.com/dimkauzh/webzen/src/backend/document"
	"github.com/dimkauzh/webzen/src/backend/global"
)

type Image struct {
  imagePath string
  width float64
  height float64
  x float64
  y float64
}

func LoadImage(imagePath string, width, height, x, y float64) Image {
  return Image{imagePath, width, height, x, y}
}

func (i *Image) Draw() {
  DrawImage(i.imagePath, i.width, i.height, i.x, i.y)
}

func DrawImage(imagePath string, width, height, x, y float64) {
	canvas := document.GetElementById("canvas")
	context := canvas.GetContext("2d")

	// Load and draw the image
	img := document.CreateElement("img")
	img.Set("src", imagePath)
	img.AddEventListener("load", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		context.DrawImage(img, x, y, width, height)
		return nil
	}))

	// Handle image loading errors
	img.AddEventListener("error", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		global.Alert("Failed to load image: " + imagePath)
		return nil
	}))
}
