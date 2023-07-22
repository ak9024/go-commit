package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-commit",
	Version: cliVersion,
	Short:   shortText,
	Long:    Long,
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
