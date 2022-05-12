package cmd

import (
	"fmt"
	"os"

	"github.com/manuhak8s/legislator/pkg/luther"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use: "destroy",
	Short: "Removes the constitution config from the kubernetes cluster.",
	Long: "All network policies become removed from the kubernetes cluster based on their configuration.",
	Run: destroy,
}


func destroy(ccmd *cobra.Command, args []string) {
	configPath, err := ccmd.Flags().GetString("path")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unknown error occured while executing destroy command.")
		return
	}

	if configPath != "" {
		luther.ExecuteDestruction(configPath)
	} else {
		fmt.Fprintln(os.Stderr, "Config path is not specified. Please enter an existent path to a config file.")
		return
	}
}