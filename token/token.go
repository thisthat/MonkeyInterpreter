// Package token provides the structure of the token used in the Lexer/Parser
package token

// We define tokens with a string which identifies their name
type TokenType string

// Token structure
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
	File    string
}

// Constant which identifies the token type
const (
	ILLEGAL = "ILLEGAL" // Unexpected symbols
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MUL    = "*"
	DIV    = "/"
	MOD    = "%"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// Pars
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "fn"
	LET      = "let"
)
