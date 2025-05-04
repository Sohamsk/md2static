// Code generated from ./internal/parser/MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

// MarkdownParserListener is a complete listener for a parse tree produced by MarkdownParserParser.
type MarkdownParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterListItem is called when entering the listItem production.
	EnterListItem(c *ListItemContext)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterInline is called when entering the inline production.
	EnterInline(c *InlineContext)

	// EnterInlineElement is called when entering the inlineElement production.
	EnterInlineElement(c *InlineElementContext)

	// EnterBold is called when entering the bold production.
	EnterBold(c *BoldContext)

	// EnterItalic is called when entering the italic production.
	EnterItalic(c *ItalicContext)

	// EnterLink is called when entering the link production.
	EnterLink(c *LinkContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitListItem is called when exiting the listItem production.
	ExitListItem(c *ListItemContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitInline is called when exiting the inline production.
	ExitInline(c *InlineContext)

	// ExitInlineElement is called when exiting the inlineElement production.
	ExitInlineElement(c *InlineElementContext)

	// ExitBold is called when exiting the bold production.
	ExitBold(c *BoldContext)

	// ExitItalic is called when exiting the italic production.
	ExitItalic(c *ItalicContext)

	// ExitLink is called when exiting the link production.
	ExitLink(c *LinkContext)
}
