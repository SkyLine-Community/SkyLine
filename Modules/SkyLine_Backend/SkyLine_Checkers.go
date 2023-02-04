package SkyLine

import (
	"fmt"
)

// Check error
func CE(e error) bool { return e != nil }

// Call buf which calls to fail, panic channel
func CallBuf(message string, verbose bool) {
	if x := recover(); x != nil {
		if verbose {
			fmt.Println(x)
		} else {
			fmt.Println(message)
		}
	}
}
