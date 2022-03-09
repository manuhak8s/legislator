package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var greetingsCmd = &cobra.Command{
	Use: "greetings",
	Short: "Sends greetings.",
	Long: "Sends greetings.",
	Run: sendGreetings,
}

func sendGreetings (cmd *cobra.Command, args []string) {
	fmt.Println("Greetings from legislator!")
}