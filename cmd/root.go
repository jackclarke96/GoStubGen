package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GoStubGen",
	Short: "GoStubGen is a tool for generating boilerplate Go code",
	Long: `GoStubGen generates interfaces, structs, and test frameworks 
	based on a provided YAML configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'GoStubGen generate -c <config.yaml>' to generate code.")
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
