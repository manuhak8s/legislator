package log

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	_ "github.com/fatih/color"
)

func triggerCircledSpinner(text string) {
	
	circledSpinner := spinner.New(spinner.CharSets[7], 100*time.Millisecond, spinner.WithWriter(os.Stderr))  
	circledSpinner.Color("yellow")
	circledSpinner.Suffix = text
	circledSpinner.Start()                                                    
	time.Sleep(4 * time.Second)
	circledSpinner.Start()
}

func LogNamespaceReading() {
	triggerCircledSpinner(" Reading namespaces from kubernetes cluster ..")
}