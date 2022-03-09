package cmd

import (
	_"fmt"

	"github.com/spf13/cobra"
)

var destroy = &cobra.Command{
	Use: "destroy",
	Short: "Removes the constitution config from the kubernetes cluster.",
	Long: "All network policies become removed from the kubernetes cluster based on their configuration.",
	Run: nil,
}