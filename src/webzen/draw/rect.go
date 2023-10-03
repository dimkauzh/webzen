package draw

import (
	"strconv"
	"webzen/src/backend/document"
)

func DrawRect(width, height, x, y float64, color [4]int) {
	canvas := document.GetElementById("canvas")
	// Get the canvas context
	context := canvas.GetContext("2d")

	// Draw the rectangle
	rgba := "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
	context.Set("fillStyle", rgba)
	context.FillRect(x, y, width, height) // Set the rectangle dimensions as needed
}

func FillBackground(color [4]int) {
	canvas := document.GetElementById("canvas")
	// Get the canvas context
	context := canvas.GetContext("2d")

	// Fill the canvas with the specified color
	rgba := "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
	context.Set("fillStyle", rgba)
	context.FillRect(0, 0, canvas.Get("width").Float(), canvas.Get("height").Float())
}
