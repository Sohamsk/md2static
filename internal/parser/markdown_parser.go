// Code generated from MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MarkdownParser struct {
	*antlr.BaseParser
}

var MarkdownParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func markdownparserParserInit() {
	staticData := &MarkdownParserParserStaticData
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
		"document", "block", "heading", "paragraph", "blockquote", "unorderedList",
		"orderedList", "listItem", "orderedListItem", "codeBlock", "horizontalRule",
		"line", "inlineElement", "bold", "boldAndItalic", "boldText", "italic",
		"italicText", "strikethrough", "strikethroughText", "code", "codeText",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 225, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 5, 0, 46, 8, 0, 10, 0, 12, 0, 49, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 60, 8, 1, 1, 2, 4, 2, 63,
		8, 2, 11, 2, 12, 2, 64, 1, 2, 1, 2, 1, 2, 5, 2, 70, 8, 2, 10, 2, 12, 2,
		73, 9, 2, 1, 3, 1, 3, 1, 3, 4, 3, 78, 8, 3, 11, 3, 12, 3, 79, 1, 3, 3,
		3, 83, 8, 3, 1, 4, 1, 4, 1, 4, 4, 4, 88, 8, 4, 11, 4, 12, 4, 89, 1, 4,
		1, 4, 1, 4, 1, 4, 4, 4, 96, 8, 4, 11, 4, 12, 4, 97, 5, 4, 100, 8, 4, 10,
		4, 12, 4, 103, 9, 4, 1, 4, 3, 4, 106, 8, 4, 1, 5, 4, 5, 109, 8, 5, 11,
		5, 12, 5, 110, 1, 6, 4, 6, 114, 8, 6, 11, 6, 12, 6, 115, 1, 7, 1, 7, 1,
		7, 4, 7, 121, 8, 7, 11, 7, 12, 7, 122, 1, 7, 1, 7, 1, 7, 3, 7, 128, 8,
		7, 1, 7, 1, 7, 4, 7, 132, 8, 7, 11, 7, 12, 7, 133, 5, 7, 136, 8, 7, 10,
		7, 12, 7, 139, 9, 7, 1, 8, 1, 8, 1, 8, 4, 8, 144, 8, 8, 11, 8, 12, 8, 145,
		1, 8, 1, 8, 1, 8, 3, 8, 151, 8, 8, 1, 8, 1, 8, 4, 8, 155, 8, 8, 11, 8,
		12, 8, 156, 5, 8, 159, 8, 8, 10, 8, 12, 8, 162, 9, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 4, 11, 171, 8, 11, 11, 11, 12, 11, 172, 1,
		11, 3, 11, 176, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 183, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 4, 15,
		194, 8, 15, 11, 15, 12, 15, 195, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 4,
		17, 203, 8, 17, 11, 17, 12, 17, 204, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		4, 19, 212, 8, 19, 11, 19, 12, 19, 213, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 4, 21, 221, 8, 21, 11, 21, 12, 21, 222, 1, 21, 0, 0, 22, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		0, 5, 1, 0, 13, 14, 4, 0, 1, 1, 5, 5, 8, 8, 21, 21, 3, 0, 1, 1, 8, 8, 21,
		21, 4, 0, 1, 1, 4, 5, 8, 8, 21, 21, 3, 0, 1, 1, 3, 3, 21, 21, 247, 0, 47,
		1, 0, 0, 0, 2, 59, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6, 77, 1, 0, 0, 0, 8,
		84, 1, 0, 0, 0, 10, 108, 1, 0, 0, 0, 12, 113, 1, 0, 0, 0, 14, 117, 1, 0,
		0, 0, 16, 140, 1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 165, 1, 0, 0, 0, 22,
		170, 1, 0, 0, 0, 24, 182, 1, 0, 0, 0, 26, 184, 1, 0, 0, 0, 28, 188, 1,
		0, 0, 0, 30, 193, 1, 0, 0, 0, 32, 197, 1, 0, 0, 0, 34, 202, 1, 0, 0, 0,
		36, 206, 1, 0, 0, 0, 38, 211, 1, 0, 0, 0, 40, 215, 1, 0, 0, 0, 42, 220,
		1, 0, 0, 0, 44, 46, 3, 2, 1, 0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 47, 1,
		0, 0, 0, 50, 51, 5, 0, 0, 1, 51, 1, 1, 0, 0, 0, 52, 60, 3, 4, 2, 0, 53,
		60, 3, 6, 3, 0, 54, 60, 3, 8, 4, 0, 55, 60, 3, 10, 5, 0, 56, 60, 3, 12,
		6, 0, 57, 60, 3, 18, 9, 0, 58, 60, 3, 20, 10, 0, 59, 52, 1, 0, 0, 0, 59,
		53, 1, 0, 0, 0, 59, 54, 1, 0, 0, 0, 59, 55, 1, 0, 0, 0, 59, 56, 1, 0, 0,
		0, 59, 57, 1, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 3, 1, 0, 0, 0, 61, 63, 5,
		15, 0, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 5, 1, 0, 0, 67, 71, 3, 22,
		11, 0, 68, 70, 5, 2, 0, 0, 69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71,
		69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 5, 1, 0, 0, 0, 73, 71, 1, 0, 0,
		0, 74, 78, 3, 22, 11, 0, 75, 78, 5, 1, 0, 0, 76, 78, 3, 24, 12, 0, 77,
		74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0,
		0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 83,
		5, 2, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 7, 1, 0, 0, 0,
		84, 87, 5, 9, 0, 0, 85, 88, 3, 22, 11, 0, 86, 88, 3, 24, 12, 0, 87, 85,
		1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 101, 1, 0, 0, 0, 91, 92, 5, 2, 0, 0, 92, 95, 5,
		9, 0, 0, 93, 96, 3, 22, 11, 0, 94, 96, 3, 24, 12, 0, 95, 93, 1, 0, 0, 0,
		95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1,
		0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0,
		101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101,
		1, 0, 0, 0, 104, 106, 5, 2, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0,
		0, 0, 106, 9, 1, 0, 0, 0, 107, 109, 3, 14, 7, 0, 108, 107, 1, 0, 0, 0,
		109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111,
		11, 1, 0, 0, 0, 112, 114, 3, 16, 8, 0, 113, 112, 1, 0, 0, 0, 114, 115,
		1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 13, 1, 0,
		0, 0, 117, 120, 5, 10, 0, 0, 118, 121, 3, 22, 11, 0, 119, 121, 3, 24, 12,
		0, 120, 118, 1, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 137, 1, 0, 0, 0, 124, 125,
		5, 2, 0, 0, 125, 127, 5, 1, 0, 0, 126, 128, 5, 10, 0, 0, 127, 126, 1, 0,
		0, 0, 127, 128, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 132, 3, 22, 11,
		0, 130, 132, 3, 24, 12, 0, 131, 129, 1, 0, 0, 0, 131, 130, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		136, 1, 0, 0, 0, 135, 124, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 15, 1, 0, 0, 0, 139, 137, 1, 0,
		0, 0, 140, 143, 5, 11, 0, 0, 141, 144, 3, 22, 11, 0, 142, 144, 3, 24, 12,
		0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 160, 1, 0, 0, 0, 147, 148,
		5, 2, 0, 0, 148, 150, 5, 1, 0, 0, 149, 151, 5, 11, 0, 0, 150, 149, 1, 0,
		0, 0, 150, 151, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 155, 3, 22, 11,
		0, 153, 155, 3, 24, 12, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0,
		155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157,
		159, 1, 0, 0, 0, 158, 147, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 17, 1, 0, 0, 0, 162, 160, 1, 0,
		0, 0, 163, 164, 7, 0, 0, 0, 164, 19, 1, 0, 0, 0, 165, 166, 5, 12, 0, 0,
		166, 21, 1, 0, 0, 0, 167, 171, 3, 24, 12, 0, 168, 171, 5, 1, 0, 0, 169,
		171, 5, 21, 0, 0, 170, 167, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 169,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0,
		0, 0, 173, 175, 1, 0, 0, 0, 174, 176, 5, 2, 0, 0, 175, 174, 1, 0, 0, 0,
		175, 176, 1, 0, 0, 0, 176, 23, 1, 0, 0, 0, 177, 183, 3, 26, 13, 0, 178,
		183, 3, 32, 16, 0, 179, 183, 3, 28, 14, 0, 180, 183, 3, 36, 18, 0, 181,
		183, 3, 40, 20, 0, 182, 177, 1, 0, 0, 0, 182, 178, 1, 0, 0, 0, 182, 179,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 25, 1, 0,
		0, 0, 184, 185, 5, 4, 0, 0, 185, 186, 3, 30, 15, 0, 186, 187, 5, 4, 0,
		0, 187, 27, 1, 0, 0, 0, 188, 189, 5, 7, 0, 0, 189, 190, 3, 34, 17, 0, 190,
		191, 5, 7, 0, 0, 191, 29, 1, 0, 0, 0, 192, 194, 7, 1, 0, 0, 193, 192, 1,
		0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0,
		0, 196, 31, 1, 0, 0, 0, 197, 198, 5, 5, 0, 0, 198, 199, 3, 34, 17, 0, 199,
		200, 5, 5, 0, 0, 200, 33, 1, 0, 0, 0, 201, 203, 7, 2, 0, 0, 202, 201, 1,
		0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0,
		0, 205, 35, 1, 0, 0, 0, 206, 207, 5, 6, 0, 0, 207, 208, 3, 38, 19, 0, 208,
		209, 5, 6, 0, 0, 209, 37, 1, 0, 0, 0, 210, 212, 7, 3, 0, 0, 211, 210, 1,
		0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0,
		0, 214, 39, 1, 0, 0, 0, 215, 216, 5, 8, 0, 0, 216, 217, 3, 42, 21, 0, 217,
		218, 5, 8, 0, 0, 218, 41, 1, 0, 0, 0, 219, 221, 7, 4, 0, 0, 220, 219, 1,
		0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0,
		0, 223, 43, 1, 0, 0, 0, 35, 47, 59, 64, 71, 77, 79, 82, 87, 89, 95, 97,
		101, 105, 110, 115, 120, 122, 127, 131, 133, 137, 143, 145, 150, 154, 156,
		160, 170, 172, 175, 182, 195, 204, 213, 222,
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

// MarkdownParserInit initializes any static state used to implement MarkdownParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMarkdownParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarkdownParserInit() {
	staticData := &MarkdownParserParserStaticData
	staticData.once.Do(markdownparserParserInit)
}

// NewMarkdownParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMarkdownParser(input antlr.TokenStream) *MarkdownParser {
	MarkdownParserInit()
	this := new(MarkdownParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MarkdownParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MarkdownParser.g4"

	return this
}

// MarkdownParser tokens.
const (
	MarkdownParserEOF                    = antlr.TokenEOF
	MarkdownParserWHITESPACE             = 1
	MarkdownParserNEWLINE                = 2
	MarkdownParserESCAPE_CHAR            = 3
	MarkdownParserBOLD_MARKER            = 4
	MarkdownParserITALIC_MARKER          = 5
	MarkdownParserSTRIKETHROUGH_MARKER   = 6
	MarkdownParserBOLD_AND_ITALIC_MARKER = 7
	MarkdownParserCODE_MARKER            = 8
	MarkdownParserBLOCKQUOTE_MARKER      = 9
	MarkdownParserUNORDERED_LIST_MARKER  = 10
	MarkdownParserORDERED_LIST_MARKER    = 11
	MarkdownParserHORIZONTAL_RULE        = 12
	MarkdownParserFENCED_CODE_BLOCK      = 13
	MarkdownParserINDENTED_CODE_BLOCK    = 14
	MarkdownParserHASH                   = 15
	MarkdownParserLBRACKET               = 16
	MarkdownParserRBRACKET               = 17
	MarkdownParserLPAREN                 = 18
	MarkdownParserRPAREN                 = 19
	MarkdownParserEXCLAMATION            = 20
	MarkdownParserTEXT                   = 21
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_document          = 0
	MarkdownParserRULE_block             = 1
	MarkdownParserRULE_heading           = 2
	MarkdownParserRULE_paragraph         = 3
	MarkdownParserRULE_blockquote        = 4
	MarkdownParserRULE_unorderedList     = 5
	MarkdownParserRULE_orderedList       = 6
	MarkdownParserRULE_listItem          = 7
	MarkdownParserRULE_orderedListItem   = 8
	MarkdownParserRULE_codeBlock         = 9
	MarkdownParserRULE_horizontalRule    = 10
	MarkdownParserRULE_line              = 11
	MarkdownParserRULE_inlineElement     = 12
	MarkdownParserRULE_bold              = 13
	MarkdownParserRULE_boldAndItalic     = 14
	MarkdownParserRULE_boldText          = 15
	MarkdownParserRULE_italic            = 16
	MarkdownParserRULE_italicText        = 17
	MarkdownParserRULE_strikethrough     = 18
	MarkdownParserRULE_strikethroughText = 19
	MarkdownParserRULE_code              = 20
	MarkdownParserRULE_codeText          = 21
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(MarkdownParserEOF, 0)
}

