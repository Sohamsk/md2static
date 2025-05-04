// Code generated from MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

func (v *BaseMarkdownParserVisitor) VisitHeading(ctx *HeadingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitParagraph(ctx *ParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitBlockquote(ctx *BlockquoteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitUnorderedList(ctx *UnorderedListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitOrderedList(ctx *OrderedListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitListItem(ctx *ListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitOrderedListItem(ctx *OrderedListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitCodeBlock(ctx *CodeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitHorizontalRule(ctx *HorizontalRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitInlineElement(ctx *InlineElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitBold(ctx *BoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitBoldAndItalic(ctx *BoldAndItalicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitBoldText(ctx *BoldTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitItalic(ctx *ItalicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitItalicText(ctx *ItalicTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitStrikethrough(ctx *StrikethroughContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitStrikethroughText(ctx *StrikethroughTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitCode(ctx *CodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMarkdownParserVisitor) VisitCodeText(ctx *CodeTextContext) interface{} {
	return v.VisitChildren(ctx)
}
