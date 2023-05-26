package scanner

import (
	"fmt"
	"strconv"
)

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANGEQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FUN
	FOR
	IF
	WHILE
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	FALSE
	VAR

	EOF
)

// just for print format
var tokenType2Label = map[TokenType]string{
	LEFT_PAREN:  "LEFT_PAREN",
	RIGHT_PAREN: "RIGHT_PAREN",
	LEFT_BRACE:  "LEFT_BRACE",
	RIGHT_BRACE: "RIGHT_BRACE",
	COMMA:       "COMMA",
	DOT:         "DOT",
	MINUS:       "MINUS",
	PLUS:        "PLUS",
	SEMICOLON:   "SEMICOLON",
	SLASH:       "SLASH",
	STAR:        "STAR",

	// One or two character tokens.
	BANG:          "BANG",
	BANGEQUAL:     "BANGEQUAL",
	EQUAL:         "EQUAL",
	EQUAL_EQUAL:   "EQUAL_EQUAL",
	GREATER:       "GREATER",
	GREATER_EQUAL: "GREATER_EQUAL",
	LESS:          "LESS",
	LESS_EQUAL:    "LESS_EQUAL",

	// Literals.
	IDENTIFIER: "IDENTIFIER",
	STRING:     "STRING",
	NUMBER:     "NUMBER",

	// Keywords.
	AND:    "AND",
	CLASS:  "CLASS",
	ELSE:   "ELSE",
	FUN:    "FUN",
	FOR:    "FOR",
	IF:     "IF",
	WHILE:  "WHILE",
	NIL:    "NIL",
	OR:     "OR",
	PRINT:  "PRINT",
	RETURN: "RETURN",
	SUPER:  "SUPER",
	THIS:   "THIS",
	TRUE:   "TRUE",
	FALSE:  "FALSE",
	VAR:    "VAR",

	EOF: "EOF",
}

var Keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"spuer":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type ErrorReporter interface {
	error(line int, message string)
}

type Scanner struct {
	Source string
	Tokens []Token

	start         int
	current       int
	line          int
	errorReporter ErrorReporter
}

func NewScanner(source string, errorReporter ErrorReporter) *Scanner {
	return &Scanner{
		start:         0,
		current:       0,
		line:          1,
		Tokens:        []Token{},
		Source:        source,
		errorReporter: errorReporter,
	}
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.Source)
}

func (sc *Scanner) scanToken() {
	c := sc.advance()
	switch c {
	case '(':
		sc.addToken(LEFT_PAREN)
	case ')':
		sc.addToken(RIGHT_PAREN)
	case '{':
		sc.addToken(LEFT_BRACE)
	case '}':
		sc.addToken(RIGHT_BRACE)
	case ',':
		sc.addToken(COMMA)
	case '.':
		sc.addToken(DOT)
	case '-':
		sc.addToken(MINUS)
	case '+':
		sc.addToken(PLUS)
	case ';':
		sc.addToken(SEMICOLON)
	case '*':
		sc.addToken(STAR)
	case '!':
		if sc.match('=') {
			sc.addToken(BANGEQUAL)
		} else {
			sc.addToken(BANG)
		}
	case '=':
		if sc.match('=') {
			sc.addToken(EQUAL_EQUAL)
		} else {
			sc.addToken(EQUAL)
		}
	case '<':
		if sc.match('=') {
			sc.addToken(LESS_EQUAL)
		} else {
			sc.addToken(LESS)
		}
	case '>':
		if sc.match('=') {
			sc.addToken(GREATER_EQUAL)
		} else {
			sc.addToken(GREATER)
		}
	case '/':
		if sc.match('/') {
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else {
			sc.addToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		sc.line++
	case '"':
		sc.xString()
	default:
		if isDigit(c) {
			sc.xNumber()
		} else if isAlpha(c) {
			sc.xIdentifier()
		} else {
			sc.errorReporter.error(sc.line, "Unexpected character.")
		}
	}
}

// just forward 1
func (sc *Scanner) advance() rune {
	str := []rune(sc.Source)
	c := str[sc.current]
	sc.current += 1
	return c
}

// match and forward 1
func (sc *Scanner) match(expected rune) bool {
	if sc.isAtEnd() {
		return false
	}
	c := []rune(sc.Source)[sc.current]
	if c != expected {
		return false
	}
	sc.current++
	return true
}

// return current + 1
func (sc *Scanner) peek() rune {
	if sc.isAtEnd() {
		return '\x00'
	}
	return []rune(sc.Source)[sc.current]
}

// return current + 2
func (sc *Scanner) peekNext() rune {
	next := sc.current + 1
	if next >= len(sc.Source) {
		return '\x00'
	}
	return []rune(sc.Source)[next]
}

func (sc *Scanner) xString() {
	for sc.peek() != '"' && sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
	}
	if sc.isAtEnd() {
		sc.errorReporter.error(sc.line, "Unterminated string.")
		return
	}
	sc.advance()

	val := sc.Source[sc.start+1 : sc.current-1]
	sc.addToken1(STRING, val)
}

func (sc *Scanner) xNumber() {
	for isDigit(sc.peek()) {
		sc.advance()
	}
	if sc.peek() == '.' && isDigit(sc.peekNext()) {
		sc.advance()
		for isDigit(sc.peek()) {
			sc.advance()
		}
	}
	val, err := strconv.ParseFloat(sc.Source[sc.start:sc.current], 64)
	if err != nil {
		sc.errorReporter.error(sc.line, "Unexpected number format.")
	}
	sc.addToken1(NUMBER, val)
}

func (sc *Scanner) xIdentifier() {
	for isAlphaNumberic(sc.peek()) {
		sc.advance()
	}
	text := sc.Source[sc.start:sc.current]
	xType, exist := Keywords[text]
	if exist {
		sc.addToken(xType)
	} else {

		sc.addToken(IDENTIFIER)
	}

}

func (sc *Scanner) addToken(xtype TokenType) {
	sc.addToken1(xtype, nil)
}

func (sc *Scanner) addToken1(xtype TokenType, literal interface{}) {
	text := sc.Source[sc.start:sc.current]
	sc.Tokens = append(sc.Tokens, Token{
		Type:    xtype,
		Lexeme:  text,
		Literal: literal,
		Line:    sc.line,
	})
}

func (sc *Scanner) ScanTokens() []Token {
	for !sc.isAtEnd() {
		sc.start = sc.current
		sc.scanToken()
	}
	sc.Tokens = append(sc.Tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: sc.line})
	return sc.Tokens
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%s  %q  %q", tokenType2Label[t.Type], t.Lexeme, t.Literal)
}
