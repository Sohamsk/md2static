// Code generated from MarkdownParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MarkdownParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MarkdownParser struct {
	*antlr.BaseParser
}

var MarkdownParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func markdownparserParserInit() {
	staticData := &MarkdownParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "'~~'", "", "'`'", "", "", "", "", "", "", "'#'",
		"'['", "']'", "'('", "')'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NEWLINE", "ESCAPE_CHAR", "BOLD_MARKER", "ITALIC_MARKER",
		"STRIKETHROUGH_MARKER", "BOLD_AND_ITALIC_MARKER", "CODE_MARKER", "BLOCKQUOTE_MARKER",
		"UNORDERED_LIST_MARKER", "ORDERED_LIST_MARKER", "HORIZONTAL_RULE", "FENCED_CODE_BLOCK",
		"INDENTED_CODE_BLOCK", "HASH", "LBRACKET", "RBRACKET", "LPAREN", "RPAREN",
		"EXCLAMATION", "TEXT",
	}
	staticData.RuleNames = []string{
		"document", "block", "h1", "h2", "h3", "h4", "h5", "h6", "heading",
		"paragraph", "blockquote", "unorderedList", "orderedList", "listItem",
		"orderedListItem", "codeBlock", "horizontalRule", "line", "inlineElement",
		"bold", "boldAndItalic", "boldText", "italic", "italicText", "strikethrough",
		"strikethroughText", "code", "codeText",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 301, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0, 61, 9, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 72, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 5, 2, 78, 8, 2, 10, 2, 12, 2, 81, 9, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 88, 8, 3, 10, 3, 12, 3, 91, 9, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 5, 4, 99, 8, 4, 10, 4, 12, 4, 102, 9, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 111, 8, 5, 10, 5, 12, 5, 114, 9, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 124, 8, 6, 10, 6, 12,
		6, 127, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 138, 8, 7, 10, 7, 12, 7, 141, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 149, 8, 8, 1, 9, 1, 9, 1, 9, 4, 9, 154, 8, 9, 11, 9, 12, 9, 155,
		1, 9, 3, 9, 159, 8, 9, 1, 10, 1, 10, 1, 10, 4, 10, 164, 8, 10, 11, 10,
		12, 10, 165, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 172, 8, 10, 11, 10, 12,
		10, 173, 5, 10, 176, 8, 10, 10, 10, 12, 10, 179, 9, 10, 1, 10, 3, 10, 182,
		8, 10, 1, 11, 4, 11, 185, 8, 11, 11, 11, 12, 11, 186, 1, 12, 4, 12, 190,
		8, 12, 11, 12, 12, 12, 191, 1, 13, 1, 13, 1, 13, 4, 13, 197, 8, 13, 11,
		13, 12, 13, 198, 1, 13, 1, 13, 1, 13, 3, 13, 204, 8, 13, 1, 13, 1, 13,
		4, 13, 208, 8, 13, 11, 13, 12, 13, 209, 5, 13, 212, 8, 13, 10, 13, 12,
		13, 215, 9, 13, 1, 14, 1, 14, 1, 14, 4, 14, 220, 8, 14, 11, 14, 12, 14,
		221, 1, 14, 1, 14, 1, 14, 3, 14, 227, 8, 14, 1, 14, 1, 14, 4, 14, 231,
		8, 14, 11, 14, 12, 14, 232, 5, 14, 235, 8, 14, 10, 14, 12, 14, 238, 9,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 4, 17, 247, 8, 17,
		11, 17, 12, 17, 248, 1, 17, 3, 17, 252, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 259, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 4, 21, 270, 8, 21, 11, 21, 12, 21, 271, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 4, 23, 279, 8, 23, 11, 23, 12, 23, 280, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 4, 25, 288, 8, 25, 11, 25, 12, 25, 289, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 27, 4, 27, 297, 8, 27, 11, 27, 12, 27, 298,
		1, 27, 0, 0, 28, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 0, 5, 1, 0, 13, 14,
		4, 0, 1, 1, 5, 5, 8, 8, 21, 21, 3, 0, 1, 1, 8, 8, 21, 21, 4, 0, 1, 1, 4,
		5, 8, 8, 21, 21, 3, 0, 1, 1, 3, 3, 21, 21, 326, 0, 59, 1, 0, 0, 0, 2, 71,
		1, 0, 0, 0, 4, 73, 1, 0, 0, 0, 6, 82, 1, 0, 0, 0, 8, 92, 1, 0, 0, 0, 10,
		103, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 128, 1, 0, 0, 0, 16, 148, 1,
		0, 0, 0, 18, 153, 1, 0, 0, 0, 20, 160, 1, 0, 0, 0, 22, 184, 1, 0, 0, 0,
		24, 189, 1, 0, 0, 0, 26, 193, 1, 0, 0, 0, 28, 216, 1, 0, 0, 0, 30, 239,
		1, 0, 0, 0, 32, 241, 1, 0, 0, 0, 34, 246, 1, 0, 0, 0, 36, 258, 1, 0, 0,
		0, 38, 260, 1, 0, 0, 0, 40, 264, 1, 0, 0, 0, 42, 269, 1, 0, 0, 0, 44, 273,
		1, 0, 0, 0, 46, 278, 1, 0, 0, 0, 48, 282, 1, 0, 0, 0, 50, 287, 1, 0, 0,
		0, 52, 291, 1, 0, 0, 0, 54, 296, 1, 0, 0, 0, 56, 58, 3, 2, 1, 0, 57, 56,
		1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5, 0, 0, 1, 63, 1, 1, 0,
		0, 0, 64, 72, 3, 16, 8, 0, 65, 72, 3, 18, 9, 0, 66, 72, 3, 20, 10, 0, 67,
		72, 3, 22, 11, 0, 68, 72, 3, 24, 12, 0, 69, 72, 3, 30, 15, 0, 70, 72, 3,
		32, 16, 0, 71, 64, 1, 0, 0, 0, 71, 65, 1, 0, 0, 0, 71, 66, 1, 0, 0, 0,
		71, 67, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 70, 1,
		0, 0, 0, 72, 3, 1, 0, 0, 0, 73, 74, 5, 15, 0, 0, 74, 75, 5, 1, 0, 0, 75,
		79, 3, 34, 17, 0, 76, 78, 5, 2, 0, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 5, 1, 0, 0, 0, 81, 79,
		1, 0, 0, 0, 82, 83, 5, 15, 0, 0, 83, 84, 5, 15, 0, 0, 84, 85, 5, 1, 0,
		0, 85, 89, 3, 34, 17, 0, 86, 88, 5, 2, 0, 0, 87, 86, 1, 0, 0, 0, 88, 91,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 7, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 92, 93, 5, 15, 0, 0, 93, 94, 5, 15, 0, 0, 94, 95, 5,
		15, 0, 0, 95, 96, 5, 1, 0, 0, 96, 100, 3, 34, 17, 0, 97, 99, 5, 2, 0, 0,
		98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 9, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 104, 5, 15,
		0, 0, 104, 105, 5, 15, 0, 0, 105, 106, 5, 15, 0, 0, 106, 107, 5, 15, 0,
		0, 107, 108, 5, 1, 0, 0, 108, 112, 3, 34, 17, 0, 109, 111, 5, 2, 0, 0,
		110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112,
		113, 1, 0, 0, 0, 113, 11, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5,
		15, 0, 0, 116, 117, 5, 15, 0, 0, 117, 118, 5, 15, 0, 0, 118, 119, 5, 15,
		0, 0, 119, 120, 5, 15, 0, 0, 120, 121, 5, 1, 0, 0, 121, 125, 3, 34, 17,
		0, 122, 124, 5, 2, 0, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 13, 1, 0, 0, 0, 127, 125, 1,
		0, 0, 0, 128, 129, 5, 15, 0, 0, 129, 130, 5, 15, 0, 0, 130, 131, 5, 15,
		0, 0, 131, 132, 5, 15, 0, 0, 132, 133, 5, 15, 0, 0, 133, 134, 5, 15, 0,
		0, 134, 135, 5, 1, 0, 0, 135, 139, 3, 34, 17, 0, 136, 138, 5, 2, 0, 0,
		137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 15, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 149, 3,
		4, 2, 0, 143, 149, 3, 6, 3, 0, 144, 149, 3, 8, 4, 0, 145, 149, 3, 10, 5,
		0, 146, 149, 3, 12, 6, 0, 147, 149, 3, 14, 7, 0, 148, 142, 1, 0, 0, 0,
		148, 143, 1, 0, 0, 0, 148, 144, 1, 0, 0, 0, 148, 145, 1, 0, 0, 0, 148,
		146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 17, 1, 0, 0, 0, 150, 154, 3,
		34, 17, 0, 151, 154, 5, 1, 0, 0, 152, 154, 3, 36, 18, 0, 153, 150, 1, 0,
		0, 0, 153, 151, 1, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0,
		155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157,
		159, 5, 2, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 19, 1,
		0, 0, 0, 160, 163, 5, 9, 0, 0, 161, 164, 3, 34, 17, 0, 162, 164, 3, 36,
		18, 0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 177, 1, 0, 0, 0, 167,
		168, 5, 2, 0, 0, 168, 171, 5, 9, 0, 0, 169, 172, 3, 34, 17, 0, 170, 172,
		3, 36, 18, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 173, 1,
		0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0,
		0, 175, 167, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 182,
		5, 2, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 21, 1, 0,
		0, 0, 183, 185, 3, 26, 13, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0,
		0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 23, 1, 0, 0, 0, 188,
		190, 3, 28, 14, 0, 189, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 189,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 25, 1, 0, 0, 0, 193, 196, 5, 10,
		0, 0, 194, 197, 3, 34, 17, 0, 195, 197, 3, 36, 18, 0, 196, 194, 1, 0, 0,
		0, 196, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 213, 1, 0, 0, 0, 200, 201, 5, 2, 0, 0, 201, 203,
		5, 1, 0, 0, 202, 204, 5, 10, 0, 0, 203, 202, 1, 0, 0, 0, 203, 204, 1, 0,
		0, 0, 204, 207, 1, 0, 0, 0, 205, 208, 3, 34, 17, 0, 206, 208, 3, 36, 18,
		0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 200,
		1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0,
		0, 0, 214, 27, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 219, 5, 11, 0, 0,
		217, 220, 3, 34, 17, 0, 218, 220, 3, 36, 18, 0, 219, 217, 1, 0, 0, 0, 219,
		218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222,
		1, 0, 0, 0, 222, 236, 1, 0, 0, 0, 223, 224, 5, 2, 0, 0, 224, 226, 5, 1,
		0, 0, 225, 227, 5, 11, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0,
		227, 230, 1, 0, 0, 0, 228, 231, 3, 34, 17, 0, 229, 231, 3, 36, 18, 0, 230,
		228, 1, 0, 0, 0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 230,
		1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 235, 1, 0, 0, 0, 234, 223, 1, 0,
		0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0,
		237, 29, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 240, 7, 0, 0, 0, 240, 31,
		1, 0, 0, 0, 241, 242, 5, 12, 0, 0, 242, 33, 1, 0, 0, 0, 243, 247, 3, 36,
		18, 0, 244, 247, 5, 1, 0, 0, 245, 247, 5, 21, 0, 0, 246, 243, 1, 0, 0,
		0, 246, 244, 1, 0, 0, 0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248,
		246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250, 252,
		5, 2, 0, 0, 251, 250, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 35, 1, 0,
		0, 0, 253, 259, 3, 38, 19, 0, 254, 259, 3, 44, 22, 0, 255, 259, 3, 40,
		20, 0, 256, 259, 3, 48, 24, 0, 257, 259, 3, 52, 26, 0, 258, 253, 1, 0,
		0, 0, 258, 254, 1, 0, 0, 0, 258, 255, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0,
		258, 257, 1, 0, 0, 0, 259, 37, 1, 0, 0, 0, 260, 261, 5, 4, 0, 0, 261, 262,
		3, 42, 21, 0, 262, 263, 5, 4, 0, 0, 263, 39, 1, 0, 0, 0, 264, 265, 5, 7,
		0, 0, 265, 266, 3, 46, 23, 0, 266, 267, 5, 7, 0, 0, 267, 41, 1, 0, 0, 0,
		268, 270, 7, 1, 0, 0, 269, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271,
		269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 43, 1, 0, 0, 0, 273, 274, 5,
		5, 0, 0, 274, 275, 3, 46, 23, 0, 275, 276, 5, 5, 0, 0, 276, 45, 1, 0, 0,
		0, 277, 279, 7, 2, 0, 0, 278, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 47, 1, 0, 0, 0, 282, 283, 5,
		6, 0, 0, 283, 284, 3, 50, 25, 0, 284, 285, 5, 6, 0, 0, 285, 49, 1, 0, 0,
		0, 286, 288, 7, 3, 0, 0, 287, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 51, 1, 0, 0, 0, 291, 292, 5,
		8, 0, 0, 292, 293, 3, 54, 27, 0, 293, 294, 5, 8, 0, 0, 294, 53, 1, 0, 0,
		0, 295, 297, 7, 4, 0, 0, 296, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298,
		296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 55, 1, 0, 0, 0, 40, 59, 71,
		79, 89, 100, 112, 125, 139, 148, 153, 155, 158, 163, 165, 171, 173, 177,
		181, 186, 191, 196, 198, 203, 207, 209, 213, 219, 221, 226, 230, 232, 236,
		246, 248, 251, 258, 271, 280, 289, 298,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MarkdownParserInit initializes any static state used to implement MarkdownParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMarkdownParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarkdownParserInit() {
	staticData := &MarkdownParserParserStaticData
	staticData.once.Do(markdownparserParserInit)
}

// NewMarkdownParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMarkdownParser(input antlr.TokenStream) *MarkdownParser {
	MarkdownParserInit()
	this := new(MarkdownParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MarkdownParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MarkdownParser.g4"

	return this
}

// MarkdownParser tokens.
const (
	MarkdownParserEOF                    = antlr.TokenEOF
	MarkdownParserWHITESPACE             = 1
	MarkdownParserNEWLINE                = 2
	MarkdownParserESCAPE_CHAR            = 3
	MarkdownParserBOLD_MARKER            = 4
	MarkdownParserITALIC_MARKER          = 5
	MarkdownParserSTRIKETHROUGH_MARKER   = 6
	MarkdownParserBOLD_AND_ITALIC_MARKER = 7
	MarkdownParserCODE_MARKER            = 8
	MarkdownParserBLOCKQUOTE_MARKER      = 9
	MarkdownParserUNORDERED_LIST_MARKER  = 10
	MarkdownParserORDERED_LIST_MARKER    = 11
	MarkdownParserHORIZONTAL_RULE        = 12
	MarkdownParserFENCED_CODE_BLOCK      = 13
	MarkdownParserINDENTED_CODE_BLOCK    = 14
	MarkdownParserHASH                   = 15
	MarkdownParserLBRACKET               = 16
	MarkdownParserRBRACKET               = 17
	MarkdownParserLPAREN                 = 18
	MarkdownParserRPAREN                 = 19
	MarkdownParserEXCLAMATION            = 20
	MarkdownParserTEXT                   = 21
)

// MarkdownParser rules.
const (
	MarkdownParserRULE_document          = 0
	MarkdownParserRULE_block             = 1
	MarkdownParserRULE_h1                = 2
	MarkdownParserRULE_h2                = 3
	MarkdownParserRULE_h3                = 4
	MarkdownParserRULE_h4                = 5
	MarkdownParserRULE_h5                = 6
	MarkdownParserRULE_h6                = 7
	MarkdownParserRULE_heading           = 8
	MarkdownParserRULE_paragraph         = 9
	MarkdownParserRULE_blockquote        = 10
	MarkdownParserRULE_unorderedList     = 11
	MarkdownParserRULE_orderedList       = 12
	MarkdownParserRULE_listItem          = 13
	MarkdownParserRULE_orderedListItem   = 14
	MarkdownParserRULE_codeBlock         = 15
	MarkdownParserRULE_horizontalRule    = 16
	MarkdownParserRULE_line              = 17
	MarkdownParserRULE_inlineElement     = 18
	MarkdownParserRULE_bold              = 19
	MarkdownParserRULE_boldAndItalic     = 20
	MarkdownParserRULE_boldText          = 21
	MarkdownParserRULE_italic            = 22
	MarkdownParserRULE_italicText        = 23
	MarkdownParserRULE_strikethrough     = 24
	MarkdownParserRULE_strikethroughText = 25
	MarkdownParserRULE_code              = 26
	MarkdownParserRULE_codeText          = 27
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(MarkdownParserEOF, 0)
}

func (s *DocumentContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *MarkdownParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarkdownParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2162674) != 0 {
		{
			p.SetState(56)
			p.Block()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(MarkdownParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Heading() IHeadingContext
	Paragraph() IParagraphContext
	Blockquote() IBlockquoteContext
	UnorderedList() IUnorderedListContext
	OrderedList() IOrderedListContext
	CodeBlock() ICodeBlockContext
	HorizontalRule() IHorizontalRuleContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Heading() IHeadingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeadingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeadingContext)
}

func (s *BlockContext) Paragraph() IParagraphContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParagraphContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParagraphContext)
}

