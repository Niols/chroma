package styles

import (
	"github.com/Niols/chroma"
)

// Dracula Style
var Dracula = Register(chroma.MustNewStyle("dracula", chroma.StyleEntries{
	chroma.Comment:                  "#6272a4",
	chroma.CommentHashbang:          "#6272a4",
	chroma.CommentMultiline:         "#6272a4",
	chroma.CommentPreproc:           "#ff79c6",
	chroma.CommentSingle:            "#6272a4",
	chroma.CommentSpecial:           "#6272a4",
	chroma.Generic:                  "#f8f8f2",
	chroma.GenericDeleted:           "#ff5555",
	chroma.GenericEmph:              "#f8f8f2 underline",
	chroma.GenericError:             "#f8f8f2",
	chroma.GenericHeading:           "#f8f8f2 bold",
	chroma.GenericInserted:          "#50fa7b bold",
	chroma.GenericOutput:            "#44475a",
	chroma.GenericPrompt:            "#f8f8f2",
	chroma.GenericStrong:            "#f8f8f2",
	chroma.GenericSubheading:        "#f8f8f2 bold",
	chroma.GenericTraceback:         "#f8f8f2",
	chroma.GenericUnderline:         "underline",
	chroma.Error:                    "#f8f8f2",
	chroma.Keyword:                  "#ff79c6",
	chroma.KeywordConstant:          "#ff79c6",
	chroma.KeywordDeclaration:       "#8be9fd italic",
	chroma.KeywordNamespace:         "#ff79c6",
	chroma.KeywordPseudo:            "#ff79c6",
	chroma.KeywordReserved:          "#ff79c6",
	chroma.KeywordType:              "#8be9fd",
	chroma.Literal:                  "#f8f8f2",
	chroma.LiteralDate:              "#f8f8f2",
	chroma.Name:                     "#f8f8f2",
	chroma.NameAttribute:            "#50fa7b",
	chroma.NameBuiltin:              "#8be9fd italic",
	chroma.NameBuiltinPseudo:        "#f8f8f2",
	chroma.NameClass:                "#50fa7b",
	chroma.NameConstant:             "#f8f8f2",
	chroma.NameDecorator:            "#f8f8f2",
	chroma.NameEntity:               "#f8f8f2",
	chroma.NameException:            "#f8f8f2",
	chroma.NameFunction:             "#50fa7b",
	chroma.NameLabel:                "#8be9fd italic",
	chroma.NameNamespace:            "#f8f8f2",
	chroma.NameOther:                "#f8f8f2",
	chroma.NameTag:                  "#ff79c6",
	chroma.NameVariable:             "#8be9fd italic",
	chroma.NameVariableClass:        "#8be9fd italic",
	chroma.NameVariableGlobal:       "#8be9fd italic",
	chroma.NameVariableInstance:     "#8be9fd italic",
	chroma.LiteralNumber:            "#bd93f9",
	chroma.LiteralNumberBin:         "#bd93f9",
	chroma.LiteralNumberFloat:       "#bd93f9",
	chroma.LiteralNumberHex:         "#bd93f9",
	chroma.LiteralNumberInteger:     "#bd93f9",
	chroma.LiteralNumberIntegerLong: "#bd93f9",
	chroma.LiteralNumberOct:         "#bd93f9",
	chroma.Operator:                 "#ff79c6",
	chroma.OperatorWord:             "#ff79c6",
	chroma.Other:                    "#f8f8f2",
	chroma.Punctuation:              "#f8f8f2",
	chroma.LiteralString:            "#f1fa8c",
	chroma.LiteralStringBacktick:    "#f1fa8c",
	chroma.LiteralStringChar:        "#f1fa8c",
	chroma.LiteralStringDoc:         "#f1fa8c",
	chroma.LiteralStringDouble:      "#f1fa8c",
	chroma.LiteralStringEscape:      "#f1fa8c",
	chroma.LiteralStringHeredoc:     "#f1fa8c",
	chroma.LiteralStringInterpol:    "#f1fa8c",
	chroma.LiteralStringOther:       "#f1fa8c",
	chroma.LiteralStringRegex:       "#f1fa8c",
	chroma.LiteralStringSingle:      "#f1fa8c",
	chroma.LiteralStringSymbol:      "#f1fa8c",
	chroma.Text:                     "#f8f8f2",
	chroma.TextWhitespace:           "#f8f8f2",
	chroma.Background:               " bg:#282a36",
}))
