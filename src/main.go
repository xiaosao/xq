package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// args := os.Args
	// argsLen := len(args)
	// myXq := xq.NewXq()
	// if argsLen > 2 {
	// 	println("Usage: xq [script]")
	// } else if argsLen == 2 {
	// 	myXq.RunFile(args[1])
	// } else {
	// 	myXq.RunPrompt()
	// }

	generateAst()
}

func generateAst() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}
	outputDir := args[1]
	// fmt.Println(outputDir)
	defineAst(outputDir, "Expr", []string{
		"Binary: Left Expr, Operator sc.Token, Right Expr",
		"Grouping: Expression Expr",
		"Literal: Value sc.Object",
		"Unary: Operator sc.Token, Right Expr",
	})
}

func defineAst(outputDir string, baseName string, types []string) {

	path := fmt.Sprintf("%s/%s.go", outputDir, baseName)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("OpenFile error.\n%q", err)
	}
	defer file.Close()

	packageTemplate := fmt.Sprintf(`
package xq

import (
	sc "github.com/xiaosao/xq.git/src/scanner" 
)

type %s struct {

}
	`, baseName)
	file.Write([]byte(packageTemplate))

	for _, t := range types {
		segments := strings.Split(t, ":")
		className := segments[0]
		fields := segments[1]
		defineStructType(file, baseName, className, fields)
	}
	defineVisitor(file, baseName, types)

}

func defineStructType(f *os.File, baseName, className, fields string) {
	fieldList := strings.Split(fields, ",")
	fieldResultStr := ""
	for _, fi := range fieldList {
		fieldSeg := strings.Split(strings.Trim(fi, " "), " ")
		fieldResultStr += fmt.Sprintf("%s %s\n", fieldSeg[0], fieldSeg[1])
	}
	structTmp := fmt.Sprintf(`
type %s struct {
	%s
	%s
}
func (t *%s)accept(visitor Visitor) {
	visitor.visit%s(t)
}
	`, className, baseName, fieldResultStr, className, className)
	f.Write([]byte(structTmp))
}

func defineVisitor(f *os.File, baseName string, types []string) {
	f.Write([]byte("type Visitor interface {"))
	for _, t := range types {
		typeName := strings.Split(t, ":")[0]
		m := fmt.Sprintf("visit%s%s (%s %s)\n", typeName, baseName, baseName, typeName)
		f.Write([]byte(m))
	}
	f.Write([]byte("}"))
}

/*
	todo:
		1. 生成 Expr 文件，只是包含了各种非终结符的结构实现
		2. 实现 visitor 模式
		3. 实现 pretty printer 实现 visitor interface
		4.
*/
