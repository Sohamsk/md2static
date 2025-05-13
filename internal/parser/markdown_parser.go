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
		"", "", "", "", "", "'#'", "", "", "", "", "", "'-'", "'*'", "'+'",
		"", "", "'***'", "'**'", "'`'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "LINE_BREAK", "HASH", "H2",
		"H3", "H4", "H5", "H6", "DASH", "ASTERISK", "PLUS", "DASHES", "EQUALS",
		"BOLD_AND_ITALIC_MARKER", "BOLD_MARKER", "CODE_MARKER", "PUNCTUATION",
		"NUMBER", "WORD",
	}
	staticData.RuleNames = []string{
		"document", "block", "paragraph", "line", "escape_char", "linebreak",
		"inline", "italic", "italic_line", "bold", "bold_line", "bold_and_italic",
		"code", "code_text", "h1", "h2", "h3", "h4", "h5", "h6", "heading",
		"unorderedList", "unorderedListItem", "listMarker", "continuationLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 235, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 5, 0, 52, 8,
		0, 10, 0, 12, 0, 55, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 63,
		8, 1, 11, 1, 12, 1, 64, 3, 1, 67, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 72, 8,
		2, 3, 2, 74, 8, 2, 4, 2, 76, 8, 2, 11, 2, 12, 2, 77, 1, 2, 3, 2, 81, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 91, 8, 3, 11,
		3, 12, 3, 92, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 103,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 4, 8, 110, 8, 8, 11, 8, 12, 8, 111,
		1, 8, 4, 8, 115, 8, 8, 11, 8, 12, 8, 116, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		4, 10, 124, 8, 10, 11, 10, 12, 10, 125, 1, 10, 4, 10, 129, 8, 10, 11, 10,
		12, 10, 130, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 4, 13, 142, 8, 13, 11, 13, 12, 13, 143, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 150, 8, 14, 1, 14, 3, 14, 153, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 159, 8, 14, 3, 14, 161, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		167, 8, 15, 1, 15, 3, 15, 170, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		176, 8, 15, 3, 15, 178, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 184,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 190, 8, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 196, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 202, 8,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 210, 8, 20, 1, 21,
		4, 21, 213, 8, 21, 11, 21, 12, 21, 214, 1, 22, 1, 22, 1, 22, 1, 22, 5,
		22, 221, 8, 22, 10, 22, 12, 22, 224, 9, 22, 1, 22, 3, 22, 227, 8, 22, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 0, 0, 25, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 0, 3, 3, 0, 1, 1, 3, 3, 21, 21, 2, 0, 2, 2, 4, 4, 1, 0, 11, 13,
		256, 0, 53, 1, 0, 0, 0, 2, 66, 1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 90, 1,
		0, 0, 0, 8, 94, 1, 0, 0, 0, 10, 96, 1, 0, 0, 0, 12, 102, 1, 0, 0, 0, 14,
		104, 1, 0, 0, 0, 16, 114, 1, 0, 0, 0, 18, 118, 1, 0, 0, 0, 20, 128, 1,
		0, 0, 0, 22, 132, 1, 0, 0, 0, 24, 136, 1, 0, 0, 0, 26, 141, 1, 0, 0, 0,
		28, 160, 1, 0, 0, 0, 30, 177, 1, 0, 0, 0, 32, 179, 1, 0, 0, 0, 34, 185,
		1, 0, 0, 0, 36, 191, 1, 0, 0, 0, 38, 197, 1, 0, 0, 0, 40, 209, 1, 0, 0,
		0, 42, 212, 1, 0, 0, 0, 44, 216, 1, 0, 0, 0, 46, 228, 1, 0, 0, 0, 48, 231,
		1, 0, 0, 0, 50, 52, 3, 2, 1, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0,
		53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 56, 57, 5, 0, 0, 1, 57, 1, 1, 0, 0, 0, 58, 67, 3, 40, 20, 0, 59,
		67, 3, 4, 2, 0, 60, 67, 3, 42, 21, 0, 61, 63, 5, 2, 0, 0, 62, 61, 1, 0,
		0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67,
		1, 0, 0, 0, 66, 58, 1, 0, 0, 0, 66, 59, 1, 0, 0, 0, 66, 60, 1, 0, 0, 0,
		66, 62, 1, 0, 0, 0, 67, 3, 1, 0, 0, 0, 68, 73, 3, 6, 3, 0, 69, 74, 3, 10,
		5, 0, 70, 72, 5, 2, 0, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74,
		1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0,
		75, 68, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 81, 3, 10, 5, 0, 80, 75, 1, 0, 0, 0, 80,
		79, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82, 91, 5, 21, 0, 0, 83, 91, 5, 1, 0,
		0, 84, 91, 5, 20, 0, 0, 85, 91, 5, 19, 0, 0, 86, 91, 5, 14, 0, 0, 87, 91,
		5, 15, 0, 0, 88, 91, 3, 8, 4, 0, 89, 91, 3, 12, 6, 0, 90, 82, 1, 0, 0,
		0, 90, 83, 1, 0, 0, 0, 90, 84, 1, 0, 0, 0, 90, 85, 1, 0, 0, 0, 90, 86,
		1, 0, 0, 0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0,
		91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 7, 1, 0,
		0, 0, 94, 95, 5, 3, 0, 0, 95, 9, 1, 0, 0, 0, 96, 97, 5, 4, 0, 0, 97, 11,
		1, 0, 0, 0, 98, 103, 3, 14, 7, 0, 99, 103, 3, 18, 9, 0, 100, 103, 3, 22,
		11, 0, 101, 103, 3, 24, 12, 0, 102, 98, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0,
		102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 13, 1, 0, 0, 0, 104, 105,
		5, 12, 0, 0, 105, 106, 3, 16, 8, 0, 106, 107, 5, 12, 0, 0, 107, 15, 1,
		0, 0, 0, 108, 110, 7, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0,
		0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113,
		115, 3, 18, 9, 0, 114, 109, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 17, 1, 0,
		0, 0, 118, 119, 5, 17, 0, 0, 119, 120, 3, 20, 10, 0, 120, 121, 5, 17, 0,
		0, 121, 19, 1, 0, 0, 0, 122, 124, 7, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124,
		125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 129,
		1, 0, 0, 0, 127, 129, 3, 14, 7, 0, 128, 123, 1, 0, 0, 0, 128, 127, 1, 0,
		0, 0, 129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0,
		131, 21, 1, 0, 0, 0, 132, 133, 5, 16, 0, 0, 133, 134, 3, 20, 10, 0, 134,
		135, 5, 16, 0, 0, 135, 23, 1, 0, 0, 0, 136, 137, 5, 18, 0, 0, 137, 138,
		3, 26, 13, 0, 138, 139, 5, 18, 0, 0, 139, 25, 1, 0, 0, 0, 140, 142, 7,
		0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0,
		0, 143, 144, 1, 0, 0, 0, 144, 27, 1, 0, 0, 0, 145, 146, 5, 5, 0, 0, 146,
		147, 5, 1, 0, 0, 147, 152, 3, 6, 3, 0, 148, 150, 5, 2, 0, 0, 149, 148,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 153, 5, 4,
		0, 0, 152, 149, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 161, 1, 0, 0, 0,
		154, 155, 3, 6, 3, 0, 155, 156, 7, 1, 0, 0, 156, 158, 5, 15, 0, 0, 157,
		159, 5, 2, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161,
		1, 0, 0, 0, 160, 145, 1, 0, 0, 0, 160, 154, 1, 0, 0, 0, 161, 29, 1, 0,
		0, 0, 162, 163, 5, 6, 0, 0, 163, 164, 5, 1, 0, 0, 164, 169, 3, 6, 3, 0,
		165, 167, 5, 2, 0, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		170, 1, 0, 0, 0, 168, 170, 5, 4, 0, 0, 169, 166, 1, 0, 0, 0, 169, 168,
		1, 0, 0, 0, 170, 178, 1, 0, 0, 0, 171, 172, 3, 6, 3, 0, 172, 173, 7, 1,
		0, 0, 173, 175, 5, 14, 0, 0, 174, 176, 5, 2, 0, 0, 175, 174, 1, 0, 0, 0,
		175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177, 162, 1, 0, 0, 0, 177,
		171, 1, 0, 0, 0, 178, 31, 1, 0, 0, 0, 179, 180, 5, 7, 0, 0, 180, 181, 5,
		1, 0, 0, 181, 183, 3, 6, 3, 0, 182, 184, 5, 2, 0, 0, 183, 182, 1, 0, 0,
		0, 183, 184, 1, 0, 0, 0, 184, 33, 1, 0, 0, 0, 185, 186, 5, 8, 0, 0, 186,
		187, 5, 1, 0, 0, 187, 189, 3, 6, 3, 0, 188, 190, 5, 2, 0, 0, 189, 188,
		1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 35, 1, 0, 0, 0, 191, 192, 5, 9,
		0, 0, 192, 193, 5, 1, 0, 0, 193, 195, 3, 6, 3, 0, 194, 196, 5, 2, 0, 0,
		195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 37, 1, 0, 0, 0, 197, 198,
		5, 10, 0, 0, 198, 199, 5, 1, 0, 0, 199, 201, 3, 6, 3, 0, 200, 202, 5, 2,
		0, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 39, 1, 0, 0, 0,
		203, 210, 3, 28, 14, 0, 204, 210, 3, 30, 15, 0, 205, 210, 3, 32, 16, 0,
		206, 210, 3, 34, 17, 0, 207, 210, 3, 36, 18, 0, 208, 210, 3, 38, 19, 0,
		209, 203, 1, 0, 0, 0, 209, 204, 1, 0, 0, 0, 209, 205, 1, 0, 0, 0, 209,
		206, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 41, 1,
		0, 0, 0, 211, 213, 3, 44, 22, 0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0,
		0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 43, 1, 0, 0, 0,
		216, 217, 3, 46, 23, 0, 217, 222, 3, 6, 3, 0, 218, 219, 5, 2, 0, 0, 219,
		221, 3, 48, 24, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220,
		1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0,
		0, 0, 225, 227, 5, 2, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0,
		227, 45, 1, 0, 0, 0, 228, 229, 7, 2, 0, 0, 229, 230, 5, 1, 0, 0, 230, 47,
		1, 0, 0, 0, 231, 232, 5, 1, 0, 0, 232, 233, 3, 6, 3, 0, 233, 49, 1, 0,
		0, 0, 33, 53, 64, 66, 71, 73, 77, 80, 90, 92, 102, 111, 114, 116, 125,
		128, 130, 143, 149, 152, 158, 160, 166, 169, 175, 177, 183, 189, 195, 201,
		209, 214, 222, 226,
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
	MarkdownParserLINE_BREAK             = 4
	MarkdownParserHASH                   = 5
	MarkdownParserH2                     = 6
	MarkdownParserH3                     = 7
	MarkdownParserH4                     = 8
	MarkdownParserH5                     = 9
	MarkdownParserH6                     = 10
	MarkdownParserDASH                   = 11
	MarkdownParserASTERISK               = 12
	MarkdownParserPLUS                   = 13
	MarkdownParserDASHES                 = 14
	MarkdownParserEQUALS                 = 15
	MarkdownParserBOLD_AND_ITALIC_MARKER = 16
	MarkdownParserBOLD_MARKER            = 17
	MarkdownParserCODE_MARKER            = 18
	MarkdownParserPUNCTUATION            = 19
	MarkdownParserNUMBER                 = 20
	MarkdownParserWORD                   = 21
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_document          = 0
	MarkdownParserRULE_block             = 1
	MarkdownParserRULE_paragraph         = 2
	MarkdownParserRULE_line              = 3
	MarkdownParserRULE_escape_char       = 4
	MarkdownParserRULE_linebreak         = 5
	MarkdownParserRULE_inline            = 6
	MarkdownParserRULE_italic            = 7
	MarkdownParserRULE_italic_line       = 8
	MarkdownParserRULE_bold              = 9
	MarkdownParserRULE_bold_line         = 10
	MarkdownParserRULE_bold_and_italic   = 11
	MarkdownParserRULE_code              = 12
	MarkdownParserRULE_code_text         = 13
	MarkdownParserRULE_h1                = 14
	MarkdownParserRULE_h2                = 15
	MarkdownParserRULE_h3                = 16
	MarkdownParserRULE_h4                = 17
	MarkdownParserRULE_h5                = 18
	MarkdownParserRULE_h6                = 19
	MarkdownParserRULE_heading           = 20
	MarkdownParserRULE_unorderedList     = 21
	MarkdownParserRULE_unorderedListItem = 22
	MarkdownParserRULE_listMarker        = 23
	MarkdownParserRULE_continuationLine  = 24
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

func (p *MarkdownParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarkdownParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194302) != 0 {
		{
			p.SetState(50)
			p.Block()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
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
	UnorderedList() IUnorderedListContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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

func (p *MarkdownParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarkdownParserRULE_block)
	var _alt int

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Heading()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Paragraph()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.UnorderedList()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(61)
					p.Match(MarkdownParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllLinebreak() []ILinebreakContext
	Linebreak(i int) ILinebreakContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *ParagraphContext) AllLinebreak() []ILinebreakContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILinebreakContext); ok {
			len++
		}
	}

	tst := make([]ILinebreakContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILinebreakContext); ok {
			tst[i] = t.(ILinebreakContext)
			i++
		}
	}

	return tst
}

