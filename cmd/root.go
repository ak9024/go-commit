package cmd

import (
	"fmt"
	"os"

	"github.com/ak9024/go-commit/internal/commit"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-commit",
	Short: "go-commit is a short of go-commit",
	Long:  "The CLI to generate commit with karma style",
	Run:   commit.Commit,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
