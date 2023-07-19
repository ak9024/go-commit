package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "go-commit",
	Long: "The CLI to generate commit with karma style",
}

func Execute() {
	rootCmd.AddCommand(commitCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
