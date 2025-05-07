package listener

import (
	"github.com/sohamsk/md2static/internal/parser"
)

// EnterH1 is called when production h1 is entered.
func (s *MarkdownListener) EnterH1(ctx *parser.H1Context) {
	s.writer.WriteString("<h1>")
}

// ExitH1 is called when production h1 is exited.
func (s *MarkdownListener) ExitH1(ctx *parser.H1Context) {
	s.writer.WriteString("</h1>\n")
}

// EnterH2 is called when production h2 is entered.
func (s *MarkdownListener) EnterH2(ctx *parser.H2Context) {
	s.writer.WriteString("<h2>")
}

// ExitH2 is called when production h2 is exited.
func (s *MarkdownListener) ExitH2(ctx *parser.H2Context) {
	s.writer.WriteString("</h2>\n")
}

// EnterH3 is called when production h3 is entered.
func (s *MarkdownListener) EnterH3(ctx *parser.H3Context) {
	s.writer.WriteString("<h3>")
}

// ExitH3 is called when production h3 is exited.
func (s *MarkdownListener) ExitH3(ctx *parser.H3Context) {
	s.writer.WriteString("</h3>\n")
}

// EnterH4 is called when production h4 is entered.
func (s *MarkdownListener) EnterH4(ctx *parser.H4Context) {
	s.writer.WriteString("<h4>")
}

// ExitH4 is called when production h4 is exited.
func (s *MarkdownListener) ExitH4(ctx *parser.H4Context) {
	s.writer.WriteString("</h4>\n")
}

// EnterH5 is called when production h5 is entered.
func (s *MarkdownListener) EnterH5(ctx *parser.H5Context) {
	s.writer.WriteString("<h5>")
}

// ExitH5 is called when production h5 is exited.
func (s *MarkdownListener) ExitH5(ctx *parser.H5Context) {
	s.writer.WriteString("</h5>\n")
}

// EnterH6 is called when production h6 is entered.
func (s *MarkdownListener) EnterH6(ctx *parser.H6Context) {
	s.writer.WriteString("<h6>")
}

// ExitH6 is called when production h6 is exited.
func (s *MarkdownListener) ExitH6(ctx *parser.H6Context) {
	s.writer.WriteString("</h6>\n")
}