func (s *BlockContext) Blockquote() IBlockquoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockquoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockquoteContext)
}

func (s *BlockContext) UnorderedList() IUnorderedListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnorderedListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnorderedListContext)
}

func (s *BlockContext) OrderedList() IOrderedListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderedListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderedListContext)
}

func (s *BlockContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *BlockContext) HorizontalRule() IHorizontalRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHorizontalRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHorizontalRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *MarkdownParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarkdownParserRULE_block)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Heading()
		}

	case MarkdownParserWHITESPACE, MarkdownParserBOLD_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserSTRIKETHROUGH_MARKER, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserCODE_MARKER, MarkdownParserTEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Paragraph()
		}

	case MarkdownParserBLOCKQUOTE_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Blockquote()
		}

	case MarkdownParserUNORDERED_LIST_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.UnorderedList()
		}

	case MarkdownParserORDERED_LIST_MARKER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.OrderedList()
		}

	case MarkdownParserFENCED_CODE_BLOCK, MarkdownParserINDENTED_CODE_BLOCK:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.CodeBlock()
		}

	case MarkdownParserHORIZONTAL_RULE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.HorizontalRule()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IH1Context is an interface to support dynamic dispatch.
type IH1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	HASH() antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsH1Context differentiates from other interfaces.
	IsH1Context()
}

