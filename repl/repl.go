package repl

import (
	"bufio"
	"fmt"
	"io"
	"seqa/token"
	"seqa/lexer"
)

const MONKEY_FACE = `
   ^    ^    
  / l__l .
(  o   ox  )
     ^
   -(@)-
`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// evaluated := evaluator.Eval(program, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }

		// io.WriteString(out, program.String())
		// io.WriteString(out, "\n")
		
		for tok := l.NextToken(); tok.Type != token.EOF && tok.Type != token.ILLEGAL; tok = l.NextToken() {
		 	fmt.Printf("%+v\n", tok)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "woops! wa ran into same monkey busioness here\n")
	io.WriteString(out, "parse errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