func (s *ParagraphContext) Linebreak(i int) ILinebreakContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinebreakContext); ok {
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

	return t.(ILinebreakContext)
}

func (s *ParagraphContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *ParagraphContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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

func (p *MarkdownParser) Paragraph() (localctx IParagraphContext) {
	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MarkdownParserRULE_paragraph)
	var _alt int

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserASTERISK, MarkdownParserDASHES, MarkdownParserEQUALS, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER, MarkdownParserPUNCTUATION, MarkdownParserNUMBER, MarkdownParserWORD:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(68)
					p.Line()
				}
				p.SetState(73)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(69)
						p.Linebreak()
					}

				case 2:
					p.SetState(71)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
						{
							p.SetState(70)
							p.Match(MarkdownParserNEWLINE)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					} else if p.HasError() { // JIM
						goto errorExit
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case MarkdownParserLINE_BREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Linebreak()
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

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllPUNCTUATION() []antlr.TerminalNode
	PUNCTUATION(i int) antlr.TerminalNode
	AllDASHES() []antlr.TerminalNode
	DASHES(i int) antlr.TerminalNode
	AllEQUALS() []antlr.TerminalNode
	EQUALS(i int) antlr.TerminalNode
	AllEscape_char() []IEscape_charContext
	Escape_char(i int) IEscape_charContext
	AllInline() []IInlineContext
	Inline(i int) IInlineContext

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

func (s *LineContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWORD)
}

