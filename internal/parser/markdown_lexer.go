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
		"", "", "", "", "", "", "'~~'", "", "'`'", "", "", "", "", "", "", "'#'",
		"'['", "']'", "'('", "')'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "BOLD_MARKER", "ITALIC_MARKER",
		"STRIKETHROUGH_MARKER", "BOLD_AND_ITALIC_MARKER", "CODE_MARKER", "BLOCKQUOTE_MARKER",
		"UNORDERED_LIST_MARKER", "ORDERED_LIST_MARKER", "HORIZONTAL_RULE", "FENCED_CODE_BLOCK",
		"INDENTED_CODE_BLOCK", "HASH", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN",
		"EXCLAMATION", "TEXT",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "BOLD_MARKER", "ITALIC_MARKER",
		"STRIKETHROUGH_MARKER", "BOLD_AND_ITALIC_MARKER", "CODE_MARKER", "BLOCKQUOTE_MARKER",
		"UNORDERED_LIST_MARKER", "ORDERED_LIST_MARKER", "HORIZONTAL_RULE", "FENCED_CODE_BLOCK",
		"INDENTED_CODE_BLOCK", "HASH", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN",
		"EXCLAMATION", "TEXT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 162, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 4, 0, 45, 8, 0, 11, 0, 12, 0, 46, 1, 1, 3, 1, 50, 8, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 61, 8, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 74, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 80, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 4, 10,
		86, 8, 10, 11, 10, 12, 10, 87, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 102, 8, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 111, 8, 12, 10, 12, 12, 12,
		114, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 124, 8, 12, 10, 12, 12, 12, 127, 9, 12, 1, 12, 1, 12, 1, 12, 3, 12,
		132, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 139, 8, 13, 10, 13,
		12, 13, 142, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 4, 20, 159, 8, 20,
		11, 20, 12, 20, 160, 3, 112, 125, 140, 0, 21, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 1, 0, 5, 2, 0,
		9, 9, 32, 32, 2, 0, 42, 42, 95, 95, 2, 0, 42, 43, 45, 45, 1, 0, 48, 57,
		8, 0, 10, 10, 13, 13, 32, 33, 35, 35, 40, 42, 62, 62, 91, 93, 95, 96, 174,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 44, 1, 0, 0, 0, 3, 49, 1, 0, 0, 0, 5,
		53, 1, 0, 0, 0, 7, 60, 1, 0, 0, 0, 9, 62, 1, 0, 0, 0, 11, 64, 1, 0, 0,
		0, 13, 73, 1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 81,
		1, 0, 0, 0, 21, 85, 1, 0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 131, 1, 0, 0,
		0, 27, 133, 1, 0, 0, 0, 29, 145, 1, 0, 0, 0, 31, 147, 1, 0, 0, 0, 33, 149,
		1, 0, 0, 0, 35, 151, 1, 0, 0, 0, 37, 153, 1, 0, 0, 0, 39, 155, 1, 0, 0,
		0, 41, 158, 1, 0, 0, 0, 43, 45, 7, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 2, 1, 0, 0, 0,
		48, 50, 5, 13, 0, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1,
		0, 0, 0, 51, 52, 5, 10, 0, 0, 52, 4, 1, 0, 0, 0, 53, 54, 5, 92, 0, 0, 54,
		55, 9, 0, 0, 0, 55, 6, 1, 0, 0, 0, 56, 57, 5, 42, 0, 0, 57, 61, 5, 42,
		0, 0, 58, 59, 5, 95, 0, 0, 59, 61, 5, 95, 0, 0, 60, 56, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 61, 8, 1, 0, 0, 0, 62, 63, 7, 1, 0, 0, 63, 10, 1, 0, 0,
		0, 64, 65, 5, 126, 0, 0, 65, 66, 5, 126, 0, 0, 66, 12, 1, 0, 0, 0, 67,
		68, 5, 42, 0, 0, 68, 69, 5, 42, 0, 0, 69, 74, 5, 42, 0, 0, 70, 71, 5, 95,
		0, 0, 71, 72, 5, 95, 0, 0, 72, 74, 5, 95, 0, 0, 73, 67, 1, 0, 0, 0, 73,
		70, 1, 0, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76, 5, 96, 0, 0, 76, 16, 1, 0,
		0, 0, 77, 79, 5, 62, 0, 0, 78, 80, 3, 1, 0, 0, 79, 78, 1, 0, 0, 0, 79,
		80, 1, 0, 0, 0, 80, 18, 1, 0, 0, 0, 81, 82, 7, 2, 0, 0, 82, 83, 3, 1, 0,
		0, 83, 20, 1, 0, 0, 0, 84, 86, 7, 3, 0, 0, 85, 84, 1, 0, 0, 0, 86, 87,
		1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 90, 5, 46, 0, 0, 90, 91, 3, 1, 0, 0, 91, 22, 1, 0, 0, 0, 92, 93, 5,
		45, 0, 0, 93, 94, 5, 45, 0, 0, 94, 102, 5, 45, 0, 0, 95, 96, 5, 95, 0,
		0, 96, 97, 5, 95, 0, 0, 97, 102, 5, 95, 0, 0, 98, 99, 5, 42, 0, 0, 99,
		100, 5, 42, 0, 0, 100, 102, 5, 42, 0, 0, 101, 92, 1, 0, 0, 0, 101, 95,
		1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 3, 3,
		1, 0, 104, 24, 1, 0, 0, 0, 105, 106, 5, 96, 0, 0, 106, 107, 5, 96, 0, 0,
		107, 108, 5, 96, 0, 0, 108, 112, 1, 0, 0, 0, 109, 111, 9, 0, 0, 0, 110,
		109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 112, 110,
		1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 96,
		0, 0, 116, 117, 5, 96, 0, 0, 117, 132, 5, 96, 0, 0, 118, 119, 5, 126, 0,
		0, 119, 120, 5, 126, 0, 0, 120, 121, 5, 126, 0, 0, 121, 125, 1, 0, 0, 0,
		122, 124, 9, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 125,
		1, 0, 0, 0, 128, 129, 5, 126, 0, 0, 129, 130, 5, 126, 0, 0, 130, 132, 5,
		126, 0, 0, 131, 105, 1, 0, 0, 0, 131, 118, 1, 0, 0, 0, 132, 26, 1, 0, 0,
		0, 133, 134, 3, 1, 0, 0, 134, 135, 3, 1, 0, 0, 135, 136, 3, 1, 0, 0, 136,
		140, 3, 1, 0, 0, 137, 139, 9, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 142,
		1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 143, 1, 0,
		0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 3, 3, 1, 0, 144, 28, 1, 0, 0, 0,
		145, 146, 5, 35, 0, 0, 146, 30, 1, 0, 0, 0, 147, 148, 5, 91, 0, 0, 148,
		32, 1, 0, 0, 0, 149, 150, 5, 93, 0, 0, 150, 34, 1, 0, 0, 0, 151, 152, 5,
		40, 0, 0, 152, 36, 1, 0, 0, 0, 153, 154, 5, 41, 0, 0, 154, 38, 1, 0, 0,
		0, 155, 156, 5, 33, 0, 0, 156, 40, 1, 0, 0, 0, 157, 159, 8, 4, 0, 0, 158,
		157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 42, 1, 0, 0, 0, 13, 0, 46, 49, 60, 73, 79, 87, 101, 112,
		125, 131, 140, 160, 0,
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
	MarkdownLexerBOLD_MARKER            = 4
	MarkdownLexerITALIC_MARKER          = 5
	MarkdownLexerSTRIKETHROUGH_MARKER   = 6
	MarkdownLexerBOLD_AND_ITALIC_MARKER = 7
	MarkdownLexerCODE_MARKER            = 8
	MarkdownLexerBLOCKQUOTE_MARKER      = 9
	MarkdownLexerUNORDERED_LIST_MARKER  = 10
	MarkdownLexerORDERED_LIST_MARKER    = 11
	MarkdownLexerHORIZONTAL_RULE        = 12
	MarkdownLexerFENCED_CODE_BLOCK      = 13
	MarkdownLexerINDENTED_CODE_BLOCK    = 14
	MarkdownLexerHASH                   = 15
	MarkdownLexerLBRACKET               = 16
	MarkdownLexerRBRACKET               = 17
	MarkdownLexerLPAREN                 = 18
	MarkdownLexerRPAREN                 = 19
	MarkdownLexerEXCLAMATION            = 20
	MarkdownLexerTEXT                   = 21
)
