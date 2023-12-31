package main

import (
	"github.com/dimkauzh/webzen/src/backend/console"
	"github.com/dimkauzh/webzen/src/backend/document"
)

func main() {
	console.Log("Hello, World!")

	p := document.CreateElement("p")

	p.Set("innerText", "Hello, World!")
	document.Body.AppendChild(p)

	pre := document.CreateElement("pre")
	pre.Set("innerText", "Test")
	pre.Style.Set("whiteSpace", "pre-wrap")
	document.Body.AppendChild(pre)
}
