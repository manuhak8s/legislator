package logger

import (
	"fmt"

	"github.com/guumaster/logsymbols"
)

// TriggerOutput is a void function that prints an output
// with icon based on the given log level:
// - success: green exclamation mark
// - fail: red cross
// - loading: blue info symbol
func TriggerOutput(logLevel string, output string) {
	switch logLevel {
	case "success":
		fmt.Println(logsymbols.Success, output)
	case "fail": 
		fmt.Println(logsymbols.Error, output)
	case "loading":
		fmt.Println(logsymbols.Info, output)
	}
}