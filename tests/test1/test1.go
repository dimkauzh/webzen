package main

import (
	"webzen/src/webzen"
	"webzen/src/webzen/draw"
)

func main() {
	webzen.Init()
	draw.FillBackground([4]int{255, 255, 188, 255})
	draw.DrawText("Hello under the world!", 21, 100, 100)
	draw.DrawRect(50, 500, 400, 400, [4]int{146, 255, 123, 255})
	draw.DrawRect(200, 200, 100, 400, [4]int{146, 255, 123, 255})
}
