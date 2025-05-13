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

// EnterParagraph is called when production paragraph is entered.
func (s *BaseMarkdownParserListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseMarkdownParserListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterLine is called when production line is entered.
func (s *BaseMarkdownParserListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseMarkdownParserListener) ExitLine(ctx *LineContext) {}

// EnterEscape_char is called when production escape_char is entered.
func (s *BaseMarkdownParserListener) EnterEscape_char(ctx *Escape_charContext) {}

// ExitEscape_char is called when production escape_char is exited.
func (s *BaseMarkdownParserListener) ExitEscape_char(ctx *Escape_charContext) {}

// EnterLinebreak is called when production linebreak is entered.
func (s *BaseMarkdownParserListener) EnterLinebreak(ctx *LinebreakContext) {}

// ExitLinebreak is called when production linebreak is exited.
func (s *BaseMarkdownParserListener) ExitLinebreak(ctx *LinebreakContext) {}

// EnterInline is called when production inline is entered.
func (s *BaseMarkdownParserListener) EnterInline(ctx *InlineContext) {}

// ExitInline is called when production inline is exited.
func (s *BaseMarkdownParserListener) ExitInline(ctx *InlineContext) {}

// EnterItalic is called when production italic is entered.
func (s *BaseMarkdownParserListener) EnterItalic(ctx *ItalicContext) {}

// ExitItalic is called when production italic is exited.
func (s *BaseMarkdownParserListener) ExitItalic(ctx *ItalicContext) {}

// EnterItalic_line is called when production italic_line is entered.
func (s *BaseMarkdownParserListener) EnterItalic_line(ctx *Italic_lineContext) {}

// ExitItalic_line is called when production italic_line is exited.
func (s *BaseMarkdownParserListener) ExitItalic_line(ctx *Italic_lineContext) {}

// EnterBold is called when production bold is entered.
func (s *BaseMarkdownParserListener) EnterBold(ctx *BoldContext) {}

// ExitBold is called when production bold is exited.
func (s *BaseMarkdownParserListener) ExitBold(ctx *BoldContext) {}

// EnterBold_line is called when production bold_line is entered.
func (s *BaseMarkdownParserListener) EnterBold_line(ctx *Bold_lineContext) {}

// ExitBold_line is called when production bold_line is exited.
func (s *BaseMarkdownParserListener) ExitBold_line(ctx *Bold_lineContext) {}

// EnterBold_and_italic is called when production bold_and_italic is entered.
func (s *BaseMarkdownParserListener) EnterBold_and_italic(ctx *Bold_and_italicContext) {}

// ExitBold_and_italic is called when production bold_and_italic is exited.
func (s *BaseMarkdownParserListener) ExitBold_and_italic(ctx *Bold_and_italicContext) {}

// EnterCode is called when production code is entered.
func (s *BaseMarkdownParserListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseMarkdownParserListener) ExitCode(ctx *CodeContext) {}

// EnterCode_text is called when production code_text is entered.
func (s *BaseMarkdownParserListener) EnterCode_text(ctx *Code_textContext) {}

// ExitCode_text is called when production code_text is exited.
func (s *BaseMarkdownParserListener) ExitCode_text(ctx *Code_textContext) {}

// EnterH1 is called when production h1 is entered.
func (s *BaseMarkdownParserListener) EnterH1(ctx *H1Context) {}

// ExitH1 is called when production h1 is exited.
func (s *BaseMarkdownParserListener) ExitH1(ctx *H1Context) {}

// EnterH2 is called when production h2 is entered.
func (s *BaseMarkdownParserListener) EnterH2(ctx *H2Context) {}

// ExitH2 is called when production h2 is exited.
func (s *BaseMarkdownParserListener) ExitH2(ctx *H2Context) {}

// EnterH3 is called when production h3 is entered.
func (s *BaseMarkdownParserListener) EnterH3(ctx *H3Context) {}

// ExitH3 is called when production h3 is exited.
func (s *BaseMarkdownParserListener) ExitH3(ctx *H3Context) {}

// EnterH4 is called when production h4 is entered.
func (s *BaseMarkdownParserListener) EnterH4(ctx *H4Context) {}

// ExitH4 is called when production h4 is exited.
func (s *BaseMarkdownParserListener) ExitH4(ctx *H4Context) {}

// EnterH5 is called when production h5 is entered.
func (s *BaseMarkdownParserListener) EnterH5(ctx *H5Context) {}

// ExitH5 is called when production h5 is exited.
func (s *BaseMarkdownParserListener) ExitH5(ctx *H5Context) {}

// EnterH6 is called when production h6 is entered.
func (s *BaseMarkdownParserListener) EnterH6(ctx *H6Context) {}

// ExitH6 is called when production h6 is exited.
func (s *BaseMarkdownParserListener) ExitH6(ctx *H6Context) {}

// EnterHeading is called when production heading is entered.
func (s *BaseMarkdownParserListener) EnterHeading(ctx *HeadingContext) {}

// ExitHeading is called when production heading is exited.
func (s *BaseMarkdownParserListener) ExitHeading(ctx *HeadingContext) {}

// EnterUnorderedList is called when production unorderedList is entered.
func (s *BaseMarkdownParserListener) EnterUnorderedList(ctx *UnorderedListContext) {}

// ExitUnorderedList is called when production unorderedList is exited.
func (s *BaseMarkdownParserListener) ExitUnorderedList(ctx *UnorderedListContext) {}

// EnterDashList is called when production dashList is entered.
func (s *BaseMarkdownParserListener) EnterDashList(ctx *DashListContext) {}

// ExitDashList is called when production dashList is exited.
func (s *BaseMarkdownParserListener) ExitDashList(ctx *DashListContext) {}

// EnterAsteriskList is called when production asteriskList is entered.
func (s *BaseMarkdownParserListener) EnterAsteriskList(ctx *AsteriskListContext) {}

// ExitAsteriskList is called when production asteriskList is exited.
func (s *BaseMarkdownParserListener) ExitAsteriskList(ctx *AsteriskListContext) {}

// EnterPlusList is called when production plusList is entered.
func (s *BaseMarkdownParserListener) EnterPlusList(ctx *PlusListContext) {}

// ExitPlusList is called when production plusList is exited.
func (s *BaseMarkdownParserListener) ExitPlusList(ctx *PlusListContext) {}

// EnterDashListItem is called when production dashListItem is entered.
func (s *BaseMarkdownParserListener) EnterDashListItem(ctx *DashListItemContext) {}

// ExitDashListItem is called when production dashListItem is exited.
func (s *BaseMarkdownParserListener) ExitDashListItem(ctx *DashListItemContext) {}

// EnterAsteriskListItem is called when production asteriskListItem is entered.
func (s *BaseMarkdownParserListener) EnterAsteriskListItem(ctx *AsteriskListItemContext) {}

// ExitAsteriskListItem is called when production asteriskListItem is exited.
func (s *BaseMarkdownParserListener) ExitAsteriskListItem(ctx *AsteriskListItemContext) {}

// EnterPlusListItem is called when production plusListItem is entered.
func (s *BaseMarkdownParserListener) EnterPlusListItem(ctx *PlusListItemContext) {}

// ExitPlusListItem is called when production plusListItem is exited.
func (s *BaseMarkdownParserListener) ExitPlusListItem(ctx *PlusListItemContext) {}

// EnterContinuationLine is called when production continuationLine is entered.
func (s *BaseMarkdownParserListener) EnterContinuationLine(ctx *ContinuationLineContext) {}

// ExitContinuationLine is called when production continuationLine is exited.
func (s *BaseMarkdownParserListener) ExitContinuationLine(ctx *ContinuationLineContext) {}
