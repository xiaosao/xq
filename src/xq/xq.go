package xq

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xiaosao/xq.git/src/scanner"
)

type Xq struct {
	hadError bool
}

func NewXq() *Xq {
	return &Xq{}
}

func (xq *Xq) RunFile(scriptPath string) {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		println("Read file wrong", err)
	}
	xq.Run(string(content))
	if xq.hadError {
		os.Exit(65)
	}
}

func (xq *Xq) RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		xq.Run(line)
		xq.hadError = false
	}
}

func (xq *Xq) Run(source string) {
	scanner := scanner.Scanner{
		Source: source,
	}
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		println(token.ToString())
	}
}

func (xq *Xq) Error(line int, message string) {
	xq.report(line, "", message)
}

func (xq *Xq) report(line int, where, message string) {
	fmt.Printf("[line %d] Error %s : %s\n", line, where, message)
}
