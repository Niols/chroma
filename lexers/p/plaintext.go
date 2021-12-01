package p

import (
	. "github.com/Niols/chroma" // nolint
	"github.com/Niols/chroma/lexers/internal"
)

var Plaintext = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "plaintext",
		Aliases:   []string{"text", "plain", "no-highlight"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
		Priority:  0.1,
	},
	internal.PlaintextRules,
))
