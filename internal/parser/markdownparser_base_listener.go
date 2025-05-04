// Code generated from ./internal/parser/MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

// BaseMarkdownParserListener is a complete listener for a parse tree produced by MarkdownParserParser.
type BaseMarkdownParserListener struct{}

var _ MarkdownParserListener = &BaseMarkdownParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarkdownParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarkdownParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarkdownParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarkdownParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseMarkdownParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseMarkdownParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMarkdownParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMarkdownParserListener) ExitBlock(ctx *BlockContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseMarkdownParserListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseMarkdownParserListener) ExitHeader(ctx *HeaderContext) {}

// EnterList is called when production list is entered.
func (s *BaseMarkdownParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseMarkdownParserListener) ExitList(ctx *ListContext) {}

// EnterListItem is called when production listItem is entered.
func (s *BaseMarkdownParserListener) EnterListItem(ctx *ListItemContext) {}

// ExitListItem is called when production listItem is exited.
func (s *BaseMarkdownParserListener) ExitListItem(ctx *ListItemContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseMarkdownParserListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseMarkdownParserListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterInline is called when production inline is entered.
func (s *BaseMarkdownParserListener) EnterInline(ctx *InlineContext) {}

// ExitInline is called when production inline is exited.
func (s *BaseMarkdownParserListener) ExitInline(ctx *InlineContext) {}

// EnterInlineElement is called when production inlineElement is entered.
func (s *BaseMarkdownParserListener) EnterInlineElement(ctx *InlineElementContext) {}

// ExitInlineElement is called when production inlineElement is exited.
func (s *BaseMarkdownParserListener) ExitInlineElement(ctx *InlineElementContext) {}

// EnterBold is called when production bold is entered.
func (s *BaseMarkdownParserListener) EnterBold(ctx *BoldContext) {}

// ExitBold is called when production bold is exited.
func (s *BaseMarkdownParserListener) ExitBold(ctx *BoldContext) {}

// EnterItalic is called when production italic is entered.
func (s *BaseMarkdownParserListener) EnterItalic(ctx *ItalicContext) {}

// ExitItalic is called when production italic is exited.
func (s *BaseMarkdownParserListener) ExitItalic(ctx *ItalicContext) {}

// EnterLink is called when production link is entered.
func (s *BaseMarkdownParserListener) EnterLink(ctx *LinkContext) {}

// ExitLink is called when production link is exited.
func (s *BaseMarkdownParserListener) ExitLink(ctx *LinkContext) {}
