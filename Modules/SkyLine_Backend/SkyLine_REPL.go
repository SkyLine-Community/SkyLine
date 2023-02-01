package SkyLine

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
)

const prompt = "SkyLine|%s(%s)>> "

func Start(in io.Reader, Out io.Writer) {
	scanner := bufio.NewScanner(in)
	Env := NewEnvironment()
	macroEnv := NewEnvironment()

	for {
		fmt.Printf(prompt, runtime.GOOS, runtime.GOARCH)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := LexNew(line)
		parser := New_Parser(l)

		program := parser.ParseProgram()
		if len(parser.ParserErrors()) != 0 {
			printParserErrors(line, Out, parser.ParserErrors())
			continue
		}
		DefineMacros(program, macroEnv)
		expanded := ExpandMacros(program, macroEnv)
		evaluated := Eval(expanded, Env)
		if evaluated == nil {
			continue
		}
		io.WriteString(Out, evaluated.Inspect())
		io.WriteString(Out, "\n")
	}
}

func printParserErrors(line string, Out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(Out, msg)
		io.WriteString(Out, "\n")
	}
}
