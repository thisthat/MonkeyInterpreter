package lexer

import (
	"github.com/thisthat/MonkeyInterpreter/token"
	"io/ioutil"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
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

func TestAdd(t *testing.T) {
	filename := "../testresources/smallPrograms/Add.monkey"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := string(buf)
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - Wrong Token %q. Expected=%q, got=%q", i, tok, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Wrong Literal. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSignleToken(t *testing.T) {
	filename := "../testresources/smallPrograms/Token.monkey"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := string(buf)
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.MOD, "%"},
		{token.PLUS, "+"},

		{token.LT, "<"},
		{token.GT, ">"},
		{token.ASSIGN, "="},
		{token.BANG, "!"},
		{token.NEQ, "!="},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},

		{token.INT, "9"},
		{token.NEQ, "!="},
		{token.INT, "10"},

		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] - Wrong Token %q. Expected=%q, got=%q", i, tok, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Test[%d] - Wrong Literal. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLineToken(t *testing.T) {
	filename := "../testresources/smallPrograms/Token.monkey"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := string(buf)
	tests := []struct {
		expectedCol  int
		expectedLine int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 1},

		{1, 2},
		{2, 2},
		{3, 2},
		{4, 2},
		{5, 2},
		//10
		{1, 4},
		{4, 4},
		{5, 4},
		{6, 4},
		{7, 4},
		{9, 4},
		{11, 4},

		{5, 5},
		{12, 5},
		{16, 5},
		//20
		{1, 6},
		{3, 6},
		{8, 6},

		{5, 7},
		{12, 7},
		{17, 7},

		{1, 8},

		{1, 10},
		{4, 10},
		{7, 10},
		//30
		{1, 11},
		{3, 11},
		{6, 11},

		{-1, -1},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Line != tt.expectedLine {
			t.Fatalf("Test[%d] - Wrong Line. Expected=%d, got=%d", i, tt.expectedLine, tok.Line)
		}
		if tok.Col != tt.expectedCol {
			t.Fatalf("Test[%d] - Wrong Column. Expected=%d, got=%d", i, tt.expectedCol, tok.Col)
		}
	}
}

func TestNotValidToken(t *testing.T) {
	input := "!=<@#>"
	tests := []struct {
		expType token.TokenType
		expLit  string
	}{
		{token.NEQ, "!="},
		{token.LT, "<"},
		{token.ILLEGAL, "@"},
		{token.ILLEGAL, "#"},
		{token.GT, ">"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Literal != tt.expLit {
			t.Fatalf("Test[%d] - Wrong Literal. Expected=%q, got=%q", i, tt.expLit, tok.Literal)
		}
		if tok.Type != tt.expType {
			t.Fatalf("Test[%d] - Wrong Type. Expected=%q, got=%q", i, tt.expType, tok.Type)
		}
	}
}
