package log

import (
	"os"
	"time"
	"fmt"

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
	fmt.Println("")
}

func LogNetworkPolicyReading() {
	triggerCircledSpinner(" Reading network policies from current namespace ..")
	fmt.Println("")
}

func LogNetworkPolicyCreating() {
	triggerCircledSpinner(" Creating network policy into current namespace ..")
	fmt.Println("")
}

func LogNetworkPolicyRemoving() {
	triggerCircledSpinner(" Removing network policies into current namespace ..")
	fmt.Println("")
}

func LogNetworkPolicyApllying() {
	triggerCircledSpinner(" Applying network policy into current namespace ..")
	fmt.Println("")
}