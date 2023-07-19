package chatgpt

import (
	"os"
)

var (
	content   = "You are act as the author of a commit message in git, your mission is to create clean and comprehensive commit messages, and also desribe the changes why and what please follow the karma style like `<type>(type): <subject> <body> <footer>`, i'll send you output of `git diff --staged` command, and you convert it into commit message. please don't add any description to the commit, only commit message and convert all message into lowwercase format"
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
)