func (s *LineContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWORD, i)
}

func (s *LineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *LineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *LineContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNUMBER)
}

func (s *LineContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNUMBER, i)
}

func (s *LineContext) AllPUNCTUATION() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserPUNCTUATION)
}

func (s *LineContext) PUNCTUATION(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserPUNCTUATION, i)
}

func (s *LineContext) AllDASHES() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserDASHES)
}

func (s *LineContext) DASHES(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserDASHES, i)
}

func (s *LineContext) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserEQUALS)
}

func (s *LineContext) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserEQUALS, i)
}

func (s *LineContext) AllEscape_char() []IEscape_charContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEscape_charContext); ok {
			len++
		}
	}

	tst := make([]IEscape_charContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEscape_charContext); ok {
			tst[i] = t.(IEscape_charContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) Escape_char(i int) IEscape_charContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEscape_charContext); ok {
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

	return t.(IEscape_charContext)
}

func (s *LineContext) AllInline() []IInlineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineContext); ok {
			len++
		}
	}

	tst := make([]IInlineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineContext); ok {
			tst[i] = t.(IInlineContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) Inline(i int) IInlineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineContext); ok {
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

	return t.(IInlineContext)
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

func (p *MarkdownParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MarkdownParserRULE_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MarkdownParserWORD:
				{
					p.SetState(82)
					p.Match(MarkdownParserWORD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserWHITESPACE:
				{
					p.SetState(83)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserNUMBER:
				{
					p.SetState(84)
					p.Match(MarkdownParserNUMBER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserPUNCTUATION:
				{
					p.SetState(85)
					p.Match(MarkdownParserPUNCTUATION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserDASHES:
				{
					p.SetState(86)
					p.Match(MarkdownParserDASHES)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserEQUALS:
				{
					p.SetState(87)
					p.Match(MarkdownParserEQUALS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserESCAPE_CHAR:
				{
					p.SetState(88)
					p.Escape_char()
				}

			case MarkdownParserASTERISK, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER:
				{
					p.SetState(89)
					p.Inline()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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

// IEscape_charContext is an interface to support dynamic dispatch.
type IEscape_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ESCAPE_CHAR() antlr.TerminalNode

	// IsEscape_charContext differentiates from other interfaces.
	IsEscape_charContext()
}

type Escape_charContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscape_charContext() *Escape_charContext {
	var p = new(Escape_charContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_escape_char
	return p
}

func InitEmptyEscape_charContext(p *Escape_charContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_escape_char
}

func (*Escape_charContext) IsEscape_charContext() {}

func NewEscape_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Escape_charContext {
	var p = new(Escape_charContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_escape_char

	return p
}

func (s *Escape_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Escape_charContext) ESCAPE_CHAR() antlr.TerminalNode {
	return s.GetToken(MarkdownParserESCAPE_CHAR, 0)
}

func (s *Escape_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Escape_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Escape_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterEscape_char(s)
	}
}

func (s *Escape_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitEscape_char(s)
	}
}

func (p *MarkdownParser) Escape_char() (localctx IEscape_charContext) {
	localctx = NewEscape_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MarkdownParserRULE_escape_char)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(MarkdownParserESCAPE_CHAR)
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

// ILinebreakContext is an interface to support dynamic dispatch.
type ILinebreakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LINE_BREAK() antlr.TerminalNode

	// IsLinebreakContext differentiates from other interfaces.
	IsLinebreakContext()
}

type LinebreakContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinebreakContext() *LinebreakContext {
	var p = new(LinebreakContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_linebreak
	return p
}

func InitEmptyLinebreakContext(p *LinebreakContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_linebreak
}

func (*LinebreakContext) IsLinebreakContext() {}

func NewLinebreakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinebreakContext {
	var p = new(LinebreakContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_linebreak

	return p
}

func (s *LinebreakContext) GetParser() antlr.Parser { return s.parser }

func (s *LinebreakContext) LINE_BREAK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserLINE_BREAK, 0)
}

func (s *LinebreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinebreakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinebreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterLinebreak(s)
	}
}

func (s *LinebreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitLinebreak(s)
	}
}

func (p *MarkdownParser) Linebreak() (localctx ILinebreakContext) {
	localctx = NewLinebreakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MarkdownParserRULE_linebreak)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(MarkdownParserLINE_BREAK)
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

// IInlineContext is an interface to support dynamic dispatch.
type IInlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Italic() IItalicContext
	Bold() IBoldContext
	Bold_and_italic() IBold_and_italicContext
	Code() ICodeContext

	// IsInlineContext differentiates from other interfaces.
	IsInlineContext()
}

type InlineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineContext() *InlineContext {
	var p = new(InlineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_inline
	return p
}

func InitEmptyInlineContext(p *InlineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_inline
}

func (*InlineContext) IsInlineContext() {}

func NewInlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineContext {
	var p = new(InlineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_inline

	return p
}

func (s *InlineContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineContext) Italic() IItalicContext {
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

func (s *InlineContext) Bold() IBoldContext {
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

func (s *InlineContext) Bold_and_italic() IBold_and_italicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBold_and_italicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBold_and_italicContext)
}

func (s *InlineContext) Code() ICodeContext {
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

func (s *InlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterInline(s)
	}
}

func (s *InlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitInline(s)
	}
}

func (p *MarkdownParser) Inline() (localctx IInlineContext) {
	localctx = NewInlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MarkdownParserRULE_inline)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserASTERISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Italic()
		}

	case MarkdownParserBOLD_MARKER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Bold()
		}

	case MarkdownParserBOLD_AND_ITALIC_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Bold_and_italic()
		}

	case MarkdownParserCODE_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
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

// IItalicContext is an interface to support dynamic dispatch.
type IItalicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IItalic_lineContext

	// SetData sets the data rule contexts.
	SetData(IItalic_lineContext)

	// Getter signatures
	AllASTERISK() []antlr.TerminalNode
	ASTERISK(i int) antlr.TerminalNode
	Italic_line() IItalic_lineContext

	// IsItalicContext differentiates from other interfaces.
	IsItalicContext()
}

type ItalicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IItalic_lineContext
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

func (s *ItalicContext) GetData() IItalic_lineContext { return s.data }

func (s *ItalicContext) SetData(v IItalic_lineContext) { s.data = v }

func (s *ItalicContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserASTERISK)
}

func (s *ItalicContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, i)
}

func (s *ItalicContext) Italic_line() IItalic_lineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalic_lineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalic_lineContext)
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

func (p *MarkdownParser) Italic() (localctx IItalicContext) {
	localctx = NewItalicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MarkdownParserRULE_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(MarkdownParserASTERISK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)

		var _x = p.Italic_line()

		localctx.(*ItalicContext).data = _x
	}
	{
		p.SetState(106)
		p.Match(MarkdownParserASTERISK)
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

// IItalic_lineContext is an interface to support dynamic dispatch.
type IItalic_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBold() []IBoldContext
	Bold(i int) IBoldContext
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllESCAPE_CHAR() []antlr.TerminalNode
	ESCAPE_CHAR(i int) antlr.TerminalNode

	// IsItalic_lineContext differentiates from other interfaces.
	IsItalic_lineContext()
}

type Italic_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItalic_lineContext() *Italic_lineContext {
	var p = new(Italic_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italic_line
	return p
}

func InitEmptyItalic_lineContext(p *Italic_lineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italic_line
}

func (*Italic_lineContext) IsItalic_lineContext() {}

func NewItalic_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Italic_lineContext {
	var p = new(Italic_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_italic_line

	return p
}

func (s *Italic_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Italic_lineContext) AllBold() []IBoldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoldContext); ok {
			len++
		}
	}

	tst := make([]IBoldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoldContext); ok {
			tst[i] = t.(IBoldContext)
			i++
		}
	}

	return tst
}

