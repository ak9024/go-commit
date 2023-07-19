package chatgpt

import (
	"fmt"
	"os"

	sdk "github.com/ak9024/go-chatgpt-sdk"
	"github.com/common-nighthawk/go-figure"
)

var OpenAIKey = os.Getenv("OPENAI_API_KEY")

// generatedCommitMessageByChatGPT(c string) to generate commit message by input in param `c`
func GeneratedCommitMessageByChatGPT(c string) (*sdk.ModelChatResponse, error) {
	_sdk := sdk.NewConfig(sdk.Config{
		OpenAIKey: OpenAIKey,
	})

	content := "You are act as the author of a commit message in git, your mission is to create clean and comprehensive commit messages, and also desribe the changes why and what please follow the karma style like `<type>(type): <subject> <body> <footer>`, i'll send you output of `git diff --staged` command, and you convert it into commit message. please don't add any description to the commit, only commit message and convert all message into lowwercase format"

	generatedFigure := figure.NewFigure("go-commit", "", true)

	if OpenAIKey == "" {
		generatedFigure.Print()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("===============================================")
		fmt.Println("Please setup OPENAI_API_KEY in your environment")
		fmt.Println("")
		fmt.Println("")
		os.Exit(0)
	}

	resp, err := _sdk.ChatCompletions(sdk.ModelChat{
		Model: "gpt-3.5-turbo",
		Messages: []sdk.Message{
			{
				Role:    "system",
				Content: content,
			},
			{
				Role:    "user",
				Content: c,
			},
		},
		MaxTokens: 200,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
