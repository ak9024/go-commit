package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// execCommand(input string) func to handle command
func ExecCommand(input string) (string, error) {
	bash := "bash"
	// execute the command with shell scripting
	_cmd := exec.Command(bash, "-c", input)

	// create a butes.Buffer to store `stdout` & `stderr`
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	_cmd.Stdout = &stdout
	_cmd.Stderr = &stderr

	// execute the command with `_cmd.Run()`
	// if err to run the script please return error
	if err := _cmd.Run(); err != nil {
		return "", err
	}

	// check if stderr string return empty string
	// return `stderr.String()`
	if stderr.String() != "" {
		return stderr.String(), nil
	}

	return stdout.String(), nil
}

// stringPrompt(input string) to handle prompt in terminal
func StringPrompt(input string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, input)
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	return strings.TrimSpace(s)
}

func PrintLogo() string {
	f := figure.NewFigure("go-commit", "", true)
	return f.String()
}

func Print(title, description string) {
	fmt.Println()
	fmt.Println()
	fmt.Println(fmt.Sprintf("[Time] %s", time.TimeOnly))
	fmt.Println(fmt.Sprintf("[Info] %s", title))
	fmt.Println()
	fmt.Println("[Output]")
	fmt.Println()
	fmt.Println(description)
	fmt.Println()
	fmt.Println()
}