func (s *Italic_lineContext) Bold(i int) IBoldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldContext); ok {
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

	return t.(IBoldContext)
}

func (s *Italic_lineContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWORD)
}

func (s *Italic_lineContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWORD, i)
}

func (s *Italic_lineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *Italic_lineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *Italic_lineContext) AllESCAPE_CHAR() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserESCAPE_CHAR)
}

func (s *Italic_lineContext) ESCAPE_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserESCAPE_CHAR, i)
}

func (s *Italic_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Italic_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Italic_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterItalic_line(s)
	}
}

func (s *Italic_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitItalic_line(s)
	}
}

func (p *MarkdownParser) Italic_line() (localctx IItalic_lineContext) {
	localctx = NewItalic_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MarkdownParserRULE_italic_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2228234) != 0) {
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserWORD:
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(108)
						_la = p.GetTokenStream().LA(1)

						if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case MarkdownParserBOLD_MARKER:
			{
				p.SetState(113)
				p.Bold()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(116)
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

// IBoldContext is an interface to support dynamic dispatch.
type IBoldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IBold_lineContext

	// SetData sets the data rule contexts.
	SetData(IBold_lineContext)

	// Getter signatures
	AllBOLD_MARKER() []antlr.TerminalNode
	BOLD_MARKER(i int) antlr.TerminalNode
	Bold_line() IBold_lineContext

	// IsBoldContext differentiates from other interfaces.
	IsBoldContext()
}

type BoldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IBold_lineContext
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

func (s *BoldContext) GetData() IBold_lineContext { return s.data }

func (s *BoldContext) SetData(v IBold_lineContext) { s.data = v }

func (s *BoldContext) AllBOLD_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_MARKER)
}

