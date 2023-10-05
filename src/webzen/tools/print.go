//go:build js && wasm
// +build js,wasm

package tools

import "github.com/dimkauzh/webzen/src/backend/console"

func Print(args ...interface{}) {
	console.Log(args)
}
