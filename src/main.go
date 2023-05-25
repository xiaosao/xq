package main

import (
	"os"

	"github.com/xiaosao/xq.git/src/xq"
)

func main() {
	args := os.Args
	argsLen := len(args)
	myXq := xq.NewXq()
	if argsLen > 2 {
		println("Usage: xq [script]")
	} else if argsLen == 2 {
		myXq.RunFile(args[1])
	} else {
		myXq.RunPrompt()
	}
}
