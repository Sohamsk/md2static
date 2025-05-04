// Code generated from ./internal/parser/MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MarkdownParserParser.
type MarkdownParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MarkdownParserParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#header.
	VisitHeader(ctx *HeaderContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#listItem.
	VisitListItem(ctx *ListItemContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#paragraph.
	VisitParagraph(ctx *ParagraphContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#inline.
	VisitInline(ctx *InlineContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#inlineElement.
	VisitInlineElement(ctx *InlineElementContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#bold.
	VisitBold(ctx *BoldContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#italic.
	VisitItalic(ctx *ItalicContext) interface{}

	// Visit a parse tree produced by MarkdownParserParser#link.
	VisitLink(ctx *LinkContext) interface{}
}
