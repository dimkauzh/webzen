package draw

import "webzen/src/backend/document"

func DrawTitleText(text string) {
	p := document.CreateElement("p")

	p.Set("innerText", "Hello, World!")
	document.Body.AppendChild(p)
}

func DrawText(text string) {
	p := document.CreateElement("pre")
	p.Set("innerText", text)
	p.Style.Set("whiteSpace", "pre-wrap")
	document.Body.AppendChild(p)
}