func (s *BoldContext) BOLD_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_MARKER, i)
}

func (s *BoldContext) Bold_line() IBold_lineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBold_lineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBold_lineContext)
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

func (p *MarkdownParser) Bold() (localctx IBoldContext) {
	localctx = NewBoldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MarkdownParserRULE_bold)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)

		var _x = p.Bold_line()

		localctx.(*BoldContext).data = _x
	}
	{
		p.SetState(120)
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

// IBold_lineContext is an interface to support dynamic dispatch.
type IBold_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllItalic() []IItalicContext
	Italic(i int) IItalicContext
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllESCAPE_CHAR() []antlr.TerminalNode
	ESCAPE_CHAR(i int) antlr.TerminalNode

	// IsBold_lineContext differentiates from other interfaces.
	IsBold_lineContext()
}

type Bold_lineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBold_lineContext() *Bold_lineContext {
	var p = new(Bold_lineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold_line
	return p
}

func InitEmptyBold_lineContext(p *Bold_lineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold_line
}

func (*Bold_lineContext) IsBold_lineContext() {}

func NewBold_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bold_lineContext {
	var p = new(Bold_lineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_bold_line

	return p
}

func (s *Bold_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Bold_lineContext) AllItalic() []IItalicContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItalicContext); ok {
			len++
		}
	}

	tst := make([]IItalicContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItalicContext); ok {
			tst[i] = t.(IItalicContext)
			i++
		}
	}

	return tst
}

func (s *Bold_lineContext) Italic(i int) IItalicContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicContext); ok {
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

	return t.(IItalicContext)
}

func (s *Bold_lineContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWORD)
}

func (s *Bold_lineContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWORD, i)
}

func (s *Bold_lineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *Bold_lineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *Bold_lineContext) AllESCAPE_CHAR() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserESCAPE_CHAR)
}

func (s *Bold_lineContext) ESCAPE_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserESCAPE_CHAR, i)
}

func (s *Bold_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bold_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bold_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBold_line(s)
	}
}

func (s *Bold_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBold_line(s)
	}
}

func (p *MarkdownParser) Bold_line() (localctx IBold_lineContext) {
	localctx = NewBold_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MarkdownParserRULE_bold_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2101258) != 0) {
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserWORD:
			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(122)
						_la = p.GetTokenStream().LA(1)

						if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(125)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case MarkdownParserASTERISK:
			{
				p.SetState(127)
				p.Italic()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(130)
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

// IBold_and_italicContext is an interface to support dynamic dispatch.
type IBold_and_italicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IBold_lineContext

	// SetData sets the data rule contexts.
	SetData(IBold_lineContext)

	// Getter signatures
	AllBOLD_AND_ITALIC_MARKER() []antlr.TerminalNode
	BOLD_AND_ITALIC_MARKER(i int) antlr.TerminalNode
	Bold_line() IBold_lineContext

	// IsBold_and_italicContext differentiates from other interfaces.
	IsBold_and_italicContext()
}

type Bold_and_italicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IBold_lineContext
}

func NewEmptyBold_and_italicContext() *Bold_and_italicContext {
	var p = new(Bold_and_italicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold_and_italic
	return p
}

func InitEmptyBold_and_italicContext(p *Bold_and_italicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold_and_italic
}

func (*Bold_and_italicContext) IsBold_and_italicContext() {}

func NewBold_and_italicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bold_and_italicContext {
	var p = new(Bold_and_italicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_bold_and_italic

	return p
}

func (s *Bold_and_italicContext) GetParser() antlr.Parser { return s.parser }

func (s *Bold_and_italicContext) GetData() IBold_lineContext { return s.data }

func (s *Bold_and_italicContext) SetData(v IBold_lineContext) { s.data = v }

func (s *Bold_and_italicContext) AllBOLD_AND_ITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_AND_ITALIC_MARKER)
}

func (s *Bold_and_italicContext) BOLD_AND_ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_AND_ITALIC_MARKER, i)
}

func (s *Bold_and_italicContext) Bold_line() IBold_lineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBold_lineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBold_lineContext)
}

func (s *Bold_and_italicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bold_and_italicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bold_and_italicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBold_and_italic(s)
	}
}

func (s *Bold_and_italicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBold_and_italic(s)
	}
}

