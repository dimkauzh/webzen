package main

import (
	// Import the webzen package
	"github.com/dimkauzh/webzen"

	// Import the webzen draw package
	"github.com/dimkauzh/webzen/src/draw"
)

func main() {
	// Initialize webzen
	webzen.Init()

	// Start loop
	for {
		// Fill background
		draw.FillBackground([4]int{255, 255, 188, 255})

		// Draw text on the screen
		draw.DrawText("Hello, world!", 21, 100, 100)

		// Draw a rectangle on the screen
		draw.DrawRect(50, 500, 400, 400, [4]int{146, 255, 123, 255})

		// Draw a second rectangle on the screen
		draw.DrawRect(200, 200, 100, 400, [4]int{146, 255, 123, 255})

		// Update
		webzen.Update()
	}
}
