// Will provide a basic fashionless regexpr free lexer
// currently from the text it seems to be a LL(1) type
package lexer

import "github.com/thisthat/MonkeyInterpreter/token"

type Lexer struct {
	input 			string
	position 		int
	readPosition	int
	ch	 			byte
	// \forall position . ch = input[position]
	// \forall position . position = readPosition + 1
}


// Create the Lexer Object
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

// Read the next char from the input
func (l *Lexer) readchar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Gimme the next Token from the input
func (l *Lexer) NextToken() token.Token  {
	var tok token.Token
	tok.Literal = "\\eof"
	tok.Type = token.EOF
	return tok
}