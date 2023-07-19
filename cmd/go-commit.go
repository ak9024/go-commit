package cmd

import (
	"fmt"
	"os"

	"github.com/ak9024/go-commit/internal/chatgpt"
	"github.com/ak9024/go-commit/utils"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func RunCommand(cmd *cobra.Command, args []string) {
	_cmd, err := utils.ExecCommand("git diff --staged")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	generatedFigure := figure.NewFigure("go-commit", "", true)

	// get `stdout` from `_cmd` and pass to GeneratedCommitMessageByChatGPT.
	out, err := chatgpt.GeneratedCommitMessageByChatGPT(_cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commitMessage := out.Choices[len(out.Choices)-1].Message.Content

	generatedFigure.Print()
	utils.PrintCommit(commitMessage)

	// ask a quetion to make sure the commit suitable with the changes
	commit := utils.StringPrompt("If ok, please confirm to commit [Y/n]? ")
	if commit != "" && commit == "Y" {
		gitCommit := fmt.Sprintf(`git commit -m "%s"`, commitMessage)
		execCommit, _ := utils.ExecCommand(gitCommit)
		fmt.Println(execCommit)
	} else {
		os.Exit(1)
	}
}
