package cmd

import (
	"github.com/ak9024/go-commit/internal/commit"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generated automate commit provide by chatgpt",
	Run:   commit.Commit,
}
