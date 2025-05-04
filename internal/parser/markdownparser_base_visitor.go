// Code generated from ./internal/parser/MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import "github.com/antlr4-go/antlr/v4"

type BaseMarkdownParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMarkdownParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitHeader(ctx *HeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitListItem(ctx *ListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitParagraph(ctx *ParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitInline(ctx *InlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitInlineElement(ctx *InlineElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitBold(ctx *BoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitItalic(ctx *ItalicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitLink(ctx *LinkContext) interface{} {
	return v.VisitChildren(ctx)
}
