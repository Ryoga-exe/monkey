package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/Ryoga-exe/monkey/internal/lexer"
	"github.com/Ryoga-exe/monkey/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
