package draw

import (
	"strconv"
	"webzen/src/backend/document"
)

func DrawText(text string, size, x, y int) {
	canvas := document.GetElementById("canvas")
	context := canvas.GetContext("2d")

	context.Set("font", strconv.Itoa(size)+"px Arial")
	context.Set("fillStyle", "black")
	context.FillText(text, float64(x), float64(y))
}
