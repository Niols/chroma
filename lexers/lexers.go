// Package lexers contains the registry of all lexers.
//
// Sub-packages contain lexer implementations.
package lexers

// nolint
import (
	"github.com/Niols/chroma"
	_ "github.com/Niols/chroma/lexers/a"
	_ "github.com/Niols/chroma/lexers/b"
	_ "github.com/Niols/chroma/lexers/c"
	_ "github.com/Niols/chroma/lexers/circular"
	_ "github.com/Niols/chroma/lexers/d"
	_ "github.com/Niols/chroma/lexers/e"
	_ "github.com/Niols/chroma/lexers/f"
	_ "github.com/Niols/chroma/lexers/g"
	_ "github.com/Niols/chroma/lexers/h"
	_ "github.com/Niols/chroma/lexers/i"
	"github.com/Niols/chroma/lexers/internal"
	_ "github.com/Niols/chroma/lexers/j"
	_ "github.com/Niols/chroma/lexers/k"
	_ "github.com/Niols/chroma/lexers/l"
	_ "github.com/Niols/chroma/lexers/m"
	_ "github.com/Niols/chroma/lexers/n"
	_ "github.com/Niols/chroma/lexers/o"
	_ "github.com/Niols/chroma/lexers/p"
	_ "github.com/Niols/chroma/lexers/q"
	_ "github.com/Niols/chroma/lexers/r"
	_ "github.com/Niols/chroma/lexers/s"
	_ "github.com/Niols/chroma/lexers/t"
	_ "github.com/Niols/chroma/lexers/v"
	_ "github.com/Niols/chroma/lexers/w"
	_ "github.com/Niols/chroma/lexers/x"
	_ "github.com/Niols/chroma/lexers/y"
	_ "github.com/Niols/chroma/lexers/z"
)

// Registry of Lexers.
var Registry = internal.Registry

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string { return internal.Names(withAliases) }

// Get a Lexer by name, alias or file extension.
func Get(name string) chroma.Lexer { return internal.Get(name) }

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer { return internal.MatchMimeType(mimeType) }

// Match returns the first lexer matching filename.
func Match(filename string) chroma.Lexer { return internal.Match(filename) }

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer { return internal.Analyse(text) }

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer { return internal.Register(lexer) }

// Fallback lexer if no other is found.
var Fallback = internal.Fallback
