package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xiaosao/xq.git/src/scanner"
)

func main() {
	args := os.Args
	argsLen := len(args)
	xq := Xq{}
	if argsLen > 2 {
		println("Usage: xq [script]")
	} else if argsLen == 2 {
		xq.runFile(args[1])
	} else {
		xq.runPrompt()
	}
}

type Xq struct{}

func (xq *Xq) runFile(scriptPath string) {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		println("Read file wrong", err)
	}
	xq.run(string(content))
}

func (xq *Xq) runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		println("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		xq.run(line)

	}
}

func (xq *Xq) run(source string) {
	scanner := scanner.Scanner{
		Source: source,
	}
	tokens := scanner.ScanTokens()
	for token := range tokens {
		println(token)
	}
}

func (xq *Xq) error(line int, message string) {
	xq.report(line, "", message)
}

func (xq *Xq) report(line int, where, message string) {
	fmt.Printf("[line %d] Error %s : %s\n", line, where, message)
}