type H1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH1Context() *H1Context {
	var p = new(H1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h1
	return p
}

func InitEmptyH1Context(p *H1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h1
}

func (*H1Context) IsH1Context() {}

func NewH1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H1Context {
	var p = new(H1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h1

	return p
}

func (s *H1Context) GetParser() antlr.Parser { return s.parser }

func (s *H1Context) GetData() ILineContext { return s.data }

func (s *H1Context) SetData(v ILineContext) { s.data = v }

func (s *H1Context) HASH() antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, 0)
}

func (s *H1Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H1Context) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *H1Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H1Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH1(s)
	}
}

func (s *H1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH1(s)
	}
}

func (p *MarkdownParser) H1() (localctx IH1Context) {
	localctx = NewH1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MarkdownParserRULE_h1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)

		var _x = p.Line()

		localctx.(*H1Context).data = _x
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(76)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IH2Context is an interface to support dynamic dispatch.
type IH2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsH2Context differentiates from other interfaces.
	IsH2Context()
}

type H2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH2Context() *H2Context {
	var p = new(H2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h2
	return p
}

func InitEmptyH2Context(p *H2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h2
}

func (*H2Context) IsH2Context() {}

func NewH2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H2Context {
	var p = new(H2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h2

	return p
}

func (s *H2Context) GetParser() antlr.Parser { return s.parser }

func (s *H2Context) GetData() ILineContext { return s.data }

func (s *H2Context) SetData(v ILineContext) { s.data = v }

func (s *H2Context) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserHASH)
}

func (s *H2Context) HASH(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, i)
}

func (s *H2Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H2Context) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *H2Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H2Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH2(s)
	}
}

func (s *H2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH2(s)
	}
}

func (p *MarkdownParser) H2() (localctx IH2Context) {
	localctx = NewH2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MarkdownParserRULE_h2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)

		var _x = p.Line()

		localctx.(*H2Context).data = _x
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(86)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IH3Context is an interface to support dynamic dispatch.
type IH3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsH3Context differentiates from other interfaces.
	IsH3Context()
}