func (s *DocumentContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarkdownParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2162674) != 0 {
		{
			p.SetState(44)
			p.Block()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(MarkdownParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Heading() IHeadingContext
	Paragraph() IParagraphContext
	Blockquote() IBlockquoteContext
	UnorderedList() IUnorderedListContext
	OrderedList() IOrderedListContext
	CodeBlock() ICodeBlockContext
	HorizontalRule() IHorizontalRuleContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Heading() IHeadingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeadingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeadingContext)
}

func (s *BlockContext) Paragraph() IParagraphContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParagraphContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParagraphContext)
}

func (s *BlockContext) Blockquote() IBlockquoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockquoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockquoteContext)
}

func (s *BlockContext) UnorderedList() IUnorderedListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnorderedListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnorderedListContext)
}

func (s *BlockContext) OrderedList() IOrderedListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderedListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderedListContext)
}

func (s *BlockContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *BlockContext) HorizontalRule() IHorizontalRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHorizontalRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHorizontalRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarkdownParserRULE_block)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Heading()
		}

	case MarkdownParserWHITESPACE, MarkdownParserBOLD_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserSTRIKETHROUGH_MARKER, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserCODE_MARKER, MarkdownParserTEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Paragraph()
		}

	case MarkdownParserBLOCKQUOTE_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Blockquote()
		}

	case MarkdownParserUNORDERED_LIST_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.UnorderedList()
		}

	case MarkdownParserORDERED_LIST_MARKER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(56)
			p.OrderedList()
		}

	case MarkdownParserFENCED_CODE_BLOCK, MarkdownParserINDENTED_CODE_BLOCK:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(57)
			p.CodeBlock()
		}

	case MarkdownParserHORIZONTAL_RULE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(58)
			p.HorizontalRule()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeadingContext is an interface to support dynamic dispatch.
type IHeadingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsHeadingContext differentiates from other interfaces.
	IsHeadingContext()
}

type HeadingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeadingContext() *HeadingContext {
	var p = new(HeadingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_heading
	return p
}

func InitEmptyHeadingContext(p *HeadingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_heading
}

func (*HeadingContext) IsHeadingContext() {}

func NewHeadingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadingContext {
	var p = new(HeadingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_heading

	return p
}

func (s *HeadingContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadingContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *HeadingContext) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *HeadingContext) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserHASH)
}

func (s *HeadingContext) HASH(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, i)
}

func (s *HeadingContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *HeadingContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *HeadingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterHeading(s)
	}
}

func (s *HeadingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitHeading(s)
	}
}

