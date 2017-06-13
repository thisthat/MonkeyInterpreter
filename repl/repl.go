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
	outter := bufio.NewWriter(out)
	const PROMPT = ">> "
	for {
		outter.Flush()
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			outter.Flush()
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			s := fmt.Sprintf("Tok: %+v\n", tok)
			outter.WriteString(s)
			outter.Flush()
		}
	}
}
