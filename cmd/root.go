package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "legislator",
	Short: "legislator manages the automated deployment of kubernetes network policies",
	Long: `Legislator is a CLI for managing the automated deployment of kubernetes network policies centralized by a constitution config file.`,
  }
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(greetingsCmd)
}