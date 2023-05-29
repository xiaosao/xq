package xq

import (
	sc "github.com/xiaosao/xq.git/src/scanner"
)

type Expr struct {
}

type Binary struct {
	Expr
	Left     Expr
	Operator sc.Token
	Right    Expr
}

func (t *Binary) accept(visitor Visitor) {
	visitor.visitBinary(t)
}

type Grouping struct {
	Expr
	Expression Expr
}

func (t *Grouping) accept(visitor Visitor) {
	visitor.visitGrouping(t)
}

type Literal struct {
	Expr
	Value sc.Object
}

func (t *Literal) accept(visitor Visitor) {
	visitor.visitLiteral(t)
}

type Unary struct {
	Expr
	Operator sc.Token
	Right    Expr
}

func (t *Unary) accept(visitor Visitor) {
	visitor.visitUnary(t)
}

type Visitor interface {
	visitBinaryExpr(Expr Binary)
	visitGroupingExpr(Expr Grouping)
	visitLiteralExpr(Expr Literal)
	visitUnaryExpr(Expr Unary)
}