func (s *HeadingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitHeading(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Heading() (localctx IHeadingContext) {
	localctx = NewHeadingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MarkdownParserRULE_heading)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MarkdownParserHASH {
		{
			p.SetState(61)
			p.Match(MarkdownParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Line()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(68)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	NEWLINE() antlr.TerminalNode

	// IsParagraphContext differentiates from other interfaces.
	IsParagraphContext()
}

type ParagraphContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParagraphContext() *ParagraphContext {
	var p = new(ParagraphContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_paragraph
	return p
}

func InitEmptyParagraphContext(p *ParagraphContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_paragraph
}

func (*ParagraphContext) IsParagraphContext() {}

func NewParagraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParagraphContext {
	var p = new(ParagraphContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_paragraph

	return p
}

func (s *ParagraphContext) GetParser() antlr.Parser { return s.parser }

func (s *ParagraphContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ParagraphContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ParagraphContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *ParagraphContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *ParagraphContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *ParagraphContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *ParagraphContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *ParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParagraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterParagraph(s)
	}
}

func (s *ParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitParagraph(s)
	}
}

func (s *ParagraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitParagraph(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Paragraph() (localctx IParagraphContext) {
	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MarkdownParserRULE_paragraph)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(74)
					p.Line()
				}

			case 2:
				{
					p.SetState(75)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				{
					p.SetState(76)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MarkdownParserNEWLINE {
		{
			p.SetState(81)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockquoteContext is an interface to support dynamic dispatch.
type IBlockquoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBLOCKQUOTE_MARKER() []antlr.TerminalNode
	BLOCKQUOTE_MARKER(i int) antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsBlockquoteContext differentiates from other interfaces.
	IsBlockquoteContext()
}

type BlockquoteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockquoteContext() *BlockquoteContext {
	var p = new(BlockquoteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_blockquote
	return p
}

func InitEmptyBlockquoteContext(p *BlockquoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_blockquote
}

func (*BlockquoteContext) IsBlockquoteContext() {}

func NewBlockquoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockquoteContext {
	var p = new(BlockquoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_blockquote

	return p
}

func (s *BlockquoteContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockquoteContext) AllBLOCKQUOTE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBLOCKQUOTE_MARKER)
}

func (s *BlockquoteContext) BLOCKQUOTE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBLOCKQUOTE_MARKER, i)
}

func (s *BlockquoteContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *BlockquoteContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *BlockquoteContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *BlockquoteContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *BlockquoteContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *BlockquoteContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *BlockquoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockquoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockquoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBlockquote(s)
	}
}

func (s *BlockquoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBlockquote(s)
	}
}

func (s *BlockquoteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitBlockquote(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Blockquote() (localctx IBlockquoteContext) {
	localctx = NewBlockquoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MarkdownParserRULE_blockquote)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(MarkdownParserBLOCKQUOTE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(85)
					p.Line()
				}

			case 2:
				{
					p.SetState(86)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(91)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(92)
				p.Match(MarkdownParserBLOCKQUOTE_MARKER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					p.SetState(95)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}

					switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
					case 1:
						{
							p.SetState(93)
							p.Line()
						}

					case 2:
						{
							p.SetState(94)
							p.InlineElement()
						}

					case antlr.ATNInvalidAltNumber:
						goto errorExit
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(97)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MarkdownParserNEWLINE {
		{
			p.SetState(104)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnorderedListContext is an interface to support dynamic dispatch.
type IUnorderedListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllListItem() []IListItemContext
	ListItem(i int) IListItemContext

	// IsUnorderedListContext differentiates from other interfaces.
	IsUnorderedListContext()
}

type UnorderedListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnorderedListContext() *UnorderedListContext {
	var p = new(UnorderedListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_unorderedList
	return p
}

func InitEmptyUnorderedListContext(p *UnorderedListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_unorderedList
}

func (*UnorderedListContext) IsUnorderedListContext() {}

func NewUnorderedListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnorderedListContext {
	var p = new(UnorderedListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_unorderedList

	return p
}

func (s *UnorderedListContext) GetParser() antlr.Parser { return s.parser }

func (s *UnorderedListContext) AllListItem() []IListItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListItemContext); ok {
			len++
		}
	}

	tst := make([]IListItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListItemContext); ok {
			tst[i] = t.(IListItemContext)
			i++
		}
	}

	return tst
}

func (s *UnorderedListContext) ListItem(i int) IListItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListItemContext)
}

func (s *UnorderedListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnorderedListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnorderedListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterUnorderedList(s)
	}
}

func (s *UnorderedListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitUnorderedList(s)
	}
}

func (s *UnorderedListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitUnorderedList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) UnorderedList() (localctx IUnorderedListContext) {
	localctx = NewUnorderedListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MarkdownParserRULE_unorderedList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(107)
				p.ListItem()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderedListContext is an interface to support dynamic dispatch.
type IOrderedListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOrderedListItem() []IOrderedListItemContext
	OrderedListItem(i int) IOrderedListItemContext

	// IsOrderedListContext differentiates from other interfaces.
	IsOrderedListContext()
}

type OrderedListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderedListContext() *OrderedListContext {
	var p = new(OrderedListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedList
	return p
}

func InitEmptyOrderedListContext(p *OrderedListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedList
}

func (*OrderedListContext) IsOrderedListContext() {}

func NewOrderedListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderedListContext {
	var p = new(OrderedListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_orderedList

	return p
}

func (s *OrderedListContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderedListContext) AllOrderedListItem() []IOrderedListItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderedListItemContext); ok {
			len++
		}
	}

	tst := make([]IOrderedListItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderedListItemContext); ok {
			tst[i] = t.(IOrderedListItemContext)
			i++
		}
	}

	return tst
}

func (s *OrderedListContext) OrderedListItem(i int) IOrderedListItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderedListItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderedListItemContext)
}

func (s *OrderedListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderedListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderedListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterOrderedList(s)
	}
}

func (s *OrderedListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitOrderedList(s)
	}
}