func (p *MarkdownParser) Bold_and_italic() (localctx IBold_and_italicContext) {
	localctx = NewBold_and_italicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MarkdownParserRULE_bold_and_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)

		var _x = p.Bold_line()

		localctx.(*Bold_and_italicContext).data = _x
	}
	{
		p.SetState(134)
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

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ICode_textContext

	// SetData sets the data rule contexts.
	SetData(ICode_textContext)

	// Getter signatures
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode
	Code_text() ICode_textContext

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ICode_textContext
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

func (s *CodeContext) GetData() ICode_textContext { return s.data }

func (s *CodeContext) SetData(v ICode_textContext) { s.data = v }

func (s *CodeContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *CodeContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *CodeContext) Code_text() ICode_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_textContext)
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

func (p *MarkdownParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MarkdownParserRULE_code)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)

		var _x = p.Code_text()

		localctx.(*CodeContext).data = _x
	}
	{
		p.SetState(138)
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

// ICode_textContext is an interface to support dynamic dispatch.
type ICode_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllESCAPE_CHAR() []antlr.TerminalNode
	ESCAPE_CHAR(i int) antlr.TerminalNode

	// IsCode_textContext differentiates from other interfaces.
	IsCode_textContext()
}

type Code_textContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCode_textContext() *Code_textContext {
	var p = new(Code_textContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_code_text
	return p
}

func InitEmptyCode_textContext(p *Code_textContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_code_text
}

func (*Code_textContext) IsCode_textContext() {}

func NewCode_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Code_textContext {
	var p = new(Code_textContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_code_text

	return p
}

func (s *Code_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Code_textContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWORD)
}

func (s *Code_textContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWORD, i)
}

func (s *Code_textContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *Code_textContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *Code_textContext) AllESCAPE_CHAR() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserESCAPE_CHAR)
}

func (s *Code_textContext) ESCAPE_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserESCAPE_CHAR, i)
}

func (s *Code_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Code_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Code_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCode_text(s)
	}
}

func (s *Code_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCode_text(s)
	}
}

func (p *MarkdownParser) Code_text() (localctx ICode_textContext) {
	localctx = NewCode_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MarkdownParserRULE_code_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
		{
			p.SetState(140)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(143)
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

// IH1Context is an interface to support dynamic dispatch.
type IH1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	HASH() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	LINE_BREAK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	EQUALS() antlr.TerminalNode

	// IsH1Context differentiates from other interfaces.
	IsH1Context()
}

type H1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH1Context() *H1Context {
	var p = new(H1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h1
	return p
}

func InitEmptyH1Context(p *H1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h1
}

func (*H1Context) IsH1Context() {}

func NewH1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H1Context {
	var p = new(H1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h1

	return p
}

func (s *H1Context) GetParser() antlr.Parser { return s.parser }

func (s *H1Context) GetData() ILineContext { return s.data }

func (s *H1Context) SetData(v ILineContext) { s.data = v }

func (s *H1Context) HASH() antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, 0)
}

func (s *H1Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H1Context) Line() ILineContext {
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

func (s *H1Context) LINE_BREAK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserLINE_BREAK, 0)
}

func (s *H1Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H1Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H1Context) EQUALS() antlr.TerminalNode {
	return s.GetToken(MarkdownParserEQUALS, 0)
}

func (s *H1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH1(s)
	}
}

func (s *H1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH1(s)
	}
}

