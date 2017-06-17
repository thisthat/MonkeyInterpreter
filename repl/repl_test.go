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

func TestREPL_Fibo(t *testing.T) {
	get := helper("../testresources/smallPrograms/Fibonacci.monkey", t)
	exp := []string{
		"let fibonacci = fn(x) if(x == 0) return 0;else if(x == 1) return 1;else return (fibonacci((x - 1)) + fibonacci((x - 2)));;",
		"",
	}
	check(t, get, exp)
}

// Not supported yet
/*func TesttREPL_Let(t *testing.T) {
	get := helper("../testresources/smallPrograms/Let.monkey", t)
	exp := []string{
		"let age = 1;",
		"",
		"",
		"",
		"",
	}
	check(t, get, exp)
}*/

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
