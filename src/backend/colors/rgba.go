//go:build js && wasm
// +build js,wasm

package colors

import "strconv"

func GetRGBA(color [4]int) string {
	return "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
}