func (s *OrderedListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitOrderedList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) OrderedList() (localctx IOrderedListContext) {
	localctx = NewOrderedListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MarkdownParserRULE_orderedList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(112)
				p.OrderedListItem()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListItemContext is an interface to support dynamic dispatch.
type IListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUNORDERED_LIST_MARKER() []antlr.TerminalNode
	UNORDERED_LIST_MARKER(i int) antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsListItemContext differentiates from other interfaces.
	IsListItemContext()
}

type ListItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListItemContext() *ListItemContext {
	var p = new(ListItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_listItem
	return p
}

func InitEmptyListItemContext(p *ListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_listItem
}

func (*ListItemContext) IsListItemContext() {}

func NewListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListItemContext {
	var p = new(ListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_listItem

	return p
}

func (s *ListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ListItemContext) AllUNORDERED_LIST_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserUNORDERED_LIST_MARKER)
}

func (s *ListItemContext) UNORDERED_LIST_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserUNORDERED_LIST_MARKER, i)
}

func (s *ListItemContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ListItemContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ListItemContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *ListItemContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *ListItemContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *ListItemContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *ListItemContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *ListItemContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *ListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterListItem(s)
	}
}

func (s *ListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitListItem(s)
	}
}

func (s *ListItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitListItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) ListItem() (localctx IListItemContext) {
	localctx = NewListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MarkdownParserRULE_listItem)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(MarkdownParserUNORDERED_LIST_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(118)
					p.Line()
				}

			case 2:
				{
					p.SetState(119)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(124)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MarkdownParserUNORDERED_LIST_MARKER {
			{
				p.SetState(126)
				p.Match(MarkdownParserUNORDERED_LIST_MARKER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(131)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(129)
						p.Line()
					}

				case 2:
					{
						p.SetState(130)
						p.InlineElement()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderedListItemContext is an interface to support dynamic dispatch.
type IOrderedListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllORDERED_LIST_MARKER() []antlr.TerminalNode
	ORDERED_LIST_MARKER(i int) antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsOrderedListItemContext differentiates from other interfaces.
	IsOrderedListItemContext()
}

type OrderedListItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderedListItemContext() *OrderedListItemContext {
	var p = new(OrderedListItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedListItem
	return p
}

func InitEmptyOrderedListItemContext(p *OrderedListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedListItem
}

func (*OrderedListItemContext) IsOrderedListItemContext() {}

func NewOrderedListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderedListItemContext {
	var p = new(OrderedListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_orderedListItem

	return p
}

func (s *OrderedListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderedListItemContext) AllORDERED_LIST_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserORDERED_LIST_MARKER)
}

func (s *OrderedListItemContext) ORDERED_LIST_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserORDERED_LIST_MARKER, i)
}

func (s *OrderedListItemContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *OrderedListItemContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *OrderedListItemContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *OrderedListItemContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *OrderedListItemContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *OrderedListItemContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *OrderedListItemContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *OrderedListItemContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *OrderedListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderedListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderedListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterOrderedListItem(s)
	}
}

func (s *OrderedListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitOrderedListItem(s)
	}
}

func (s *OrderedListItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitOrderedListItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) OrderedListItem() (localctx IOrderedListItemContext) {
	localctx = NewOrderedListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MarkdownParserRULE_orderedListItem)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(MarkdownParserORDERED_LIST_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(141)
					p.Line()
				}

			case 2:
				{
					p.SetState(142)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(147)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MarkdownParserORDERED_LIST_MARKER {
			{
				p.SetState(149)
				p.Match(MarkdownParserORDERED_LIST_MARKER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(154)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(152)
						p.Line()
					}

				case 2:
					{
						p.SetState(153)
						p.InlineElement()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeBlockContext is an interface to support dynamic dispatch.
type ICodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FENCED_CODE_BLOCK() antlr.TerminalNode
	INDENTED_CODE_BLOCK() antlr.TerminalNode

	// IsCodeBlockContext differentiates from other interfaces.
	IsCodeBlockContext()
}

type CodeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContext() *CodeBlockContext {
	var p = new(CodeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeBlock
	return p
}

func InitEmptyCodeBlockContext(p *CodeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeBlock
}

func (*CodeBlockContext) IsCodeBlockContext() {}

func NewCodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContext {
	var p = new(CodeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_codeBlock

	return p
}

func (s *CodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContext) FENCED_CODE_BLOCK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserFENCED_CODE_BLOCK, 0)
}

func (s *CodeBlockContext) INDENTED_CODE_BLOCK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserINDENTED_CODE_BLOCK, 0)
}

func (s *CodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCodeBlock(s)
	}
}

func (s *CodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCodeBlock(s)
	}
}

func (s *CodeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitCodeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) CodeBlock() (localctx ICodeBlockContext) {
	localctx = NewCodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MarkdownParserRULE_codeBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MarkdownParserFENCED_CODE_BLOCK || _la == MarkdownParserINDENTED_CODE_BLOCK) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHorizontalRuleContext is an interface to support dynamic dispatch.
type IHorizontalRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HORIZONTAL_RULE() antlr.TerminalNode

	// IsHorizontalRuleContext differentiates from other interfaces.
	IsHorizontalRuleContext()
}

type HorizontalRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHorizontalRuleContext() *HorizontalRuleContext {
	var p = new(HorizontalRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_horizontalRule
	return p
}

func InitEmptyHorizontalRuleContext(p *HorizontalRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_horizontalRule
}

func (*HorizontalRuleContext) IsHorizontalRuleContext() {}

func NewHorizontalRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HorizontalRuleContext {
	var p = new(HorizontalRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_horizontalRule

	return p
}

func (s *HorizontalRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *HorizontalRuleContext) HORIZONTAL_RULE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserHORIZONTAL_RULE, 0)
}

