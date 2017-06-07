package lexer

import  (
	"testing"
	"github.com/thisthat/MonkeyInterpreter/token"
)

func TestNextToken(t *testing.T){
	input := "=+(){},;"
	tests := []struct {
		expectedType 	token.TokenType
		expectedLiteral	string
	}{
		{token.ASSIGN,"="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input);
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - Wrong Token. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Wrong Literal. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}