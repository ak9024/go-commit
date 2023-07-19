package chatgpt

import (
	"os"

	sdk "github.com/ak9024/go-chatgpt-sdk"
)

// generatedCommitMessageByChatGPT(c string) to generate commit message by input in param `c`
func GeneratedCommitMessageByChatGPT(c string) (*sdk.ModelChatResponse, error) {
	_sdk := sdk.NewConfig(sdk.Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	})

	content := "You are act as the author of a commit message in git, your mission is to create clean and comprehensive commit messages, and also desribe the why and what please follow the karma style like `<type>(type): <subject> <body> <footer>`, i'll send you output of `git diff --staged` command, and you convert it into commit message. please don't add any description to the commit, only commit message and convert all message into lowwercase format"

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
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
