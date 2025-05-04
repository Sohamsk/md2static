// Code generated from MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

// BaseMarkdownParserListener is a complete listener for a parse tree produced by MarkdownParser.
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

// EnterHeading is called when production heading is entered.
func (s *BaseMarkdownParserListener) EnterHeading(ctx *HeadingContext) {}

// ExitHeading is called when production heading is exited.
func (s *BaseMarkdownParserListener) ExitHeading(ctx *HeadingContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseMarkdownParserListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseMarkdownParserListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterBlockquote is called when production blockquote is entered.
func (s *BaseMarkdownParserListener) EnterBlockquote(ctx *BlockquoteContext) {}

// ExitBlockquote is called when production blockquote is exited.
func (s *BaseMarkdownParserListener) ExitBlockquote(ctx *BlockquoteContext) {}

// EnterUnorderedList is called when production unorderedList is entered.
func (s *BaseMarkdownParserListener) EnterUnorderedList(ctx *UnorderedListContext) {}

// ExitUnorderedList is called when production unorderedList is exited.
func (s *BaseMarkdownParserListener) ExitUnorderedList(ctx *UnorderedListContext) {}

// EnterOrderedList is called when production orderedList is entered.
func (s *BaseMarkdownParserListener) EnterOrderedList(ctx *OrderedListContext) {}

// ExitOrderedList is called when production orderedList is exited.
func (s *BaseMarkdownParserListener) ExitOrderedList(ctx *OrderedListContext) {}

// EnterListItem is called when production listItem is entered.
func (s *BaseMarkdownParserListener) EnterListItem(ctx *ListItemContext) {}

// ExitListItem is called when production listItem is exited.
func (s *BaseMarkdownParserListener) ExitListItem(ctx *ListItemContext) {}

// EnterOrderedListItem is called when production orderedListItem is entered.
func (s *BaseMarkdownParserListener) EnterOrderedListItem(ctx *OrderedListItemContext) {}

// ExitOrderedListItem is called when production orderedListItem is exited.
func (s *BaseMarkdownParserListener) ExitOrderedListItem(ctx *OrderedListItemContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseMarkdownParserListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseMarkdownParserListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterHorizontalRule is called when production horizontalRule is entered.
func (s *BaseMarkdownParserListener) EnterHorizontalRule(ctx *HorizontalRuleContext) {}

// ExitHorizontalRule is called when production horizontalRule is exited.
func (s *BaseMarkdownParserListener) ExitHorizontalRule(ctx *HorizontalRuleContext) {}

// EnterLine is called when production line is entered.
func (s *BaseMarkdownParserListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseMarkdownParserListener) ExitLine(ctx *LineContext) {}

// EnterInlineElement is called when production inlineElement is entered.
func (s *BaseMarkdownParserListener) EnterInlineElement(ctx *InlineElementContext) {}

// ExitInlineElement is called when production inlineElement is exited.
func (s *BaseMarkdownParserListener) ExitInlineElement(ctx *InlineElementContext) {}

// EnterBold is called when production bold is entered.
func (s *BaseMarkdownParserListener) EnterBold(ctx *BoldContext) {}

// ExitBold is called when production bold is exited.
func (s *BaseMarkdownParserListener) ExitBold(ctx *BoldContext) {}

// EnterBoldAndItalic is called when production boldAndItalic is entered.
func (s *BaseMarkdownParserListener) EnterBoldAndItalic(ctx *BoldAndItalicContext) {}

// ExitBoldAndItalic is called when production boldAndItalic is exited.
func (s *BaseMarkdownParserListener) ExitBoldAndItalic(ctx *BoldAndItalicContext) {}

// EnterBoldText is called when production boldText is entered.
func (s *BaseMarkdownParserListener) EnterBoldText(ctx *BoldTextContext) {}

// ExitBoldText is called when production boldText is exited.
func (s *BaseMarkdownParserListener) ExitBoldText(ctx *BoldTextContext) {}

// EnterItalic is called when production italic is entered.
func (s *BaseMarkdownParserListener) EnterItalic(ctx *ItalicContext) {}

// ExitItalic is called when production italic is exited.
func (s *BaseMarkdownParserListener) ExitItalic(ctx *ItalicContext) {}

// EnterItalicText is called when production italicText is entered.
func (s *BaseMarkdownParserListener) EnterItalicText(ctx *ItalicTextContext) {}

// ExitItalicText is called when production italicText is exited.
func (s *BaseMarkdownParserListener) ExitItalicText(ctx *ItalicTextContext) {}

// EnterStrikethrough is called when production strikethrough is entered.
func (s *BaseMarkdownParserListener) EnterStrikethrough(ctx *StrikethroughContext) {}

// ExitStrikethrough is called when production strikethrough is exited.
func (s *BaseMarkdownParserListener) ExitStrikethrough(ctx *StrikethroughContext) {}

// EnterStrikethroughText is called when production strikethroughText is entered.
func (s *BaseMarkdownParserListener) EnterStrikethroughText(ctx *StrikethroughTextContext) {}

// ExitStrikethroughText is called when production strikethroughText is exited.
func (s *BaseMarkdownParserListener) ExitStrikethroughText(ctx *StrikethroughTextContext) {}

// EnterCode is called when production code is entered.
func (s *BaseMarkdownParserListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseMarkdownParserListener) ExitCode(ctx *CodeContext) {}

// EnterCodeText is called when production codeText is entered.
func (s *BaseMarkdownParserListener) EnterCodeText(ctx *CodeTextContext) {}

// ExitCodeText is called when production codeText is exited.
func (s *BaseMarkdownParserListener) ExitCodeText(ctx *CodeTextContext) {}
