package commit

import (
	"fmt"
	"os"
	"strings"

	"github.com/ak9024/go-commit/pkg/chatgpt"
	"github.com/ak9024/go-commit/utils"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func Commit(cmd *cobra.Command, args []string) {
	_cmd, err := utils.ExecCommand("git diff --staged")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	generatedFigure := figure.NewFigure("go-commit", "", true)

	if _cmd == "" {
		generatedFigure.Print()
		utils.Print("Oops", "Please doing git add <file> first or make sure git init in your repository")
		os.Exit(0)
	}

	// get `stdout` from `_cmd` and pass to chatgpt.GeneratedCommitMessageByChatGPT.
	out, err := chatgpt.GeneratedCommitMessageByChatGPT(_cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	commitMessage := out.Choices[len(out.Choices)-1].Message.Content

	generatedFigure.Print()
	utils.Print("Please review your commit message", commitMessage)

	// ask a quetion to make sure the commit suitable with the changes
	commit := utils.StringPrompt("If ok, please confirm to commit [y/n]? ")

	if commit != "" && strings.EqualFold(commit, "y") {
		gitCommit := fmt.Sprintf(`git commit -m "%s"`, commitMessage)
		execCommit, err := utils.ExecCommand(gitCommit)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(execCommit)
	} else {
		os.Exit(0)
	}

}