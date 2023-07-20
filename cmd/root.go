package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-commit",
	Short: "The CLI (Command Line Interface) to help you generated commit automaticly",
	Long:  "The CLI (Command Line Interface) to help you generated commit automaticly, using karma style for git convention",
}

func Execute() {
	rootCmd.AddCommand(commitCmd)

	// to enable suggestions
	rootCmd.SuggestionsMinimumDistance = 1

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
