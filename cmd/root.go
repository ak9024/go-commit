package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gc",
	Short: "gc is a short of go-commit",
	Long:  "The CLI to help you create a commit message help by chatgpt",
	Run:   RunCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
