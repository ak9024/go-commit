package commit

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ak9024/go-commit/pkg/chatgpt"
	"github.com/ak9024/go-commit/utils"
	"github.com/briandowns/spinner"
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

	s := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	s.Start()
	s.Prefix = "waiting.. "
	s.HideCursor = true

	// get `stdout` from `_cmd` and pass to chatgpt.GeneratedCommitMessageByChatGPT.
	out, err := chatgpt.GeneratedCommitMessageByChatGPT(_cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	s.Stop()

	commitMessage := out.Choices[len(out.Choices)-1].Message.Content

	generatedFigure.Print()
	utils.Print("Please review your commit message", commitMessage)

	// ask a quetion to make sure the commit suitable with the changes
	commit := utils.StringPrompt("[Ask] If ok, please confirm to commit [y/n]? ")

	if commit != "" && strings.EqualFold(commit, "y") {
		gitCommit := fmt.Sprintf(`git commit -m "%s"`, commitMessage)
		execCommit, _ := utils.ExecCommand(gitCommit)
		fmt.Println(execCommit)
	} else {
		os.Exit(0)
	}

}
