//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("WebAssembly Loaded!")

	// JavaScript callback example
	js.Global().Set("sayHello", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		message := args[0].String()
		js.Global().Get("alert").Invoke("Hello from Go: " + message)
		return nil
	}))

	select {}
}
