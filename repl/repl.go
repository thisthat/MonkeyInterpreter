package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/thisthat/MonkeyInterpreter/lexer"
	"github.com/thisthat/MonkeyInterpreter/parser"
)

// Start starts the Read Evaluate Print Loop
func Start(in io.Reader, out io.Writer) {
	StartHelp(in,out,true)
}

// StartHelp starts the Read Evaluate Print Loop
func StartHelp(in io.Reader, out io.Writer, logo bool) {
	scanner := bufio.NewScanner(in)
	const PROMPT = ">> "
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors(), logo)
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string, logo bool) {
	if logo {
		io.WriteString(out, MonkeyFace)
		io.WriteString(out, ErrorMsg)
	}
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

// MonkeyFace is an ASCII logo of the language
const MonkeyFace = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// ErrorMsg is used to print an error msg
const ErrorMsg = `
Woops! We ran into some monkey business here!
parser errors:
`