func (p *MarkdownParser) H1() (localctx IH1Context) {
	localctx = NewH1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MarkdownParserRULE_h1)
	var _la int

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(MarkdownParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)

			var _x = p.Line()

			localctx.(*H1Context).data = _x
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			p.SetState(149)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(148)
					p.Match(MarkdownParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		case 2:
			{
				p.SetState(151)
				p.Match(MarkdownParserLINE_BREAK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserASTERISK, MarkdownParserDASHES, MarkdownParserEQUALS, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER, MarkdownParserPUNCTUATION, MarkdownParserNUMBER, MarkdownParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)

			var _x = p.Line()

			localctx.(*H1Context).data = _x
		}
		{
			p.SetState(155)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserNEWLINE || _la == MarkdownParserLINE_BREAK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(156)
			p.Match(MarkdownParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(157)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IH2Context is an interface to support dynamic dispatch.
type IH2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	H2() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	LINE_BREAK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	DASHES() antlr.TerminalNode

	// IsH2Context differentiates from other interfaces.
	IsH2Context()
}

type H2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH2Context() *H2Context {
	var p = new(H2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h2
	return p
}

func InitEmptyH2Context(p *H2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h2
}

func (*H2Context) IsH2Context() {}

func NewH2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H2Context {
	var p = new(H2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h2

	return p
}

func (s *H2Context) GetParser() antlr.Parser { return s.parser }

func (s *H2Context) GetData() ILineContext { return s.data }

func (s *H2Context) SetData(v ILineContext) { s.data = v }

func (s *H2Context) H2() antlr.TerminalNode {
	return s.GetToken(MarkdownParserH2, 0)
}

func (s *H2Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H2Context) Line() ILineContext {
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

func (s *H2Context) LINE_BREAK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserLINE_BREAK, 0)
}

func (s *H2Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H2Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H2Context) DASHES() antlr.TerminalNode {
	return s.GetToken(MarkdownParserDASHES, 0)
}

func (s *H2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH2(s)
	}
}

func (s *H2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH2(s)
	}
}

func (p *MarkdownParser) H2() (localctx IH2Context) {
	localctx = NewH2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MarkdownParserRULE_h2)
	var _la int

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserH2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(MarkdownParserH2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _x = p.Line()

			localctx.(*H2Context).data = _x
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
		case 1:
			p.SetState(166)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(165)
					p.Match(MarkdownParserNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		case 2:
			{
				p.SetState(168)
				p.Match(MarkdownParserLINE_BREAK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserASTERISK, MarkdownParserDASHES, MarkdownParserEQUALS, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER, MarkdownParserPUNCTUATION, MarkdownParserNUMBER, MarkdownParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)

			var _x = p.Line()

			localctx.(*H2Context).data = _x
		}
		{
			p.SetState(172)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserNEWLINE || _la == MarkdownParserLINE_BREAK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(173)
			p.Match(MarkdownParserDASHES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
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

// IH3Context is an interface to support dynamic dispatch.
type IH3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	H3() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	NEWLINE() antlr.TerminalNode

	// IsH3Context differentiates from other interfaces.
	IsH3Context()
}

type H3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH3Context() *H3Context {
	var p = new(H3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h3
	return p
}

func InitEmptyH3Context(p *H3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h3
}

func (*H3Context) IsH3Context() {}

func NewH3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H3Context {
	var p = new(H3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h3

	return p
}

func (s *H3Context) GetParser() antlr.Parser { return s.parser }

func (s *H3Context) GetData() ILineContext { return s.data }

func (s *H3Context) SetData(v ILineContext) { s.data = v }

func (s *H3Context) H3() antlr.TerminalNode {
	return s.GetToken(MarkdownParserH3, 0)
}

func (s *H3Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H3Context) Line() ILineContext {
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

func (s *H3Context) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *H3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH3(s)
	}
}

func (s *H3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH3(s)
	}
}

func (p *MarkdownParser) H3() (localctx IH3Context) {
	localctx = NewH3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MarkdownParserRULE_h3)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(MarkdownParserH3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)

		var _x = p.Line()

		localctx.(*H3Context).data = _x
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(182)
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

// IH4Context is an interface to support dynamic dispatch.
type IH4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	H4() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	NEWLINE() antlr.TerminalNode

	// IsH4Context differentiates from other interfaces.
	IsH4Context()
}

type H4Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH4Context() *H4Context {
	var p = new(H4Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h4
	return p
}

func InitEmptyH4Context(p *H4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h4
}

func (*H4Context) IsH4Context() {}

func NewH4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H4Context {
	var p = new(H4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h4

	return p
}

func (s *H4Context) GetParser() antlr.Parser { return s.parser }

func (s *H4Context) GetData() ILineContext { return s.data }

func (s *H4Context) SetData(v ILineContext) { s.data = v }

func (s *H4Context) H4() antlr.TerminalNode {
	return s.GetToken(MarkdownParserH4, 0)
}

func (s *H4Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H4Context) Line() ILineContext {
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

func (s *H4Context) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *H4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH4(s)
	}
}

func (s *H4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH4(s)
	}
}

func (p *MarkdownParser) H4() (localctx IH4Context) {
	localctx = NewH4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MarkdownParserRULE_h4)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(MarkdownParserH4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)

		var _x = p.Line()

		localctx.(*H4Context).data = _x
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(188)
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

// IH5Context is an interface to support dynamic dispatch.
type IH5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	H5() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	NEWLINE() antlr.TerminalNode

	// IsH5Context differentiates from other interfaces.
	IsH5Context()
}

type H5Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH5Context() *H5Context {
	var p = new(H5Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h5
	return p
}

func InitEmptyH5Context(p *H5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h5
}

func (*H5Context) IsH5Context() {}

func NewH5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H5Context {
	var p = new(H5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h5

	return p
}

func (s *H5Context) GetParser() antlr.Parser { return s.parser }

func (s *H5Context) GetData() ILineContext { return s.data }

func (s *H5Context) SetData(v ILineContext) { s.data = v }

func (s *H5Context) H5() antlr.TerminalNode {
	return s.GetToken(MarkdownParserH5, 0)
}

func (s *H5Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H5Context) Line() ILineContext {
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

func (s *H5Context) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *H5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH5(s)
	}
}

func (s *H5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH5(s)
	}
}

func (p *MarkdownParser) H5() (localctx IH5Context) {
	localctx = NewH5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MarkdownParserRULE_h5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(MarkdownParserH5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)

		var _x = p.Line()

		localctx.(*H5Context).data = _x
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(194)
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

// IH6Context is an interface to support dynamic dispatch.
type IH6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	H6() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	NEWLINE() antlr.TerminalNode

	// IsH6Context differentiates from other interfaces.
	IsH6Context()
}

type H6Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH6Context() *H6Context {
	var p = new(H6Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h6
	return p
}

func InitEmptyH6Context(p *H6Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h6
}

func (*H6Context) IsH6Context() {}

func NewH6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H6Context {
	var p = new(H6Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h6

	return p
}

func (s *H6Context) GetParser() antlr.Parser { return s.parser }

func (s *H6Context) GetData() ILineContext { return s.data }

func (s *H6Context) SetData(v ILineContext) { s.data = v }

func (s *H6Context) H6() antlr.TerminalNode {
	return s.GetToken(MarkdownParserH6, 0)
}

func (s *H6Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H6Context) Line() ILineContext {
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

func (s *H6Context) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *H6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH6(s)
	}
}

func (s *H6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH6(s)
	}
}

func (p *MarkdownParser) H6() (localctx IH6Context) {
	localctx = NewH6Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MarkdownParserRULE_h6)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(MarkdownParserH6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)

		var _x = p.Line()

		localctx.(*H6Context).data = _x
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(200)
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

// IHeadingContext is an interface to support dynamic dispatch.
type IHeadingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	H1() IH1Context
	H2() IH2Context
	H3() IH3Context
	H4() IH4Context
	H5() IH5Context
	H6() IH6Context

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

func (s *HeadingContext) H1() IH1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH1Context)
}

func (s *HeadingContext) H2() IH2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH2Context)
}

func (s *HeadingContext) H3() IH3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH3Context)
}

func (s *HeadingContext) H4() IH4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH4Context)
}

