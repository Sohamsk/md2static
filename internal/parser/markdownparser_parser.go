// Code generated from ./internal/parser/MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type MarkdownParserParser struct {
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
		"", "", "", "", "", "'**'", "'*'", "'__'", "'_'", "'['", "']'", "'('",
		"')'",
	}
	staticData.SymbolicNames = []string{
		"", "H1", "H2", "H3", "LIST_MARKER", "DOUBLE_ASTERISK", "SINGLE_ASTERISK",
		"DOUBLE_UNDERSCORE", "SINGLE_UNDERSCORE", "LBRACK", "RBRACK", "LPAREN",
		"RPAREN", "TEXT", "WS", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"document", "block", "header", "list", "listItem", "paragraph", "inline",
		"inlineElement", "bold", "italic", "link",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 96, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 4, 0, 24, 8, 0, 11, 0, 12, 0, 25, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 3, 1, 33, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 4, 3, 50, 8, 3, 11, 3, 12, 3, 51,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 4, 6, 63, 8, 6, 11,
		6, 12, 6, 64, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 71, 8, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 79, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 87, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 0,
		0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 0, 96, 0, 23, 1, 0, 0,
		0, 2, 32, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 53, 1, 0,
		0, 0, 10, 58, 1, 0, 0, 0, 12, 62, 1, 0, 0, 0, 14, 70, 1, 0, 0, 0, 16, 78,
		1, 0, 0, 0, 18, 86, 1, 0, 0, 0, 20, 88, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0,
		23, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1,
		0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 5, 0, 0, 1, 28, 1, 1, 0, 0, 0, 29,
		33, 3, 4, 2, 0, 30, 33, 3, 6, 3, 0, 31, 33, 3, 10, 5, 0, 32, 29, 1, 0,
		0, 0, 32, 30, 1, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 3, 1, 0, 0, 0, 34, 35,
		5, 1, 0, 0, 35, 36, 3, 12, 6, 0, 36, 37, 5, 15, 0, 0, 37, 47, 1, 0, 0,
		0, 38, 39, 5, 2, 0, 0, 39, 40, 3, 12, 6, 0, 40, 41, 5, 15, 0, 0, 41, 47,
		1, 0, 0, 0, 42, 43, 5, 3, 0, 0, 43, 44, 3, 12, 6, 0, 44, 45, 5, 15, 0,
		0, 45, 47, 1, 0, 0, 0, 46, 34, 1, 0, 0, 0, 46, 38, 1, 0, 0, 0, 46, 42,
		1, 0, 0, 0, 47, 5, 1, 0, 0, 0, 48, 50, 3, 8, 4, 0, 49, 48, 1, 0, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 7, 1, 0,
		0, 0, 53, 54, 5, 4, 0, 0, 54, 55, 5, 14, 0, 0, 55, 56, 3, 12, 6, 0, 56,
		57, 5, 15, 0, 0, 57, 9, 1, 0, 0, 0, 58, 59, 3, 12, 6, 0, 59, 60, 5, 15,
		0, 0, 60, 11, 1, 0, 0, 0, 61, 63, 3, 14, 7, 0, 62, 61, 1, 0, 0, 0, 63,
		64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 13, 1, 0, 0,
		0, 66, 71, 3, 16, 8, 0, 67, 71, 3, 18, 9, 0, 68, 71, 3, 20, 10, 0, 69,
		71, 5, 13, 0, 0, 70, 66, 1, 0, 0, 0, 70, 67, 1, 0, 0, 0, 70, 68, 1, 0,
		0, 0, 70, 69, 1, 0, 0, 0, 71, 15, 1, 0, 0, 0, 72, 73, 5, 5, 0, 0, 73, 74,
		5, 13, 0, 0, 74, 79, 5, 5, 0, 0, 75, 76, 5, 7, 0, 0, 76, 77, 5, 13, 0,
		0, 77, 79, 5, 7, 0, 0, 78, 72, 1, 0, 0, 0, 78, 75, 1, 0, 0, 0, 79, 17,
		1, 0, 0, 0, 80, 81, 5, 6, 0, 0, 81, 82, 5, 13, 0, 0, 82, 87, 5, 6, 0, 0,
		83, 84, 5, 8, 0, 0, 84, 85, 5, 13, 0, 0, 85, 87, 5, 8, 0, 0, 86, 80, 1,
		0, 0, 0, 86, 83, 1, 0, 0, 0, 87, 19, 1, 0, 0, 0, 88, 89, 5, 9, 0, 0, 89,
		90, 5, 13, 0, 0, 90, 91, 5, 10, 0, 0, 91, 92, 5, 11, 0, 0, 92, 93, 5, 13,
		0, 0, 93, 94, 5, 12, 0, 0, 94, 21, 1, 0, 0, 0, 8, 25, 32, 46, 51, 64, 70,
		78, 86,
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

// MarkdownParserParserInit initializes any static state used to implement MarkdownParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMarkdownParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarkdownParserParserInit() {
	staticData := &MarkdownParserParserStaticData
	staticData.once.Do(markdownparserParserInit)
}

// NewMarkdownParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMarkdownParserParser(input antlr.TokenStream) *MarkdownParserParser {
	MarkdownParserParserInit()
	this := new(MarkdownParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MarkdownParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MarkdownParser.g4"

	return this
}

// MarkdownParserParser tokens.
const (
	MarkdownParserParserEOF               = antlr.TokenEOF
	MarkdownParserParserH1                = 1
	MarkdownParserParserH2                = 2
	MarkdownParserParserH3                = 3
	MarkdownParserParserLIST_MARKER       = 4
	MarkdownParserParserDOUBLE_ASTERISK   = 5
	MarkdownParserParserSINGLE_ASTERISK   = 6
	MarkdownParserParserDOUBLE_UNDERSCORE = 7
	MarkdownParserParserSINGLE_UNDERSCORE = 8
	MarkdownParserParserLBRACK            = 9
	MarkdownParserParserRBRACK            = 10
	MarkdownParserParserLPAREN            = 11
	MarkdownParserParserRPAREN            = 12
	MarkdownParserParserTEXT              = 13
	MarkdownParserParserWS                = 14
	MarkdownParserParserNEWLINE           = 15
)

// MarkdownParserParser rules.
const (
	MarkdownParserParserRULE_document      = 0
	MarkdownParserParserRULE_block         = 1
	MarkdownParserParserRULE_header        = 2
	MarkdownParserParserRULE_list          = 3
	MarkdownParserParserRULE_listItem      = 4
	MarkdownParserParserRULE_paragraph     = 5
	MarkdownParserParserRULE_inline        = 6
	MarkdownParserParserRULE_inlineElement = 7
	MarkdownParserParserRULE_bold          = 8
	MarkdownParserParserRULE_italic        = 9
	MarkdownParserParserRULE_link          = 10
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
	p.RuleIndex = MarkdownParserParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserEOF, 0)
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

func (p *MarkdownParserParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarkdownParserParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9214) != 0) {
		{
			p.SetState(22)
			p.Block()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(27)
		p.Match(MarkdownParserParserEOF)
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
	Header() IHeaderContext
	List() IListContext
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
	p.RuleIndex = MarkdownParserParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Header() IHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *BlockContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
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

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParserParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarkdownParserParserRULE_block)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserParserH1, MarkdownParserParserH2, MarkdownParserParserH3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Header()
		}

	case MarkdownParserParserLIST_MARKER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.List()
		}

	case MarkdownParserParserDOUBLE_ASTERISK, MarkdownParserParserSINGLE_ASTERISK, MarkdownParserParserDOUBLE_UNDERSCORE, MarkdownParserParserSINGLE_UNDERSCORE, MarkdownParserParserLBRACK, MarkdownParserParserTEXT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)
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

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	H1() antlr.TerminalNode
	Inline() IInlineContext
	NEWLINE() antlr.TerminalNode
	H2() antlr.TerminalNode
	H3() antlr.TerminalNode

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_header
	return p
}

func InitEmptyHeaderContext(p *HeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_header
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) H1() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserH1, 0)
}

func (s *HeaderContext) Inline() IInlineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineContext)
}

func (s *HeaderContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserNEWLINE, 0)
}

func (s *HeaderContext) H2() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserH2, 0)
}

