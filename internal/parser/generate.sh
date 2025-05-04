#!/bin/sh
alias antlr4='java -Xmx500M -cp "../../external/antlr.jar:$CLASSPATH" org.antlr.v4.Tool' || echo "Command 1 failed"
antlr4 -Dlanguage=Go -visitor -package parser *.g4 || echo "Command 2 failed"
