// Package lexer will provide a basic fashionless regexpr free lexer
// currently from the text it seems to be a LL(1) type
package lexer

import "github.com/thisthat/MonkeyInterpreter/token"

// Lexer Structure which holds the status
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	// \forall position . ch = input[position]
	// \forall position . position = readPosition + 1
}

// New create the Lexer Object
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

// readchar read the next char from the input
func (l *Lexer) readchar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken gives the next Token from the input
func (l *Lexer) NextToken() token.Token {
	var t token.TokenType
	switch l.ch {
	case '=':
		t = token.ASSIGN
	case ';':
		t = token.SEMICOLON
	case '(':
		t = token.LPAREN
	case ')':
		t = token.RPAREN
	case ',':
		t = token.COMMA
	case '+':
		t = token.PLUS
	case '{':
		t = token.LBRACE
	case '}':
		t = token.RBRACE
	case 0:
		t = token.EOF
	}
	tok := token.Token{Type: t, Literal: string(l.ch)}
	l.readchar()
	return tok
}
