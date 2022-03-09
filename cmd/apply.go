package cmd

import (
	_"fmt"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use: "apply",
	Short: "Applies the constitution config to the kubernetes cluster.",
	Long: "All network policies become deployed to the kubernetes cluster based on their configuration.",
	Run: nil,
}