func (s *HeaderContext) H3() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserH3, 0)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (s *HeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParserParser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MarkdownParserParserRULE_header)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserParserH1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(MarkdownParserParserH1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Inline()
		}
		{
			p.SetState(36)
			p.Match(MarkdownParserParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MarkdownParserParserH2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(MarkdownParserParserH2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Inline()
		}
		{
			p.SetState(40)
			p.Match(MarkdownParserParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MarkdownParserParserH3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Match(MarkdownParserParserH3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Inline()
		}
		{
			p.SetState(44)
			p.Match(MarkdownParserParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllListItem() []IListItemContext
	ListItem(i int) IListItemContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllListItem() []IListItemContext {
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

func (s *ListContext) ListItem(i int) IListItemContext {
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

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParserParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MarkdownParserParserRULE_list)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(48)
				p.ListItem()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	LIST_MARKER() antlr.TerminalNode
	WS() antlr.TerminalNode
	Inline() IInlineContext
	NEWLINE() antlr.TerminalNode

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
	p.RuleIndex = MarkdownParserParserRULE_listItem
	return p
}

func InitEmptyListItemContext(p *ListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_listItem
}

func (*ListItemContext) IsListItemContext() {}

func NewListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListItemContext {
	var p = new(ListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_listItem

	return p
}

func (s *ListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ListItemContext) LIST_MARKER() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserLIST_MARKER, 0)
}

func (s *ListItemContext) WS() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserWS, 0)
}

func (s *ListItemContext) Inline() IInlineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineContext)
}

func (s *ListItemContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserNEWLINE, 0)
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

func (p *MarkdownParserParser) ListItem() (localctx IListItemContext) {
	localctx = NewListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MarkdownParserParserRULE_listItem)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(MarkdownParserParserLIST_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(MarkdownParserParserWS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Inline()
	}
	{
		p.SetState(56)
		p.Match(MarkdownParserParserNEWLINE)
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

// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Inline() IInlineContext
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
	p.RuleIndex = MarkdownParserParserRULE_paragraph
	return p
}

func InitEmptyParagraphContext(p *ParagraphContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_paragraph
}

func (*ParagraphContext) IsParagraphContext() {}

func NewParagraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParagraphContext {
	var p = new(ParagraphContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_paragraph

	return p
}

func (s *ParagraphContext) GetParser() antlr.Parser { return s.parser }

func (s *ParagraphContext) Inline() IInlineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineContext)
}

func (s *ParagraphContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserNEWLINE, 0)
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

func (p *MarkdownParserParser) Paragraph() (localctx IParagraphContext) {
	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MarkdownParserParserRULE_paragraph)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Inline()
	}
	{
		p.SetState(59)
		p.Match(MarkdownParserParserNEWLINE)
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
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext

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
	p.RuleIndex = MarkdownParserParserRULE_inline
	return p
}

func InitEmptyInlineContext(p *InlineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_inline
}

func (*InlineContext) IsInlineContext() {}

func NewInlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineContext {
	var p = new(InlineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_inline

	return p
}

func (s *InlineContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineContext) AllInlineElement() []IInlineElementContext {
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

func (s *InlineContext) InlineElement(i int) IInlineElementContext {
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

func (s *InlineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitInline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParserParser) Inline() (localctx IInlineContext) {
	localctx = NewInlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MarkdownParserParserRULE_inline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9184) != 0) {
		{
			p.SetState(61)
			p.InlineElement()
		}

		p.SetState(64)
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

// IInlineElementContext is an interface to support dynamic dispatch.
type IInlineElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bold() IBoldContext
	Italic() IItalicContext
	Link() ILinkContext
	TEXT() antlr.TerminalNode

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
	p.RuleIndex = MarkdownParserParserRULE_inlineElement
	return p
}

func InitEmptyInlineElementContext(p *InlineElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_inlineElement
}

func (*InlineElementContext) IsInlineElementContext() {}

func NewInlineElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineElementContext {
	var p = new(InlineElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_inlineElement

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

func (s *InlineElementContext) Link() ILinkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinkContext)
}

func (s *InlineElementContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserTEXT, 0)
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

func (p *MarkdownParserParser) InlineElement() (localctx IInlineElementContext) {
	localctx = NewInlineElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MarkdownParserParserRULE_inlineElement)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserParserDOUBLE_ASTERISK, MarkdownParserParserDOUBLE_UNDERSCORE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Bold()
		}

	case MarkdownParserParserSINGLE_ASTERISK, MarkdownParserParserSINGLE_UNDERSCORE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Italic()
		}

	case MarkdownParserParserLBRACK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Link()
		}

	case MarkdownParserParserTEXT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.Match(MarkdownParserParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	// Getter signatures
	AllDOUBLE_ASTERISK() []antlr.TerminalNode
	DOUBLE_ASTERISK(i int) antlr.TerminalNode
	TEXT() antlr.TerminalNode
	AllDOUBLE_UNDERSCORE() []antlr.TerminalNode
	DOUBLE_UNDERSCORE(i int) antlr.TerminalNode

	// IsBoldContext differentiates from other interfaces.
	IsBoldContext()
}

type BoldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoldContext() *BoldContext {
	var p = new(BoldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_bold
	return p
}

func InitEmptyBoldContext(p *BoldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_bold
}

func (*BoldContext) IsBoldContext() {}

func NewBoldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldContext {
	var p = new(BoldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_bold

	return p
}

func (s *BoldContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldContext) AllDOUBLE_ASTERISK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserParserDOUBLE_ASTERISK)
}

func (s *BoldContext) DOUBLE_ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserDOUBLE_ASTERISK, i)
}

func (s *BoldContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserTEXT, 0)
}

func (s *BoldContext) AllDOUBLE_UNDERSCORE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserParserDOUBLE_UNDERSCORE)
}

func (s *BoldContext) DOUBLE_UNDERSCORE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserDOUBLE_UNDERSCORE, i)
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

func (p *MarkdownParserParser) Bold() (localctx IBoldContext) {
	localctx = NewBoldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MarkdownParserParserRULE_bold)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserParserDOUBLE_ASTERISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(MarkdownParserParserDOUBLE_ASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(MarkdownParserParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(MarkdownParserParserDOUBLE_ASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MarkdownParserParserDOUBLE_UNDERSCORE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(MarkdownParserParserDOUBLE_UNDERSCORE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(MarkdownParserParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(MarkdownParserParserDOUBLE_UNDERSCORE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	// Getter signatures
	AllSINGLE_ASTERISK() []antlr.TerminalNode
	SINGLE_ASTERISK(i int) antlr.TerminalNode
	TEXT() antlr.TerminalNode
	AllSINGLE_UNDERSCORE() []antlr.TerminalNode
	SINGLE_UNDERSCORE(i int) antlr.TerminalNode

	// IsItalicContext differentiates from other interfaces.
	IsItalicContext()
}

type ItalicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItalicContext() *ItalicContext {
	var p = new(ItalicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_italic
	return p
}

func InitEmptyItalicContext(p *ItalicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_italic
}

func (*ItalicContext) IsItalicContext() {}

func NewItalicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItalicContext {
	var p = new(ItalicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_italic

	return p
}

func (s *ItalicContext) GetParser() antlr.Parser { return s.parser }

func (s *ItalicContext) AllSINGLE_ASTERISK() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserParserSINGLE_ASTERISK)
}

func (s *ItalicContext) SINGLE_ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserSINGLE_ASTERISK, i)
}

func (s *ItalicContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserTEXT, 0)
}

func (s *ItalicContext) AllSINGLE_UNDERSCORE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserParserSINGLE_UNDERSCORE)
}

func (s *ItalicContext) SINGLE_UNDERSCORE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserSINGLE_UNDERSCORE, i)
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

func (p *MarkdownParserParser) Italic() (localctx IItalicContext) {
	localctx = NewItalicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MarkdownParserParserRULE_italic)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserParserSINGLE_ASTERISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(MarkdownParserParserSINGLE_ASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(MarkdownParserParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(MarkdownParserParserSINGLE_ASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MarkdownParserParserSINGLE_UNDERSCORE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(MarkdownParserParserSINGLE_UNDERSCORE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(MarkdownParserParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(MarkdownParserParserSINGLE_UNDERSCORE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ILinkContext is an interface to support dynamic dispatch.
type ILinkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsLinkContext differentiates from other interfaces.
	IsLinkContext()
}

type LinkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinkContext() *LinkContext {
	var p = new(LinkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_link
	return p
}

func InitEmptyLinkContext(p *LinkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserParserRULE_link
}

func (*LinkContext) IsLinkContext() {}

func NewLinkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinkContext {
	var p = new(LinkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserParserRULE_link

	return p
}

func (s *LinkContext) GetParser() antlr.Parser { return s.parser }

func (s *LinkContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserLBRACK, 0)
}

func (s *LinkContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserParserTEXT)
}

func (s *LinkContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserTEXT, i)
}

func (s *LinkContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserRBRACK, 0)
}

func (s *LinkContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserLPAREN, 0)
}

func (s *LinkContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MarkdownParserParserRPAREN, 0)
}

func (s *LinkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterLink(s)
	}
}

func (s *LinkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitLink(s)
	}
}

func (s *LinkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MarkdownParserVisitor:
		return t.VisitLink(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MarkdownParserParser) Link() (localctx ILinkContext) {
	localctx = NewLinkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MarkdownParserParserRULE_link)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(MarkdownParserParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(MarkdownParserParserTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(MarkdownParserParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(MarkdownParserParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(MarkdownParserParserTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(MarkdownParserParserRPAREN)
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
