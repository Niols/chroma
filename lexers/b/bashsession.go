package b

import (
	. "github.com/Niols/chroma" // nolint
	"github.com/Niols/chroma/lexers/internal"
)

// BashSession lexer.
var BashSession = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "BashSession",
		Aliases:   []string{"bash-session", "console", "shell-session"},
		Filenames: []string{".sh-session"},
		MimeTypes: []string{"text/x-sh"},
		EnsureNL:  true,
	},
	bashsessionRules,
))

func bashsessionRules() Rules {
	return Rules{
		"root": {
			{`(^[#$%>]\s*)(.*\n?)`, ByGroups(GenericPrompt, Using(Bash)), nil},
			{`^.+\n?`, GenericOutput, nil},
		},
	}
}
