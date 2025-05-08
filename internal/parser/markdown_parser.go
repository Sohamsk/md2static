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
		"", "", "", "", "", "'#'", "", "", "", "", "", "", "", "'***'", "'*'",
		"'**'", "'`'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "LINE_BREAK", "HASH", "H2",
		"H3", "H4", "H5", "H6", "DASHES", "EQUALS", "BOLD_AND_ITALIC_MARKER",
		"ITALIC_MARKER", "BOLD_MARKER", "CODE_MARKER", "PUNCTUATION", "NUMBER",
		"WORD",
	}
	staticData.RuleNames = []string{
		"document", "block", "line", "escape_char", "linebreak", "inline", "italic",
		"italic_line", "bold", "bold_line", "bold_and_italic", "code", "code_text",
		"h1", "h2", "h3", "h4", "h5", "h6", "paragraph", "heading",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 195, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		4, 1, 54, 8, 1, 11, 1, 12, 1, 55, 1, 1, 3, 1, 59, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 69, 8, 2, 11, 2, 12, 2, 70, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 81, 8, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 4, 7, 88, 8, 7, 11, 7, 12, 7, 89, 1, 7, 4, 7, 93, 8, 7,
		11, 7, 12, 7, 94, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 4, 9, 102, 8, 9, 11, 9,
		12, 9, 103, 1, 9, 4, 9, 107, 8, 9, 11, 9, 12, 9, 108, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 120, 8, 12, 11, 12,
		12, 12, 121, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 128, 8, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 134, 8, 13, 3, 13, 136, 8, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 142, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 148, 8, 14,
		3, 14, 150, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 156, 8, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 162, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 168, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 174, 8, 18, 1, 19, 1,
		19, 1, 19, 3, 19, 179, 8, 19, 3, 19, 181, 8, 19, 4, 19, 183, 8, 19, 11,
		19, 12, 19, 184, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 193,
		8, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 0, 2, 3, 0, 1, 1, 3, 3, 19, 19, 2, 0, 2,
		2, 4, 4, 214, 0, 45, 1, 0, 0, 0, 2, 58, 1, 0, 0, 0, 4, 68, 1, 0, 0, 0,
		6, 72, 1, 0, 0, 0, 8, 74, 1, 0, 0, 0, 10, 80, 1, 0, 0, 0, 12, 82, 1, 0,
		0, 0, 14, 92, 1, 0, 0, 0, 16, 96, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0, 20,
		110, 1, 0, 0, 0, 22, 114, 1, 0, 0, 0, 24, 119, 1, 0, 0, 0, 26, 135, 1,
		0, 0, 0, 28, 149, 1, 0, 0, 0, 30, 151, 1, 0, 0, 0, 32, 157, 1, 0, 0, 0,
		34, 163, 1, 0, 0, 0, 36, 169, 1, 0, 0, 0, 38, 182, 1, 0, 0, 0, 40, 192,
		1, 0, 0, 0, 42, 44, 3, 2, 1, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0,
		45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47, 45, 1,
		0, 0, 0, 48, 49, 5, 0, 0, 1, 49, 1, 1, 0, 0, 0, 50, 59, 3, 40, 20, 0, 51,
		59, 3, 38, 19, 0, 52, 54, 5, 2, 0, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0,
		0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 59,
		5, 4, 0, 0, 58, 50, 1, 0, 0, 0, 58, 51, 1, 0, 0, 0, 58, 53, 1, 0, 0, 0,
		58, 57, 1, 0, 0, 0, 59, 3, 1, 0, 0, 0, 60, 69, 5, 19, 0, 0, 61, 69, 5,
		1, 0, 0, 62, 69, 5, 18, 0, 0, 63, 69, 5, 17, 0, 0, 64, 69, 5, 11, 0, 0,
		65, 69, 5, 12, 0, 0, 66, 69, 3, 6, 3, 0, 67, 69, 3, 10, 5, 0, 68, 60, 1,
		0, 0, 0, 68, 61, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 63, 1, 0, 0, 0, 68,
		64, 1, 0, 0, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1, 0, 0,
		0, 69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 5, 1,
		0, 0, 0, 72, 73, 5, 3, 0, 0, 73, 7, 1, 0, 0, 0, 74, 75, 5, 4, 0, 0, 75,
		9, 1, 0, 0, 0, 76, 81, 3, 12, 6, 0, 77, 81, 3, 16, 8, 0, 78, 81, 3, 20,
		10, 0, 79, 81, 3, 22, 11, 0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 80,
		78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 11, 1, 0, 0, 0, 82, 83, 5, 14,
		0, 0, 83, 84, 3, 14, 7, 0, 84, 85, 5, 14, 0, 0, 85, 13, 1, 0, 0, 0, 86,
		88, 7, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0,
		0, 89, 90, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 93, 3, 16, 8, 0, 92, 87,
		1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 15, 1, 0, 0, 0, 96, 97, 5, 15, 0, 0, 97, 98, 3,
		18, 9, 0, 98, 99, 5, 15, 0, 0, 99, 17, 1, 0, 0, 0, 100, 102, 7, 0, 0, 0,
		101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 107, 3, 12, 6, 0, 106, 101,
		1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 19, 1, 0, 0, 0, 110, 111, 5, 13, 0, 0,
		111, 112, 3, 18, 9, 0, 112, 113, 5, 13, 0, 0, 113, 21, 1, 0, 0, 0, 114,
		115, 5, 16, 0, 0, 115, 116, 3, 24, 12, 0, 116, 117, 5, 16, 0, 0, 117, 23,
		1, 0, 0, 0, 118, 120, 7, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 121, 1, 0,
		0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 25, 1, 0, 0, 0,
		123, 124, 5, 5, 0, 0, 124, 125, 5, 1, 0, 0, 125, 127, 3, 4, 2, 0, 126,
		128, 5, 2, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 136,
		1, 0, 0, 0, 129, 130, 3, 4, 2, 0, 130, 131, 7, 1, 0, 0, 131, 133, 5, 12,
		0, 0, 132, 134, 5, 2, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 136, 1, 0, 0, 0, 135, 123, 1, 0, 0, 0, 135, 129, 1, 0, 0, 0, 136,
		27, 1, 0, 0, 0, 137, 138, 5, 6, 0, 0, 138, 139, 5, 1, 0, 0, 139, 141, 3,
		4, 2, 0, 140, 142, 5, 2, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0,
		0, 142, 150, 1, 0, 0, 0, 143, 144, 3, 4, 2, 0, 144, 145, 7, 1, 0, 0, 145,
		147, 5, 11, 0, 0, 146, 148, 5, 2, 0, 0, 147, 146, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 137, 1, 0, 0, 0, 149, 143, 1, 0,
		0, 0, 150, 29, 1, 0, 0, 0, 151, 152, 5, 7, 0, 0, 152, 153, 5, 1, 0, 0,
		153, 155, 3, 4, 2, 0, 154, 156, 5, 2, 0, 0, 155, 154, 1, 0, 0, 0, 155,
		156, 1, 0, 0, 0, 156, 31, 1, 0, 0, 0, 157, 158, 5, 8, 0, 0, 158, 159, 5,
		1, 0, 0, 159, 161, 3, 4, 2, 0, 160, 162, 5, 2, 0, 0, 161, 160, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 33, 1, 0, 0, 0, 163, 164, 5, 9, 0, 0, 164,
		165, 5, 1, 0, 0, 165, 167, 3, 4, 2, 0, 166, 168, 5, 2, 0, 0, 167, 166,
		1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 35, 1, 0, 0, 0, 169, 170, 5, 10,
		0, 0, 170, 171, 5, 1, 0, 0, 171, 173, 3, 4, 2, 0, 172, 174, 5, 2, 0, 0,
		173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 37, 1, 0, 0, 0, 175, 180,
		3, 4, 2, 0, 176, 181, 3, 8, 4, 0, 177, 179, 5, 2, 0, 0, 178, 177, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180, 176, 1, 0, 0, 0,
		180, 178, 1, 0, 0, 0, 181, 183, 1, 0, 0, 0, 182, 175, 1, 0, 0, 0, 183,
		184, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 39, 1,
		0, 0, 0, 186, 193, 3, 26, 13, 0, 187, 193, 3, 28, 14, 0, 188, 193, 3, 30,
		15, 0, 189, 193, 3, 32, 16, 0, 190, 193, 3, 34, 17, 0, 191, 193, 3, 36,
		18, 0, 192, 186, 1, 0, 0, 0, 192, 187, 1, 0, 0, 0, 192, 188, 1, 0, 0, 0,
		192, 189, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193,
		41, 1, 0, 0, 0, 27, 45, 55, 58, 68, 70, 80, 89, 92, 94, 103, 106, 108,
		121, 127, 133, 135, 141, 147, 149, 155, 161, 167, 173, 178, 180, 184, 192,
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
	MarkdownParserDASHES                 = 11
	MarkdownParserEQUALS                 = 12
	MarkdownParserBOLD_AND_ITALIC_MARKER = 13
	MarkdownParserITALIC_MARKER          = 14
	MarkdownParserBOLD_MARKER            = 15
	MarkdownParserCODE_MARKER            = 16
	MarkdownParserPUNCTUATION            = 17
	MarkdownParserNUMBER                 = 18
	MarkdownParserWORD                   = 19
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_document        = 0
	MarkdownParserRULE_block           = 1
	MarkdownParserRULE_line            = 2
	MarkdownParserRULE_escape_char     = 3
	MarkdownParserRULE_linebreak       = 4
	MarkdownParserRULE_inline          = 5
	MarkdownParserRULE_italic          = 6
	MarkdownParserRULE_italic_line     = 7
	MarkdownParserRULE_bold            = 8
	MarkdownParserRULE_bold_line       = 9
	MarkdownParserRULE_bold_and_italic = 10
	MarkdownParserRULE_code            = 11
	MarkdownParserRULE_code_text       = 12
	MarkdownParserRULE_h1              = 13
	MarkdownParserRULE_h2              = 14
	MarkdownParserRULE_h3              = 15
	MarkdownParserRULE_h4              = 16
	MarkdownParserRULE_h5              = 17
	MarkdownParserRULE_h6              = 18
	MarkdownParserRULE_paragraph       = 19
	MarkdownParserRULE_heading         = 20
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1048574) != 0 {
		{
			p.SetState(42)
			p.Block()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	LINE_BREAK() antlr.TerminalNode

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

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *BlockContext) LINE_BREAK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserLINE_BREAK, 0)
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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Heading()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Paragraph()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(52)
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

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Match(MarkdownParserLINE_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
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
	p.EnterRule(localctx, 4, MarkdownParserRULE_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MarkdownParserWORD:
				{
					p.SetState(60)
					p.Match(MarkdownParserWORD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserWHITESPACE:
				{
					p.SetState(61)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserNUMBER:
				{
					p.SetState(62)
					p.Match(MarkdownParserNUMBER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserPUNCTUATION:
				{
					p.SetState(63)
					p.Match(MarkdownParserPUNCTUATION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserDASHES:
				{
					p.SetState(64)
					p.Match(MarkdownParserDASHES)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserEQUALS:
				{
					p.SetState(65)
					p.Match(MarkdownParserEQUALS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserESCAPE_CHAR:
				{
					p.SetState(66)
					p.Escape_char()
				}

			case MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER:
				{
					p.SetState(67)
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

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, MarkdownParserRULE_escape_char)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
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
	p.EnterRule(localctx, 8, MarkdownParserRULE_linebreak)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
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
	p.EnterRule(localctx, 10, MarkdownParserRULE_inline)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserITALIC_MARKER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Italic()
		}

	case MarkdownParserBOLD_MARKER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Bold()
		}

	case MarkdownParserBOLD_AND_ITALIC_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Bold_and_italic()
		}

	case MarkdownParserCODE_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
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
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
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

func (s *ItalicContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *ItalicContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
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
	p.EnterRule(localctx, 12, MarkdownParserRULE_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(MarkdownParserITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)

		var _x = p.Italic_line()

		localctx.(*ItalicContext).data = _x
	}
	{
		p.SetState(84)
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
	p.EnterRule(localctx, 14, MarkdownParserRULE_italic_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&557066) != 0) {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserWORD:
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(86)
						_la = p.GetTokenStream().LA(1)

						if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&524298) != 0) {
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

				p.SetState(89)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case MarkdownParserBOLD_MARKER:
			{
				p.SetState(91)
				p.Bold()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(94)
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
	p.EnterRule(localctx, 16, MarkdownParserRULE_bold)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)

		var _x = p.Bold_line()

		localctx.(*BoldContext).data = _x
	}
	{
		p.SetState(98)
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
	p.EnterRule(localctx, 18, MarkdownParserRULE_bold_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&540682) != 0) {
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserWORD:
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(100)
						_la = p.GetTokenStream().LA(1)

						if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&524298) != 0) {
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

				p.SetState(103)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case MarkdownParserITALIC_MARKER:
			{
				p.SetState(105)
				p.Italic()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(108)
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
	p.EnterRule(localctx, 20, MarkdownParserRULE_bold_and_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)

		var _x = p.Bold_line()

		localctx.(*Bold_and_italicContext).data = _x
	}
	{
		p.SetState(112)
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
	p.EnterRule(localctx, 22, MarkdownParserRULE_code)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)

		var _x = p.Code_text()

		localctx.(*CodeContext).data = _x
	}
	{
		p.SetState(116)
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
	p.EnterRule(localctx, 24, MarkdownParserRULE_code_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&524298) != 0) {
		{
			p.SetState(118)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&524298) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(121)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	LINE_BREAK() antlr.TerminalNode

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

func (s *H1Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H1Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H1Context) EQUALS() antlr.TerminalNode {
	return s.GetToken(MarkdownParserEQUALS, 0)
}

func (s *H1Context) LINE_BREAK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserLINE_BREAK, 0)
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
	p.EnterRule(localctx, 26, MarkdownParserRULE_h1)
	var _la int

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(MarkdownParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)

			var _x = p.Line()

			localctx.(*H1Context).data = _x
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(126)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserDASHES, MarkdownParserEQUALS, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER, MarkdownParserPUNCTUATION, MarkdownParserNUMBER, MarkdownParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)

			var _x = p.Line()

			localctx.(*H1Context).data = _x
		}
		{
			p.SetState(130)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserNEWLINE || _la == MarkdownParserLINE_BREAK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(131)
			p.Match(MarkdownParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(132)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	DASHES() antlr.TerminalNode
	LINE_BREAK() antlr.TerminalNode

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

func (s *H2Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H2Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H2Context) DASHES() antlr.TerminalNode {
	return s.GetToken(MarkdownParserDASHES, 0)
}

func (s *H2Context) LINE_BREAK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserLINE_BREAK, 0)
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
	p.EnterRule(localctx, 28, MarkdownParserRULE_h2)
	var _la int

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserH2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(MarkdownParserH2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)

			var _x = p.Line()

			localctx.(*H2Context).data = _x
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(140)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MarkdownParserWHITESPACE, MarkdownParserESCAPE_CHAR, MarkdownParserDASHES, MarkdownParserEQUALS, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER, MarkdownParserPUNCTUATION, MarkdownParserNUMBER, MarkdownParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)

			var _x = p.Line()

			localctx.(*H2Context).data = _x
		}
		{
			p.SetState(144)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MarkdownParserNEWLINE || _la == MarkdownParserLINE_BREAK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(145)
			p.Match(MarkdownParserDASHES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(146)
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
	p.EnterRule(localctx, 30, MarkdownParserRULE_h3)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(MarkdownParserH3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)

		var _x = p.Line()

		localctx.(*H3Context).data = _x
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(154)
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
	p.EnterRule(localctx, 32, MarkdownParserRULE_h4)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(MarkdownParserH4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)

		var _x = p.Line()

		localctx.(*H4Context).data = _x
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(160)
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
	p.EnterRule(localctx, 34, MarkdownParserRULE_h5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(MarkdownParserH5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)

		var _x = p.Line()

		localctx.(*H5Context).data = _x
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(166)
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
	p.EnterRule(localctx, 36, MarkdownParserRULE_h6)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(MarkdownParserH6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)

		var _x = p.Line()

		localctx.(*H6Context).data = _x
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(172)
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
	p.EnterRule(localctx, 38, MarkdownParserRULE_paragraph)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(175)
				p.Line()
			}
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(176)
					p.Linebreak()
				}

			case 2:
				p.SetState(178)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(177)
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

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.H1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.H2()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.H3()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.H4()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(190)
			p.H5()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(191)
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
