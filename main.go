package main

import (
	"bufio"
	"os"

	"github.com/xiaosao/xq.git/src/scanner"
)

func main() {
	args := os.Args
	argsLen := len(args)
	if argsLen > 2 {
		println("Usage: xq [script]")
	} else if argsLen == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(scriptPath string) {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		println("Read file wrong", err)
	}
	run(string(content))
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		println("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		run(line)

	}
}

func run(source string) {
	scanner := scanner.Scanner{
		Source: source,
	}
	tokens := scanner.ScanTokens()
	for token := range tokens {
		println(token)
	}
}
