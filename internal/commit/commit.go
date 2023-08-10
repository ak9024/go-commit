package commit

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ak9024/go-commit/pkg/chatgpt"
	"github.com/ak9024/go-commit/utils"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func Commit(cmd *cobra.Command, args []string) {
	_cmd, err := utils.ExecCommand("git diff --staged")
	if err != nil {
		fmt.Println(err)

		return
	}

	if _cmd == "" {
		fmt.Println(utils.PrintLogo())
		utils.Print("Invalid", "Can't be execute git diff --staged, please execute git add before commit")

		return
	}

	// add spinner to handle buffering while async process
	s := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	s.Start()
	s.HideCursor = true

	// get `stdout` from `_cmd` and pass to chatgpt.GeneratedCommitMessageByChatGPT.
	out, err := chatgpt.GeneratedCommitMessageByChatGPT(_cmd)
	if err != nil {
		fmt.Println(err)

		return
	}

	// stop the spinner if the data return
	s.Stop()

	// get message generate from chatgpt
	commitMessage := out.Choices[len(out.Choices)-1].Message.Content

	// present the result, before commit happened
	fmt.Println(utils.PrintLogo())
	utils.Print("Please review your commit message:", commitMessage)

	// ask a quetion to confirm, before commit execute
	confirmCommit := utils.StringPrompt("[Confirm] If ok, please confirm to execute commit [y/n]? ")

	// validate just process commit, if user aggre for the next action (commit)
	if confirmCommit != "" && strings.EqualFold(confirmCommit, "y") {
		// compose the command for commit
		gitCommit := fmt.Sprintf(`"%s"`, commitMessage)
		// exec the command for commit
		execCommit, _ := utils.ExecCommand(fmt.Sprintf("git commit -m %s", gitCommit))
		fmt.Print(execCommit)
		return
	}

	// close all the process if user ignore to commit
	os.Exit(0)
}
