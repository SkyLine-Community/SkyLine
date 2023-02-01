package SkyLine

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func RunAndParseNoFile(data string) error {
	parser := New_Parser(LexNew(data))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}
	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *ErrorsTrace {
					fmt.Println(ErrorSymBolMap[CODE_NO_FUNCTIONS_OR_SYMBOLS_LOADED]("FILE_SECTOR_1:WAS_NOT_FILE_USED_FLAG_(-e)"))
				}
			}
		}
	}()
	_, x := io.WriteString(os.Stdout, result.Inspect()+"\n")
	return x
}

func RunAndParseFiles(filename string) error {
	FileCurrent.New(filename)
	f, x := os.Open(filename)
	if x != nil {
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		fmt.Println(ErrorSymBolMap[CODE_FILE_INTEGRITY_FILE_INVALID_FILE_DOES_NOT_EXIST](filename))
		os.Exit(0)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if line == nil {
		fmt.Println(ErrorSymBolMap[CODE_FILE_INTEGRITY_FILE_INVALID_MUST_HAVE_CODE_OR_LOGIC_INSIDE_FILE_NULL](filename))
	}
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println(ErrorSymBolMap[CODE_FILE_FAILED_USING_INPUT_OUTPUT_READER_AND_UTILITY_FILE_ISSUE](filename, fmt.Sprint(x)))
		os.Exit(1)
	}
	parser := New_Parser(LexNew(string(data)))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}

	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *ErrorsTrace {
					fmt.Println(ErrorSymBolMap[CODE_NO_FUNCTIONS_OR_SYMBOLS_LOADED](filename))
				}
			}
		}
	}()
	_, x = io.WriteString(os.Stdout, result.Inspect()+"\n")
	return x
}
