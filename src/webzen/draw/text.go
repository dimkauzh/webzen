package draw

import "webzen/src/backend/document"

func DrawText(text string) {
	p := document.CreateElement("p")

	p.Set("innerText", "Hello, World!")
	document.Body.AppendChild(p)

}
