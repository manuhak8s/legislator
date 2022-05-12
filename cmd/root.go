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
	// demo and test commands
	rootCmd.AddCommand(greetingsCmd)
	rootCmd.AddCommand(namespaceCmd)
	rootCmd.AddCommand(listPolicyCmd)
	rootCmd.AddCommand(deletePolicyCmd)
	rootCmd.AddCommand(createPolicyCmd)

	// core cli commands
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(destroyCmd)

	// flags
	applyCmd.PersistentFlags().String("path", "", "Path to an existent config file.")
	destroyCmd.PersistentFlags().String("path", "", "Path to an existent config file.")
}