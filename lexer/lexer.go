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

	l.skipWhites()

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
	case '-':
		t = token.MINUS
	case '/':
		t = token.DIV
	case '*':
		t = token.MUL
	case '%':
		t = token.MOD
	case '{':
		t = token.LBRACE
	case '}':
		t = token.RBRACE
	case 0:
		t = token.EOF
	default:
		//handle strings
		if isLetter(l.ch) {
			lit := l.readIdentifier()
			return token.Token{Type: token.LookupTypeIdent(lit), Literal: lit}
		} else if isDigit(l.ch) {
			lit := l.readNumber()
			return token.Token{Type: token.INT, Literal: lit}
		} else {
			t = token.ILLEGAL
		}
	}
	lit := string(l.ch)
	if lit == "\x00" {
		lit = ""
	}
	tok := token.Token{Type: t, Literal: lit}
	l.readchar()
	return tok
}
func (l *Lexer) skipWhites() {
	for isWhite(l.ch) {
		l.readchar()
	}

}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readchar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readchar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isWhite(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}
