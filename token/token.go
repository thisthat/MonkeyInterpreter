// Package token provides the structure of the token used in the Lexer/Parser
package token

// TokenType defines tokens with a string which identifies their name
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
	LT     = "<"
	GT     = ">"
	BANG   = "!"

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
	TRUE     = "true"
	FALSE    = "false"
	RETURN   = "return"
	IF       = "if"
	ELSE     = "else"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

// LookupTypeIdent returns the token type for the keywords or IDENT
func LookupTypeIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
