// Code generated from ./internal/parser/MarkdownLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "", "", "", "", "'**'", "'*'", "'__'", "'_'", "'['", "']'", "'('",
		"')'",
	}
	staticData.SymbolicNames = []string{
		"", "H1", "H2", "H3", "LIST_MARKER", "DOUBLE_ASTERISK", "SINGLE_ASTERISK",
		"DOUBLE_UNDERSCORE", "SINGLE_UNDERSCORE", "LBRACK", "RBRACK", "LPAREN",
		"RPAREN", "TEXT", "WS", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"H1", "H2", "H3", "LIST_MARKER", "DOUBLE_ASTERISK", "SINGLE_ASTERISK",
		"DOUBLE_UNDERSCORE", "SINGLE_UNDERSCORE", "LBRACK", "RBRACK", "LPAREN",
		"RPAREN", "TEXT", "WS", "NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 80, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 4, 12,
		67, 8, 12, 11, 12, 12, 12, 68, 1, 13, 4, 13, 72, 8, 13, 11, 13, 12, 13,
		73, 1, 14, 3, 14, 77, 8, 14, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 1, 0, 2, 2, 0, 42, 43, 45, 45, 7, 0, 10, 10, 13, 13, 35,
		35, 40, 42, 91, 91, 93, 93, 95, 95, 82, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 34, 1, 0, 0, 0, 5, 39,
		1, 0, 0, 0, 7, 45, 1, 0, 0, 0, 9, 47, 1, 0, 0, 0, 11, 50, 1, 0, 0, 0, 13,
		52, 1, 0, 0, 0, 15, 55, 1, 0, 0, 0, 17, 57, 1, 0, 0, 0, 19, 59, 1, 0, 0,
		0, 21, 61, 1, 0, 0, 0, 23, 63, 1, 0, 0, 0, 25, 66, 1, 0, 0, 0, 27, 71,
		1, 0, 0, 0, 29, 76, 1, 0, 0, 0, 31, 32, 5, 35, 0, 0, 32, 33, 3, 27, 13,
		0, 33, 2, 1, 0, 0, 0, 34, 35, 5, 35, 0, 0, 35, 36, 5, 35, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 38, 3, 27, 13, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 35, 0,
		0, 40, 41, 5, 35, 0, 0, 41, 42, 5, 35, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44,
		3, 27, 13, 0, 44, 6, 1, 0, 0, 0, 45, 46, 7, 0, 0, 0, 46, 8, 1, 0, 0, 0,
		47, 48, 5, 42, 0, 0, 48, 49, 5, 42, 0, 0, 49, 10, 1, 0, 0, 0, 50, 51, 5,
		42, 0, 0, 51, 12, 1, 0, 0, 0, 52, 53, 5, 95, 0, 0, 53, 54, 5, 95, 0, 0,
		54, 14, 1, 0, 0, 0, 55, 56, 5, 95, 0, 0, 56, 16, 1, 0, 0, 0, 57, 58, 5,
		91, 0, 0, 58, 18, 1, 0, 0, 0, 59, 60, 5, 93, 0, 0, 60, 20, 1, 0, 0, 0,
		61, 62, 5, 40, 0, 0, 62, 22, 1, 0, 0, 0, 63, 64, 5, 41, 0, 0, 64, 24, 1,
		0, 0, 0, 65, 67, 8, 1, 0, 0, 66, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68,
		66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 26, 1, 0, 0, 0, 70, 72, 5, 32,
		0, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74,
		1, 0, 0, 0, 74, 28, 1, 0, 0, 0, 75, 77, 5, 13, 0, 0, 76, 75, 1, 0, 0, 0,
		76, 77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 5, 10, 0, 0, 79, 30, 1,
		0, 0, 0, 4, 0, 68, 73, 76, 0,
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
	MarkdownLexerH1                = 1
	MarkdownLexerH2                = 2
	MarkdownLexerH3                = 3
	MarkdownLexerLIST_MARKER       = 4
	MarkdownLexerDOUBLE_ASTERISK   = 5
	MarkdownLexerSINGLE_ASTERISK   = 6
	MarkdownLexerDOUBLE_UNDERSCORE = 7
	MarkdownLexerSINGLE_UNDERSCORE = 8
	MarkdownLexerLBRACK            = 9
	MarkdownLexerRBRACK            = 10
	MarkdownLexerLPAREN            = 11
	MarkdownLexerRPAREN            = 12
	MarkdownLexerTEXT              = 13
	MarkdownLexerWS                = 14
	MarkdownLexerNEWLINE           = 15
)
