// Code generated from MarkdownLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MarkdownLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MarkdownLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func markdownlexerLexerInit() {
	staticData := &MarkdownLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'#'", "", "", "", "", "", "'***'", "'*'", "'**'", "'`'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "HASH", "H2", "H3", "H4",
		"H5", "H6", "BOLD_AND_ITALIC_MARKER", "ITALIC_MARKER", "BOLD_MARKER",
		"CODE_MARKER", "PUNCTUATION", "NUMBER", "WORD",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "HASH", "H2", "H3", "H4", "H5",
		"H6", "BOLD_AND_ITALIC_MARKER", "ITALIC_MARKER", "BOLD_MARKER", "CODE_MARKER",
		"PUNCTUATION", "NUMBER", "WORD",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 97, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 4, 0, 35, 8, 0, 11, 0, 12, 0, 36, 1, 1, 3, 1, 40, 8, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 3, 13, 86, 8, 13,
		1, 14, 4, 14, 89, 8, 14, 11, 14, 12, 14, 90, 1, 15, 4, 15, 94, 8, 15, 11,
		15, 12, 15, 95, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 1, 0,
		4, 2, 0, 9, 9, 32, 32, 10, 0, 33, 34, 39, 41, 44, 44, 46, 46, 58, 59, 63,
		63, 91, 91, 93, 93, 123, 123, 125, 125, 1, 0, 48, 57, 5, 0, 9, 10, 13,
		13, 32, 32, 35, 35, 42, 42, 100, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 34, 1, 0, 0, 0, 3, 39, 1, 0,
		0, 0, 5, 43, 1, 0, 0, 0, 7, 46, 1, 0, 0, 0, 9, 48, 1, 0, 0, 0, 11, 51,
		1, 0, 0, 0, 13, 55, 1, 0, 0, 0, 15, 60, 1, 0, 0, 0, 17, 66, 1, 0, 0, 0,
		19, 73, 1, 0, 0, 0, 21, 77, 1, 0, 0, 0, 23, 79, 1, 0, 0, 0, 25, 82, 1,
		0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 88, 1, 0, 0, 0, 31, 93, 1, 0, 0, 0, 33,
		35, 7, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0,
		0, 36, 37, 1, 0, 0, 0, 37, 2, 1, 0, 0, 0, 38, 40, 5, 13, 0, 0, 39, 38,
		1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 5, 10, 0, 0,
		42, 4, 1, 0, 0, 0, 43, 44, 5, 92, 0, 0, 44, 45, 9, 0, 0, 0, 45, 6, 1, 0,
		0, 0, 46, 47, 5, 35, 0, 0, 47, 8, 1, 0, 0, 0, 48, 49, 3, 7, 3, 0, 49, 50,
		3, 7, 3, 0, 50, 10, 1, 0, 0, 0, 51, 52, 3, 7, 3, 0, 52, 53, 3, 7, 3, 0,
		53, 54, 3, 7, 3, 0, 54, 12, 1, 0, 0, 0, 55, 56, 3, 7, 3, 0, 56, 57, 3,
		7, 3, 0, 57, 58, 3, 7, 3, 0, 58, 59, 3, 7, 3, 0, 59, 14, 1, 0, 0, 0, 60,
		61, 3, 7, 3, 0, 61, 62, 3, 7, 3, 0, 62, 63, 3, 7, 3, 0, 63, 64, 3, 7, 3,
		0, 64, 65, 3, 7, 3, 0, 65, 16, 1, 0, 0, 0, 66, 67, 3, 7, 3, 0, 67, 68,
		3, 7, 3, 0, 68, 69, 3, 7, 3, 0, 69, 70, 3, 7, 3, 0, 70, 71, 3, 7, 3, 0,
		71, 72, 3, 7, 3, 0, 72, 18, 1, 0, 0, 0, 73, 74, 5, 42, 0, 0, 74, 75, 5,
		42, 0, 0, 75, 76, 5, 42, 0, 0, 76, 20, 1, 0, 0, 0, 77, 78, 5, 42, 0, 0,
		78, 22, 1, 0, 0, 0, 79, 80, 5, 42, 0, 0, 80, 81, 5, 42, 0, 0, 81, 24, 1,
		0, 0, 0, 82, 83, 5, 96, 0, 0, 83, 26, 1, 0, 0, 0, 84, 86, 7, 1, 0, 0, 85,
		84, 1, 0, 0, 0, 86, 28, 1, 0, 0, 0, 87, 89, 7, 2, 0, 0, 88, 87, 1, 0, 0,
		0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 30,
		1, 0, 0, 0, 92, 94, 8, 3, 0, 0, 93, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0,
		95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 32, 1, 0, 0, 0, 6, 0, 36, 39,
		85, 90, 95, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MarkdownLexerInit initializes any static state used to implement MarkdownLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMarkdownLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarkdownLexerInit() {
	staticData := &MarkdownLexerLexerStaticData
	staticData.once.Do(markdownlexerLexerInit)
}

// NewMarkdownLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMarkdownLexer(input antlr.CharStream) *MarkdownLexer {
	MarkdownLexerInit()
	l := new(MarkdownLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MarkdownLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MarkdownLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MarkdownLexer tokens.
const (
	MarkdownLexerWHITESPACE             = 1
	MarkdownLexerNEWLINE                = 2
	MarkdownLexerESCAPE_CHAR            = 3
	MarkdownLexerHASH                   = 4
	MarkdownLexerH2                     = 5
	MarkdownLexerH3                     = 6
	MarkdownLexerH4                     = 7
	MarkdownLexerH5                     = 8
	MarkdownLexerH6                     = 9
	MarkdownLexerBOLD_AND_ITALIC_MARKER = 10
	MarkdownLexerITALIC_MARKER          = 11
	MarkdownLexerBOLD_MARKER            = 12
	MarkdownLexerCODE_MARKER            = 13
	MarkdownLexerPUNCTUATION            = 14
	MarkdownLexerNUMBER                 = 15
	MarkdownLexerWORD                   = 16
)