func (s *HorizontalRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HorizontalRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HorizontalRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterHorizontalRule(s)
	}
}

func (s *HorizontalRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitHorizontalRule(s)
	}
}

func (s *HorizontalRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitHorizontalRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) HorizontalRule() (localctx IHorizontalRuleContext) {
	localctx = NewHorizontalRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MarkdownParserRULE_horizontalRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(MarkdownParserHORIZONTAL_RULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	NEWLINE() antlr.TerminalNode

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *LineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *LineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *LineContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *LineContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitLine(s)
	}
}

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MarkdownParserRULE_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MarkdownParserBOLD_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserSTRIKETHROUGH_MARKER, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserCODE_MARKER:
				{
					p.SetState(167)
					p.InlineElement()
				}

			case MarkdownParserWHITESPACE:
				{
					p.SetState(168)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserTEXT:
				{
					p.SetState(169)
					p.Match(MarkdownParserTEXT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(174)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInlineElementContext is an interface to support dynamic dispatch.
type IInlineElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bold() IBoldContext
	Italic() IItalicContext
	BoldAndItalic() IBoldAndItalicContext
	Strikethrough() IStrikethroughContext
	Code() ICodeContext

	// IsInlineElementContext differentiates from other interfaces.
	IsInlineElementContext()
}

type InlineElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineElementContext() *InlineElementContext {
	var p = new(InlineElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_inlineElement
	return p
}

func InitEmptyInlineElementContext(p *InlineElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_inlineElement
}

func (*InlineElementContext) IsInlineElementContext() {}

func NewInlineElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineElementContext {
	var p = new(InlineElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_inlineElement

	return p
}

func (s *InlineElementContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineElementContext) Bold() IBoldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoldContext)
}

func (s *InlineElementContext) Italic() IItalicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalicContext)
}

func (s *InlineElementContext) BoldAndItalic() IBoldAndItalicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldAndItalicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoldAndItalicContext)
}

func (s *InlineElementContext) Strikethrough() IStrikethroughContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrikethroughContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrikethroughContext)
}

func (s *InlineElementContext) Code() ICodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *InlineElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterInlineElement(s)
	}
}

func (s *InlineElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitInlineElement(s)
	}
}

func (s *InlineElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitInlineElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) InlineElement() (localctx IInlineElementContext) {
	localctx = NewInlineElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MarkdownParserRULE_inlineElement)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserBOLD_MARKER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Bold()
		}

	case MarkdownParserITALIC_MARKER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.Italic()
		}

	case MarkdownParserBOLD_AND_ITALIC_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(179)
			p.BoldAndItalic()
		}

	case MarkdownParserSTRIKETHROUGH_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(180)
			p.Strikethrough()
		}

	case MarkdownParserCODE_MARKER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(181)
			p.Code()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoldContext is an interface to support dynamic dispatch.
type IBoldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetText returns the text rule contexts.
	GetText() IBoldTextContext

	// SetText sets the text rule contexts.
	SetText(IBoldTextContext)

	// Getter signatures
	AllBOLD_MARKER() []antlr.TerminalNode
	BOLD_MARKER(i int) antlr.TerminalNode
	BoldText() IBoldTextContext

	// IsBoldContext differentiates from other interfaces.
	IsBoldContext()
}

type BoldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   IBoldTextContext
}

func NewEmptyBoldContext() *BoldContext {
	var p = new(BoldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold
	return p
}

func InitEmptyBoldContext(p *BoldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold
}

func (*BoldContext) IsBoldContext() {}

func NewBoldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldContext {
	var p = new(BoldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_bold

	return p
}

func (s *BoldContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldContext) GetText() IBoldTextContext { return s.text }

func (s *BoldContext) SetText(v IBoldTextContext) { s.text = v }

func (s *BoldContext) AllBOLD_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_MARKER)
}

func (s *BoldContext) BOLD_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_MARKER, i)
}

func (s *BoldContext) BoldText() IBoldTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoldTextContext)
}

func (s *BoldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBold(s)
	}
}

func (s *BoldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBold(s)
	}
}

func (s *BoldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitBold(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Bold() (localctx IBoldContext) {
	localctx = NewBoldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MarkdownParserRULE_bold)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)

		var _x = p.BoldText()

		localctx.(*BoldContext).text = _x
	}
	{
		p.SetState(186)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoldAndItalicContext is an interface to support dynamic dispatch.
type IBoldAndItalicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetText returns the text rule contexts.
	GetText() IItalicTextContext

	// SetText sets the text rule contexts.
	SetText(IItalicTextContext)

	// Getter signatures
	AllBOLD_AND_ITALIC_MARKER() []antlr.TerminalNode
	BOLD_AND_ITALIC_MARKER(i int) antlr.TerminalNode
	ItalicText() IItalicTextContext

	// IsBoldAndItalicContext differentiates from other interfaces.
	IsBoldAndItalicContext()
}

type BoldAndItalicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   IItalicTextContext
}

