package SkyLine

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
Carrier is an executor for other files, it is NOT a linker
Carrier for example will take a given file and start a new environment to control it
for example if you have two files main.csc and test.csc, and you want to execute functions
in test.csc before any other function executes in main.csc you can actually end up using
carry|"test.csc"|
before any other function executes in main.csc or is called, carry in the end will stop
the interpreter from executing any other function or parsing any other data until it is
done executing and parsing the contents of the file within the carry keyword
*/

func ReadImportedCarrierFile(filename string) (bool, error) {
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		log.Fatal(x)
	}
	parser := New_Parser(LexNew(string(data)))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return false, errors.New(parser.ParserErrors()[0])
	}

	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return false, nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *ErrorsTrace {
					fmt.Println("SkyLine parser: no functions or variables loaded...")
				}
			}
		}
	}()
	_, x = io.WriteString(os.Stdout, result.Inspect()+"\n")
	if x != nil {
		return false, x
	}
	return true, nil
}
