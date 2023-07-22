package cmd

import (
	"fmt"

	"github.com/ak9024/go-commit/utils"
)

var (
	cliVersion = "v0.2.0"
	shortText  = "The CLI (Command Line Interface) helps you generate commits automatically"
	longText   = "The CLI (Command Line Interface) helps you generate commits automatically, using karma style for git convention"
	Long       = fmt.Sprintf(`
%s
%s
	`,
		utils.PrintLogo(),
		longText,
	)
)