type H3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH3Context() *H3Context {
	var p = new(H3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h3
	return p
}

func InitEmptyH3Context(p *H3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h3
}

func (*H3Context) IsH3Context() {}

func NewH3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H3Context {
	var p = new(H3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h3

	return p
}

func (s *H3Context) GetParser() antlr.Parser { return s.parser }

func (s *H3Context) GetData() ILineContext { return s.data }

func (s *H3Context) SetData(v ILineContext) { s.data = v }

func (s *H3Context) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserHASH)
}

func (s *H3Context) HASH(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, i)
}

func (s *H3Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H3Context) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *H3Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H3Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH3(s)
	}
}

func (s *H3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH3(s)
	}
}

func (p *MarkdownParser) H3() (localctx IH3Context) {
	localctx = NewH3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MarkdownParserRULE_h3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)

		var _x = p.Line()

		localctx.(*H3Context).data = _x
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(97)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IH4Context is an interface to support dynamic dispatch.
type IH4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsH4Context differentiates from other interfaces.
	IsH4Context()
}

type H4Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH4Context() *H4Context {
	var p = new(H4Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h4
	return p
}

func InitEmptyH4Context(p *H4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h4
}

func (*H4Context) IsH4Context() {}

func NewH4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H4Context {
	var p = new(H4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h4

	return p
}

func (s *H4Context) GetParser() antlr.Parser { return s.parser }

func (s *H4Context) GetData() ILineContext { return s.data }

func (s *H4Context) SetData(v ILineContext) { s.data = v }

func (s *H4Context) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserHASH)
}

func (s *H4Context) HASH(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, i)
}

func (s *H4Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H4Context) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *H4Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H4Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH4(s)
	}
}

func (s *H4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH4(s)
	}
}

func (p *MarkdownParser) H4() (localctx IH4Context) {
	localctx = NewH4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MarkdownParserRULE_h4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)

		var _x = p.Line()

		localctx.(*H4Context).data = _x
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(109)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IH5Context is an interface to support dynamic dispatch.
type IH5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsH5Context differentiates from other interfaces.
	IsH5Context()
}

type H5Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH5Context() *H5Context {
	var p = new(H5Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h5
	return p
}

func InitEmptyH5Context(p *H5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h5
}

func (*H5Context) IsH5Context() {}

func NewH5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H5Context {
	var p = new(H5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h5

	return p
}

func (s *H5Context) GetParser() antlr.Parser { return s.parser }

func (s *H5Context) GetData() ILineContext { return s.data }

func (s *H5Context) SetData(v ILineContext) { s.data = v }

func (s *H5Context) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserHASH)
}

func (s *H5Context) HASH(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, i)
}

func (s *H5Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H5Context) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *H5Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H5Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH5(s)
	}
}

func (s *H5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH5(s)
	}
}

func (p *MarkdownParser) H5() (localctx IH5Context) {
	localctx = NewH5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MarkdownParserRULE_h5)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)

		var _x = p.Line()

		localctx.(*H5Context).data = _x
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(122)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IH6Context is an interface to support dynamic dispatch.
type IH6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ILineContext

	// SetData sets the data rule contexts.
	SetData(ILineContext)

	// Getter signatures
	AllHASH() []antlr.TerminalNode
	HASH(i int) antlr.TerminalNode
	WHITESPACE() antlr.TerminalNode
	Line() ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsH6Context differentiates from other interfaces.
	IsH6Context()
}

type H6Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ILineContext
}

func NewEmptyH6Context() *H6Context {
	var p = new(H6Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h6
	return p
}

func InitEmptyH6Context(p *H6Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_h6
}

func (*H6Context) IsH6Context() {}

func NewH6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H6Context {
	var p = new(H6Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_h6

	return p
}

func (s *H6Context) GetParser() antlr.Parser { return s.parser }

func (s *H6Context) GetData() ILineContext { return s.data }

func (s *H6Context) SetData(v ILineContext) { s.data = v }

func (s *H6Context) AllHASH() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserHASH)
}

func (s *H6Context) HASH(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserHASH, i)
}

func (s *H6Context) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, 0)
}

func (s *H6Context) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *H6Context) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *H6Context) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *H6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterH6(s)
	}
}

func (s *H6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitH6(s)
	}
}

func (p *MarkdownParser) H6() (localctx IH6Context) {
	localctx = NewH6Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MarkdownParserRULE_h6)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(MarkdownParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(MarkdownParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)

		var _x = p.Line()

		localctx.(*H6Context).data = _x
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(136)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeadingContext is an interface to support dynamic dispatch.
type IHeadingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	H1() IH1Context
	H2() IH2Context
	H3() IH3Context
	H4() IH4Context
	H5() IH5Context
	H6() IH6Context

	// IsHeadingContext differentiates from other interfaces.
	IsHeadingContext()
}

type HeadingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeadingContext() *HeadingContext {
	var p = new(HeadingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_heading
	return p
}

func InitEmptyHeadingContext(p *HeadingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_heading
}

func (*HeadingContext) IsHeadingContext() {}

func NewHeadingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadingContext {
	var p = new(HeadingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_heading

	return p
}

func (s *HeadingContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadingContext) H1() IH1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH1Context)
}

func (s *HeadingContext) H2() IH2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH2Context)
}

func (s *HeadingContext) H3() IH3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH3Context)
}

func (s *HeadingContext) H4() IH4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH4Context)
}

func (s *HeadingContext) H5() IH5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH5Context)
}

func (s *HeadingContext) H6() IH6Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IH6Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IH6Context)
}

func (s *HeadingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterHeading(s)
	}
}

func (s *HeadingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitHeading(s)
	}
}

func (p *MarkdownParser) Heading() (localctx IHeadingContext) {
	localctx = NewHeadingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MarkdownParserRULE_heading)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.H1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.H2()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.H3()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(145)
			p.H4()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			p.H5()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(147)
			p.H6()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	NEWLINE() antlr.TerminalNode

	// IsParagraphContext differentiates from other interfaces.
	IsParagraphContext()
}

type ParagraphContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParagraphContext() *ParagraphContext {
	var p = new(ParagraphContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_paragraph
	return p
}

func InitEmptyParagraphContext(p *ParagraphContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_paragraph
}

func (*ParagraphContext) IsParagraphContext() {}

func NewParagraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParagraphContext {
	var p = new(ParagraphContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_paragraph

	return p
}

func (s *ParagraphContext) GetParser() antlr.Parser { return s.parser }

func (s *ParagraphContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ParagraphContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ParagraphContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *ParagraphContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *ParagraphContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *ParagraphContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *ParagraphContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *ParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParagraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterParagraph(s)
	}
}

func (s *ParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitParagraph(s)
	}
}