func NewEmptyBoldAndItalicContext() *BoldAndItalicContext {
	var p = new(BoldAndItalicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldAndItalic
	return p
}

func InitEmptyBoldAndItalicContext(p *BoldAndItalicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldAndItalic
}

func (*BoldAndItalicContext) IsBoldAndItalicContext() {}

func NewBoldAndItalicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldAndItalicContext {
	var p = new(BoldAndItalicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_boldAndItalic

	return p
}

func (s *BoldAndItalicContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldAndItalicContext) GetText() IItalicTextContext { return s.text }

func (s *BoldAndItalicContext) SetText(v IItalicTextContext) { s.text = v }

func (s *BoldAndItalicContext) AllBOLD_AND_ITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_AND_ITALIC_MARKER)
}

func (s *BoldAndItalicContext) BOLD_AND_ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_AND_ITALIC_MARKER, i)
}

func (s *BoldAndItalicContext) ItalicText() IItalicTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalicTextContext)
}

func (s *BoldAndItalicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldAndItalicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldAndItalicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBoldAndItalic(s)
	}
}

func (s *BoldAndItalicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBoldAndItalic(s)
	}
}

func (s *BoldAndItalicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitBoldAndItalic(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) BoldAndItalic() (localctx IBoldAndItalicContext) {
	localctx = NewBoldAndItalicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MarkdownParserRULE_boldAndItalic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)

		var _x = p.ItalicText()

		localctx.(*BoldAndItalicContext).text = _x
	}
	{
		p.SetState(190)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoldTextContext is an interface to support dynamic dispatch.
type IBoldTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode

	// IsBoldTextContext differentiates from other interfaces.
	IsBoldTextContext()
}

type BoldTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoldTextContext() *BoldTextContext {
	var p = new(BoldTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldText
	return p
}

func InitEmptyBoldTextContext(p *BoldTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldText
}

func (*BoldTextContext) IsBoldTextContext() {}

func NewBoldTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldTextContext {
	var p = new(BoldTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_boldText

	return p
}

func (s *BoldTextContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *BoldTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *BoldTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *BoldTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *BoldTextContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *BoldTextContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
}

func (s *BoldTextContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *BoldTextContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *BoldTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBoldText(s)
	}
}

func (s *BoldTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBoldText(s)
	}
}

func (s *BoldTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitBoldText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) BoldText() (localctx IBoldTextContext) {
	localctx = NewBoldTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MarkdownParserRULE_boldText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097442) != 0) {
		{
			p.SetState(192)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097442) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItalicContext is an interface to support dynamic dispatch.
type IItalicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetText returns the text rule contexts.
	GetText() IItalicTextContext

	// SetText sets the text rule contexts.
	SetText(IItalicTextContext)

	// Getter signatures
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
	ItalicText() IItalicTextContext

	// IsItalicContext differentiates from other interfaces.
	IsItalicContext()
}

type ItalicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   IItalicTextContext
}

func NewEmptyItalicContext() *ItalicContext {
	var p = new(ItalicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italic
	return p
}

func InitEmptyItalicContext(p *ItalicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italic
}

func (*ItalicContext) IsItalicContext() {}

func NewItalicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItalicContext {
	var p = new(ItalicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_italic

	return p
}

func (s *ItalicContext) GetParser() antlr.Parser { return s.parser }

func (s *ItalicContext) GetText() IItalicTextContext { return s.text }

func (s *ItalicContext) SetText(v IItalicTextContext) { s.text = v }

func (s *ItalicContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *ItalicContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
}

func (s *ItalicContext) ItalicText() IItalicTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalicTextContext)
}

func (s *ItalicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItalicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItalicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterItalic(s)
	}
}

func (s *ItalicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitItalic(s)
	}
}

func (s *ItalicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitItalic(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Italic() (localctx IItalicContext) {
	localctx = NewItalicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MarkdownParserRULE_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(MarkdownParserITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)

		var _x = p.ItalicText()

		localctx.(*ItalicContext).text = _x
	}
	{
		p.SetState(199)
		p.Match(MarkdownParserITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItalicTextContext is an interface to support dynamic dispatch.
type IItalicTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode

	// IsItalicTextContext differentiates from other interfaces.
	IsItalicTextContext()
}

type ItalicTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItalicTextContext() *ItalicTextContext {
	var p = new(ItalicTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italicText
	return p
}

func InitEmptyItalicTextContext(p *ItalicTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italicText
}

func (*ItalicTextContext) IsItalicTextContext() {}

func NewItalicTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItalicTextContext {
	var p = new(ItalicTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_italicText

	return p
}

func (s *ItalicTextContext) GetParser() antlr.Parser { return s.parser }

func (s *ItalicTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *ItalicTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *ItalicTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *ItalicTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *ItalicTextContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *ItalicTextContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *ItalicTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItalicTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItalicTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterItalicText(s)
	}
}

func (s *ItalicTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitItalicText(s)
	}
}

func (s *ItalicTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitItalicText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) ItalicText() (localctx IItalicTextContext) {
	localctx = NewItalicTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MarkdownParserRULE_italicText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097410) != 0) {
		{
			p.SetState(201)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097410) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrikethroughContext is an interface to support dynamic dispatch.
type IStrikethroughContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetText returns the text rule contexts.
	GetText() IStrikethroughTextContext

	// SetText sets the text rule contexts.
	SetText(IStrikethroughTextContext)

	// Getter signatures
	AllSTRIKETHROUGH_MARKER() []antlr.TerminalNode
	STRIKETHROUGH_MARKER(i int) antlr.TerminalNode
	StrikethroughText() IStrikethroughTextContext

	// IsStrikethroughContext differentiates from other interfaces.
	IsStrikethroughContext()
}

type StrikethroughContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   IStrikethroughTextContext
}

func NewEmptyStrikethroughContext() *StrikethroughContext {
	var p = new(StrikethroughContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethrough
	return p
}

func InitEmptyStrikethroughContext(p *StrikethroughContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethrough
}

func (*StrikethroughContext) IsStrikethroughContext() {}

func NewStrikethroughContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrikethroughContext {
	var p = new(StrikethroughContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_strikethrough

	return p
}

func (s *StrikethroughContext) GetParser() antlr.Parser { return s.parser }

func (s *StrikethroughContext) GetText() IStrikethroughTextContext { return s.text }

func (s *StrikethroughContext) SetText(v IStrikethroughTextContext) { s.text = v }

func (s *StrikethroughContext) AllSTRIKETHROUGH_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSTRIKETHROUGH_MARKER)
}

func (s *StrikethroughContext) STRIKETHROUGH_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSTRIKETHROUGH_MARKER, i)
}

