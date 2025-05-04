#!/bin/sh
PARSER_DIR="$(pwd)"
alias antlr4='java -Xmx500M -cp "$PARSER_DIR/../../external/antlr.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -package parser *.g4
