// Code generated from MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MarkdownParser.
type MarkdownParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MarkdownParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by MarkdownParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MarkdownParser#heading.
	VisitHeading(ctx *HeadingContext) interface{}

	// Visit a parse tree produced by MarkdownParser#paragraph.
	VisitParagraph(ctx *ParagraphContext) interface{}

	// Visit a parse tree produced by MarkdownParser#blockquote.
	VisitBlockquote(ctx *BlockquoteContext) interface{}

	// Visit a parse tree produced by MarkdownParser#unorderedList.
	VisitUnorderedList(ctx *UnorderedListContext) interface{}

	// Visit a parse tree produced by MarkdownParser#orderedList.
	VisitOrderedList(ctx *OrderedListContext) interface{}

	// Visit a parse tree produced by MarkdownParser#listItem.
	VisitListItem(ctx *ListItemContext) interface{}

	// Visit a parse tree produced by MarkdownParser#orderedListItem.
	VisitOrderedListItem(ctx *OrderedListItemContext) interface{}

	// Visit a parse tree produced by MarkdownParser#codeBlock.
	VisitCodeBlock(ctx *CodeBlockContext) interface{}

	// Visit a parse tree produced by MarkdownParser#horizontalRule.
	VisitHorizontalRule(ctx *HorizontalRuleContext) interface{}

	// Visit a parse tree produced by MarkdownParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by MarkdownParser#inlineElement.
	VisitInlineElement(ctx *InlineElementContext) interface{}

	// Visit a parse tree produced by MarkdownParser#bold.
	VisitBold(ctx *BoldContext) interface{}

	// Visit a parse tree produced by MarkdownParser#boldAndItalic.
	VisitBoldAndItalic(ctx *BoldAndItalicContext) interface{}

	// Visit a parse tree produced by MarkdownParser#boldText.
	VisitBoldText(ctx *BoldTextContext) interface{}

	// Visit a parse tree produced by MarkdownParser#italic.
	VisitItalic(ctx *ItalicContext) interface{}

	// Visit a parse tree produced by MarkdownParser#italicText.
	VisitItalicText(ctx *ItalicTextContext) interface{}

	// Visit a parse tree produced by MarkdownParser#strikethrough.
	VisitStrikethrough(ctx *StrikethroughContext) interface{}

	// Visit a parse tree produced by MarkdownParser#strikethroughText.
	VisitStrikethroughText(ctx *StrikethroughTextContext) interface{}

	// Visit a parse tree produced by MarkdownParser#code.
	VisitCode(ctx *CodeContext) interface{}

	// Visit a parse tree produced by MarkdownParser#codeText.
	VisitCodeText(ctx *CodeTextContext) interface{}
}
