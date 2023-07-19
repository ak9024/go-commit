package chatgpt

import (
	"os"

	sdk "github.com/ak9024/go-chatgpt-sdk"
	"github.com/ak9024/go-commit/utils"
	"github.com/common-nighthawk/go-figure"
)

// generatedCommitMessageByChatGPT(c string) to generate commit message by input in param `c`
func GeneratedCommitMessageByChatGPT(c string) (*sdk.ModelChatResponse, error) {
	_sdk := sdk.NewConfig(sdk.Config{
		OpenAIKey: OpenAIKey,
	})

	generatedFigure := figure.NewFigure("go-commit", "", true)

	if OpenAIKey == "" {
		generatedFigure.Print()
		utils.Print("Oops", "Please setup OPENAI_API_KEY in your environment")
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