func (s *HeadingContext) H5() IH5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH5Context)
}

func (s *HeadingContext) H6() IH6Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH6Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH6Context)
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

func (p *MarkdownParser) Heading() (localctx IHeadingContext) {
	localctx = NewHeadingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MarkdownParserRULE_heading)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.H1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.H2()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.H3()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(206)
			p.H4()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(207)
			p.H5()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(208)
			p.H6()
		}

	case antlr.ATNInvalidAltNumber:
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

// IUnorderedListContext is an interface to support dynamic dispatch.
type IUnorderedListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnorderedListItem() []IUnorderedListItemContext
	UnorderedListItem(i int) IUnorderedListItemContext

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

func (s *UnorderedListContext) AllUnorderedListItem() []IUnorderedListItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnorderedListItemContext); ok {
			len++
		}
	}

	tst := make([]IUnorderedListItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnorderedListItemContext); ok {
			tst[i] = t.(IUnorderedListItemContext)
			i++
		}
	}

	return tst
}

func (s *UnorderedListContext) UnorderedListItem(i int) IUnorderedListItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnorderedListItemContext); ok {
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

	return t.(IUnorderedListItemContext)
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

func (p *MarkdownParser) UnorderedList() (localctx IUnorderedListContext) {
	localctx = NewUnorderedListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MarkdownParserRULE_unorderedList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(211)
				p.UnorderedListItem()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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

// IUnorderedListItemContext is an interface to support dynamic dispatch.
type IUnorderedListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListMarker() IListMarkerContext
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllContinuationLine() []IContinuationLineContext
	ContinuationLine(i int) IContinuationLineContext

	// IsUnorderedListItemContext differentiates from other interfaces.
	IsUnorderedListItemContext()
}

type UnorderedListItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnorderedListItemContext() *UnorderedListItemContext {
	var p = new(UnorderedListItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_unorderedListItem
	return p
}

func InitEmptyUnorderedListItemContext(p *UnorderedListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_unorderedListItem
}

func (*UnorderedListItemContext) IsUnorderedListItemContext() {}

func NewUnorderedListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnorderedListItemContext {
	var p = new(UnorderedListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_unorderedListItem

	return p
}

func (s *UnorderedListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *UnorderedListItemContext) ListMarker() IListMarkerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListMarkerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListMarkerContext)
}

func (s *UnorderedListItemContext) Line() ILineContext {
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

func (s *UnorderedListItemContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *UnorderedListItemContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *UnorderedListItemContext) AllContinuationLine() []IContinuationLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContinuationLineContext); ok {
			len++
		}
	}

	tst := make([]IContinuationLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContinuationLineContext); ok {
			tst[i] = t.(IContinuationLineContext)
			i++
		}
	}

	return tst
}

func (s *UnorderedListItemContext) ContinuationLine(i int) IContinuationLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuationLineContext); ok {
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

	return t.(IContinuationLineContext)
}

func (s *UnorderedListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnorderedListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnorderedListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterUnorderedListItem(s)
	}
}

func (s *UnorderedListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitUnorderedListItem(s)
	}
}

func (p *MarkdownParser) UnorderedListItem() (localctx IUnorderedListItemContext) {
	localctx = NewUnorderedListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MarkdownParserRULE_unorderedListItem)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.ListMarker()
	}
	{
		p.SetState(217)
		p.Line()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(218)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(219)
				p.ContinuationLine()
			}

		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(225)
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

// IListMarkerContext is an interface to support dynamic dispatch.
type IListMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHITESPACE() antlr.TerminalNode
	DASH() antlr.TerminalNode
	ASTERISK() antlr.TerminalNode
	PLUS() antlr.TerminalNode

	// IsListMarkerContext differentiates from other interfaces.
	IsListMarkerContext()
}

type ListMarkerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListMarkerContext() *ListMarkerContext {
	var p = new(ListMarkerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_listMarker
	return p
}

func InitEmptyListMarkerContext(p *ListMarkerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_listMarker
}

func (*ListMarkerContext) IsListMarkerContext() {}

func NewListMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListMarkerContext {
	var p = new(ListMarkerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_listMarker

	return p
}

func (s *ListMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *ListMarkerContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *ListMarkerContext) DASH() antlr.TerminalNode {
	return s.GetToken(MarkdownParserDASH, 0)
}

func (s *ListMarkerContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserASTERISK, 0)
}

func (s *ListMarkerContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MarkdownParserPLUS, 0)
}

func (s *ListMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterListMarker(s)
	}
}

func (s *ListMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitListMarker(s)
	}
}

func (p *MarkdownParser) ListMarker() (localctx IListMarkerContext) {
	localctx = NewListMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MarkdownParserRULE_listMarker)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(229)
		p.Match(MarkdownParserWHITESPACE)
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

// IContinuationLineContext is an interface to support dynamic dispatch.
type IContinuationLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext

	// IsContinuationLineContext differentiates from other interfaces.
	IsContinuationLineContext()
}

type ContinuationLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinuationLineContext() *ContinuationLineContext {
	var p = new(ContinuationLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_continuationLine
	return p
}

func InitEmptyContinuationLineContext(p *ContinuationLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_continuationLine
}

func (*ContinuationLineContext) IsContinuationLineContext() {}

func NewContinuationLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuationLineContext {
	var p = new(ContinuationLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_continuationLine

	return p
}

func (s *ContinuationLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuationLineContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *ContinuationLineContext) Line() ILineContext {
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

func (s *ContinuationLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuationLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuationLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterContinuationLine(s)
	}
}

func (s *ContinuationLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitContinuationLine(s)
	}
}

func (p *MarkdownParser) ContinuationLine() (localctx IContinuationLineContext) {
	localctx = NewContinuationLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MarkdownParserRULE_continuationLine)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Line()
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