func (p *MarkdownParser) Paragraph() (localctx IParagraphContext) {
	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MarkdownParserRULE_paragraph)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(150)
					p.Line()
				}

			case 2:
				{
					p.SetState(151)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				{
					p.SetState(152)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MarkdownParserNEWLINE {
		{
			p.SetState(157)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockquoteContext is an interface to support dynamic dispatch.
type IBlockquoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBLOCKQUOTE_MARKER() []antlr.TerminalNode
	BLOCKQUOTE_MARKER(i int) antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsBlockquoteContext differentiates from other interfaces.
	IsBlockquoteContext()
}

type BlockquoteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockquoteContext() *BlockquoteContext {
	var p = new(BlockquoteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_blockquote
	return p
}

func InitEmptyBlockquoteContext(p *BlockquoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_blockquote
}

func (*BlockquoteContext) IsBlockquoteContext() {}

func NewBlockquoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockquoteContext {
	var p = new(BlockquoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_blockquote

	return p
}

func (s *BlockquoteContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockquoteContext) AllBLOCKQUOTE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBLOCKQUOTE_MARKER)
}

func (s *BlockquoteContext) BLOCKQUOTE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBLOCKQUOTE_MARKER, i)
}

func (s *BlockquoteContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *BlockquoteContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *BlockquoteContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *BlockquoteContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *BlockquoteContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *BlockquoteContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *BlockquoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockquoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockquoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBlockquote(s)
	}
}

func (s *BlockquoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBlockquote(s)
	}
}

func (p *MarkdownParser) Blockquote() (localctx IBlockquoteContext) {
	localctx = NewBlockquoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MarkdownParserRULE_blockquote)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(MarkdownParserBLOCKQUOTE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(161)
					p.Line()
				}

			case 2:
				{
					p.SetState(162)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(167)
				p.Match(MarkdownParserNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(168)
				p.Match(MarkdownParserBLOCKQUOTE_MARKER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					p.SetState(171)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}

					switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
					case 1:
						{
							p.SetState(169)
							p.Line()
						}

					case 2:
						{
							p.SetState(170)
							p.InlineElement()
						}

					case antlr.ATNInvalidAltNumber:
						goto errorExit
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(173)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MarkdownParserNEWLINE {
		{
			p.SetState(180)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnorderedListContext is an interface to support dynamic dispatch.
type IUnorderedListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllListItem() []IListItemContext
	ListItem(i int) IListItemContext

	// IsUnorderedListContext differentiates from other interfaces.
	IsUnorderedListContext()
}

type UnorderedListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnorderedListContext() *UnorderedListContext {
	var p = new(UnorderedListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_unorderedList
	return p
}

func InitEmptyUnorderedListContext(p *UnorderedListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_unorderedList
}

func (*UnorderedListContext) IsUnorderedListContext() {}

func NewUnorderedListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnorderedListContext {
	var p = new(UnorderedListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_unorderedList

	return p
}

func (s *UnorderedListContext) GetParser() antlr.Parser { return s.parser }

func (s *UnorderedListContext) AllListItem() []IListItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListItemContext); ok {
			len++
		}
	}

	tst := make([]IListItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListItemContext); ok {
			tst[i] = t.(IListItemContext)
			i++
		}
	}

	return tst
}

func (s *UnorderedListContext) ListItem(i int) IListItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListItemContext)
}

func (s *UnorderedListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnorderedListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnorderedListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterUnorderedList(s)
	}
}

func (s *UnorderedListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitUnorderedList(s)
	}
}

func (p *MarkdownParser) UnorderedList() (localctx IUnorderedListContext) {
	localctx = NewUnorderedListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MarkdownParserRULE_unorderedList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(183)
				p.ListItem()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderedListContext is an interface to support dynamic dispatch.
type IOrderedListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOrderedListItem() []IOrderedListItemContext
	OrderedListItem(i int) IOrderedListItemContext

	// IsOrderedListContext differentiates from other interfaces.
	IsOrderedListContext()
}

type OrderedListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderedListContext() *OrderedListContext {
	var p = new(OrderedListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedList
	return p
}

func InitEmptyOrderedListContext(p *OrderedListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedList
}

func (*OrderedListContext) IsOrderedListContext() {}

func NewOrderedListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderedListContext {
	var p = new(OrderedListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_orderedList

	return p
}

func (s *OrderedListContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderedListContext) AllOrderedListItem() []IOrderedListItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderedListItemContext); ok {
			len++
		}
	}

	tst := make([]IOrderedListItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderedListItemContext); ok {
			tst[i] = t.(IOrderedListItemContext)
			i++
		}
	}

	return tst
}

func (s *OrderedListContext) OrderedListItem(i int) IOrderedListItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderedListItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderedListItemContext)
}

func (s *OrderedListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderedListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderedListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterOrderedList(s)
	}
}

func (s *OrderedListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitOrderedList(s)
	}
}

func (p *MarkdownParser) OrderedList() (localctx IOrderedListContext) {
	localctx = NewOrderedListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MarkdownParserRULE_orderedList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(188)
				p.OrderedListItem()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListItemContext is an interface to support dynamic dispatch.
type IListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUNORDERED_LIST_MARKER() []antlr.TerminalNode
	UNORDERED_LIST_MARKER(i int) antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsListItemContext differentiates from other interfaces.
	IsListItemContext()
}

type ListItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListItemContext() *ListItemContext {
	var p = new(ListItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_listItem
	return p
}

func InitEmptyListItemContext(p *ListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_listItem
}

func (*ListItemContext) IsListItemContext() {}

func NewListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListItemContext {
	var p = new(ListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_listItem

	return p
}

func (s *ListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ListItemContext) AllUNORDERED_LIST_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserUNORDERED_LIST_MARKER)
}

func (s *ListItemContext) UNORDERED_LIST_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserUNORDERED_LIST_MARKER, i)
}

func (s *ListItemContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ListItemContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ListItemContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *ListItemContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *ListItemContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *ListItemContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *ListItemContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *ListItemContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *ListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterListItem(s)
	}
}

func (s *ListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitListItem(s)
	}
}

func (p *MarkdownParser) ListItem() (localctx IListItemContext) {
	localctx = NewListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MarkdownParserRULE_listItem)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(MarkdownParserUNORDERED_LIST_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(194)
					p.Line()
				}

			case 2:
				{
					p.SetState(195)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(200)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MarkdownParserUNORDERED_LIST_MARKER {
			{
				p.SetState(202)
				p.Match(MarkdownParserUNORDERED_LIST_MARKER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(207)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(205)
						p.Line()
					}

				case 2:
					{
						p.SetState(206)
						p.InlineElement()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderedListItemContext is an interface to support dynamic dispatch.
type IOrderedListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllORDERED_LIST_MARKER() []antlr.TerminalNode
	ORDERED_LIST_MARKER(i int) antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsOrderedListItemContext differentiates from other interfaces.
	IsOrderedListItemContext()
}

type OrderedListItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderedListItemContext() *OrderedListItemContext {
	var p = new(OrderedListItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedListItem
	return p
}

func InitEmptyOrderedListItemContext(p *OrderedListItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_orderedListItem
}

func (*OrderedListItemContext) IsOrderedListItemContext() {}

func NewOrderedListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderedListItemContext {
	var p = new(OrderedListItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_orderedListItem

	return p
}

func (s *OrderedListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderedListItemContext) AllORDERED_LIST_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserORDERED_LIST_MARKER)
}

func (s *OrderedListItemContext) ORDERED_LIST_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserORDERED_LIST_MARKER, i)
}

func (s *OrderedListItemContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *OrderedListItemContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *OrderedListItemContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *OrderedListItemContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *OrderedListItemContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserNEWLINE)
}

