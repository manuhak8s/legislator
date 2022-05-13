package logger

import (
	"fmt"

	"github.com/guumaster/logsymbols"
)

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