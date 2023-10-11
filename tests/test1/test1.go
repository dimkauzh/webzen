package main

import (
	"github.com/dimkauzh/webzen/src/webzen"
	"github.com/dimkauzh/webzen/src/webzen/draw"
	"github.com/dimkauzh/webzen/src/webzen/keys"
	"github.com/dimkauzh/webzen/src/webzen/tools"
	"github.com/dimkauzh/webzen/src/webzen/ui"
)

func main() {
	webzen.Init()
	for {
		draw.FillBackground([4]int{255, 255, 188, 255})
		draw.DrawText("Hello under the world!", 21, 100, 100)
		draw.DrawRect(50, 500, 400, 400, [4]int{146, 255, 123, 255})
		draw.DrawRect(200, 200, 100, 400, [4]int{146, 255, 123, 255})

		ui.NewButton("texting rrnrrh", 10, 100, 100, 200, 100, func() {
			tools.Print("Button pressed")
		})

		if keys.KeyPressedOnce("a") {
			tools.Print("A key pressed")
		}
		webzen.Update()
	}
}
