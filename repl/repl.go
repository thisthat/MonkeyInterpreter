package repl

import (
	"bufio"
	"fmt"
	"github.com/thisthat/MonkeyInterpreter/lexer"
	"github.com/thisthat/MonkeyInterpreter/token"
	"io"
)

// Start starts the Read Evaluate Print Loop
func Start(in io.Reader, out io.Writer) {
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("Tok: %+v\n", tok)
		}
	}
}
