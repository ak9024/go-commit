package chatgpt

import (
	"fmt"
	"os"
)

var (
	command    = "git diff --staged"
	maxLines   = "60"
	karmaStyle = `
	<type>(<scope>): <subject>
	<BLANK_LINE>
	<body>
	<BLANK_LINE>
	<footer>
	`
	exampleKarmaStyle = `
	fix(middleware): ensure Range headers adhere more closely to RFC 2616
	
	Add one new dependency, use 'range-parser' (Express dependency) to compute
	range. It is more well-tested in the wild.
	`
	content = fmt.Sprintf(`
	You are the author of a commit message in git. 
	Your mission is to create clean and comprehensive commit messages in the conventional commit message and explain WHAT were the changes and WHY the changes were done. 
	I'll send you an output "%s" command, and you convert it into commit message.
	Use git karma style like this %s, 
	Please follow karma style example %s.
	Use the present tense.
	Lines must not be longer than %s characters.
	`,
		command,
		karmaStyle,
		exampleKarmaStyle,
		maxLines)
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
)
