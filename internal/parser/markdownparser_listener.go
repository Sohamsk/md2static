// Code generated from MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

// MarkdownParserListener is a complete listener for a parse tree produced by MarkdownParser.
type MarkdownParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterEscape_char is called when entering the escape_char production.
	EnterEscape_char(c *Escape_charContext)

	// EnterLinebreak is called when entering the linebreak production.
	EnterLinebreak(c *LinebreakContext)

	// EnterInline is called when entering the inline production.
	EnterInline(c *InlineContext)

	// EnterItalic is called when entering the italic production.
	EnterItalic(c *ItalicContext)

	// EnterItalic_line is called when entering the italic_line production.
	EnterItalic_line(c *Italic_lineContext)

	// EnterBold is called when entering the bold production.
	EnterBold(c *BoldContext)

	// EnterBold_line is called when entering the bold_line production.
	EnterBold_line(c *Bold_lineContext)

	// EnterBold_and_italic is called when entering the bold_and_italic production.
	EnterBold_and_italic(c *Bold_and_italicContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterCode_text is called when entering the code_text production.
	EnterCode_text(c *Code_textContext)

	// EnterH1 is called when entering the h1 production.
	EnterH1(c *H1Context)

	// EnterH2 is called when entering the h2 production.
	EnterH2(c *H2Context)

	// EnterH3 is called when entering the h3 production.
	EnterH3(c *H3Context)

	// EnterH4 is called when entering the h4 production.
	EnterH4(c *H4Context)

	// EnterH5 is called when entering the h5 production.
	EnterH5(c *H5Context)

	// EnterH6 is called when entering the h6 production.
	EnterH6(c *H6Context)

	// EnterHeading is called when entering the heading production.
	EnterHeading(c *HeadingContext)

	// EnterUnorderedList is called when entering the unorderedList production.
	EnterUnorderedList(c *UnorderedListContext)

	// EnterUnorderedListItem is called when entering the unorderedListItem production.
	EnterUnorderedListItem(c *UnorderedListItemContext)

	// EnterListMarker is called when entering the listMarker production.
	EnterListMarker(c *ListMarkerContext)

	// EnterContinuationLine is called when entering the continuationLine production.
	EnterContinuationLine(c *ContinuationLineContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitEscape_char is called when exiting the escape_char production.
	ExitEscape_char(c *Escape_charContext)

	// ExitLinebreak is called when exiting the linebreak production.
	ExitLinebreak(c *LinebreakContext)

	// ExitInline is called when exiting the inline production.
	ExitInline(c *InlineContext)

	// ExitItalic is called when exiting the italic production.
	ExitItalic(c *ItalicContext)

	// ExitItalic_line is called when exiting the italic_line production.
	ExitItalic_line(c *Italic_lineContext)

	// ExitBold is called when exiting the bold production.
	ExitBold(c *BoldContext)

	// ExitBold_line is called when exiting the bold_line production.
	ExitBold_line(c *Bold_lineContext)

	// ExitBold_and_italic is called when exiting the bold_and_italic production.
	ExitBold_and_italic(c *Bold_and_italicContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitCode_text is called when exiting the code_text production.
	ExitCode_text(c *Code_textContext)

	// ExitH1 is called when exiting the h1 production.
	ExitH1(c *H1Context)

	// ExitH2 is called when exiting the h2 production.
	ExitH2(c *H2Context)

	// ExitH3 is called when exiting the h3 production.
	ExitH3(c *H3Context)

	// ExitH4 is called when exiting the h4 production.
	ExitH4(c *H4Context)

	// ExitH5 is called when exiting the h5 production.
	ExitH5(c *H5Context)

	// ExitH6 is called when exiting the h6 production.
	ExitH6(c *H6Context)

	// ExitHeading is called when exiting the heading production.
	ExitHeading(c *HeadingContext)

	// ExitUnorderedList is called when exiting the unorderedList production.
	ExitUnorderedList(c *UnorderedListContext)

	// ExitUnorderedListItem is called when exiting the unorderedListItem production.
	ExitUnorderedListItem(c *UnorderedListItemContext)

	// ExitListMarker is called when exiting the listMarker production.
	ExitListMarker(c *ListMarkerContext)

	// ExitContinuationLine is called when exiting the continuationLine production.
	ExitContinuationLine(c *ContinuationLineContext)
}
