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
		"", "", "", "", "'#'", "", "", "", "", "", "'***'", "'*'", "'**'", "'`'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "HASH", "H2", "H3", "H4",
		"H5", "H6", "BOLD_AND_ITALIC_MARKER", "ITALIC_MARKER", "BOLD_MARKER",
		"CODE_MARKER", "PUNCTUATION", "NUMBER", "WORD",
	}
	staticData.RuleNames = []string{
		"document", "block", "line", "inline", "italic", "italic_line", "bold",
		"bold_line", "bold_and_italic", "code", "code_text", "h1", "h2", "h3",
		"h4", "h5", "h6", "paragraph", "heading",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 180, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 5, 0, 40, 8, 0, 10, 0,
		12, 0, 43, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 49, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 4, 2, 56, 8, 2, 11, 2, 12, 2, 57, 1, 2, 5, 2, 61, 8, 2,
		10, 2, 12, 2, 64, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 70, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 4, 5, 77, 8, 5, 11, 5, 12, 5, 78, 1, 5, 4, 5, 82,
		8, 5, 11, 5, 12, 5, 83, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 4, 7, 91, 8, 7, 11,
		7, 12, 7, 92, 1, 7, 4, 7, 96, 8, 7, 11, 7, 12, 7, 97, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 4, 10, 109, 8, 10, 11, 10, 12, 10,
		110, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 117, 8, 11, 10, 11, 12, 11, 120,
		9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 126, 8, 12, 10, 12, 12, 12, 129,
		9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 135, 8, 13, 10, 13, 12, 13, 138,
		9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 144, 8, 14, 10, 14, 12, 14, 147,
		9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 153, 8, 15, 10, 15, 12, 15, 156,
		9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 162, 8, 16, 10, 16, 12, 16, 165,
		9, 16, 1, 17, 4, 17, 168, 8, 17, 11, 17, 12, 17, 169, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 3, 18, 178, 8, 18, 1, 18, 0, 0, 19, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 0, 2, 2, 0,
		1, 1, 16, 16, 3, 0, 1, 1, 3, 3, 16, 16, 190, 0, 41, 1, 0, 0, 0, 2, 48,
		1, 0, 0, 0, 4, 55, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0, 8, 71, 1, 0, 0, 0, 10,
		81, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 99, 1, 0, 0,
		0, 18, 103, 1, 0, 0, 0, 20, 108, 1, 0, 0, 0, 22, 112, 1, 0, 0, 0, 24, 121,
		1, 0, 0, 0, 26, 130, 1, 0, 0, 0, 28, 139, 1, 0, 0, 0, 30, 148, 1, 0, 0,
		0, 32, 157, 1, 0, 0, 0, 34, 167, 1, 0, 0, 0, 36, 177, 1, 0, 0, 0, 38, 40,
		3, 2, 1, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0,
		41, 42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 5,
		0, 0, 1, 45, 1, 1, 0, 0, 0, 46, 49, 3, 36, 18, 0, 47, 49, 3, 34, 17, 0,
		48, 46, 1, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49, 3, 1, 0, 0, 0, 50, 56, 5, 16,
		0, 0, 51, 56, 5, 1, 0, 0, 52, 56, 5, 15, 0, 0, 53, 56, 5, 14, 0, 0, 54,
		56, 3, 6, 3, 0, 55, 50, 1, 0, 0, 0, 55, 51, 1, 0, 0, 0, 55, 52, 1, 0, 0,
		0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55,
		1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 62, 1, 0, 0, 0, 59, 61, 5, 2, 0, 0,
		60, 59, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1,
		0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 70, 3, 8, 4, 0, 66,
		70, 3, 12, 6, 0, 67, 70, 3, 16, 8, 0, 68, 70, 3, 18, 9, 0, 69, 65, 1, 0,
		0, 0, 69, 66, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 7,
		1, 0, 0, 0, 71, 72, 5, 11, 0, 0, 72, 73, 3, 10, 5, 0, 73, 74, 5, 11, 0,
		0, 74, 9, 1, 0, 0, 0, 75, 77, 7, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80,
		82, 3, 12, 6, 0, 81, 76, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 83, 1, 0,
		0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 11, 1, 0, 0, 0, 85, 86,
		5, 12, 0, 0, 86, 87, 3, 14, 7, 0, 87, 88, 5, 12, 0, 0, 88, 13, 1, 0, 0,
		0, 89, 91, 7, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 96, 3, 8, 4, 0,
		95, 90, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 15, 1, 0, 0, 0, 99, 100, 5, 10, 0, 0,
		100, 101, 3, 14, 7, 0, 101, 102, 5, 10, 0, 0, 102, 17, 1, 0, 0, 0, 103,
		104, 5, 13, 0, 0, 104, 105, 3, 20, 10, 0, 105, 106, 5, 13, 0, 0, 106, 19,
		1, 0, 0, 0, 107, 109, 7, 1, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 21, 1, 0, 0, 0,
		112, 113, 5, 4, 0, 0, 113, 114, 5, 1, 0, 0, 114, 118, 3, 4, 2, 0, 115,
		117, 5, 2, 0, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 23, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 121, 122, 5, 5, 0, 0, 122, 123, 5, 1, 0, 0, 123, 127, 3, 4, 2, 0,
		124, 126, 5, 2, 0, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 25, 1, 0, 0, 0, 129, 127, 1,
		0, 0, 0, 130, 131, 5, 6, 0, 0, 131, 132, 5, 1, 0, 0, 132, 136, 3, 4, 2,
		0, 133, 135, 5, 2, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136,
		134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 27, 1, 0, 0, 0, 138, 136, 1,
		0, 0, 0, 139, 140, 5, 7, 0, 0, 140, 141, 5, 1, 0, 0, 141, 145, 3, 4, 2,
		0, 142, 144, 5, 2, 0, 0, 143, 142, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145,
		143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 29, 1, 0, 0, 0, 147, 145, 1,
		0, 0, 0, 148, 149, 5, 8, 0, 0, 149, 150, 5, 1, 0, 0, 150, 154, 3, 4, 2,
		0, 151, 153, 5, 2, 0, 0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 31, 1, 0, 0, 0, 156, 154, 1,
		0, 0, 0, 157, 158, 5, 9, 0, 0, 158, 159, 5, 1, 0, 0, 159, 163, 3, 4, 2,
		0, 160, 162, 5, 2, 0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163,
		161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 33, 1, 0, 0, 0, 165, 163, 1,
		0, 0, 0, 166, 168, 3, 4, 2, 0, 167, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0,
		0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 35, 1, 0, 0, 0, 171,
		178, 3, 22, 11, 0, 172, 178, 3, 24, 12, 0, 173, 178, 3, 26, 13, 0, 174,
		178, 3, 28, 14, 0, 175, 178, 3, 30, 15, 0, 176, 178, 3, 32, 16, 0, 177,
		171, 1, 0, 0, 0, 177, 172, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0, 177, 174,
		1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 37, 1, 0,
		0, 0, 21, 41, 48, 55, 57, 62, 69, 78, 81, 83, 92, 95, 97, 110, 118, 127,
		136, 145, 154, 163, 169, 177,
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
	MarkdownParserHASH                   = 4
	MarkdownParserH2                     = 5
	MarkdownParserH3                     = 6
	MarkdownParserH4                     = 7
	MarkdownParserH5                     = 8
	MarkdownParserH6                     = 9
	MarkdownParserBOLD_AND_ITALIC_MARKER = 10
	MarkdownParserITALIC_MARKER          = 11
	MarkdownParserBOLD_MARKER            = 12
	MarkdownParserCODE_MARKER            = 13
	MarkdownParserPUNCTUATION            = 14
	MarkdownParserNUMBER                 = 15
	MarkdownParserWORD                   = 16
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_document        = 0
	MarkdownParserRULE_block           = 1
	MarkdownParserRULE_line            = 2
	MarkdownParserRULE_inline          = 3
	MarkdownParserRULE_italic          = 4
	MarkdownParserRULE_italic_line     = 5
	MarkdownParserRULE_bold            = 6
	MarkdownParserRULE_bold_line       = 7
	MarkdownParserRULE_bold_and_italic = 8
	MarkdownParserRULE_code            = 9
	MarkdownParserRULE_code_text       = 10
	MarkdownParserRULE_h1              = 11
	MarkdownParserRULE_h2              = 12
	MarkdownParserRULE_h3              = 13
	MarkdownParserRULE_h4              = 14
	MarkdownParserRULE_h5              = 15
	MarkdownParserRULE_h6              = 16
	MarkdownParserRULE_paragraph       = 17
	MarkdownParserRULE_heading         = 18
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&131058) != 0 {
		{
			p.SetState(38)
			p.Block()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserHASH, MarkdownParserH2, MarkdownParserH3, MarkdownParserH4, MarkdownParserH5, MarkdownParserH6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Heading()
		}

	case MarkdownParserWHITESPACE, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER, MarkdownParserPUNCTUATION, MarkdownParserNUMBER, MarkdownParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Paragraph()
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
	AllInline() []IInlineContext
	Inline(i int) IInlineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *LineContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *LineContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MarkdownParserWORD:
				{
					p.SetState(50)
					p.Match(MarkdownParserWORD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserWHITESPACE:
				{
					p.SetState(51)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserNUMBER:
				{
					p.SetState(52)
					p.Match(MarkdownParserNUMBER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserPUNCTUATION:
				{
					p.SetState(53)
					p.Match(MarkdownParserPUNCTUATION)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserBOLD_MARKER, MarkdownParserCODE_MARKER:
				{
					p.SetState(54)
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

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(59)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
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
	p.EnterRule(localctx, 6, MarkdownParserRULE_inline)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserITALIC_MARKER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Italic()
		}

	case MarkdownParserBOLD_MARKER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Bold()
		}

	case MarkdownParserBOLD_AND_ITALIC_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Bold_and_italic()
		}

	case MarkdownParserCODE_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
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
	p.EnterRule(localctx, 8, MarkdownParserRULE_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(MarkdownParserITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)

		var _x = p.Italic_line()

		localctx.(*ItalicContext).data = _x
	}
	{
		p.SetState(73)
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
	p.EnterRule(localctx, 10, MarkdownParserRULE_italic_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69634) != 0) {
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserWHITESPACE, MarkdownParserWORD:
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(75)
						_la = p.GetTokenStream().LA(1)

						if !(_la == MarkdownParserWHITESPACE || _la == MarkdownParserWORD) {
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

				p.SetState(78)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case MarkdownParserBOLD_MARKER:
			{
				p.SetState(80)
				p.Bold()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(83)
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
	p.EnterRule(localctx, 12, MarkdownParserRULE_bold)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)

		var _x = p.Bold_line()

		localctx.(*BoldContext).data = _x
	}
	{
		p.SetState(87)
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
	p.EnterRule(localctx, 14, MarkdownParserRULE_bold_line)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67586) != 0) {
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MarkdownParserWHITESPACE, MarkdownParserWORD:
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(89)
						_la = p.GetTokenStream().LA(1)

						if !(_la == MarkdownParserWHITESPACE || _la == MarkdownParserWORD) {
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

				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case MarkdownParserITALIC_MARKER:
			{
				p.SetState(94)
				p.Italic()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(97)
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
	p.EnterRule(localctx, 16, MarkdownParserRULE_bold_and_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)

		var _x = p.Bold_line()

		localctx.(*Bold_and_italicContext).data = _x
	}
	{
		p.SetState(101)
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
	p.EnterRule(localctx, 18, MarkdownParserRULE_code)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)

		var _x = p.Code_text()

		localctx.(*CodeContext).data = _x
	}
	{
		p.SetState(105)
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
	p.EnterRule(localctx, 20, MarkdownParserRULE_code_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65546) != 0) {
		{
			p.SetState(107)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65546) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(110)
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
	p.EnterRule(localctx, 22, MarkdownParserRULE_h1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)

		var _x = p.Line()

		localctx.(*H1Context).data = _x
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(115)
			p.Match(MarkdownParserNEWLINE)
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
	p.EnterRule(localctx, 24, MarkdownParserRULE_h2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(MarkdownParserH2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)

		var _x = p.Line()

		localctx.(*H2Context).data = _x
	}
	p.SetState(127)
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

		p.SetState(129)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *H3Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H3Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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
	p.EnterRule(localctx, 26, MarkdownParserRULE_h3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(MarkdownParserH3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)

		var _x = p.Line()

		localctx.(*H3Context).data = _x
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(133)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(138)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *H4Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H4Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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
	p.EnterRule(localctx, 28, MarkdownParserRULE_h4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(MarkdownParserH4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)

		var _x = p.Line()

		localctx.(*H4Context).data = _x
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(142)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(147)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *H5Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H5Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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
	p.EnterRule(localctx, 30, MarkdownParserRULE_h5)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(MarkdownParserH5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)

		var _x = p.Line()

		localctx.(*H5Context).data = _x
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(151)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(156)
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
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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

func (s *H6Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H6Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
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
	p.EnterRule(localctx, 32, MarkdownParserRULE_h6)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(MarkdownParserH6)
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

		localctx.(*H6Context).data = _x
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(160)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(165)
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
	p.EnterRule(localctx, 34, MarkdownParserRULE_paragraph)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(166)
				p.Line()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, MarkdownParserRULE_heading)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.H1()
		}

	case MarkdownParserH2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.H2()
		}

	case MarkdownParserH3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.H3()
		}

	case MarkdownParserH4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
			p.H4()
		}

	case MarkdownParserH5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(175)
			p.H5()
		}

	case MarkdownParserH6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(176)
			p.H6()
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
