package SkyLine

import (
	"io/ioutil"
	"log"
	"strings"
)

func Output(char string, indentby int, increaseindent int) string {
	charliner := strings.Repeat(" ", indentby-1)
	charliner += strings.Repeat(" ", increaseindent) + char
	return charliner
}

// Tracert will basically return the column and line number of a token found in a file, for example this will be used
// if skyline needs to find the error of which a placement is expected or token is expected.
/*
func ScanFileRawOpenForTokenLiteral(filename, token string, increaseindent int) (LineNumber, ColumnNumber int, LineOfCode string, CharLiner string) {
	file, err := os.Open(filename)
	if CE(err) {
		log.Fatal(err)
	}
	defer file.Close()
	lineNum := 1
	colNum := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			if r == rune(token[0]) {
				if line[i:i+len(token)] == token {
					ColumnNumber = colNum + 1
					LineNumber = lineNum
					LineOfCode = line
					CharLiner = Output("^", colNum+i, increaseindent)
					break
				}
			}
		}
		lineNum++
		colNum += len(line) + 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return LineNumber, ColumnNumber, LineOfCode, CharLiner
}
*/

// Scan last few lines within the code file given a certain line number
// For example if you get line 15 as an error the program will scan from the 10th line until the 15th line
// in other words report the last 5 lines before the error
func GrabLast5LinesBasedOnIntegerInputFromLineTracerErrorSystem(filename string, erroredline int) (Last5 []string, errorline string) {
	bytes, x := ioutil.ReadFile(filename)
	if x != nil {
		log.Fatal(x)
	}
	Content := string(bytes)
	lines := strings.Split(Content, "\n")
	if len(lines) <= 6 {
		return nil, ""
	}
	GetLastLineBeforeError := erroredline - 2
	for i := 0; i < 6; i++ {
		Last5 = append(Last5, string(lines[GetLastLineBeforeError-i]))
	}
	ReverseArrayForFileTraceback(Last5)
	return Last5, lines[erroredline-1]
}
