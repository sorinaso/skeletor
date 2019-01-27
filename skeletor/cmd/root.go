package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "skeletor",
	Short: "Skeletor: application that generates skeletons.",
	Long: `Skeletor is an application that generate skeletons

You can use this application for generate base skeletons from templates
and update them after.`,
}
