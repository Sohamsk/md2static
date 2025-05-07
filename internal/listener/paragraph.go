package listener

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/sohamsk/md2static/internal/parser"
)

// EnterParagraph is called when production paragraph is entered.
func (s *MarkdownListener) EnterParagraph(ctx *parser.ParagraphContext) {
	var content string
	for _, child := range ctx.GetChildren() {
		for _, grandChild := range child.GetChildren() {
			if inline, ok := grandChild.(*parser.InlineContext); ok {
				for _, inlineItem := range inline.GetChildren() {
					switch v := inlineItem.(type) {
					case *parser.BoldContext:
						content += "<b>" + v.GetData().GetText() + "</b>"
					case *parser.ItalicContext:
						content += "<i>" + v.GetData().GetText() + "</i>"
					case *parser.Bold_and_italicContext:
						content += "<b><i>" + v.GetData().GetText() + "</i></b>"
					case *parser.CodeContext:
						content += "<code>" + v.GetData().GetText() + "</code>"
					}
				}
			} else if _, ok := grandChild.(*parser.LinebreakContext); ok {
				content += "<br>"
			} else if e, ok := grandChild.(*parser.Escape_charContext); ok {
				content += strings.TrimLeft(e.GetText(), "\\")
			} else {
				tn, _ := grandChild.(antlr.TerminalNode)
				content += tn.GetText()
			}
		}
	}
	s.writer.WriteString("<p>" + content)
}

// ExitParagraph is called when production paragraph is exited.
func (s *MarkdownListener) ExitParagraph(ctx *parser.ParagraphContext) {
	s.writer.WriteString("</p>")
}
