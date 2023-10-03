package main

import (
	"webzen/src/webzen"
	"webzen/src/webzen/draw"
)

func main() {
	webzen.Init()
	draw.DrawTitleText("Hello, World!")
	draw.DrawText("Hello under the world!")
	draw.DrawRect(500, 500, 400, 400, [4]int{146, 255, 123, 255})
}