func (s *OrderedListItemContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, i)
}

func (s *OrderedListItemContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *OrderedListItemContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *OrderedListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderedListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderedListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterOrderedListItem(s)
	}
}

func (s *OrderedListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitOrderedListItem(s)
	}
}

func (p *MarkdownParser) OrderedListItem() (localctx IOrderedListItemContext) {
	localctx = NewOrderedListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MarkdownParserRULE_orderedListItem)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(MarkdownParserORDERED_LIST_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(217)
					p.Line()
				}

			case 2:
				{
					p.SetState(218)
					p.InlineElement()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MarkdownParserNEWLINE {
		{
			p.SetState(223)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(MarkdownParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MarkdownParserORDERED_LIST_MARKER {
			{
				p.SetState(225)
				p.Match(MarkdownParserORDERED_LIST_MARKER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(230)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(228)
						p.Line()
					}

				case 2:
					{
						p.SetState(229)
						p.InlineElement()
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeBlockContext is an interface to support dynamic dispatch.
type ICodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FENCED_CODE_BLOCK() antlr.TerminalNode
	INDENTED_CODE_BLOCK() antlr.TerminalNode

	// IsCodeBlockContext differentiates from other interfaces.
	IsCodeBlockContext()
}

type CodeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContext() *CodeBlockContext {
	var p = new(CodeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeBlock
	return p
}

func InitEmptyCodeBlockContext(p *CodeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeBlock
}

func (*CodeBlockContext) IsCodeBlockContext() {}

func NewCodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContext {
	var p = new(CodeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_codeBlock

	return p
}

func (s *CodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContext) FENCED_CODE_BLOCK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserFENCED_CODE_BLOCK, 0)
}

func (s *CodeBlockContext) INDENTED_CODE_BLOCK() antlr.TerminalNode {
	return s.GetToken(MarkdownParserINDENTED_CODE_BLOCK, 0)
}

func (s *CodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCodeBlock(s)
	}
}

func (s *CodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCodeBlock(s)
	}
}

func (p *MarkdownParser) CodeBlock() (localctx ICodeBlockContext) {
	localctx = NewCodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MarkdownParserRULE_codeBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MarkdownParserFENCED_CODE_BLOCK || _la == MarkdownParserINDENTED_CODE_BLOCK) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHorizontalRuleContext is an interface to support dynamic dispatch.
type IHorizontalRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HORIZONTAL_RULE() antlr.TerminalNode

	// IsHorizontalRuleContext differentiates from other interfaces.
	IsHorizontalRuleContext()
}

type HorizontalRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHorizontalRuleContext() *HorizontalRuleContext {
	var p = new(HorizontalRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_horizontalRule
	return p
}

func InitEmptyHorizontalRuleContext(p *HorizontalRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_horizontalRule
}

func (*HorizontalRuleContext) IsHorizontalRuleContext() {}

func NewHorizontalRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HorizontalRuleContext {
	var p = new(HorizontalRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_horizontalRule

	return p
}

func (s *HorizontalRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *HorizontalRuleContext) HORIZONTAL_RULE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserHORIZONTAL_RULE, 0)
}

func (s *HorizontalRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HorizontalRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HorizontalRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterHorizontalRule(s)
	}
}

func (s *HorizontalRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitHorizontalRule(s)
	}
}

func (p *MarkdownParser) HorizontalRule() (localctx IHorizontalRuleContext) {
	localctx = NewHorizontalRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MarkdownParserRULE_horizontalRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(MarkdownParserHORIZONTAL_RULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInlineElement() []IInlineElementContext
	InlineElement(i int) IInlineElementContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	NEWLINE() antlr.TerminalNode

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) AllInlineElement() []IInlineElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInlineElementContext); ok {
			len++
		}
	}

	tst := make([]IInlineElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInlineElementContext); ok {
			tst[i] = t.(IInlineElementContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) InlineElement(i int) IInlineElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineElementContext)
}

func (s *LineContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *LineContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *LineContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *LineContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MarkdownParserNEWLINE, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *MarkdownParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MarkdownParserRULE_line)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MarkdownParserBOLD_MARKER, MarkdownParserITALIC_MARKER, MarkdownParserSTRIKETHROUGH_MARKER, MarkdownParserBOLD_AND_ITALIC_MARKER, MarkdownParserCODE_MARKER:
				{
					p.SetState(243)
					p.InlineElement()
				}

			case MarkdownParserWHITESPACE:
				{
					p.SetState(244)
					p.Match(MarkdownParserWHITESPACE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case MarkdownParserTEXT:
				{
					p.SetState(245)
					p.Match(MarkdownParserTEXT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(250)
			p.Match(MarkdownParserNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInlineElementContext is an interface to support dynamic dispatch.
type IInlineElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bold() IBoldContext
	Italic() IItalicContext
	BoldAndItalic() IBoldAndItalicContext
	Strikethrough() IStrikethroughContext
	Code() ICodeContext

	// IsInlineElementContext differentiates from other interfaces.
	IsInlineElementContext()
}

type InlineElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineElementContext() *InlineElementContext {
	var p = new(InlineElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_inlineElement
	return p
}

func InitEmptyInlineElementContext(p *InlineElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_inlineElement
}

func (*InlineElementContext) IsInlineElementContext() {}

func NewInlineElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineElementContext {
	var p = new(InlineElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_inlineElement

	return p
}

func (s *InlineElementContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineElementContext) Bold() IBoldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoldContext)
}

func (s *InlineElementContext) Italic() IItalicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalicContext)
}

func (s *InlineElementContext) BoldAndItalic() IBoldAndItalicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldAndItalicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoldAndItalicContext)
}

func (s *InlineElementContext) Strikethrough() IStrikethroughContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrikethroughContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrikethroughContext)
}

func (s *InlineElementContext) Code() ICodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *InlineElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterInlineElement(s)
	}
}

func (s *InlineElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitInlineElement(s)
	}
}

func (p *MarkdownParser) InlineElement() (localctx IInlineElementContext) {
	localctx = NewInlineElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MarkdownParserRULE_inlineElement)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MarkdownParserBOLD_MARKER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Bold()
		}

	case MarkdownParserITALIC_MARKER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Italic()
		}

	case MarkdownParserBOLD_AND_ITALIC_MARKER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.BoldAndItalic()
		}

	case MarkdownParserSTRIKETHROUGH_MARKER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
			p.Strikethrough()
		}

	case MarkdownParserCODE_MARKER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(257)
			p.Code()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoldContext is an interface to support dynamic dispatch.
type IBoldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IBoldTextContext

	// SetData sets the data rule contexts.
	SetData(IBoldTextContext)

	// Getter signatures
	AllBOLD_MARKER() []antlr.TerminalNode
	BOLD_MARKER(i int) antlr.TerminalNode
	BoldText() IBoldTextContext

	// IsBoldContext differentiates from other interfaces.
	IsBoldContext()
}

type BoldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IBoldTextContext
}

func NewEmptyBoldContext() *BoldContext {
	var p = new(BoldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold
	return p
}

func InitEmptyBoldContext(p *BoldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_bold
}

func (*BoldContext) IsBoldContext() {}

func NewBoldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldContext {
	var p = new(BoldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_bold

	return p
}

func (s *BoldContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldContext) GetData() IBoldTextContext { return s.data }

func (s *BoldContext) SetData(v IBoldTextContext) { s.data = v }

func (s *BoldContext) AllBOLD_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_MARKER)
}

func (s *BoldContext) BOLD_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_MARKER, i)
}

func (s *BoldContext) BoldText() IBoldTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoldTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoldTextContext)
}

func (s *BoldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBold(s)
	}
}

func (s *BoldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBold(s)
	}
}

func (p *MarkdownParser) Bold() (localctx IBoldContext) {
	localctx = NewBoldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MarkdownParserRULE_bold)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)

		var _x = p.BoldText()

		localctx.(*BoldContext).data = _x
	}
	{
		p.SetState(262)
		p.Match(MarkdownParserBOLD_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoldAndItalicContext is an interface to support dynamic dispatch.
type IBoldAndItalicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IItalicTextContext

	// SetData sets the data rule contexts.
	SetData(IItalicTextContext)

	// Getter signatures
	AllBOLD_AND_ITALIC_MARKER() []antlr.TerminalNode
	BOLD_AND_ITALIC_MARKER(i int) antlr.TerminalNode
	ItalicText() IItalicTextContext

	// IsBoldAndItalicContext differentiates from other interfaces.
	IsBoldAndItalicContext()
}

type BoldAndItalicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IItalicTextContext
}

func NewEmptyBoldAndItalicContext() *BoldAndItalicContext {
	var p = new(BoldAndItalicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldAndItalic
	return p
}

func InitEmptyBoldAndItalicContext(p *BoldAndItalicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldAndItalic
}

func (*BoldAndItalicContext) IsBoldAndItalicContext() {}

func NewBoldAndItalicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldAndItalicContext {
	var p = new(BoldAndItalicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_boldAndItalic

	return p
}

func (s *BoldAndItalicContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldAndItalicContext) GetData() IItalicTextContext { return s.data }

func (s *BoldAndItalicContext) SetData(v IItalicTextContext) { s.data = v }

func (s *BoldAndItalicContext) AllBOLD_AND_ITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_AND_ITALIC_MARKER)
}

func (s *BoldAndItalicContext) BOLD_AND_ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_AND_ITALIC_MARKER, i)
}

func (s *BoldAndItalicContext) ItalicText() IItalicTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalicTextContext)
}

func (s *BoldAndItalicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldAndItalicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldAndItalicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBoldAndItalic(s)
	}
}

func (s *BoldAndItalicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBoldAndItalic(s)
	}
}

func (p *MarkdownParser) BoldAndItalic() (localctx IBoldAndItalicContext) {
	localctx = NewBoldAndItalicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MarkdownParserRULE_boldAndItalic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)

		var _x = p.ItalicText()

		localctx.(*BoldAndItalicContext).data = _x
	}
	{
		p.SetState(266)
		p.Match(MarkdownParserBOLD_AND_ITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoldTextContext is an interface to support dynamic dispatch.
type IBoldTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode

	// IsBoldTextContext differentiates from other interfaces.
	IsBoldTextContext()
}

type BoldTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoldTextContext() *BoldTextContext {
	var p = new(BoldTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldText
	return p
}

func InitEmptyBoldTextContext(p *BoldTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_boldText
}

func (*BoldTextContext) IsBoldTextContext() {}

func NewBoldTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldTextContext {
	var p = new(BoldTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_boldText

	return p
}

func (s *BoldTextContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *BoldTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *BoldTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *BoldTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *BoldTextContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *BoldTextContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
}

func (s *BoldTextContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *BoldTextContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *BoldTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterBoldText(s)
	}
}

func (s *BoldTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitBoldText(s)
	}
}

func (p *MarkdownParser) BoldText() (localctx IBoldTextContext) {
	localctx = NewBoldTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MarkdownParserRULE_boldText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097442) != 0) {
		{
			p.SetState(268)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097442) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItalicContext is an interface to support dynamic dispatch.
type IItalicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IItalicTextContext

	// SetData sets the data rule contexts.
	SetData(IItalicTextContext)

	// Getter signatures
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
	ItalicText() IItalicTextContext

	// IsItalicContext differentiates from other interfaces.
	IsItalicContext()
}

type ItalicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IItalicTextContext
}

func NewEmptyItalicContext() *ItalicContext {
	var p = new(ItalicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italic
	return p
}

func InitEmptyItalicContext(p *ItalicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italic
}

func (*ItalicContext) IsItalicContext() {}

func NewItalicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItalicContext {
	var p = new(ItalicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_italic

	return p
}

func (s *ItalicContext) GetParser() antlr.Parser { return s.parser }

func (s *ItalicContext) GetData() IItalicTextContext { return s.data }

func (s *ItalicContext) SetData(v IItalicTextContext) { s.data = v }

func (s *ItalicContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *ItalicContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
}

func (s *ItalicContext) ItalicText() IItalicTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItalicTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItalicTextContext)
}

func (s *ItalicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItalicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItalicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterItalic(s)
	}
}

func (s *ItalicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitItalic(s)
	}
}

func (p *MarkdownParser) Italic() (localctx IItalicContext) {
	localctx = NewItalicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MarkdownParserRULE_italic)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(MarkdownParserITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)

		var _x = p.ItalicText()

		localctx.(*ItalicContext).data = _x
	}
	{
		p.SetState(275)
		p.Match(MarkdownParserITALIC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IItalicTextContext is an interface to support dynamic dispatch.
type IItalicTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode

	// IsItalicTextContext differentiates from other interfaces.
	IsItalicTextContext()
}

type ItalicTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItalicTextContext() *ItalicTextContext {
	var p = new(ItalicTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italicText
	return p
}

func InitEmptyItalicTextContext(p *ItalicTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_italicText
}

func (*ItalicTextContext) IsItalicTextContext() {}

func NewItalicTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItalicTextContext {
	var p = new(ItalicTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_italicText

	return p
}

func (s *ItalicTextContext) GetParser() antlr.Parser { return s.parser }

func (s *ItalicTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *ItalicTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *ItalicTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *ItalicTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *ItalicTextContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *ItalicTextContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *ItalicTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItalicTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItalicTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterItalicText(s)
	}
}

func (s *ItalicTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitItalicText(s)
	}
}

func (p *MarkdownParser) ItalicText() (localctx IItalicTextContext) {
	localctx = NewItalicTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MarkdownParserRULE_italicText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097410) != 0) {
		{
			p.SetState(277)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097410) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrikethroughContext is an interface to support dynamic dispatch.
type IStrikethroughContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() IStrikethroughTextContext

	// SetData sets the data rule contexts.
	SetData(IStrikethroughTextContext)

	// Getter signatures
	AllSTRIKETHROUGH_MARKER() []antlr.TerminalNode
	STRIKETHROUGH_MARKER(i int) antlr.TerminalNode
	StrikethroughText() IStrikethroughTextContext

	// IsStrikethroughContext differentiates from other interfaces.
	IsStrikethroughContext()
}

type StrikethroughContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   IStrikethroughTextContext
}

func NewEmptyStrikethroughContext() *StrikethroughContext {
	var p = new(StrikethroughContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethrough
	return p
}

func InitEmptyStrikethroughContext(p *StrikethroughContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethrough
}

func (*StrikethroughContext) IsStrikethroughContext() {}

func NewStrikethroughContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrikethroughContext {
	var p = new(StrikethroughContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_strikethrough

	return p
}

func (s *StrikethroughContext) GetParser() antlr.Parser { return s.parser }

func (s *StrikethroughContext) GetData() IStrikethroughTextContext { return s.data }

func (s *StrikethroughContext) SetData(v IStrikethroughTextContext) { s.data = v }

func (s *StrikethroughContext) AllSTRIKETHROUGH_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserSTRIKETHROUGH_MARKER)
}

func (s *StrikethroughContext) STRIKETHROUGH_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserSTRIKETHROUGH_MARKER, i)
}

func (s *StrikethroughContext) StrikethroughText() IStrikethroughTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrikethroughTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrikethroughTextContext)
}

func (s *StrikethroughContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrikethroughContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrikethroughContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterStrikethrough(s)
	}
}

func (s *StrikethroughContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitStrikethrough(s)
	}
}

func (p *MarkdownParser) Strikethrough() (localctx IStrikethroughContext) {
	localctx = NewStrikethroughContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MarkdownParserRULE_strikethrough)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(MarkdownParserSTRIKETHROUGH_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)

		var _x = p.StrikethroughText()

		localctx.(*StrikethroughContext).data = _x
	}
	{
		p.SetState(284)
		p.Match(MarkdownParserSTRIKETHROUGH_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrikethroughTextContext is an interface to support dynamic dispatch.
type IStrikethroughTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllITALIC_MARKER() []antlr.TerminalNode
	ITALIC_MARKER(i int) antlr.TerminalNode
	AllBOLD_MARKER() []antlr.TerminalNode
	BOLD_MARKER(i int) antlr.TerminalNode
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode

	// IsStrikethroughTextContext differentiates from other interfaces.
	IsStrikethroughTextContext()
}

type StrikethroughTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrikethroughTextContext() *StrikethroughTextContext {
	var p = new(StrikethroughTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethroughText
	return p
}

func InitEmptyStrikethroughTextContext(p *StrikethroughTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_strikethroughText
}

func (*StrikethroughTextContext) IsStrikethroughTextContext() {}

func NewStrikethroughTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrikethroughTextContext {
	var p = new(StrikethroughTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_strikethroughText

	return p
}

func (s *StrikethroughTextContext) GetParser() antlr.Parser { return s.parser }

func (s *StrikethroughTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *StrikethroughTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *StrikethroughTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *StrikethroughTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *StrikethroughTextContext) AllITALIC_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserITALIC_MARKER)
}

func (s *StrikethroughTextContext) ITALIC_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserITALIC_MARKER, i)
}

func (s *StrikethroughTextContext) AllBOLD_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserBOLD_MARKER)
}

func (s *StrikethroughTextContext) BOLD_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserBOLD_MARKER, i)
}

func (s *StrikethroughTextContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *StrikethroughTextContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *StrikethroughTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrikethroughTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrikethroughTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterStrikethroughText(s)
	}
}

func (s *StrikethroughTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitStrikethroughText(s)
	}
}

func (p *MarkdownParser) StrikethroughText() (localctx IStrikethroughTextContext) {
	localctx = NewStrikethroughTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MarkdownParserRULE_strikethroughText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097458) != 0) {
		{
			p.SetState(286)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097458) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetData returns the data rule contexts.
	GetData() ICodeTextContext

	// SetData sets the data rule contexts.
	SetData(ICodeTextContext)

	// Getter signatures
	AllCODE_MARKER() []antlr.TerminalNode
	CODE_MARKER(i int) antlr.TerminalNode
	CodeText() ICodeTextContext

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	data   ICodeTextContext
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_code
	return p
}

func InitEmptyCodeContext(p *CodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_code
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) GetData() ICodeTextContext { return s.data }

func (s *CodeContext) SetData(v ICodeTextContext) { s.data = v }

func (s *CodeContext) AllCODE_MARKER() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserCODE_MARKER)
}

func (s *CodeContext) CODE_MARKER(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserCODE_MARKER, i)
}

func (s *CodeContext) CodeText() ICodeTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeTextContext)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCode(s)
	}
}

func (s *CodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCode(s)
	}
}

func (p *MarkdownParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MarkdownParserRULE_code)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)

		var _x = p.CodeText()

		localctx.(*CodeContext).data = _x
	}
	{
		p.SetState(293)
		p.Match(MarkdownParserCODE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeTextContext is an interface to support dynamic dispatch.
type ICodeTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllESCAPE_CHAR() []antlr.TerminalNode
	ESCAPE_CHAR(i int) antlr.TerminalNode

	// IsCodeTextContext differentiates from other interfaces.
	IsCodeTextContext()
}

type CodeTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeTextContext() *CodeTextContext {
	var p = new(CodeTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeText
	return p
}

func InitEmptyCodeTextContext(p *CodeTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MarkdownParserRULE_codeText
}

func (*CodeTextContext) IsCodeTextContext() {}

func NewCodeTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeTextContext {
	var p = new(CodeTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarkdownParserRULE_codeText

	return p
}

func (s *CodeTextContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeTextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserTEXT)
}

func (s *CodeTextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserTEXT, i)
}

func (s *CodeTextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserWHITESPACE)
}

func (s *CodeTextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserWHITESPACE, i)
}

func (s *CodeTextContext) AllESCAPE_CHAR() []antlr.TerminalNode {
	return s.GetTokens(MarkdownParserESCAPE_CHAR)
}

func (s *CodeTextContext) ESCAPE_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(MarkdownParserESCAPE_CHAR, i)
}

func (s *CodeTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.EnterCodeText(s)
	}
}

func (s *CodeTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarkdownParserListener); ok {
		listenerT.ExitCodeText(s)
	}
}

func (p *MarkdownParser) CodeText() (localctx ICodeTextContext) {
	localctx = NewCodeTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MarkdownParserRULE_codeText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
		{
			p.SetState(295)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097162) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
