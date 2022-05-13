package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/manuhak8s/legislator/pkg/luther"
)

var applyCmd = &cobra.Command{
	Use: "apply",
	Short: "Applies the constitution config to the kubernetes cluster.",
	Long: "All network policies become deployed to the kubernetes cluster based on their configuration.",
	Run: apply,
}

func apply(ccmd *cobra.Command, args []string) {
	configPath, err := ccmd.Flags().GetString("path")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unknown error occured while executing apply command.")
		return
	}

	if configPath == "" {
		fmt.Fprintln(os.Stderr, "Config path is not specified. Please enter an existent path to a config file.")
		return
	} else {
		luther.ExecuteLegislation(configPath)
	}
}