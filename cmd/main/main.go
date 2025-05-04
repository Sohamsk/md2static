package main

import (
	"bufio"
	"bytes"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/sohamsk/md2static/internal/listener"
	"github.com/sohamsk/md2static/internal/parser"
)

func main() {
	var buffer bytes.Buffer

	inputfileName := os.Args[1]

	input, err := antlr.NewFileStream(inputfileName)
	if err != nil {
		log.Panic("File error")
	}

	lexer := parser.NewMarkdownLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMarkdownParser(stream)
	p.BuildParseTrees = true
	tree := p.Document()

	writer := bufio.NewWriter(&buffer)
	md := listener.NewMarkdownListener(writer)
	antlr.ParseTreeWalkerDefault.Walk(md, tree)
	writer.Flush()

	err = os.Mkdir("output", 0755)
	if err != nil {
		log.Panic(err)
	}
	f, err := os.Create("./output/index.html")
	if err != nil {
		log.Panic(err)
	}
	f.Write(buffer.Bytes())
	f.Sync()
}