func (s *StrikethroughContext) StrikethroughText() IStrikethroughTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrikethroughTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrikethroughTextContext)
}

func (s *StrikethroughContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrikethroughContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrikethroughContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterStrikethrough(s)
	}
}

func (s *StrikethroughContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitStrikethrough(s)
	}
}

func (s *StrikethroughContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitStrikethrough(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Strikethrough() (localctx IStrikethroughContext) {
	localctx = NewStrikethroughContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MarkdownParserRULE_strikethrough)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(MarkdownParserSTRIKETHROUGH_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)

		var _x = p.StrikethroughText()

		localctx.(*StrikethroughContext).text = _x
	}
	{
		p.SetState(208)
		p.Match(MarkdownParserSTRIKETHROUGH_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrikethroughTextContext is an interface to support dynamic dispatch.
type IStrikethroughTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
	AllBOLD_MARKER() []antlr.TerminalNode
	BOLD_MARKER(i int) antlr.TerminalNode
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode

	// IsStrikethroughTextContext differentiates from other interfaces.
	IsStrikethroughTextContext()
}

type StrikethroughTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrikethroughTextContext() *StrikethroughTextContext {
	var p = new(StrikethroughTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethroughText
	return p
}

func InitEmptyStrikethroughTextContext(p *StrikethroughTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethroughText
}

func (*StrikethroughTextContext) IsStrikethroughTextContext() {}

func NewStrikethroughTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrikethroughTextContext {
	var p = new(StrikethroughTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_strikethroughText

	return p
}

func (s *StrikethroughTextContext) GetParser() antlr.Parser { return s.parser }

func (s *StrikethroughTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *StrikethroughTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *StrikethroughTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *StrikethroughTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *StrikethroughTextContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *StrikethroughTextContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
}

func (s *StrikethroughTextContext) AllBOLD_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_MARKER)
}

func (s *StrikethroughTextContext) BOLD_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_MARKER, i)
}

func (s *StrikethroughTextContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *StrikethroughTextContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *StrikethroughTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrikethroughTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrikethroughTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterStrikethroughText(s)
	}
}

func (s *StrikethroughTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitStrikethroughText(s)
	}
}

func (s *StrikethroughTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitStrikethroughText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) StrikethroughText() (localctx IStrikethroughTextContext) {
	localctx = NewStrikethroughTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MarkdownParserRULE_strikethroughText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097458) != 0) {
		{
			p.SetState(210)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097458) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetText returns the text rule contexts.
	GetText() ICodeTextContext

	// SetText sets the text rule contexts.
	SetText(ICodeTextContext)

	// Getter signatures
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode
	CodeText() ICodeTextContext

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   ICodeTextContext
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_code
	return p
}

func InitEmptyCodeContext(p *CodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_code
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) GetText() ICodeTextContext { return s.text }

func (s *CodeContext) SetText(v ICodeTextContext) { s.text = v }

func (s *CodeContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *CodeContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *CodeContext) CodeText() ICodeTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeTextContext)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCode(s)
	}
}

func (s *CodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCode(s)
	}
}

func (s *CodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MarkdownParserRULE_code)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)

		var _x = p.CodeText()

		localctx.(*CodeContext).text = _x
	}
	{
		p.SetState(217)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeTextContext is an interface to support dynamic dispatch.
type ICodeTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllESCAPE_CHAR() []antlr.TerminalNode
	ESCAPE_CHAR(i int) antlr.TerminalNode

	// IsCodeTextContext differentiates from other interfaces.
	IsCodeTextContext()
}

type CodeTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeTextContext() *CodeTextContext {
	var p = new(CodeTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeText
	return p
}

func InitEmptyCodeTextContext(p *CodeTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeText
}

func (*CodeTextContext) IsCodeTextContext() {}

func NewCodeTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeTextContext {
	var p = new(CodeTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_codeText

	return p
}

func (s *CodeTextContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *CodeTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *CodeTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *CodeTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *CodeTextContext) AllESCAPE_CHAR() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserESCAPE_CHAR)
}

func (s *CodeTextContext) ESCAPE_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserESCAPE_CHAR, i)
}

func (s *CodeTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCodeText(s)
	}
}

func (s *CodeTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCodeText(s)
	}
}

func (s *CodeTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitCodeText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParser) CodeText() (localctx ICodeTextContext) {
	localctx = NewCodeTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MarkdownParserRULE_codeText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
		{
			p.SetState(219)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
