package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("hello world")
	js.Global().Set("whatever", whateverFactory())
	<-make(chan struct{})
}

func whateverFactory() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		return "whatever, man"
	})
	return jsonFunc
}
