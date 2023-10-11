//go:build js && wasm
// +build js,wasm

package ui

import (
	"strconv"
	"syscall/js"

	"github.com/dimkauzh/webzen/src/backend/colors"
	"github.com/dimkauzh/webzen/src/backend/document"
)

var buttonClicked = false

func NewButton(text string, textSize int, x, y, width, height float64, clickCallback func()) {
	canvas := document.GetElementById("webzen")
	context := canvas.GetContext("2d")

	rgba := colors.GetRGBA([4]int{146, 255, 123, 255})

	context.Set("fillStyle", rgba)
	context.FillRect(x, y, width, height)

	textX := x + width/2 - float64(textSize*len(text)/4)
	textY := y + height/2 + float64(textSize)/3

	context.Set("font", strconv.Itoa(textSize)+"px Arial")
	context.Set("fillStyle", "black")
	context.FillText(text, textX, textY)

	if !buttonClicked {
		canvas.AddEventListener("click", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
			// Check if the click event occurred within the button area
			mouseEvent := p[0]
			mouseX := mouseEvent.Get("offsetX").Int()
			mouseY := mouseEvent.Get("offsetY").Int()

			if mouseX >= int(x) && mouseX <= int(x)+int(width) && mouseY >= int(y) && mouseY <= int(y)+int(height) {
				// Call the provided clickCallback function when the button area is clicked
				clickCallback()
			}
			return nil
		}))

		buttonClicked = true
	}
}
