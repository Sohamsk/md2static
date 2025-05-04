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

	// EnterHeading is called when entering the heading production.
	EnterHeading(c *HeadingContext)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterBlockquote is called when entering the blockquote production.
	EnterBlockquote(c *BlockquoteContext)

	// EnterUnorderedList is called when entering the unorderedList production.
	EnterUnorderedList(c *UnorderedListContext)

	// EnterOrderedList is called when entering the orderedList production.
	EnterOrderedList(c *OrderedListContext)

	// EnterListItem is called when entering the listItem production.
	EnterListItem(c *ListItemContext)

	// EnterOrderedListItem is called when entering the orderedListItem production.
	EnterOrderedListItem(c *OrderedListItemContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterHorizontalRule is called when entering the horizontalRule production.
	EnterHorizontalRule(c *HorizontalRuleContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterInlineElement is called when entering the inlineElement production.
	EnterInlineElement(c *InlineElementContext)

	// EnterBold is called when entering the bold production.
	EnterBold(c *BoldContext)

	// EnterBoldAndItalic is called when entering the boldAndItalic production.
	EnterBoldAndItalic(c *BoldAndItalicContext)

	// EnterBoldText is called when entering the boldText production.
	EnterBoldText(c *BoldTextContext)

	// EnterItalic is called when entering the italic production.
	EnterItalic(c *ItalicContext)

	// EnterItalicText is called when entering the italicText production.
	EnterItalicText(c *ItalicTextContext)

	// EnterStrikethrough is called when entering the strikethrough production.
	EnterStrikethrough(c *StrikethroughContext)

	// EnterStrikethroughText is called when entering the strikethroughText production.
	EnterStrikethroughText(c *StrikethroughTextContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterCodeText is called when entering the codeText production.
	EnterCodeText(c *CodeTextContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitHeading is called when exiting the heading production.
	ExitHeading(c *HeadingContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitBlockquote is called when exiting the blockquote production.
	ExitBlockquote(c *BlockquoteContext)

	// ExitUnorderedList is called when exiting the unorderedList production.
	ExitUnorderedList(c *UnorderedListContext)

	// ExitOrderedList is called when exiting the orderedList production.
	ExitOrderedList(c *OrderedListContext)

	// ExitListItem is called when exiting the listItem production.
	ExitListItem(c *ListItemContext)

	// ExitOrderedListItem is called when exiting the orderedListItem production.
	ExitOrderedListItem(c *OrderedListItemContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitHorizontalRule is called when exiting the horizontalRule production.
	ExitHorizontalRule(c *HorizontalRuleContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitInlineElement is called when exiting the inlineElement production.
	ExitInlineElement(c *InlineElementContext)

	// ExitBold is called when exiting the bold production.
	ExitBold(c *BoldContext)

	// ExitBoldAndItalic is called when exiting the boldAndItalic production.
	ExitBoldAndItalic(c *BoldAndItalicContext)

	// ExitBoldText is called when exiting the boldText production.
	ExitBoldText(c *BoldTextContext)

	// ExitItalic is called when exiting the italic production.
	ExitItalic(c *ItalicContext)

	// ExitItalicText is called when exiting the italicText production.
	ExitItalicText(c *ItalicTextContext)

	// ExitStrikethrough is called when exiting the strikethrough production.
	ExitStrikethrough(c *StrikethroughContext)

	// ExitStrikethroughText is called when exiting the strikethroughText production.
	ExitStrikethroughText(c *StrikethroughTextContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitCodeText is called when exiting the codeText production.
	ExitCodeText(c *CodeTextContext)
}
