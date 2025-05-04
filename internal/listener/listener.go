package listener

import (
	"bufio"

	"github.com/sohamsk/md2static/internal/parser"
)

type MarkdownListener struct {
	*parser.BaseMarkdownParserListener
	writer *bufio.Writer
}

func NewMarkdownListener(writer *bufio.Writer) *MarkdownListener {
	l := new(MarkdownListener)
	l.writer = writer
	return l
}

// EnterDocument is called when production document is entered.
func (s *MarkdownListener) EnterDocument(ctx *parser.DocumentContext) {
	s.writer.WriteString("<html>")
	s.writer.WriteString("<body>")
	// can give extra context from say a configuration tag
	// we can give things such as title
}

// ExitDocument is called when production document is exited.
func (s *MarkdownListener) ExitDocument(ctx *parser.DocumentContext) {
	s.writer.WriteString("</body>")
	s.writer.WriteString("</html>")
}
