package xq

import (
	sc "github.com/xiaosao/xq.git/src/scanner"
)

type Expr struct {
}

type Binary struct {
	Expr
	left     Expr
	operator sc.Token
	right    Expr
}

type Grouping struct {
	Expr
	expression Expr
}

type Literal struct {
	Expr
	value sc.Object
}

type Unary struct {
	Expr
	operator sc.Token
	right    Expr
}
