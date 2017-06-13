package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	output := ""
	o := bytes.NewBufferString(output)
	Start(bytes.NewReader([]byte(input)), o)
	tests := []struct {
		expected string
	}{
		{"Tok: {Type:= Literal:= Line:1 Col:1}"},
		{"Tok: {Type:+ Literal:+ Line:1 Col:2}"},
		{"Tok: {Type:( Literal:( Line:1 Col:3}"},
		{"Tok: {Type:) Literal:) Line:1 Col:4}"},
		{"Tok: {Type:{ Literal:{ Line:1 Col:5}"},
		{"Tok: {Type:} Literal:} Line:1 Col:6}"},
		{"Tok: {Type:, Literal:, Line:1 Col:7}"},
		{"Tok: {Type:; Literal:; Line:1 Col:8}"},
	}
	results := strings.Split(o.String(), "\n")

	if len(tests) != len(results)-1 {
		t.Fatalf("Expected %d test, get %d", len(tests), len(results)-1)
	}

	for i, r := range tests {
		if r.expected != results[i] {
			t.Fatalf("%d - Expected %s test, get %s", i, r.expected, results[i])
		}
	}

}
