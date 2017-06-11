// Package lexer will provide a basic fashionless regexpr free lexer
// currently from the text it seems to be a LL(1) type
package lexer

import (
	"github.com/thisthat/MonkeyInterpreter/token"
)

// Lexer Structure which holds the status
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	col          int
	// \forall position . ch = input[position]
	// \forall position . position = readPosition + 1
}

// New create the Lexer Object
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, col: 0}
	l.readchar()
	return l
}

// readchar read the next char from the input and move forward
func (l *Lexer) readchar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.col++
}

// peekChar read the next char in the input without then moving to the next one
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]

}

// NextToken gives the next Token from the input
func (l *Lexer) NextToken() token.Token {
	var t token.TokenType

	l.skipWhites()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			l.readchar()
			return token.Token{Type: token.EQ, Literal: lit, Line: l.line, Col: l.col - len(lit)}
		}
		t = token.ASSIGN

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readchar()
			lit := string(ch) + string(l.ch)
			l.readchar()
			return token.Token{Type: token.NEQ, Literal: lit, Line: l.line, Col: l.col - len(lit)}
		}
		t = token.BANG

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
	case '<':
		t = token.LT
	case '>':
		t = token.GT
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
			return token.Token{Type: token.LookupTypeIdent(lit), Literal: lit, Line: l.line, Col: l.col - len(lit)}
		} else if isDigit(l.ch) {
			lit := l.readNumber()
			return token.Token{Type: token.INT, Literal: lit, Line: l.line, Col: l.col - len(lit)}
		} else {
			t = token.ILLEGAL
		}
	}
	lit := string(l.ch)
	ll := l.line
	cc := l.col
	if lit == "\x00" {
		lit = ""
		ll = -1
		cc = -1
	}
	tok := token.Token{Type: t, Literal: lit, Line: ll, Col: cc}
	l.readchar()
	return tok
}

func (l *Lexer) skipWhites() {
	for isWhite(l.ch) {
		if isNewLine(l.ch) {
			l.line++
			l.col = 0
		}
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
	return ch == ' ' || ch == '\t' || ch == '\r' || isNewLine(ch)
}
func isNewLine(ch byte) bool {
	return ch == '\n'
}
