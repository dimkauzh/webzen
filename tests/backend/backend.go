package main

import (
	"webzen/src/backend/console"
	"webzen/src/backend/document"
)

func main() {
	console.Log("Hello, World!")

	p := document.CreateElement("p")

	p.Set("innerText", "Hello, World!")
	document.Body.AppendChild(p)
}

// func main() {
// 	p := js.Global().Get("document").Call("createElement", "p")
// 	p.Set("innerText", "Hello, World!")
// 	js.Global().Get("document").Get("body").Call("appendChild", p)

// 	pre := js.Global().Get("document").Call("createElement", "pre")
// 	pre.Set("innerText", fmt.Sprintf("args: %q\nevn: %q", os.Args, os.Environ()))
// 	pre.Get("style").Set("whiteSpace", "pre-wrap")
// 	js.Global().Get("document").Get("body").Call("appendChild", pre)
// }
