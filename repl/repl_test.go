package repl

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestREPL_add(t *testing.T) {
	get := helper("../testresources/smallPrograms/Add.monkey", t)
	exp := []string{
		"let five = 5;",
		"let ten = 10;",
		"let add = fn(x, y) (x + y);",
		"let result = add(five, ten);",
		"",
	}
	check(t, get, exp)
}

func TestREPL_fibo(t *testing.T) {
	get := helper("../testresources/smallPrograms/Fibonacci.monkey", t)
	exp := []string{
		"let fibonacci = fn(x) if(x == 0) return 0;else if(x == 1) return 1;else return (fibonacci((x - 1)) + fibonacci((x - 2)));;",
		"",
	}
	check(t, get, exp)
}

// Not supported yet it will fail
func TestREPL_Let(t *testing.T) {
	get := helperV2("../testresources/smallPrograms/Let.monkey", t)
	exp := []string{
		"let age = 1;",
		"\tno prefix parse function for ILLEGAL found",
		"\tno prefix parse function for ILLEGAL found",
		"let result = (10 * (20 / 2));",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for , found",
		"no prefix parse function for , found",
		"no prefix parse function for , found",
		"no prefix parse function for , found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for { found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for , found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for ILLEGAL found",
		"no prefix parse function for } found",
		"",
	}
	check(t, get, exp)
}


func TestREPL_logo(t *testing.T) {
	input := "let"
	output := ""
	o := bytes.NewBufferString(output)
	Start(bytes.NewReader([]byte(input)), o)
	exp := MonkeyFace + ErrorMsg +
		"\texpected next token to be IDENT, got EOF instead\n"
	if exp != o.String() {
		t.Fatalf("We do not get the expected error message")
	}
}

func check(t *testing.T, get []string, exp []string) {
	assert.Equalf(t, len(get), len(exp), "Expected size: %d - get %d", len(exp), len(get))
	for i := range exp {
		assert.EqualValues(t, len(get), len(exp), "Input %d:. Expected [%s] - Get [%s]", i, exp[i], get[i])
	}
}

func helper(filename string, t *testing.T) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := string(buf)
	output := ""
	o := bytes.NewBufferString(output)
	Start(bytes.NewReader([]byte(input)), o)
	//t.Fatalf("%s\n", o)
	return strings.Split(o.String(), "\n")
}

func helperV2(filename string, t *testing.T) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	input := string(buf)
	output := ""
	o := bytes.NewBufferString(output)
	StartHelp(bytes.NewReader([]byte(input)), o, false)
	//t.Fatalf("%s\n", o)
	return strings.Split(o.String(), "\n")
}
