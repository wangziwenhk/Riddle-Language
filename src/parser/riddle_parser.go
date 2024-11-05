// Code generated from E:/Riddle-Go/RiddleParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RiddleParser

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

type RiddleParser struct {
	*antlr.BaseParser
}

var RiddleParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func riddleparserParserInit() {
	staticData := &RiddleParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'var'", "'val'", "'for'", "'while'", "'if'", "'else'", "'fun'",
		"'return'", "'import'", "'package'", "'class'", "'public'", "'protected'",
		"'Private'", "'override'", "'true'", "'false'", "'static'", "'const'",
		"'null'", "'try'", "'catch'", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"':'", "';'", "','", "'=='", "'='", "'>'", "'<'", "'<<'", "'>>'", "'>>>'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'&'", "'|'", "'^'", "'.'",
		"'\"'", "'''", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "Var", "Val", "For", "While", "If", "Else", "Func", "Return", "Import",
		"Package", "Class", "Public", "Protected", "Private", "Override", "True",
		"False", "Static", "Const", "Null", "Try", "Catch", "LeftBracket", "RightBracket",
		"LeftSquare", "RightSquare", "LeftCurly", "RightCurly", "Colon", "Semi",
		"Comma", "Equal", "Assign", "Greater", "Less", "LeftLeft", "RightRight",
		"RightRightRight", "Add", "Sub", "Star", "Div", "Mod", "Not", "And",
		"Or", "Xor", "Dot", "DoubleQuotes", "Quotes", "Endl", "Identifier",
		"Hexadecimal", "Decimal", "Octal", "Binary", "Float", "IntegerSequence",
		"HEX_DIGIT", "OCTAL_DIGIT", "BINARY_DIGIT", "DIGIT", "STRING", "LINE_COMMENT",
		"BLOCK_COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "null_cnt", "statement_ed", "statement", "bodyExpr", "packStatement",
		"importStatement", "varDefineStatement", "argsExpr", "defineArgs", "funcDefine",
		"forStatement", "whileStatement", "ifStatement", "returnStatement",
		"classDefine", "tryExpr", "catchExpr", "exprPtr", "exprPtrParser", "expression",
		"id", "number", "boolean", "string", "float", "integer", "templateArg",
		"templateArgs", "typeName",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 477, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 5, 0, 62, 8, 0,
		10, 0, 12, 0, 65, 9, 0, 1, 0, 3, 0, 68, 8, 0, 1, 1, 1, 1, 1, 2, 1, 2, 3,
		2, 74, 8, 2, 1, 2, 3, 2, 77, 8, 2, 1, 2, 1, 2, 3, 2, 81, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 95,
		8, 3, 1, 4, 1, 4, 5, 4, 99, 8, 4, 10, 4, 12, 4, 102, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 127, 8, 7,
		1, 8, 1, 8, 1, 8, 5, 8, 132, 8, 8, 10, 8, 12, 8, 135, 9, 8, 1, 8, 3, 8,
		138, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 145, 8, 9, 10, 9, 12, 9,
		148, 9, 9, 1, 9, 1, 9, 1, 9, 3, 9, 153, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 163, 8, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 3, 11, 170, 8, 11, 1, 11, 1, 11, 3, 11, 174, 8, 11, 1, 11, 1,
		11, 3, 11, 178, 8, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 205, 8, 13,
		1, 14, 1, 14, 3, 14, 209, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 3, 16, 218, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 234, 8,
		18, 1, 18, 1, 18, 1, 18, 5, 18, 239, 8, 18, 10, 18, 12, 18, 242, 9, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 347,
		8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		5, 20, 411, 8, 20, 10, 20, 12, 20, 414, 9, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 3, 22, 420, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 426, 8, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 3, 26, 441, 8, 26, 1, 27, 1, 27, 3, 27, 445, 8, 27, 1, 28,
		1, 28, 1, 28, 5, 28, 450, 8, 28, 10, 28, 12, 28, 453, 9, 28, 1, 28, 3,
		28, 456, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29,
		465, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 472, 8, 29, 10, 29,
		12, 29, 475, 9, 29, 1, 29, 0, 3, 36, 40, 58, 30, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 0, 1, 2, 0, 30, 30, 51, 51, 534, 0, 67, 1, 0, 0, 0,
		2, 69, 1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 94, 1, 0, 0, 0, 8, 96, 1, 0, 0,
		0, 10, 105, 1, 0, 0, 0, 12, 108, 1, 0, 0, 0, 14, 126, 1, 0, 0, 0, 16, 137,
		1, 0, 0, 0, 18, 152, 1, 0, 0, 0, 20, 154, 1, 0, 0, 0, 22, 166, 1, 0, 0,
		0, 24, 182, 1, 0, 0, 0, 26, 204, 1, 0, 0, 0, 28, 206, 1, 0, 0, 0, 30, 210,
		1, 0, 0, 0, 32, 214, 1, 0, 0, 0, 34, 221, 1, 0, 0, 0, 36, 233, 1, 0, 0,
		0, 38, 243, 1, 0, 0, 0, 40, 346, 1, 0, 0, 0, 42, 415, 1, 0, 0, 0, 44, 419,
		1, 0, 0, 0, 46, 425, 1, 0, 0, 0, 48, 427, 1, 0, 0, 0, 50, 429, 1, 0, 0,
		0, 52, 440, 1, 0, 0, 0, 54, 444, 1, 0, 0, 0, 56, 455, 1, 0, 0, 0, 58, 464,
		1, 0, 0, 0, 60, 62, 3, 4, 2, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0,
		63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 68, 1, 0, 0, 0, 65, 63, 1,
		0, 0, 0, 66, 68, 5, 0, 0, 1, 67, 63, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68,
		1, 1, 0, 0, 0, 69, 70, 7, 0, 0, 0, 70, 3, 1, 0, 0, 0, 71, 73, 3, 6, 3,
		0, 72, 74, 5, 30, 0, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76,
		1, 0, 0, 0, 75, 77, 5, 51, 0, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0,
		77, 81, 1, 0, 0, 0, 78, 81, 5, 30, 0, 0, 79, 81, 5, 51, 0, 0, 80, 71, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82,
		95, 3, 10, 5, 0, 83, 95, 3, 12, 6, 0, 84, 95, 3, 30, 15, 0, 85, 95, 3,
		20, 10, 0, 86, 95, 3, 14, 7, 0, 87, 95, 3, 22, 11, 0, 88, 95, 3, 24, 12,
		0, 89, 95, 3, 26, 13, 0, 90, 95, 3, 28, 14, 0, 91, 95, 3, 32, 16, 0, 92,
		95, 3, 40, 20, 0, 93, 95, 3, 8, 4, 0, 94, 82, 1, 0, 0, 0, 94, 83, 1, 0,
		0, 0, 94, 84, 1, 0, 0, 0, 94, 85, 1, 0, 0, 0, 94, 86, 1, 0, 0, 0, 94, 87,
		1, 0, 0, 0, 94, 88, 1, 0, 0, 0, 94, 89, 1, 0, 0, 0, 94, 90, 1, 0, 0, 0,
		94, 91, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 7, 1, 0,
		0, 0, 96, 100, 5, 27, 0, 0, 97, 99, 3, 4, 2, 0, 98, 97, 1, 0, 0, 0, 99,
		102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1,
		0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 104, 5, 28, 0, 0, 104, 9, 1, 0, 0,
		0, 105, 106, 5, 10, 0, 0, 106, 107, 3, 42, 21, 0, 107, 11, 1, 0, 0, 0,
		108, 109, 5, 9, 0, 0, 109, 110, 3, 42, 21, 0, 110, 13, 1, 0, 0, 0, 111,
		112, 5, 1, 0, 0, 112, 113, 5, 52, 0, 0, 113, 114, 5, 29, 0, 0, 114, 127,
		3, 58, 29, 0, 115, 116, 5, 1, 0, 0, 116, 117, 5, 52, 0, 0, 117, 118, 5,
		33, 0, 0, 118, 127, 3, 40, 20, 0, 119, 120, 5, 1, 0, 0, 120, 121, 5, 52,
		0, 0, 121, 122, 5, 29, 0, 0, 122, 123, 3, 58, 29, 0, 123, 124, 5, 33, 0,
		0, 124, 125, 3, 40, 20, 0, 125, 127, 1, 0, 0, 0, 126, 111, 1, 0, 0, 0,
		126, 115, 1, 0, 0, 0, 126, 119, 1, 0, 0, 0, 127, 15, 1, 0, 0, 0, 128, 129,
		3, 40, 20, 0, 129, 130, 5, 31, 0, 0, 130, 132, 1, 0, 0, 0, 131, 128, 1,
		0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0,
		0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 138, 3, 40, 20, 0,
		137, 133, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 17, 1, 0, 0, 0, 139, 140,
		5, 52, 0, 0, 140, 141, 5, 29, 0, 0, 141, 142, 3, 58, 29, 0, 142, 143, 5,
		31, 0, 0, 143, 145, 1, 0, 0, 0, 144, 139, 1, 0, 0, 0, 145, 148, 1, 0, 0,
		0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148,
		146, 1, 0, 0, 0, 149, 150, 5, 52, 0, 0, 150, 151, 5, 29, 0, 0, 151, 153,
		3, 58, 29, 0, 152, 146, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 19, 1, 0,
		0, 0, 154, 155, 5, 7, 0, 0, 155, 156, 5, 52, 0, 0, 156, 157, 5, 23, 0,
		0, 157, 158, 3, 18, 9, 0, 158, 162, 5, 24, 0, 0, 159, 160, 5, 40, 0, 0,
		160, 161, 5, 34, 0, 0, 161, 163, 3, 58, 29, 0, 162, 159, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 3, 8, 4, 0, 165, 21, 1,
		0, 0, 0, 166, 167, 5, 3, 0, 0, 167, 169, 5, 23, 0, 0, 168, 170, 3, 6, 3,
		0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171,
		173, 5, 30, 0, 0, 172, 174, 3, 6, 3, 0, 173, 172, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 5, 30, 0, 0, 176, 178, 3, 6,
		3, 0, 177, 176, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0,
		179, 180, 5, 24, 0, 0, 180, 181, 3, 4, 2, 0, 181, 23, 1, 0, 0, 0, 182,
		183, 5, 4, 0, 0, 183, 184, 5, 23, 0, 0, 184, 185, 3, 40, 20, 0, 185, 186,
		5, 24, 0, 0, 186, 187, 3, 4, 2, 0, 187, 25, 1, 0, 0, 0, 188, 189, 5, 5,
		0, 0, 189, 190, 5, 23, 0, 0, 190, 191, 3, 40, 20, 0, 191, 192, 5, 24, 0,
		0, 192, 193, 3, 4, 2, 0, 193, 194, 6, 13, -1, 0, 194, 205, 1, 0, 0, 0,
		195, 196, 5, 5, 0, 0, 196, 197, 5, 23, 0, 0, 197, 198, 3, 40, 20, 0, 198,
		199, 5, 24, 0, 0, 199, 200, 3, 4, 2, 0, 200, 201, 5, 6, 0, 0, 201, 202,
		3, 4, 2, 0, 202, 203, 6, 13, -1, 0, 203, 205, 1, 0, 0, 0, 204, 188, 1,
		0, 0, 0, 204, 195, 1, 0, 0, 0, 205, 27, 1, 0, 0, 0, 206, 208, 5, 8, 0,
		0, 207, 209, 3, 4, 2, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		29, 1, 0, 0, 0, 210, 211, 5, 11, 0, 0, 211, 212, 3, 42, 21, 0, 212, 213,
		3, 8, 4, 0, 213, 31, 1, 0, 0, 0, 214, 215, 5, 21, 0, 0, 215, 217, 3, 8,
		4, 0, 216, 218, 3, 2, 1, 0, 217, 216, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0,
		218, 219, 1, 0, 0, 0, 219, 220, 3, 34, 17, 0, 220, 33, 1, 0, 0, 0, 221,
		222, 5, 22, 0, 0, 222, 223, 5, 23, 0, 0, 223, 224, 3, 14, 7, 0, 224, 225,
		5, 24, 0, 0, 225, 35, 1, 0, 0, 0, 226, 227, 6, 18, -1, 0, 227, 228, 5,
		52, 0, 0, 228, 229, 5, 23, 0, 0, 229, 230, 3, 16, 8, 0, 230, 231, 5, 24,
		0, 0, 231, 234, 1, 0, 0, 0, 232, 234, 5, 52, 0, 0, 233, 226, 1, 0, 0, 0,
		233, 232, 1, 0, 0, 0, 234, 240, 1, 0, 0, 0, 235, 236, 10, 1, 0, 0, 236,
		237, 5, 48, 0, 0, 237, 239, 3, 36, 18, 2, 238, 235, 1, 0, 0, 0, 239, 242,
		1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 37, 1, 0,
		0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 3, 36, 18, 0, 244, 39, 1, 0, 0, 0,
		245, 246, 6, 20, -1, 0, 246, 247, 5, 35, 0, 0, 247, 248, 3, 58, 29, 0,
		248, 249, 5, 34, 0, 0, 249, 250, 5, 23, 0, 0, 250, 251, 3, 38, 19, 0, 251,
		252, 5, 24, 0, 0, 252, 347, 1, 0, 0, 0, 253, 254, 5, 23, 0, 0, 254, 255,
		3, 40, 20, 0, 255, 256, 5, 24, 0, 0, 256, 347, 1, 0, 0, 0, 257, 258, 5,
		44, 0, 0, 258, 347, 3, 40, 20, 44, 259, 260, 5, 39, 0, 0, 260, 347, 3,
		40, 20, 43, 261, 262, 5, 40, 0, 0, 262, 347, 3, 40, 20, 42, 263, 264, 5,
		39, 0, 0, 264, 265, 5, 39, 0, 0, 265, 347, 3, 38, 19, 0, 266, 267, 3, 38,
		19, 0, 267, 268, 5, 39, 0, 0, 268, 269, 5, 39, 0, 0, 269, 347, 1, 0, 0,
		0, 270, 271, 5, 40, 0, 0, 271, 272, 5, 40, 0, 0, 272, 347, 3, 38, 19, 0,
		273, 274, 3, 38, 19, 0, 274, 275, 5, 40, 0, 0, 275, 276, 5, 40, 0, 0, 276,
		347, 1, 0, 0, 0, 277, 347, 3, 36, 18, 0, 278, 279, 3, 38, 19, 0, 279, 280,
		5, 33, 0, 0, 280, 281, 3, 40, 20, 17, 281, 347, 1, 0, 0, 0, 282, 283, 3,
		38, 19, 0, 283, 284, 5, 39, 0, 0, 284, 285, 5, 33, 0, 0, 285, 286, 3, 40,
		20, 16, 286, 347, 1, 0, 0, 0, 287, 288, 3, 38, 19, 0, 288, 289, 5, 40,
		0, 0, 289, 290, 5, 33, 0, 0, 290, 291, 3, 40, 20, 15, 291, 347, 1, 0, 0,
		0, 292, 293, 3, 38, 19, 0, 293, 294, 5, 41, 0, 0, 294, 295, 5, 33, 0, 0,
		295, 296, 3, 40, 20, 14, 296, 347, 1, 0, 0, 0, 297, 298, 3, 38, 19, 0,
		298, 299, 5, 42, 0, 0, 299, 300, 5, 33, 0, 0, 300, 301, 3, 40, 20, 13,
		301, 347, 1, 0, 0, 0, 302, 303, 3, 38, 19, 0, 303, 304, 5, 43, 0, 0, 304,
		305, 5, 33, 0, 0, 305, 306, 3, 40, 20, 12, 306, 347, 1, 0, 0, 0, 307, 308,
		3, 38, 19, 0, 308, 309, 5, 39, 0, 0, 309, 310, 5, 33, 0, 0, 310, 311, 3,
		40, 20, 11, 311, 347, 1, 0, 0, 0, 312, 313, 3, 38, 19, 0, 313, 314, 5,
		45, 0, 0, 314, 315, 5, 33, 0, 0, 315, 316, 3, 40, 20, 10, 316, 347, 1,
		0, 0, 0, 317, 318, 3, 38, 19, 0, 318, 319, 5, 46, 0, 0, 319, 320, 5, 33,
		0, 0, 320, 321, 3, 40, 20, 9, 321, 347, 1, 0, 0, 0, 322, 323, 3, 38, 19,
		0, 323, 324, 5, 47, 0, 0, 324, 325, 5, 33, 0, 0, 325, 326, 3, 40, 20, 8,
		326, 347, 1, 0, 0, 0, 327, 328, 3, 38, 19, 0, 328, 329, 5, 36, 0, 0, 329,
		330, 5, 33, 0, 0, 330, 331, 3, 40, 20, 7, 331, 347, 1, 0, 0, 0, 332, 333,
		3, 38, 19, 0, 333, 334, 5, 37, 0, 0, 334, 335, 5, 33, 0, 0, 335, 336, 3,
		40, 20, 6, 336, 347, 1, 0, 0, 0, 337, 338, 3, 38, 19, 0, 338, 339, 5, 38,
		0, 0, 339, 340, 5, 33, 0, 0, 340, 341, 3, 40, 20, 5, 341, 347, 1, 0, 0,
		0, 342, 347, 3, 48, 24, 0, 343, 347, 3, 44, 22, 0, 344, 347, 3, 46, 23,
		0, 345, 347, 5, 20, 0, 0, 346, 245, 1, 0, 0, 0, 346, 253, 1, 0, 0, 0, 346,
		257, 1, 0, 0, 0, 346, 259, 1, 0, 0, 0, 346, 261, 1, 0, 0, 0, 346, 263,
		1, 0, 0, 0, 346, 266, 1, 0, 0, 0, 346, 270, 1, 0, 0, 0, 346, 273, 1, 0,
		0, 0, 346, 277, 1, 0, 0, 0, 346, 278, 1, 0, 0, 0, 346, 282, 1, 0, 0, 0,
		346, 287, 1, 0, 0, 0, 346, 292, 1, 0, 0, 0, 346, 297, 1, 0, 0, 0, 346,
		302, 1, 0, 0, 0, 346, 307, 1, 0, 0, 0, 346, 312, 1, 0, 0, 0, 346, 317,
		1, 0, 0, 0, 346, 322, 1, 0, 0, 0, 346, 327, 1, 0, 0, 0, 346, 332, 1, 0,
		0, 0, 346, 337, 1, 0, 0, 0, 346, 342, 1, 0, 0, 0, 346, 343, 1, 0, 0, 0,
		346, 344, 1, 0, 0, 0, 346, 345, 1, 0, 0, 0, 347, 412, 1, 0, 0, 0, 348,
		349, 10, 36, 0, 0, 349, 350, 5, 41, 0, 0, 350, 411, 3, 40, 20, 37, 351,
		352, 10, 35, 0, 0, 352, 353, 5, 42, 0, 0, 353, 411, 3, 40, 20, 36, 354,
		355, 10, 34, 0, 0, 355, 356, 5, 43, 0, 0, 356, 411, 3, 40, 20, 35, 357,
		358, 10, 33, 0, 0, 358, 359, 5, 39, 0, 0, 359, 411, 3, 40, 20, 34, 360,
		361, 10, 32, 0, 0, 361, 362, 5, 40, 0, 0, 362, 411, 3, 40, 20, 33, 363,
		364, 10, 31, 0, 0, 364, 365, 5, 36, 0, 0, 365, 411, 3, 40, 20, 32, 366,
		367, 10, 30, 0, 0, 367, 368, 5, 37, 0, 0, 368, 411, 3, 40, 20, 31, 369,
		370, 10, 29, 0, 0, 370, 371, 5, 38, 0, 0, 371, 411, 3, 40, 20, 30, 372,
		373, 10, 28, 0, 0, 373, 374, 5, 34, 0, 0, 374, 411, 3, 40, 20, 29, 375,
		376, 10, 27, 0, 0, 376, 377, 5, 35, 0, 0, 377, 411, 3, 40, 20, 28, 378,
		379, 10, 26, 0, 0, 379, 380, 5, 34, 0, 0, 380, 381, 5, 33, 0, 0, 381, 411,
		3, 40, 20, 27, 382, 383, 10, 25, 0, 0, 383, 384, 5, 35, 0, 0, 384, 385,
		5, 33, 0, 0, 385, 411, 3, 40, 20, 26, 386, 387, 10, 24, 0, 0, 387, 388,
		5, 32, 0, 0, 388, 411, 3, 40, 20, 25, 389, 390, 10, 23, 0, 0, 390, 391,
		5, 44, 0, 0, 391, 392, 5, 33, 0, 0, 392, 411, 3, 40, 20, 24, 393, 394,
		10, 22, 0, 0, 394, 395, 5, 45, 0, 0, 395, 411, 3, 40, 20, 23, 396, 397,
		10, 21, 0, 0, 397, 398, 5, 47, 0, 0, 398, 411, 3, 40, 20, 22, 399, 400,
		10, 20, 0, 0, 400, 401, 5, 46, 0, 0, 401, 411, 3, 40, 20, 21, 402, 403,
		10, 19, 0, 0, 403, 404, 5, 45, 0, 0, 404, 405, 5, 45, 0, 0, 405, 411, 3,
		40, 20, 20, 406, 407, 10, 18, 0, 0, 407, 408, 5, 46, 0, 0, 408, 409, 5,
		46, 0, 0, 409, 411, 3, 40, 20, 19, 410, 348, 1, 0, 0, 0, 410, 351, 1, 0,
		0, 0, 410, 354, 1, 0, 0, 0, 410, 357, 1, 0, 0, 0, 410, 360, 1, 0, 0, 0,
		410, 363, 1, 0, 0, 0, 410, 366, 1, 0, 0, 0, 410, 369, 1, 0, 0, 0, 410,
		372, 1, 0, 0, 0, 410, 375, 1, 0, 0, 0, 410, 378, 1, 0, 0, 0, 410, 382,
		1, 0, 0, 0, 410, 386, 1, 0, 0, 0, 410, 389, 1, 0, 0, 0, 410, 393, 1, 0,
		0, 0, 410, 396, 1, 0, 0, 0, 410, 399, 1, 0, 0, 0, 410, 402, 1, 0, 0, 0,
		410, 406, 1, 0, 0, 0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412,
		413, 1, 0, 0, 0, 413, 41, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 415, 416, 5,
		52, 0, 0, 416, 43, 1, 0, 0, 0, 417, 420, 3, 52, 26, 0, 418, 420, 3, 50,
		25, 0, 419, 417, 1, 0, 0, 0, 419, 418, 1, 0, 0, 0, 420, 45, 1, 0, 0, 0,
		421, 422, 5, 16, 0, 0, 422, 426, 6, 23, -1, 0, 423, 424, 5, 17, 0, 0, 424,
		426, 6, 23, -1, 0, 425, 421, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426, 47,
		1, 0, 0, 0, 427, 428, 5, 63, 0, 0, 428, 49, 1, 0, 0, 0, 429, 430, 5, 57,
		0, 0, 430, 431, 6, 25, -1, 0, 431, 51, 1, 0, 0, 0, 432, 433, 5, 54, 0,
		0, 433, 441, 6, 26, -1, 0, 434, 435, 5, 53, 0, 0, 435, 441, 6, 26, -1,
		0, 436, 437, 5, 56, 0, 0, 437, 441, 6, 26, -1, 0, 438, 439, 5, 55, 0, 0,
		439, 441, 6, 26, -1, 0, 440, 432, 1, 0, 0, 0, 440, 434, 1, 0, 0, 0, 440,
		436, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 441, 53, 1, 0, 0, 0, 442, 445, 3,
		40, 20, 0, 443, 445, 3, 58, 29, 0, 444, 442, 1, 0, 0, 0, 444, 443, 1, 0,
		0, 0, 445, 55, 1, 0, 0, 0, 446, 447, 3, 54, 27, 0, 447, 448, 5, 31, 0,
		0, 448, 450, 1, 0, 0, 0, 449, 446, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451,
		449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 454, 1, 0, 0, 0, 453, 451,
		1, 0, 0, 0, 454, 456, 3, 54, 27, 0, 455, 451, 1, 0, 0, 0, 455, 456, 1,
		0, 0, 0, 456, 57, 1, 0, 0, 0, 457, 458, 6, 29, -1, 0, 458, 465, 3, 42,
		21, 0, 459, 460, 3, 42, 21, 0, 460, 461, 5, 35, 0, 0, 461, 462, 3, 56,
		28, 0, 462, 463, 5, 34, 0, 0, 463, 465, 1, 0, 0, 0, 464, 457, 1, 0, 0,
		0, 464, 459, 1, 0, 0, 0, 465, 473, 1, 0, 0, 0, 466, 467, 10, 1, 0, 0, 467,
		468, 5, 25, 0, 0, 468, 469, 3, 40, 20, 0, 469, 470, 5, 26, 0, 0, 470, 472,
		1, 0, 0, 0, 471, 466, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473, 471, 1, 0,
		0, 0, 473, 474, 1, 0, 0, 0, 474, 59, 1, 0, 0, 0, 475, 473, 1, 0, 0, 0,
		32, 63, 67, 73, 76, 80, 94, 100, 126, 133, 137, 146, 152, 162, 169, 173,
		177, 204, 208, 217, 233, 240, 346, 410, 412, 419, 425, 440, 444, 451, 455,
		464, 473,
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

// RiddleParserInit initializes any static state used to implement RiddleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRiddleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RiddleParserInit() {
	staticData := &RiddleParserParserStaticData
	staticData.once.Do(riddleparserParserInit)
}

// NewRiddleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRiddleParser(input antlr.TokenStream) *RiddleParser {
	RiddleParserInit()
	this := new(RiddleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RiddleParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RiddleParser.g4"

	return this
}

// RiddleParser tokens.
const (
	RiddleParserEOF             = antlr.TokenEOF
	RiddleParserVar             = 1
	RiddleParserVal             = 2
	RiddleParserFor             = 3
	RiddleParserWhile           = 4
	RiddleParserIf              = 5
	RiddleParserElse            = 6
	RiddleParserFunc            = 7
	RiddleParserReturn          = 8
	RiddleParserImport          = 9
	RiddleParserPackage         = 10
	RiddleParserClass           = 11
	RiddleParserPublic          = 12
	RiddleParserProtected       = 13
	RiddleParserPrivate         = 14
	RiddleParserOverride        = 15
	RiddleParserTrue            = 16
	RiddleParserFalse           = 17
	RiddleParserStatic          = 18
	RiddleParserConst           = 19
	RiddleParserNull            = 20
	RiddleParserTry             = 21
	RiddleParserCatch           = 22
	RiddleParserLeftBracket     = 23
	RiddleParserRightBracket    = 24
	RiddleParserLeftSquare      = 25
	RiddleParserRightSquare     = 26
	RiddleParserLeftCurly       = 27
	RiddleParserRightCurly      = 28
	RiddleParserColon           = 29
	RiddleParserSemi            = 30
	RiddleParserComma           = 31
	RiddleParserEqual           = 32
	RiddleParserAssign          = 33
	RiddleParserGreater         = 34
	RiddleParserLess            = 35
	RiddleParserLeftLeft        = 36
	RiddleParserRightRight      = 37
	RiddleParserRightRightRight = 38
	RiddleParserAdd             = 39
	RiddleParserSub             = 40
	RiddleParserStar            = 41
	RiddleParserDiv             = 42
	RiddleParserMod             = 43
	RiddleParserNot             = 44
	RiddleParserAnd             = 45
	RiddleParserOr              = 46
	RiddleParserXor             = 47
	RiddleParserDot             = 48
	RiddleParserDoubleQuotes    = 49
	RiddleParserQuotes          = 50
	RiddleParserEndl            = 51
	RiddleParserIdentifier      = 52
	RiddleParserHexadecimal     = 53
	RiddleParserDecimal         = 54
	RiddleParserOctal           = 55
	RiddleParserBinary          = 56
	RiddleParserFloat           = 57
	RiddleParserIntegerSequence = 58
	RiddleParserHEX_DIGIT       = 59
	RiddleParserOCTAL_DIGIT     = 60
	RiddleParserBINARY_DIGIT    = 61
	RiddleParserDIGIT           = 62
	RiddleParserSTRING          = 63
	RiddleParserLINE_COMMENT    = 64
	RiddleParserBLOCK_COMMENT   = 65
	RiddleParserWHITESPACE      = 66
)

// RiddleParser rules.
const (
	RiddleParserRULE_program            = 0
	RiddleParserRULE_null_cnt           = 1
	RiddleParserRULE_statement_ed       = 2
	RiddleParserRULE_statement          = 3
	RiddleParserRULE_bodyExpr           = 4
	RiddleParserRULE_packStatement      = 5
	RiddleParserRULE_importStatement    = 6
	RiddleParserRULE_varDefineStatement = 7
	RiddleParserRULE_argsExpr           = 8
	RiddleParserRULE_defineArgs         = 9
	RiddleParserRULE_funcDefine         = 10
	RiddleParserRULE_forStatement       = 11
	RiddleParserRULE_whileStatement     = 12
	RiddleParserRULE_ifStatement        = 13
	RiddleParserRULE_returnStatement    = 14
	RiddleParserRULE_classDefine        = 15
	RiddleParserRULE_tryExpr            = 16
	RiddleParserRULE_catchExpr          = 17
	RiddleParserRULE_exprPtr            = 18
	RiddleParserRULE_exprPtrParser      = 19
	RiddleParserRULE_expression         = 20
	RiddleParserRULE_id                 = 21
	RiddleParserRULE_number             = 22
	RiddleParserRULE_boolean            = 23
	RiddleParserRULE_string             = 24
	RiddleParserRULE_float              = 25
	RiddleParserRULE_integer            = 26
	RiddleParserRULE_templateArg        = 27
	RiddleParserRULE_templateArgs       = 28
	RiddleParserRULE_typeName           = 29
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement_ed() []IStatement_edContext
	Statement_ed(i int) IStatement_edContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement_ed() []IStatement_edContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatement_edContext); ok {
			len++
		}
	}

	tst := make([]IStatement_edContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatement_edContext); ok {
			tst[i] = t.(IStatement_edContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement_ed(i int) IStatement_edContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_edContext); ok {
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

	return t.(IStatement_edContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(RiddleParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RiddleParserRULE_program)
	var _la int

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8937374183483830342) != 0 {
			{
				p.SetState(60)
				p.Statement_ed()
			}

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(RiddleParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// INull_cntContext is an interface to support dynamic dispatch.
type INull_cntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Semi() antlr.TerminalNode
	Endl() antlr.TerminalNode

	// IsNull_cntContext differentiates from other interfaces.
	IsNull_cntContext()
}

type Null_cntContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNull_cntContext() *Null_cntContext {
	var p = new(Null_cntContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_null_cnt
	return p
}

func InitEmptyNull_cntContext(p *Null_cntContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_null_cnt
}

func (*Null_cntContext) IsNull_cntContext() {}

func NewNull_cntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Null_cntContext {
	var p = new(Null_cntContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_null_cnt

	return p
}

func (s *Null_cntContext) GetParser() antlr.Parser { return s.parser }

func (s *Null_cntContext) Semi() antlr.TerminalNode {
	return s.GetToken(RiddleParserSemi, 0)
}

func (s *Null_cntContext) Endl() antlr.TerminalNode {
	return s.GetToken(RiddleParserEndl, 0)
}

func (s *Null_cntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Null_cntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Null_cntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNull_cnt(s)
	}
}

func (s *Null_cntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNull_cnt(s)
	}
}

func (s *Null_cntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNull_cnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Null_cnt() (localctx INull_cntContext) {
	localctx = NewNull_cntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RiddleParserRULE_null_cnt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RiddleParserSemi || _la == RiddleParserEndl) {
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

// IStatement_edContext is an interface to support dynamic dispatch.
type IStatement_edContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	Semi() antlr.TerminalNode
	Endl() antlr.TerminalNode

	// IsStatement_edContext differentiates from other interfaces.
	IsStatement_edContext()
}

type Statement_edContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_edContext() *Statement_edContext {
	var p = new(Statement_edContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_statement_ed
	return p
}

func InitEmptyStatement_edContext(p *Statement_edContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_statement_ed
}

func (*Statement_edContext) IsStatement_edContext() {}

func NewStatement_edContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_edContext {
	var p = new(Statement_edContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_statement_ed

	return p
}

func (s *Statement_edContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_edContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_edContext) Semi() antlr.TerminalNode {
	return s.GetToken(RiddleParserSemi, 0)
}

func (s *Statement_edContext) Endl() antlr.TerminalNode {
	return s.GetToken(RiddleParserEndl, 0)
}

func (s *Statement_edContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_edContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_edContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterStatement_ed(s)
	}
}

func (s *Statement_edContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitStatement_ed(s)
	}
}

func (s *Statement_edContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitStatement_ed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Statement_ed() (localctx IStatement_edContext) {
	localctx = NewStatement_edContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RiddleParserRULE_statement_ed)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RiddleParserVar, RiddleParserFor, RiddleParserWhile, RiddleParserIf, RiddleParserFunc, RiddleParserReturn, RiddleParserImport, RiddleParserPackage, RiddleParserClass, RiddleParserTrue, RiddleParserFalse, RiddleParserNull, RiddleParserTry, RiddleParserLeftBracket, RiddleParserLeftCurly, RiddleParserLess, RiddleParserAdd, RiddleParserSub, RiddleParserNot, RiddleParserIdentifier, RiddleParserHexadecimal, RiddleParserDecimal, RiddleParserOctal, RiddleParserBinary, RiddleParserFloat, RiddleParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Statement()
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(72)
				p.Match(RiddleParserSemi)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(75)
				p.Match(RiddleParserEndl)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case RiddleParserSemi:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Match(RiddleParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RiddleParserEndl:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Match(RiddleParserEndl)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PackStatement() IPackStatementContext
	ImportStatement() IImportStatementContext
	ClassDefine() IClassDefineContext
	FuncDefine() IFuncDefineContext
	VarDefineStatement() IVarDefineStatementContext
	ForStatement() IForStatementContext
	WhileStatement() IWhileStatementContext
	IfStatement() IIfStatementContext
	ReturnStatement() IReturnStatementContext
	TryExpr() ITryExprContext
	Expression() IExpressionContext
	BodyExpr() IBodyExprContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) PackStatement() IPackStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackStatementContext)
}

func (s *StatementContext) ImportStatement() IImportStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *StatementContext) ClassDefine() IClassDefineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDefineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDefineContext)
}

func (s *StatementContext) FuncDefine() IFuncDefineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefineContext)
}

func (s *StatementContext) VarDefineStatement() IVarDefineStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDefineStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDefineStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) TryExpr() ITryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryExprContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) BodyExpr() IBodyExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyExprContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RiddleParserRULE_statement)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RiddleParserPackage:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.PackStatement()
		}

	case RiddleParserImport:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.ImportStatement()
		}

	case RiddleParserClass:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.ClassDefine()
		}

	case RiddleParserFunc:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.FuncDefine()
		}

	case RiddleParserVar:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)
			p.VarDefineStatement()
		}

	case RiddleParserFor:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(87)
			p.ForStatement()
		}

	case RiddleParserWhile:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(88)
			p.WhileStatement()
		}

	case RiddleParserIf:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(89)
			p.IfStatement()
		}

	case RiddleParserReturn:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(90)
			p.ReturnStatement()
		}

	case RiddleParserTry:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(91)
			p.TryExpr()
		}

	case RiddleParserTrue, RiddleParserFalse, RiddleParserNull, RiddleParserLeftBracket, RiddleParserLess, RiddleParserAdd, RiddleParserSub, RiddleParserNot, RiddleParserIdentifier, RiddleParserHexadecimal, RiddleParserDecimal, RiddleParserOctal, RiddleParserBinary, RiddleParserFloat, RiddleParserSTRING:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(92)
			p.expression(0)
		}

	case RiddleParserLeftCurly:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(93)
			p.BodyExpr()
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

// IBodyExprContext is an interface to support dynamic dispatch.
type IBodyExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftCurly() antlr.TerminalNode
	RightCurly() antlr.TerminalNode
	AllStatement_ed() []IStatement_edContext
	Statement_ed(i int) IStatement_edContext

	// IsBodyExprContext differentiates from other interfaces.
	IsBodyExprContext()
}

type BodyExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyExprContext() *BodyExprContext {
	var p = new(BodyExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_bodyExpr
	return p
}

func InitEmptyBodyExprContext(p *BodyExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_bodyExpr
}

func (*BodyExprContext) IsBodyExprContext() {}

func NewBodyExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyExprContext {
	var p = new(BodyExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_bodyExpr

	return p
}

func (s *BodyExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyExprContext) LeftCurly() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftCurly, 0)
}

func (s *BodyExprContext) RightCurly() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightCurly, 0)
}

func (s *BodyExprContext) AllStatement_ed() []IStatement_edContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatement_edContext); ok {
			len++
		}
	}

	tst := make([]IStatement_edContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatement_edContext); ok {
			tst[i] = t.(IStatement_edContext)
			i++
		}
	}

	return tst
}

func (s *BodyExprContext) Statement_ed(i int) IStatement_edContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_edContext); ok {
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

	return t.(IStatement_edContext)
}

func (s *BodyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBodyExpr(s)
	}
}

func (s *BodyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBodyExpr(s)
	}
}

func (s *BodyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBodyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) BodyExpr() (localctx IBodyExprContext) {
	localctx = NewBodyExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RiddleParserRULE_bodyExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(RiddleParserLeftCurly)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8937374183483830342) != 0 {
		{
			p.SetState(97)
			p.Statement_ed()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(RiddleParserRightCurly)
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

// IPackStatementContext is an interface to support dynamic dispatch.
type IPackStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPackName returns the packName rule contexts.
	GetPackName() IIdContext

	// SetPackName sets the packName rule contexts.
	SetPackName(IIdContext)

	// Getter signatures
	Package() antlr.TerminalNode
	Id() IIdContext

	// IsPackStatementContext differentiates from other interfaces.
	IsPackStatementContext()
}

type PackStatementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	packName IIdContext
}

func NewEmptyPackStatementContext() *PackStatementContext {
	var p = new(PackStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_packStatement
	return p
}

func InitEmptyPackStatementContext(p *PackStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_packStatement
}

func (*PackStatementContext) IsPackStatementContext() {}

func NewPackStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackStatementContext {
	var p = new(PackStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_packStatement

	return p
}

func (s *PackStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackStatementContext) GetPackName() IIdContext { return s.packName }

func (s *PackStatementContext) SetPackName(v IIdContext) { s.packName = v }

func (s *PackStatementContext) Package() antlr.TerminalNode {
	return s.GetToken(RiddleParserPackage, 0)
}

func (s *PackStatementContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PackStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterPackStatement(s)
	}
}

func (s *PackStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitPackStatement(s)
	}
}

func (s *PackStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitPackStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) PackStatement() (localctx IPackStatementContext) {
	localctx = NewPackStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RiddleParserRULE_packStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(RiddleParserPackage)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)

		var _x = p.Id()

		localctx.(*PackStatementContext).packName = _x
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

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLibName returns the libName rule contexts.
	GetLibName() IIdContext

	// SetLibName sets the libName rule contexts.
	SetLibName(IIdContext)

	// Getter signatures
	Import() antlr.TerminalNode
	Id() IIdContext

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	libName IIdContext
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) GetLibName() IIdContext { return s.libName }

func (s *ImportStatementContext) SetLibName(v IIdContext) { s.libName = v }

func (s *ImportStatementContext) Import() antlr.TerminalNode {
	return s.GetToken(RiddleParserImport, 0)
}

func (s *ImportStatementContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RiddleParserRULE_importStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(RiddleParserImport)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)

		var _x = p.Id()

		localctx.(*ImportStatementContext).libName = _x
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

// IVarDefineStatementContext is an interface to support dynamic dispatch.
type IVarDefineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetType_ returns the type_ rule contexts.
	GetType_() ITypeNameContext

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetType_ sets the type_ rule contexts.
	SetType_(ITypeNameContext)

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// Getter signatures
	Var() antlr.TerminalNode
	Colon() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	TypeName() ITypeNameContext
	Assign() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVarDefineStatementContext differentiates from other interfaces.
	IsVarDefineStatementContext()
}

type VarDefineStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	type_  ITypeNameContext
	value  IExpressionContext
}

func NewEmptyVarDefineStatementContext() *VarDefineStatementContext {
	var p = new(VarDefineStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_varDefineStatement
	return p
}

func InitEmptyVarDefineStatementContext(p *VarDefineStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_varDefineStatement
}

func (*VarDefineStatementContext) IsVarDefineStatementContext() {}

func NewVarDefineStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefineStatementContext {
	var p = new(VarDefineStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_varDefineStatement

	return p
}

func (s *VarDefineStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefineStatementContext) GetName() antlr.Token { return s.name }

func (s *VarDefineStatementContext) SetName(v antlr.Token) { s.name = v }

func (s *VarDefineStatementContext) GetType_() ITypeNameContext { return s.type_ }

func (s *VarDefineStatementContext) GetValue() IExpressionContext { return s.value }

func (s *VarDefineStatementContext) SetType_(v ITypeNameContext) { s.type_ = v }

func (s *VarDefineStatementContext) SetValue(v IExpressionContext) { s.value = v }

func (s *VarDefineStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(RiddleParserVar, 0)
}

func (s *VarDefineStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(RiddleParserColon, 0)
}

func (s *VarDefineStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RiddleParserIdentifier, 0)
}

func (s *VarDefineStatementContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *VarDefineStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *VarDefineStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDefineStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefineStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefineStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterVarDefineStatement(s)
	}
}

func (s *VarDefineStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitVarDefineStatement(s)
	}
}

func (s *VarDefineStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitVarDefineStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) VarDefineStatement() (localctx IVarDefineStatementContext) {
	localctx = NewVarDefineStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RiddleParserRULE_varDefineStatement)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(RiddleParserVar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)

			var _m = p.Match(RiddleParserIdentifier)

			localctx.(*VarDefineStatementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(RiddleParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)

			var _x = p.typeName(0)

			localctx.(*VarDefineStatementContext).type_ = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(RiddleParserVar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)

			var _m = p.Match(RiddleParserIdentifier)

			localctx.(*VarDefineStatementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.expression(0)

			localctx.(*VarDefineStatementContext).value = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Match(RiddleParserVar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)

			var _m = p.Match(RiddleParserIdentifier)

			localctx.(*VarDefineStatementContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(RiddleParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)

			var _x = p.typeName(0)

			localctx.(*VarDefineStatementContext).type_ = _x
		}
		{
			p.SetState(123)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)

			var _x = p.expression(0)

			localctx.(*VarDefineStatementContext).value = _x
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

// IArgsExprContext is an interface to support dynamic dispatch.
type IArgsExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArgsExprContext differentiates from other interfaces.
	IsArgsExprContext()
}

type ArgsExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsExprContext() *ArgsExprContext {
	var p = new(ArgsExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_argsExpr
	return p
}

func InitEmptyArgsExprContext(p *ArgsExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_argsExpr
}

func (*ArgsExprContext) IsArgsExprContext() {}

func NewArgsExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsExprContext {
	var p = new(ArgsExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_argsExpr

	return p
}

func (s *ArgsExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgsExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgsExprContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserComma)
}

func (s *ArgsExprContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserComma, i)
}

func (s *ArgsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterArgsExpr(s)
	}
}

func (s *ArgsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitArgsExpr(s)
	}
}

func (s *ArgsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitArgsExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ArgsExpr() (localctx IArgsExprContext) {
	localctx = NewArgsExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RiddleParserRULE_argsExpr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8939625984507576320) != 0 {
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(128)
					p.expression(0)
				}
				{
					p.SetState(129)
					p.Match(RiddleParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.expression(0)
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

// IDefineArgsContext is an interface to support dynamic dispatch.
type IDefineArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllColon() []antlr.TerminalNode
	Colon(i int) antlr.TerminalNode
	AllTypeName() []ITypeNameContext
	TypeName(i int) ITypeNameContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsDefineArgsContext differentiates from other interfaces.
	IsDefineArgsContext()
}

type DefineArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineArgsContext() *DefineArgsContext {
	var p = new(DefineArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_defineArgs
	return p
}

func InitEmptyDefineArgsContext(p *DefineArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_defineArgs
}

func (*DefineArgsContext) IsDefineArgsContext() {}

func NewDefineArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineArgsContext {
	var p = new(DefineArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_defineArgs

	return p
}

func (s *DefineArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineArgsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserIdentifier)
}

func (s *DefineArgsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserIdentifier, i)
}

func (s *DefineArgsContext) AllColon() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserColon)
}

func (s *DefineArgsContext) Colon(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserColon, i)
}

func (s *DefineArgsContext) AllTypeName() []ITypeNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeNameContext); ok {
			len++
		}
	}

	tst := make([]ITypeNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeNameContext); ok {
			tst[i] = t.(ITypeNameContext)
			i++
		}
	}

	return tst
}

func (s *DefineArgsContext) TypeName(i int) ITypeNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
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

	return t.(ITypeNameContext)
}

func (s *DefineArgsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserComma)
}

func (s *DefineArgsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserComma, i)
}

func (s *DefineArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterDefineArgs(s)
	}
}

func (s *DefineArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitDefineArgs(s)
	}
}

func (s *DefineArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitDefineArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) DefineArgs() (localctx IDefineArgsContext) {
	localctx = NewDefineArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RiddleParserRULE_defineArgs)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RiddleParserIdentifier {
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(139)
					p.Match(RiddleParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(140)
					p.Match(RiddleParserColon)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(141)
					p.typeName(0)
				}
				{
					p.SetState(142)
					p.Match(RiddleParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		{
			p.SetState(149)
			p.Match(RiddleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(RiddleParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.typeName(0)
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

// IFuncDefineContext is an interface to support dynamic dispatch.
type IFuncDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFuncName returns the funcName token.
	GetFuncName() antlr.Token

	// SetFuncName sets the funcName token.
	SetFuncName(antlr.Token)

	// GetArgs returns the args rule contexts.
	GetArgs() IDefineArgsContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeNameContext

	// GetBody returns the body rule contexts.
	GetBody() IBodyExprContext

	// SetArgs sets the args rule contexts.
	SetArgs(IDefineArgsContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeNameContext)

	// SetBody sets the body rule contexts.
	SetBody(IBodyExprContext)

	// Getter signatures
	Func() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	DefineArgs() IDefineArgsContext
	BodyExpr() IBodyExprContext
	Sub() antlr.TerminalNode
	Greater() antlr.TerminalNode
	TypeName() ITypeNameContext

	// IsFuncDefineContext differentiates from other interfaces.
	IsFuncDefineContext()
}

type FuncDefineContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	funcName   antlr.Token
	args       IDefineArgsContext
	returnType ITypeNameContext
	body       IBodyExprContext
}

func NewEmptyFuncDefineContext() *FuncDefineContext {
	var p = new(FuncDefineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_funcDefine
	return p
}

func InitEmptyFuncDefineContext(p *FuncDefineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_funcDefine
}

func (*FuncDefineContext) IsFuncDefineContext() {}

func NewFuncDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefineContext {
	var p = new(FuncDefineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_funcDefine

	return p
}

func (s *FuncDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefineContext) GetFuncName() antlr.Token { return s.funcName }

func (s *FuncDefineContext) SetFuncName(v antlr.Token) { s.funcName = v }

func (s *FuncDefineContext) GetArgs() IDefineArgsContext { return s.args }

func (s *FuncDefineContext) GetReturnType() ITypeNameContext { return s.returnType }

func (s *FuncDefineContext) GetBody() IBodyExprContext { return s.body }

func (s *FuncDefineContext) SetArgs(v IDefineArgsContext) { s.args = v }

func (s *FuncDefineContext) SetReturnType(v ITypeNameContext) { s.returnType = v }

func (s *FuncDefineContext) SetBody(v IBodyExprContext) { s.body = v }

func (s *FuncDefineContext) Func() antlr.TerminalNode {
	return s.GetToken(RiddleParserFunc, 0)
}

func (s *FuncDefineContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *FuncDefineContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *FuncDefineContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RiddleParserIdentifier, 0)
}

func (s *FuncDefineContext) DefineArgs() IDefineArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefineArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefineArgsContext)
}

func (s *FuncDefineContext) BodyExpr() IBodyExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyExprContext)
}

func (s *FuncDefineContext) Sub() antlr.TerminalNode {
	return s.GetToken(RiddleParserSub, 0)
}

func (s *FuncDefineContext) Greater() antlr.TerminalNode {
	return s.GetToken(RiddleParserGreater, 0)
}

func (s *FuncDefineContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *FuncDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterFuncDefine(s)
	}
}

func (s *FuncDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitFuncDefine(s)
	}
}

func (s *FuncDefineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitFuncDefine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) FuncDefine() (localctx IFuncDefineContext) {
	localctx = NewFuncDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RiddleParserRULE_funcDefine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(RiddleParserFunc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)

		var _m = p.Match(RiddleParserIdentifier)

		localctx.(*FuncDefineContext).funcName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(RiddleParserLeftBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)

		var _x = p.DefineArgs()

		localctx.(*FuncDefineContext).args = _x
	}
	{
		p.SetState(158)
		p.Match(RiddleParserRightBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RiddleParserSub {
		{
			p.SetState(159)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(RiddleParserGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)

			var _x = p.typeName(0)

			localctx.(*FuncDefineContext).returnType = _x
		}

	}
	{
		p.SetState(164)

		var _x = p.BodyExpr()

		localctx.(*FuncDefineContext).body = _x
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

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IStatementContext

	// GetTermCond returns the termCond rule contexts.
	GetTermCond() IStatementContext

	// GetSelfVar returns the selfVar rule contexts.
	GetSelfVar() IStatementContext

	// GetBody returns the body rule contexts.
	GetBody() IStatement_edContext

	// SetInit sets the init rule contexts.
	SetInit(IStatementContext)

	// SetTermCond sets the termCond rule contexts.
	SetTermCond(IStatementContext)

	// SetSelfVar sets the selfVar rule contexts.
	SetSelfVar(IStatementContext)

	// SetBody sets the body rule contexts.
	SetBody(IStatement_edContext)

	// Getter signatures
	For() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	Statement_ed() IStatement_edContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	init     IStatementContext
	termCond IStatementContext
	selfVar  IStatementContext
	body     IStatement_edContext
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) GetInit() IStatementContext { return s.init }

func (s *ForStatementContext) GetTermCond() IStatementContext { return s.termCond }

func (s *ForStatementContext) GetSelfVar() IStatementContext { return s.selfVar }

func (s *ForStatementContext) GetBody() IStatement_edContext { return s.body }

func (s *ForStatementContext) SetInit(v IStatementContext) { s.init = v }

func (s *ForStatementContext) SetTermCond(v IStatementContext) { s.termCond = v }

func (s *ForStatementContext) SetSelfVar(v IStatementContext) { s.selfVar = v }

func (s *ForStatementContext) SetBody(v IStatement_edContext) { s.body = v }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(RiddleParserFor, 0)
}

func (s *ForStatementContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *ForStatementContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserSemi)
}

func (s *ForStatementContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserSemi, i)
}

func (s *ForStatementContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *ForStatementContext) Statement_ed() IStatement_edContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_edContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_edContext)
}

func (s *ForStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RiddleParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(RiddleParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(RiddleParserLeftBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8939625984371257414) != 0 {
		{
			p.SetState(168)

			var _x = p.Statement()

			localctx.(*ForStatementContext).init = _x
		}

	}
	{
		p.SetState(171)
		p.Match(RiddleParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8939625984371257414) != 0 {
		{
			p.SetState(172)

			var _x = p.Statement()

			localctx.(*ForStatementContext).termCond = _x
		}

	}
	{
		p.SetState(175)
		p.Match(RiddleParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8939625984371257414) != 0 {
		{
			p.SetState(176)

			var _x = p.Statement()

			localctx.(*ForStatementContext).selfVar = _x
		}

	}
	{
		p.SetState(179)
		p.Match(RiddleParserRightBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)

		var _x = p.Statement_ed()

		localctx.(*ForStatementContext).body = _x
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

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRunCond returns the runCond rule contexts.
	GetRunCond() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IStatement_edContext

	// SetRunCond sets the runCond rule contexts.
	SetRunCond(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IStatement_edContext)

	// Getter signatures
	While() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	Expression() IExpressionContext
	Statement_ed() IStatement_edContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	runCond IExpressionContext
	body    IStatement_edContext
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) GetRunCond() IExpressionContext { return s.runCond }

func (s *WhileStatementContext) GetBody() IStatement_edContext { return s.body }

func (s *WhileStatementContext) SetRunCond(v IExpressionContext) { s.runCond = v }

func (s *WhileStatementContext) SetBody(v IStatement_edContext) { s.body = v }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(RiddleParserWhile, 0)
}

func (s *WhileStatementContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *WhileStatementContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Statement_ed() IStatement_edContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_edContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_edContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RiddleParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(RiddleParserWhile)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(RiddleParserLeftBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)

		var _x = p.expression(0)

		localctx.(*WhileStatementContext).runCond = _x
	}
	{
		p.SetState(185)
		p.Match(RiddleParserRightBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)

		var _x = p.Statement_ed()

		localctx.(*WhileStatementContext).body = _x
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IStatement_edContext

	// GetElseBody returns the elseBody rule contexts.
	GetElseBody() IStatement_edContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IStatement_edContext)

	// SetElseBody sets the elseBody rule contexts.
	SetElseBody(IStatement_edContext)

	// GetHasElse returns the hasElse attribute.
	GetHasElse() bool

	// SetHasElse sets the hasElse attribute.
	SetHasElse(bool)

	// Getter signatures
	If() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	RightBracket() antlr.TerminalNode
	Expression() IExpressionContext
	AllStatement_ed() []IStatement_edContext
	Statement_ed(i int) IStatement_edContext
	Else() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	hasElse  bool
	cond     IExpressionContext
	body     IStatement_edContext
	elseBody IStatement_edContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetCond() IExpressionContext { return s.cond }

func (s *IfStatementContext) GetBody() IStatement_edContext { return s.body }

func (s *IfStatementContext) GetElseBody() IStatement_edContext { return s.elseBody }

func (s *IfStatementContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *IfStatementContext) SetBody(v IStatement_edContext) { s.body = v }

func (s *IfStatementContext) SetElseBody(v IStatement_edContext) { s.elseBody = v }

func (s *IfStatementContext) GetHasElse() bool { return s.hasElse }

func (s *IfStatementContext) SetHasElse(v bool) { s.hasElse = v }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(RiddleParserIf, 0)
}

func (s *IfStatementContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *IfStatementContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllStatement_ed() []IStatement_edContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatement_edContext); ok {
			len++
		}
	}

	tst := make([]IStatement_edContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatement_edContext); ok {
			tst[i] = t.(IStatement_edContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement_ed(i int) IStatement_edContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_edContext); ok {
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

	return t.(IStatement_edContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(RiddleParserElse, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RiddleParserRULE_ifStatement)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(RiddleParserIf)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(RiddleParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)

			var _x = p.expression(0)

			localctx.(*IfStatementContext).cond = _x
		}
		{
			p.SetState(191)
			p.Match(RiddleParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)

			var _x = p.Statement_ed()

			localctx.(*IfStatementContext).body = _x
		}
		localctx.(*IfStatementContext).hasElse = false

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(RiddleParserIf)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(RiddleParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)

			var _x = p.expression(0)

			localctx.(*IfStatementContext).cond = _x
		}
		{
			p.SetState(198)
			p.Match(RiddleParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)

			var _x = p.Statement_ed()

			localctx.(*IfStatementContext).body = _x
		}
		{
			p.SetState(200)
			p.Match(RiddleParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)

			var _x = p.Statement_ed()

			localctx.(*IfStatementContext).elseBody = _x
		}
		localctx.(*IfStatementContext).hasElse = true

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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetResult returns the result rule contexts.
	GetResult() IStatement_edContext

	// SetResult sets the result rule contexts.
	SetResult(IStatement_edContext)

	// Getter signatures
	Return() antlr.TerminalNode
	Statement_ed() IStatement_edContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result IStatement_edContext
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) GetResult() IStatement_edContext { return s.result }

func (s *ReturnStatementContext) SetResult(v IStatement_edContext) { s.result = v }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(RiddleParserReturn, 0)
}

func (s *ReturnStatementContext) Statement_ed() IStatement_edContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_edContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_edContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RiddleParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(RiddleParserReturn)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(207)

			var _x = p.Statement_ed()

			localctx.(*ReturnStatementContext).result = _x
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

// IClassDefineContext is an interface to support dynamic dispatch.
type IClassDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetClassName returns the className rule contexts.
	GetClassName() IIdContext

	// GetBody returns the body rule contexts.
	GetBody() IBodyExprContext

	// SetClassName sets the className rule contexts.
	SetClassName(IIdContext)

	// SetBody sets the body rule contexts.
	SetBody(IBodyExprContext)

	// Getter signatures
	Class() antlr.TerminalNode
	Id() IIdContext
	BodyExpr() IBodyExprContext

	// IsClassDefineContext differentiates from other interfaces.
	IsClassDefineContext()
}

type ClassDefineContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	className IIdContext
	body      IBodyExprContext
}

func NewEmptyClassDefineContext() *ClassDefineContext {
	var p = new(ClassDefineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_classDefine
	return p
}

func InitEmptyClassDefineContext(p *ClassDefineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_classDefine
}

func (*ClassDefineContext) IsClassDefineContext() {}

func NewClassDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefineContext {
	var p = new(ClassDefineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_classDefine

	return p
}

func (s *ClassDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefineContext) GetClassName() IIdContext { return s.className }

func (s *ClassDefineContext) GetBody() IBodyExprContext { return s.body }

func (s *ClassDefineContext) SetClassName(v IIdContext) { s.className = v }

func (s *ClassDefineContext) SetBody(v IBodyExprContext) { s.body = v }

func (s *ClassDefineContext) Class() antlr.TerminalNode {
	return s.GetToken(RiddleParserClass, 0)
}

func (s *ClassDefineContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ClassDefineContext) BodyExpr() IBodyExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyExprContext)
}

func (s *ClassDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterClassDefine(s)
	}
}

func (s *ClassDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitClassDefine(s)
	}
}

func (s *ClassDefineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitClassDefine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ClassDefine() (localctx IClassDefineContext) {
	localctx = NewClassDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RiddleParserRULE_classDefine)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(RiddleParserClass)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)

		var _x = p.Id()

		localctx.(*ClassDefineContext).className = _x
	}
	{
		p.SetState(212)

		var _x = p.BodyExpr()

		localctx.(*ClassDefineContext).body = _x
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

// ITryExprContext is an interface to support dynamic dispatch.
type ITryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTryBody returns the tryBody rule contexts.
	GetTryBody() IBodyExprContext

	// SetTryBody sets the tryBody rule contexts.
	SetTryBody(IBodyExprContext)

	// Getter signatures
	Try() antlr.TerminalNode
	CatchExpr() ICatchExprContext
	BodyExpr() IBodyExprContext
	Null_cnt() INull_cntContext

	// IsTryExprContext differentiates from other interfaces.
	IsTryExprContext()
}

type TryExprContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	tryBody IBodyExprContext
}

func NewEmptyTryExprContext() *TryExprContext {
	var p = new(TryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_tryExpr
	return p
}

func InitEmptyTryExprContext(p *TryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_tryExpr
}

func (*TryExprContext) IsTryExprContext() {}

func NewTryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryExprContext {
	var p = new(TryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_tryExpr

	return p
}

func (s *TryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TryExprContext) GetTryBody() IBodyExprContext { return s.tryBody }

func (s *TryExprContext) SetTryBody(v IBodyExprContext) { s.tryBody = v }

func (s *TryExprContext) Try() antlr.TerminalNode {
	return s.GetToken(RiddleParserTry, 0)
}

func (s *TryExprContext) CatchExpr() ICatchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatchExprContext)
}

func (s *TryExprContext) BodyExpr() IBodyExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyExprContext)
}

func (s *TryExprContext) Null_cnt() INull_cntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INull_cntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INull_cntContext)
}

func (s *TryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterTryExpr(s)
	}
}

func (s *TryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitTryExpr(s)
	}
}

func (s *TryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitTryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) TryExpr() (localctx ITryExprContext) {
	localctx = NewTryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RiddleParserRULE_tryExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(RiddleParserTry)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)

		var _x = p.BodyExpr()

		localctx.(*TryExprContext).tryBody = _x
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RiddleParserSemi || _la == RiddleParserEndl {
		{
			p.SetState(216)
			p.Null_cnt()
		}

	}
	{
		p.SetState(219)
		p.CatchExpr()
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

// ICatchExprContext is an interface to support dynamic dispatch.
type ICatchExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Catch() antlr.TerminalNode
	LeftBracket() antlr.TerminalNode
	VarDefineStatement() IVarDefineStatementContext
	RightBracket() antlr.TerminalNode

	// IsCatchExprContext differentiates from other interfaces.
	IsCatchExprContext()
}

type CatchExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatchExprContext() *CatchExprContext {
	var p = new(CatchExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_catchExpr
	return p
}

func InitEmptyCatchExprContext(p *CatchExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_catchExpr
}

func (*CatchExprContext) IsCatchExprContext() {}

func NewCatchExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchExprContext {
	var p = new(CatchExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_catchExpr

	return p
}

func (s *CatchExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchExprContext) Catch() antlr.TerminalNode {
	return s.GetToken(RiddleParserCatch, 0)
}

func (s *CatchExprContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *CatchExprContext) VarDefineStatement() IVarDefineStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDefineStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDefineStatementContext)
}

func (s *CatchExprContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *CatchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterCatchExpr(s)
	}
}

func (s *CatchExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitCatchExpr(s)
	}
}

func (s *CatchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitCatchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) CatchExpr() (localctx ICatchExprContext) {
	localctx = NewCatchExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RiddleParserRULE_catchExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(RiddleParserCatch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(RiddleParserLeftBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.VarDefineStatement()
	}
	{
		p.SetState(224)
		p.Match(RiddleParserRightBracket)
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

// IExprPtrContext is an interface to support dynamic dispatch.
type IExprPtrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprPtrContext differentiates from other interfaces.
	IsExprPtrContext()
}

type ExprPtrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprPtrContext() *ExprPtrContext {
	var p = new(ExprPtrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_exprPtr
	return p
}

func InitEmptyExprPtrContext(p *ExprPtrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_exprPtr
}

func (*ExprPtrContext) IsExprPtrContext() {}

func NewExprPtrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprPtrContext {
	var p = new(ExprPtrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_exprPtr

	return p
}

func (s *ExprPtrContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprPtrContext) CopyAll(ctx *ExprPtrContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprPtrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPtrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncExprContext struct {
	ExprPtrContext
	funcName antlr.Token
	args     IArgsExprContext
}

func NewFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncExprContext {
	var p = new(FuncExprContext)

	InitEmptyExprPtrContext(&p.ExprPtrContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprPtrContext))

	return p
}

func (s *FuncExprContext) GetFuncName() antlr.Token { return s.funcName }

func (s *FuncExprContext) SetFuncName(v antlr.Token) { s.funcName = v }

func (s *FuncExprContext) GetArgs() IArgsExprContext { return s.args }

func (s *FuncExprContext) SetArgs(v IArgsExprContext) { s.args = v }

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *FuncExprContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *FuncExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RiddleParserIdentifier, 0)
}

func (s *FuncExprContext) ArgsExpr() IArgsExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsExprContext)
}

func (s *FuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterFuncExpr(s)
	}
}

func (s *FuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitFuncExpr(s)
	}
}

func (s *FuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ObjectExprContext struct {
	ExprPtrContext
}

func NewObjectExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectExprContext {
	var p = new(ObjectExprContext)

	InitEmptyExprPtrContext(&p.ExprPtrContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprPtrContext))

	return p
}

func (s *ObjectExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RiddleParserIdentifier, 0)
}

func (s *ObjectExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterObjectExpr(s)
	}
}

func (s *ObjectExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitObjectExpr(s)
	}
}

func (s *ObjectExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitObjectExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlendExprContext struct {
	ExprPtrContext
	parents IExprPtrContext
	childs  IExprPtrContext
}

func NewBlendExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlendExprContext {
	var p = new(BlendExprContext)

	InitEmptyExprPtrContext(&p.ExprPtrContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprPtrContext))

	return p
}

func (s *BlendExprContext) GetParents() IExprPtrContext { return s.parents }

func (s *BlendExprContext) GetChilds() IExprPtrContext { return s.childs }

func (s *BlendExprContext) SetParents(v IExprPtrContext) { s.parents = v }

func (s *BlendExprContext) SetChilds(v IExprPtrContext) { s.childs = v }

func (s *BlendExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlendExprContext) Dot() antlr.TerminalNode {
	return s.GetToken(RiddleParserDot, 0)
}

func (s *BlendExprContext) AllExprPtr() []IExprPtrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprPtrContext); ok {
			len++
		}
	}

	tst := make([]IExprPtrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprPtrContext); ok {
			tst[i] = t.(IExprPtrContext)
			i++
		}
	}

	return tst
}

func (s *BlendExprContext) ExprPtr(i int) IExprPtrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrContext); ok {
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

	return t.(IExprPtrContext)
}

func (s *BlendExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBlendExpr(s)
	}
}

func (s *BlendExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBlendExpr(s)
	}
}

func (s *BlendExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBlendExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ExprPtr() (localctx IExprPtrContext) {
	return p.exprPtr(0)
}

func (p *RiddleParser) exprPtr(_p int) (localctx IExprPtrContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprPtrContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprPtrContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, RiddleParserRULE_exprPtr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(227)

			var _m = p.Match(RiddleParserIdentifier)

			localctx.(*FuncExprContext).funcName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(RiddleParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.ArgsExpr()

			localctx.(*FuncExprContext).args = _x
		}
		{
			p.SetState(230)
			p.Match(RiddleParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewObjectExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(RiddleParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBlendExprContext(p, NewExprPtrContext(p, _parentctx, _parentState))
			localctx.(*BlendExprContext).parents = _prevctx

			p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_exprPtr)
			p.SetState(235)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(236)
				p.Match(RiddleParserDot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(237)

				var _x = p.exprPtr(2)

				localctx.(*BlendExprContext).childs = _x
			}

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprPtrParserContext is an interface to support dynamic dispatch.
type IExprPtrParserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprPtr() IExprPtrContext

	// IsExprPtrParserContext differentiates from other interfaces.
	IsExprPtrParserContext()
}

type ExprPtrParserContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprPtrParserContext() *ExprPtrParserContext {
	var p = new(ExprPtrParserContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_exprPtrParser
	return p
}

func InitEmptyExprPtrParserContext(p *ExprPtrParserContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_exprPtrParser
}

func (*ExprPtrParserContext) IsExprPtrParserContext() {}

func NewExprPtrParserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprPtrParserContext {
	var p = new(ExprPtrParserContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_exprPtrParser

	return p
}

func (s *ExprPtrParserContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprPtrParserContext) ExprPtr() IExprPtrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrContext)
}

func (s *ExprPtrParserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPtrParserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprPtrParserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterExprPtrParser(s)
	}
}

func (s *ExprPtrParserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitExprPtrParser(s)
	}
}

func (s *ExprPtrParserContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitExprPtrParser(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) ExprPtrParser() (localctx IExprPtrParserContext) {
	localctx = NewExprPtrParserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RiddleParserRULE_exprPtrParser)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.exprPtr(0)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelfSubRightExprContext struct {
	ExpressionContext
	expr IExprPtrParserContext
}

func NewSelfSubRightExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfSubRightExprContext {
	var p = new(SelfSubRightExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelfSubRightExprContext) GetExpr() IExprPtrParserContext { return s.expr }

func (s *SelfSubRightExprContext) SetExpr(v IExprPtrParserContext) { s.expr = v }

func (s *SelfSubRightExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfSubRightExprContext) AllSub() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserSub)
}

func (s *SelfSubRightExprContext) Sub(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserSub, i)
}

func (s *SelfSubRightExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *SelfSubRightExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterSelfSubRightExpr(s)
	}
}

func (s *SelfSubRightExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitSelfSubRightExpr(s)
	}
}

func (s *SelfSubRightExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitSelfSubRightExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewAndAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndAssignExprContext {
	var p = new(AndAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *AndAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *AndAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *AndAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AndAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndAssignExprContext) And() antlr.TerminalNode {
	return s.GetToken(RiddleParserAnd, 0)
}

func (s *AndAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *AndAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *AndAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAndAssignExpr(s)
	}
}

func (s *AndAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAndAssignExpr(s)
	}
}

func (s *AndAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAndAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullExprContext struct {
	ExpressionContext
}

func NewNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprContext {
	var p = new(NullExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExprContext) Null() antlr.TerminalNode {
	return s.GetToken(RiddleParserNull, 0)
}

func (s *NullExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNullExpr(s)
	}
}

func (s *NullExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNullExpr(s)
	}
}

func (s *NullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExprContext {
	var p = new(ModExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ModExprContext) GetLeft() IExpressionContext { return s.left }

func (s *ModExprContext) GetRight() IExpressionContext { return s.right }

func (s *ModExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ModExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExprContext) Mod() antlr.TerminalNode {
	return s.GetToken(RiddleParserMod, 0)
}

func (s *ModExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ModExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ModExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterModExpr(s)
	}
}

func (s *ModExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitModExpr(s)
	}
}

func (s *ModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitModExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CastExprContext struct {
	ExpressionContext
	type_ ITypeNameContext
	value IExprPtrParserContext
}

func NewCastExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExprContext {
	var p = new(CastExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CastExprContext) GetType_() ITypeNameContext { return s.type_ }

func (s *CastExprContext) GetValue() IExprPtrParserContext { return s.value }

func (s *CastExprContext) SetType_(v ITypeNameContext) { s.type_ = v }

func (s *CastExprContext) SetValue(v IExprPtrParserContext) { s.value = v }

func (s *CastExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExprContext) Less() antlr.TerminalNode {
	return s.GetToken(RiddleParserLess, 0)
}

func (s *CastExprContext) Greater() antlr.TerminalNode {
	return s.GetToken(RiddleParserGreater, 0)
}

func (s *CastExprContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *CastExprContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *CastExprContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CastExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *CastExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterCastExpr(s)
	}
}

func (s *CastExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitCastExpr(s)
	}
}

func (s *CastExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitCastExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LShrExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewLShrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LShrExprContext {
	var p = new(LShrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LShrExprContext) GetLeft() IExpressionContext { return s.left }

func (s *LShrExprContext) GetRight() IExpressionContext { return s.right }

func (s *LShrExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *LShrExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LShrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LShrExprContext) RightRightRight() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightRightRight, 0)
}

func (s *LShrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LShrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LShrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterLShrExpr(s)
	}
}

func (s *LShrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitLShrExpr(s)
	}
}

func (s *LShrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitLShrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegativeExprContext struct {
	ExpressionContext
	expr IExpressionContext
}

func NewNegativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegativeExprContext {
	var p = new(NegativeExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NegativeExprContext) GetExpr() IExpressionContext { return s.expr }

func (s *NegativeExprContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *NegativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeExprContext) Sub() antlr.TerminalNode {
	return s.GetToken(RiddleParserSub, 0)
}

func (s *NegativeExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNegativeExpr(s)
	}
}

func (s *NegativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNegativeExpr(s)
	}
}

func (s *NegativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNegativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExprContext struct {
	ExpressionContext
}

func NewNumberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExprContext {
	var p = new(NumberExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *NumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNumberExpr(s)
	}
}

func (s *NumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNumberExpr(s)
	}
}

func (s *NumberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNumberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewMulAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulAssignExprContext {
	var p = new(MulAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *MulAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *MulAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *MulAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *MulAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulAssignExprContext) Star() antlr.TerminalNode {
	return s.GetToken(RiddleParserStar, 0)
}

func (s *MulAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *MulAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *MulAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterMulAssignExpr(s)
	}
}

func (s *MulAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitMulAssignExpr(s)
	}
}

func (s *MulAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitMulAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type XorAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewXorAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XorAssignExprContext {
	var p = new(XorAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *XorAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *XorAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *XorAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *XorAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *XorAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorAssignExprContext) Xor() antlr.TerminalNode {
	return s.GetToken(RiddleParserXor, 0)
}

func (s *XorAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *XorAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *XorAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *XorAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterXorAssignExpr(s)
	}
}

func (s *XorAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitXorAssignExpr(s)
	}
}

func (s *XorAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitXorAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SubExprContext) GetLeft() IExpressionContext { return s.left }

func (s *SubExprContext) GetRight() IExpressionContext { return s.right }

func (s *SubExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *SubExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) Sub() antlr.TerminalNode {
	return s.GetToken(RiddleParserSub, 0)
}

func (s *SubExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SubExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

func (s *SubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterEqualExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewGreaterEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterEqualExprContext {
	var p = new(GreaterEqualExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *GreaterEqualExprContext) GetLeft() IExpressionContext { return s.left }

func (s *GreaterEqualExprContext) GetRight() IExpressionContext { return s.right }

func (s *GreaterEqualExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *GreaterEqualExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *GreaterEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterEqualExprContext) Greater() antlr.TerminalNode {
	return s.GetToken(RiddleParserGreater, 0)
}

func (s *GreaterEqualExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *GreaterEqualExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GreaterEqualExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *GreaterEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterGreaterEqualExpr(s)
	}
}

func (s *GreaterEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitGreaterEqualExpr(s)
	}
}

func (s *GreaterEqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitGreaterEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewAddAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddAssignExprContext {
	var p = new(AddAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *AddAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *AddAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *AddAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AddAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddAssignExprContext) Add() antlr.TerminalNode {
	return s.GetToken(RiddleParserAdd, 0)
}

func (s *AddAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *AddAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *AddAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAddAssignExpr(s)
	}
}

func (s *AddAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAddAssignExpr(s)
	}
}

func (s *AddAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAddAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewOrAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrAssignExprContext {
	var p = new(OrAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *OrAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *OrAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *OrAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OrAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrAssignExprContext) Or() antlr.TerminalNode {
	return s.GetToken(RiddleParserOr, 0)
}

func (s *OrAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *OrAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *OrAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterOrAssignExpr(s)
	}
}

func (s *OrAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitOrAssignExpr(s)
	}
}

func (s *OrAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitOrAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitXorExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewBitXorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXorExprContext {
	var p = new(BitXorExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitXorExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BitXorExprContext) GetRight() IExpressionContext { return s.right }

func (s *BitXorExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BitXorExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BitXorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorExprContext) Xor() antlr.TerminalNode {
	return s.GetToken(RiddleParserXor, 0)
}

func (s *BitXorExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitXorExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BitXorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBitXorExpr(s)
	}
}

func (s *BitXorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBitXorExpr(s)
	}
}

func (s *BitXorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBitXorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PtrExprContext struct {
	ExpressionContext
}

func NewPtrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PtrExprContext {
	var p = new(PtrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PtrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PtrExprContext) ExprPtr() IExprPtrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrContext)
}

func (s *PtrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterPtrExpr(s)
	}
}

func (s *PtrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitPtrExpr(s)
	}
}

func (s *PtrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitPtrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExprContext struct {
	ExpressionContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewGreaterExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterExprContext {
	var p = new(GreaterExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *GreaterExprContext) GetLeft() IExpressionContext { return s.left }

func (s *GreaterExprContext) GetRight() IExpressionContext { return s.right }

func (s *GreaterExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *GreaterExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *GreaterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterExprContext) Greater() antlr.TerminalNode {
	return s.GetToken(RiddleParserGreater, 0)
}

func (s *GreaterExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GreaterExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *GreaterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterGreaterExpr(s)
	}
}

func (s *GreaterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitGreaterExpr(s)
	}
}

func (s *GreaterExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitGreaterExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewDivAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivAssignExprContext {
	var p = new(DivAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DivAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *DivAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *DivAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *DivAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *DivAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivAssignExprContext) Div() antlr.TerminalNode {
	return s.GetToken(RiddleParserDiv, 0)
}

func (s *DivAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *DivAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *DivAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DivAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterDivAssignExpr(s)
	}
}

func (s *DivAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitDivAssignExpr(s)
	}
}

func (s *DivAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitDivAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetLeft() IExpressionContext { return s.left }

func (s *MulExprContext) GetRight() IExpressionContext { return s.right }

func (s *MulExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *MulExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) Star() antlr.TerminalNode {
	return s.GetToken(RiddleParserStar, 0)
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotEqualExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewNotEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualExprContext {
	var p = new(NotEqualExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotEqualExprContext) GetLeft() IExpressionContext { return s.left }

func (s *NotEqualExprContext) GetRight() IExpressionContext { return s.right }

func (s *NotEqualExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *NotEqualExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *NotEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualExprContext) Not() antlr.TerminalNode {
	return s.GetToken(RiddleParserNot, 0)
}

func (s *NotEqualExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *NotEqualExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *NotEqualExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *NotEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNotEqualExpr(s)
	}
}

func (s *NotEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNotEqualExpr(s)
	}
}

func (s *NotEqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNotEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfSubLeftExprContext struct {
	ExpressionContext
	expr IExprPtrParserContext
}

func NewSelfSubLeftExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfSubLeftExprContext {
	var p = new(SelfSubLeftExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelfSubLeftExprContext) GetExpr() IExprPtrParserContext { return s.expr }

func (s *SelfSubLeftExprContext) SetExpr(v IExprPtrParserContext) { s.expr = v }

func (s *SelfSubLeftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfSubLeftExprContext) AllSub() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserSub)
}

func (s *SelfSubLeftExprContext) Sub(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserSub, i)
}

func (s *SelfSubLeftExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *SelfSubLeftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterSelfSubLeftExpr(s)
	}
}

func (s *SelfSubLeftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitSelfSubLeftExpr(s)
	}
}

func (s *SelfSubLeftExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitSelfSubLeftExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExprContext {
	var p = new(DivExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DivExprContext) GetLeft() IExpressionContext { return s.left }

func (s *DivExprContext) GetRight() IExpressionContext { return s.right }

func (s *DivExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *DivExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *DivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExprContext) Div() antlr.TerminalNode {
	return s.GetToken(RiddleParserDiv, 0)
}

func (s *DivExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DivExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *DivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterDivExpr(s)
	}
}

func (s *DivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitDivExpr(s)
	}
}

func (s *DivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitAndExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewBitAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExprContext {
	var p = new(BitAndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitAndExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BitAndExprContext) GetRight() IExpressionContext { return s.right }

func (s *BitAndExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BitAndExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BitAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExprContext) And() antlr.TerminalNode {
	return s.GetToken(RiddleParserAnd, 0)
}

func (s *BitAndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitAndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BitAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBitAndExpr(s)
	}
}

func (s *BitAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBitAndExpr(s)
	}
}

func (s *BitAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *AssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *AssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *AssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *AssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *AssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfAddRightExprContext struct {
	ExpressionContext
	expr IExprPtrParserContext
}

func NewSelfAddRightExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfAddRightExprContext {
	var p = new(SelfAddRightExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelfAddRightExprContext) GetExpr() IExprPtrParserContext { return s.expr }

func (s *SelfAddRightExprContext) SetExpr(v IExprPtrParserContext) { s.expr = v }

func (s *SelfAddRightExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfAddRightExprContext) AllAdd() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserAdd)
}

func (s *SelfAddRightExprContext) Add(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserAdd, i)
}

func (s *SelfAddRightExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *SelfAddRightExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterSelfAddRightExpr(s)
	}
}

func (s *SelfAddRightExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitSelfAddRightExpr(s)
	}
}

func (s *SelfAddRightExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitSelfAddRightExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewSubAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubAssignExprContext {
	var p = new(SubAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SubAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *SubAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *SubAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *SubAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *SubAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAssignExprContext) Sub() antlr.TerminalNode {
	return s.GetToken(RiddleParserSub, 0)
}

func (s *SubAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *SubAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *SubAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterSubAssignExpr(s)
	}
}

func (s *SubAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitSubAssignExpr(s)
	}
}

func (s *SubAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitSubAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketExprContext struct {
	ExpressionContext
	expr IExpressionContext
}

func NewBracketExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExprContext {
	var p = new(BracketExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BracketExprContext) GetExpr() IExpressionContext { return s.expr }

func (s *BracketExprContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *BracketExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExprContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftBracket, 0)
}

func (s *BracketExprContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightBracket, 0)
}

func (s *BracketExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBracketExpr(s)
	}
}

func (s *BracketExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBracketExpr(s)
	}
}

func (s *BracketExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBracketExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanExprContext struct {
	ExpressionContext
}

func NewBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExprContext {
	var p = new(BooleanExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) Boolean() IBooleanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanContext)
}

func (s *BooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBooleanExpr(s)
	}
}

func (s *BooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBooleanExpr(s)
	}
}

func (s *BooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AShrExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewAShrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AShrExprContext {
	var p = new(AShrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AShrExprContext) GetLeft() IExpressionContext { return s.left }

func (s *AShrExprContext) GetRight() IExpressionContext { return s.right }

func (s *AShrExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AShrExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AShrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AShrExprContext) RightRight() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightRight, 0)
}

func (s *AShrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AShrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AShrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAShrExpr(s)
	}
}

func (s *AShrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAShrExpr(s)
	}
}

func (s *AShrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAShrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewLessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessExprContext {
	var p = new(LessExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LessExprContext) GetLeft() IExpressionContext { return s.left }

func (s *LessExprContext) GetRight() IExpressionContext { return s.right }

func (s *LessExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *LessExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessExprContext) Less() antlr.TerminalNode {
	return s.GetToken(RiddleParserLess, 0)
}

func (s *LessExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LessExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterLessExpr(s)
	}
}

func (s *LessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitLessExpr(s)
	}
}

func (s *LessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitLessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetLeft() IExpressionContext { return s.left }

func (s *OrExprContext) GetRight() IExpressionContext { return s.right }

func (s *OrExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *OrExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserOr)
}

func (s *OrExprContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserOr, i)
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfAddLeftExprContext struct {
	ExpressionContext
	expr IExprPtrParserContext
}

func NewSelfAddLeftExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfAddLeftExprContext {
	var p = new(SelfAddLeftExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelfAddLeftExprContext) GetExpr() IExprPtrParserContext { return s.expr }

func (s *SelfAddLeftExprContext) SetExpr(v IExprPtrParserContext) { s.expr = v }

func (s *SelfAddLeftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfAddLeftExprContext) AllAdd() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserAdd)
}

func (s *SelfAddLeftExprContext) Add(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserAdd, i)
}

func (s *SelfAddLeftExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *SelfAddLeftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterSelfAddLeftExpr(s)
	}
}

func (s *SelfAddLeftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitSelfAddLeftExpr(s)
	}
}

func (s *SelfAddLeftExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitSelfAddLeftExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LShrAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewLShrAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LShrAssignExprContext {
	var p = new(LShrAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LShrAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *LShrAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *LShrAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *LShrAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LShrAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LShrAssignExprContext) RightRightRight() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightRightRight, 0)
}

func (s *LShrAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *LShrAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *LShrAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LShrAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterLShrAssignExpr(s)
	}
}

func (s *LShrAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitLShrAssignExpr(s)
	}
}

func (s *LShrAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitLShrAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessEqualExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewLessEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessEqualExprContext {
	var p = new(LessEqualExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LessEqualExprContext) GetLeft() IExpressionContext { return s.left }

func (s *LessEqualExprContext) GetRight() IExpressionContext { return s.right }

func (s *LessEqualExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *LessEqualExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LessEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessEqualExprContext) Less() antlr.TerminalNode {
	return s.GetToken(RiddleParserLess, 0)
}

func (s *LessEqualExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *LessEqualExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LessEqualExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LessEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterLessEqualExpr(s)
	}
}

func (s *LessEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitLessEqualExpr(s)
	}
}

func (s *LessEqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitLessEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AShrAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewAShrAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AShrAssignExprContext {
	var p = new(AShrAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AShrAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *AShrAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *AShrAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *AShrAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AShrAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AShrAssignExprContext) RightRight() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightRight, 0)
}

func (s *AShrAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *AShrAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *AShrAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AShrAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAShrAssignExpr(s)
	}
}

func (s *AShrAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAShrAssignExpr(s)
	}
}

func (s *AShrAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAShrAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitOrExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewBitOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExprContext {
	var p = new(BitOrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitOrExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BitOrExprContext) GetRight() IExpressionContext { return s.right }

func (s *BitOrExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BitOrExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BitOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExprContext) Or() antlr.TerminalNode {
	return s.GetToken(RiddleParserOr, 0)
}

func (s *BitOrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitOrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BitOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBitOrExpr(s)
	}
}

func (s *BitOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBitOrExpr(s)
	}
}

func (s *BitOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExpressionContext
	expr IExpressionContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExprContext) GetExpr() IExpressionContext { return s.expr }

func (s *NotExprContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Not() antlr.TerminalNode {
	return s.GetToken(RiddleParserNot, 0)
}

func (s *NotExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetLeft() IExpressionContext { return s.left }

func (s *AddExprContext) GetRight() IExpressionContext { return s.right }

func (s *AddExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AddExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) Add() antlr.TerminalNode {
	return s.GetToken(RiddleParserAdd, 0)
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShlAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewShlAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShlAssignExprContext {
	var p = new(ShlAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ShlAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *ShlAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *ShlAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *ShlAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ShlAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShlAssignExprContext) LeftLeft() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftLeft, 0)
}

func (s *ShlAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *ShlAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *ShlAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShlAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterShlAssignExpr(s)
	}
}

func (s *ShlAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitShlAssignExpr(s)
	}
}

func (s *ShlAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitShlAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModAssignExprContext struct {
	ExpressionContext
	left  IExprPtrParserContext
	right IExpressionContext
}

func NewModAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModAssignExprContext {
	var p = new(ModAssignExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ModAssignExprContext) GetLeft() IExprPtrParserContext { return s.left }

func (s *ModAssignExprContext) GetRight() IExpressionContext { return s.right }

func (s *ModAssignExprContext) SetLeft(v IExprPtrParserContext) { s.left = v }

func (s *ModAssignExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ModAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModAssignExprContext) Mod() antlr.TerminalNode {
	return s.GetToken(RiddleParserMod, 0)
}

func (s *ModAssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(RiddleParserAssign, 0)
}

func (s *ModAssignExprContext) ExprPtrParser() IExprPtrParserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprPtrParserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprPtrParserContext)
}

func (s *ModAssignExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterModAssignExpr(s)
	}
}

func (s *ModAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitModAssignExpr(s)
	}
}

func (s *ModAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitModAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PositiveExprContext struct {
	ExpressionContext
	expr IExpressionContext
}

func NewPositiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PositiveExprContext {
	var p = new(PositiveExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PositiveExprContext) GetExpr() IExpressionContext { return s.expr }

func (s *PositiveExprContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *PositiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositiveExprContext) Add() antlr.TerminalNode {
	return s.GetToken(RiddleParserAdd, 0)
}

func (s *PositiveExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PositiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterPositiveExpr(s)
	}
}

func (s *PositiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitPositiveExpr(s)
	}
}

func (s *PositiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitPositiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShlExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewShlExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShlExprContext {
	var p = new(ShlExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ShlExprContext) GetLeft() IExpressionContext { return s.left }

func (s *ShlExprContext) GetRight() IExpressionContext { return s.right }

func (s *ShlExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ShlExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ShlExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShlExprContext) LeftLeft() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftLeft, 0)
}

func (s *ShlExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ShlExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ShlExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterShlExpr(s)
	}
}

func (s *ShlExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitShlExpr(s)
	}
}

func (s *ShlExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitShlExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualExprContext {
	var p = new(EqualExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqualExprContext) GetLeft() IExpressionContext { return s.left }

func (s *EqualExprContext) GetRight() IExpressionContext { return s.right }

func (s *EqualExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *EqualExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *EqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExprContext) Equal() antlr.TerminalNode {
	return s.GetToken(RiddleParserEqual, 0)
}

func (s *EqualExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *EqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterEqualExpr(s)
	}
}

func (s *EqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitEqualExpr(s)
	}
}

func (s *EqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetLeft() IExpressionContext { return s.left }

func (s *AndExprContext) GetRight() IExpressionContext { return s.right }

func (s *AndExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AndExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllAnd() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserAnd)
}

func (s *AndExprContext) And(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserAnd, i)
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *RiddleParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, RiddleParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCastExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(246)
			p.Match(RiddleParserLess)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)

			var _x = p.typeName(0)

			localctx.(*CastExprContext).type_ = _x
		}
		{
			p.SetState(248)
			p.Match(RiddleParserGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Match(RiddleParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)

			var _x = p.ExprPtrParser()

			localctx.(*CastExprContext).value = _x
		}
		{
			p.SetState(251)
			p.Match(RiddleParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewBracketExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(253)
			p.Match(RiddleParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)

			var _x = p.expression(0)

			localctx.(*BracketExprContext).expr = _x
		}
		{
			p.SetState(255)
			p.Match(RiddleParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(257)
			p.Match(RiddleParserNot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)

			var _x = p.expression(44)

			localctx.(*NotExprContext).expr = _x
		}

	case 4:
		localctx = NewPositiveExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _x = p.expression(43)

			localctx.(*PositiveExprContext).expr = _x
		}

	case 5:
		localctx = NewNegativeExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)

			var _x = p.expression(42)

			localctx.(*NegativeExprContext).expr = _x
		}

	case 6:
		localctx = NewSelfAddLeftExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(263)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)

			var _x = p.ExprPtrParser()

			localctx.(*SelfAddLeftExprContext).expr = _x
		}

	case 7:
		localctx = NewSelfAddRightExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)

			var _x = p.ExprPtrParser()

			localctx.(*SelfAddRightExprContext).expr = _x
		}
		{
			p.SetState(267)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewSelfSubLeftExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(270)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)

			var _x = p.ExprPtrParser()

			localctx.(*SelfSubLeftExprContext).expr = _x
		}

	case 9:
		localctx = NewSelfSubRightExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)

			var _x = p.ExprPtrParser()

			localctx.(*SelfSubRightExprContext).expr = _x
		}
		{
			p.SetState(274)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewPtrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.exprPtr(0)
		}

	case 11:
		localctx = NewAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)

			var _x = p.ExprPtrParser()

			localctx.(*AssignExprContext).left = _x
		}
		{
			p.SetState(279)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)

			var _x = p.expression(17)

			localctx.(*AssignExprContext).right = _x
		}

	case 12:
		localctx = NewAddAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(282)

			var _x = p.ExprPtrParser()

			localctx.(*AddAssignExprContext).left = _x
		}
		{
			p.SetState(283)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)

			var _x = p.expression(16)

			localctx.(*AddAssignExprContext).right = _x
		}

	case 13:
		localctx = NewSubAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(287)

			var _x = p.ExprPtrParser()

			localctx.(*SubAssignExprContext).left = _x
		}
		{
			p.SetState(288)
			p.Match(RiddleParserSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)

			var _x = p.expression(15)

			localctx.(*SubAssignExprContext).right = _x
		}

	case 14:
		localctx = NewMulAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(292)

			var _x = p.ExprPtrParser()

			localctx.(*MulAssignExprContext).left = _x
		}
		{
			p.SetState(293)
			p.Match(RiddleParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)

			var _x = p.expression(14)

			localctx.(*MulAssignExprContext).right = _x
		}

	case 15:
		localctx = NewDivAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(297)

			var _x = p.ExprPtrParser()

			localctx.(*DivAssignExprContext).left = _x
		}
		{
			p.SetState(298)
			p.Match(RiddleParserDiv)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)

			var _x = p.expression(13)

			localctx.(*DivAssignExprContext).right = _x
		}

	case 16:
		localctx = NewModAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(302)

			var _x = p.ExprPtrParser()

			localctx.(*ModAssignExprContext).left = _x
		}
		{
			p.SetState(303)
			p.Match(RiddleParserMod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)

			var _x = p.expression(12)

			localctx.(*ModAssignExprContext).right = _x
		}

	case 17:
		localctx = NewAddAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)

			var _x = p.ExprPtrParser()

			localctx.(*AddAssignExprContext).left = _x
		}
		{
			p.SetState(308)
			p.Match(RiddleParserAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)

			var _x = p.expression(11)

			localctx.(*AddAssignExprContext).right = _x
		}

	case 18:
		localctx = NewAndAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(312)

			var _x = p.ExprPtrParser()

			localctx.(*AndAssignExprContext).left = _x
		}
		{
			p.SetState(313)
			p.Match(RiddleParserAnd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)

			var _x = p.expression(10)

			localctx.(*AndAssignExprContext).right = _x
		}

	case 19:
		localctx = NewOrAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(317)

			var _x = p.ExprPtrParser()

			localctx.(*OrAssignExprContext).left = _x
		}
		{
			p.SetState(318)
			p.Match(RiddleParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)

			var _x = p.expression(9)

			localctx.(*OrAssignExprContext).right = _x
		}

	case 20:
		localctx = NewXorAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(322)

			var _x = p.ExprPtrParser()

			localctx.(*XorAssignExprContext).left = _x
		}
		{
			p.SetState(323)
			p.Match(RiddleParserXor)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)

			var _x = p.expression(8)

			localctx.(*XorAssignExprContext).right = _x
		}

	case 21:
		localctx = NewShlAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(327)

			var _x = p.ExprPtrParser()

			localctx.(*ShlAssignExprContext).left = _x
		}
		{
			p.SetState(328)
			p.Match(RiddleParserLeftLeft)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)

			var _x = p.expression(7)

			localctx.(*ShlAssignExprContext).right = _x
		}

	case 22:
		localctx = NewAShrAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(332)

			var _x = p.ExprPtrParser()

			localctx.(*AShrAssignExprContext).left = _x
		}
		{
			p.SetState(333)
			p.Match(RiddleParserRightRight)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)

			var _x = p.expression(6)

			localctx.(*AShrAssignExprContext).right = _x
		}

	case 23:
		localctx = NewLShrAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(337)

			var _x = p.ExprPtrParser()

			localctx.(*LShrAssignExprContext).left = _x
		}
		{
			p.SetState(338)
			p.Match(RiddleParserRightRightRight)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(RiddleParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)

			var _x = p.expression(5)

			localctx.(*LShrAssignExprContext).right = _x
		}

	case 24:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(342)
			p.String_()
		}

	case 25:
		localctx = NewNumberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(343)
			p.Number()
		}

	case 26:
		localctx = NewBooleanExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(344)
			p.Boolean()
		}

	case 27:
		localctx = NewNullExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(345)
			p.Match(RiddleParserNull)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*MulExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(348)

				if !(p.Precpred(p.GetParserRuleContext(), 36)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 36)", ""))
					goto errorExit
				}
				{
					p.SetState(349)
					p.Match(RiddleParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(350)

					var _x = p.expression(37)

					localctx.(*MulExprContext).right = _x
				}

			case 2:
				localctx = NewDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*DivExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 35)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 35)", ""))
					goto errorExit
				}
				{
					p.SetState(352)
					p.Match(RiddleParserDiv)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(353)

					var _x = p.expression(36)

					localctx.(*DivExprContext).right = _x
				}

			case 3:
				localctx = NewModExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ModExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(354)

				if !(p.Precpred(p.GetParserRuleContext(), 34)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 34)", ""))
					goto errorExit
				}
				{
					p.SetState(355)
					p.Match(RiddleParserMod)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(356)

					var _x = p.expression(35)

					localctx.(*ModExprContext).right = _x
				}

			case 4:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(357)

				if !(p.Precpred(p.GetParserRuleContext(), 33)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 33)", ""))
					goto errorExit
				}
				{
					p.SetState(358)
					p.Match(RiddleParserAdd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(359)

					var _x = p.expression(34)

					localctx.(*AddExprContext).right = _x
				}

			case 5:
				localctx = NewSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*SubExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
					goto errorExit
				}
				{
					p.SetState(361)
					p.Match(RiddleParserSub)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(362)

					var _x = p.expression(33)

					localctx.(*SubExprContext).right = _x
				}

			case 6:
				localctx = NewShlExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ShlExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(363)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
					goto errorExit
				}
				{
					p.SetState(364)
					p.Match(RiddleParserLeftLeft)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(365)

					var _x = p.expression(32)

					localctx.(*ShlExprContext).right = _x
				}

			case 7:
				localctx = NewAShrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AShrExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(366)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
					goto errorExit
				}
				{
					p.SetState(367)
					p.Match(RiddleParserRightRight)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(368)

					var _x = p.expression(31)

					localctx.(*AShrExprContext).right = _x
				}

			case 8:
				localctx = NewLShrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LShrExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(369)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
				}
				{
					p.SetState(370)
					p.Match(RiddleParserRightRightRight)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(371)

					var _x = p.expression(30)

					localctx.(*LShrExprContext).right = _x
				}

			case 9:
				localctx = NewGreaterExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*GreaterExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(372)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(373)
					p.Match(RiddleParserGreater)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(374)

					var _x = p.expression(29)

					localctx.(*GreaterExprContext).right = _x
				}

			case 10:
				localctx = NewLessExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LessExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(375)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(376)
					p.Match(RiddleParserLess)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(377)

					var _x = p.expression(28)

					localctx.(*LessExprContext).right = _x
				}

			case 11:
				localctx = NewGreaterEqualExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*GreaterEqualExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(378)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				{
					p.SetState(379)
					p.Match(RiddleParserGreater)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(380)
					p.Match(RiddleParserAssign)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(381)

					var _x = p.expression(27)

					localctx.(*GreaterEqualExprContext).right = _x
				}

			case 12:
				localctx = NewLessEqualExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*LessEqualExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(382)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				{
					p.SetState(383)
					p.Match(RiddleParserLess)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(384)
					p.Match(RiddleParserAssign)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(385)

					var _x = p.expression(26)

					localctx.(*LessEqualExprContext).right = _x
				}

			case 13:
				localctx = NewEqualExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*EqualExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(387)
					p.Match(RiddleParserEqual)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(388)

					var _x = p.expression(25)

					localctx.(*EqualExprContext).right = _x
				}

			case 14:
				localctx = NewNotEqualExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*NotEqualExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(389)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(390)
					p.Match(RiddleParserNot)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(391)
					p.Match(RiddleParserAssign)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(392)

					var _x = p.expression(24)

					localctx.(*NotEqualExprContext).right = _x
				}

			case 15:
				localctx = NewBitAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BitAndExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(393)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(394)
					p.Match(RiddleParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(395)

					var _x = p.expression(23)

					localctx.(*BitAndExprContext).right = _x
				}

			case 16:
				localctx = NewBitXorExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BitXorExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(396)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(397)
					p.Match(RiddleParserXor)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(398)

					var _x = p.expression(22)

					localctx.(*BitXorExprContext).right = _x
				}

			case 17:
				localctx = NewBitOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BitOrExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(399)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(400)
					p.Match(RiddleParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(401)

					var _x = p.expression(21)

					localctx.(*BitOrExprContext).right = _x
				}

			case 18:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AndExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(403)
					p.Match(RiddleParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(404)
					p.Match(RiddleParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(405)

					var _x = p.expression(20)

					localctx.(*AndExprContext).right = _x
				}

			case 19:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OrExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_expression)
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(407)
					p.Match(RiddleParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(408)
					p.Match(RiddleParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(409)

					var _x = p.expression(19)

					localctx.(*OrExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RiddleParserIdentifier, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RiddleParserRULE_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		p.Match(RiddleParserIdentifier)
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	Float() IFloatContext

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NumberContext) Float() IFloatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RiddleParserRULE_number)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RiddleParserHexadecimal, RiddleParserDecimal, RiddleParserOctal, RiddleParserBinary:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(417)
			p.Integer()
		}

	case RiddleParserFloat:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Float()
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

// IBooleanContext is an interface to support dynamic dispatch.
type IBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value attribute.
	GetValue() bool

	// SetValue sets the value attribute.
	SetValue(bool)

	// Getter signatures
	True() antlr.TerminalNode
	False() antlr.TerminalNode

	// IsBooleanContext differentiates from other interfaces.
	IsBooleanContext()
}

type BooleanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	value  bool
}

func NewEmptyBooleanContext() *BooleanContext {
	var p = new(BooleanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_boolean
	return p
}

func InitEmptyBooleanContext(p *BooleanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_boolean
}

func (*BooleanContext) IsBooleanContext() {}

func NewBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanContext {
	var p = new(BooleanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_boolean

	return p
}

func (s *BooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanContext) GetValue() bool { return s.value }

func (s *BooleanContext) SetValue(v bool) { s.value = v }

func (s *BooleanContext) True() antlr.TerminalNode {
	return s.GetToken(RiddleParserTrue, 0)
}

func (s *BooleanContext) False() antlr.TerminalNode {
	return s.GetToken(RiddleParserFalse, 0)
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Boolean() (localctx IBooleanContext) {
	localctx = NewBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RiddleParserRULE_boolean)
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RiddleParserTrue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.Match(RiddleParserTrue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*BooleanContext).value = true

	case RiddleParserFalse:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(423)
			p.Match(RiddleParserFalse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*BooleanContext).value = false

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

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(RiddleParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RiddleParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.Match(RiddleParserSTRING)
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

// IFloatContext is an interface to support dynamic dispatch.
type IFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_Float returns the _Float token.
	Get_Float() antlr.Token

	// Set_Float sets the _Float token.
	Set_Float(antlr.Token)

	// GetValue returns the value attribute.
	GetValue() float64

	// SetValue sets the value attribute.
	SetValue(float64)

	// Getter signatures
	Float() antlr.TerminalNode

	// IsFloatContext differentiates from other interfaces.
	IsFloatContext()
}

type FloatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	value  float64
	_Float antlr.Token
}

func NewEmptyFloatContext() *FloatContext {
	var p = new(FloatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_float
	return p
}

func InitEmptyFloatContext(p *FloatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_float
}

func (*FloatContext) IsFloatContext() {}

func NewFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatContext {
	var p = new(FloatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_float

	return p
}

func (s *FloatContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatContext) Get_Float() antlr.Token { return s._Float }

func (s *FloatContext) Set_Float(v antlr.Token) { s._Float = v }

func (s *FloatContext) GetValue() float64 { return s.value }

func (s *FloatContext) SetValue(v float64) { s.value = v }

func (s *FloatContext) Float() antlr.TerminalNode {
	return s.GetToken(RiddleParserFloat, 0)
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Float() (localctx IFloatContext) {
	localctx = NewFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RiddleParserRULE_float)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)

		var _m = p.Match(RiddleParserFloat)

		localctx.(*FloatContext)._Float = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*FloatContext).value, _ = strconv.ParseFloat((func() string {
		if localctx.(*FloatContext).Get_Float() == nil {
			return ""
		} else {
			return localctx.(*FloatContext).Get_Float().GetText()
		}
	}()), 64)

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

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_Decimal returns the _Decimal token.
	Get_Decimal() antlr.Token

	// Get_Hexadecimal returns the _Hexadecimal token.
	Get_Hexadecimal() antlr.Token

	// Get_Binary returns the _Binary token.
	Get_Binary() antlr.Token

	// Get_Octal returns the _Octal token.
	Get_Octal() antlr.Token

	// Set_Decimal sets the _Decimal token.
	Set_Decimal(antlr.Token)

	// Set_Hexadecimal sets the _Hexadecimal token.
	Set_Hexadecimal(antlr.Token)

	// Set_Binary sets the _Binary token.
	Set_Binary(antlr.Token)

	// Set_Octal sets the _Octal token.
	Set_Octal(antlr.Token)

	// GetValue returns the value attribute.
	GetValue() int

	// SetValue sets the value attribute.
	SetValue(int)

	// Getter signatures
	Decimal() antlr.TerminalNode
	Hexadecimal() antlr.TerminalNode
	Binary() antlr.TerminalNode
	Octal() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	value        int
	_Decimal     antlr.Token
	_Hexadecimal antlr.Token
	_Binary      antlr.Token
	_Octal       antlr.Token
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) Get_Decimal() antlr.Token { return s._Decimal }

func (s *IntegerContext) Get_Hexadecimal() antlr.Token { return s._Hexadecimal }

func (s *IntegerContext) Get_Binary() antlr.Token { return s._Binary }

func (s *IntegerContext) Get_Octal() antlr.Token { return s._Octal }

func (s *IntegerContext) Set_Decimal(v antlr.Token) { s._Decimal = v }

func (s *IntegerContext) Set_Hexadecimal(v antlr.Token) { s._Hexadecimal = v }

func (s *IntegerContext) Set_Binary(v antlr.Token) { s._Binary = v }

func (s *IntegerContext) Set_Octal(v antlr.Token) { s._Octal = v }

func (s *IntegerContext) GetValue() int { return s.value }

func (s *IntegerContext) SetValue(v int) { s.value = v }

func (s *IntegerContext) Decimal() antlr.TerminalNode {
	return s.GetToken(RiddleParserDecimal, 0)
}

func (s *IntegerContext) Hexadecimal() antlr.TerminalNode {
	return s.GetToken(RiddleParserHexadecimal, 0)
}

func (s *IntegerContext) Binary() antlr.TerminalNode {
	return s.GetToken(RiddleParserBinary, 0)
}

func (s *IntegerContext) Octal() antlr.TerminalNode {
	return s.GetToken(RiddleParserOctal, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RiddleParserRULE_integer)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RiddleParserDecimal:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)

			var _m = p.Match(RiddleParserDecimal)

			localctx.(*IntegerContext)._Decimal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		value, _ := strconv.Atoi((func() string {
			if localctx.(*IntegerContext).Get_Decimal() == nil {
				return ""
			} else {
				return localctx.(*IntegerContext).Get_Decimal().GetText()
			}
		}()))
		localctx.(*IntegerContext).value = int(value)

	case RiddleParserHexadecimal:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)

			var _m = p.Match(RiddleParserHexadecimal)

			localctx.(*IntegerContext)._Hexadecimal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		value, _ := strconv.ParseInt((func() string {
			if localctx.(*IntegerContext).Get_Hexadecimal() == nil {
				return ""
			} else {
				return localctx.(*IntegerContext).Get_Hexadecimal().GetText()
			}
		}())[2:], 16, 32)
		localctx.(*IntegerContext).value = int(value)

	case RiddleParserBinary:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(436)

			var _m = p.Match(RiddleParserBinary)

			localctx.(*IntegerContext)._Binary = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		value, _ := strconv.ParseInt((func() string {
			if localctx.(*IntegerContext).Get_Binary() == nil {
				return ""
			} else {
				return localctx.(*IntegerContext).Get_Binary().GetText()
			}
		}())[2:], 2, 32)
		localctx.(*IntegerContext).value = int(value)

	case RiddleParserOctal:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(438)

			var _m = p.Match(RiddleParserOctal)

			localctx.(*IntegerContext)._Octal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		value, _ := strconv.ParseInt((func() string {
			if localctx.(*IntegerContext).Get_Octal() == nil {
				return ""
			} else {
				return localctx.(*IntegerContext).Get_Octal().GetText()
			}
		}())[1:], 8, 32)
		localctx.(*IntegerContext).value = int(value)

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

// ITemplateArgContext is an interface to support dynamic dispatch.
type ITemplateArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	TypeName() ITypeNameContext

	// IsTemplateArgContext differentiates from other interfaces.
	IsTemplateArgContext()
}

type TemplateArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateArgContext() *TemplateArgContext {
	var p = new(TemplateArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_templateArg
	return p
}

func InitEmptyTemplateArgContext(p *TemplateArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_templateArg
}

func (*TemplateArgContext) IsTemplateArgContext() {}

func NewTemplateArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateArgContext {
	var p = new(TemplateArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_templateArg

	return p
}

func (s *TemplateArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TemplateArgContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TemplateArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterTemplateArg(s)
	}
}

func (s *TemplateArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitTemplateArg(s)
	}
}

func (s *TemplateArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitTemplateArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) TemplateArg() (localctx ITemplateArgContext) {
	localctx = NewTemplateArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RiddleParserRULE_templateArg)
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.typeName(0)
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

// ITemplateArgsContext is an interface to support dynamic dispatch.
type ITemplateArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTemplateArg() []ITemplateArgContext
	TemplateArg(i int) ITemplateArgContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsTemplateArgsContext differentiates from other interfaces.
	IsTemplateArgsContext()
}

type TemplateArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateArgsContext() *TemplateArgsContext {
	var p = new(TemplateArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_templateArgs
	return p
}

func InitEmptyTemplateArgsContext(p *TemplateArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_templateArgs
}

func (*TemplateArgsContext) IsTemplateArgsContext() {}

func NewTemplateArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateArgsContext {
	var p = new(TemplateArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_templateArgs

	return p
}

func (s *TemplateArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateArgsContext) AllTemplateArg() []ITemplateArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITemplateArgContext); ok {
			len++
		}
	}

	tst := make([]ITemplateArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITemplateArgContext); ok {
			tst[i] = t.(ITemplateArgContext)
			i++
		}
	}

	return tst
}

func (s *TemplateArgsContext) TemplateArg(i int) ITemplateArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemplateArgContext); ok {
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

	return t.(ITemplateArgContext)
}

func (s *TemplateArgsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(RiddleParserComma)
}

func (s *TemplateArgsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(RiddleParserComma, i)
}

func (s *TemplateArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterTemplateArgs(s)
	}
}

func (s *TemplateArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitTemplateArgs(s)
	}
}

func (s *TemplateArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitTemplateArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) TemplateArgs() (localctx ITemplateArgsContext) {
	localctx = NewTemplateArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RiddleParserRULE_templateArgs)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8939625984507576320) != 0 {
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(446)
					p.TemplateArg()
				}
				{
					p.SetState(447)
					p.Match(RiddleParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(453)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.TemplateArg()
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

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBaseType returns the baseType rule contexts.
	GetBaseType() ITypeNameContext

	// GetName returns the name rule contexts.
	GetName() IIdContext

	// GetArgs returns the args rule contexts.
	GetArgs() ITemplateArgsContext

	// GetSize returns the size rule contexts.
	GetSize() IExpressionContext

	// SetBaseType sets the baseType rule contexts.
	SetBaseType(ITypeNameContext)

	// SetName sets the name rule contexts.
	SetName(IIdContext)

	// SetArgs sets the args rule contexts.
	SetArgs(ITemplateArgsContext)

	// SetSize sets the size rule contexts.
	SetSize(IExpressionContext)

	// Getter signatures
	Id() IIdContext
	Less() antlr.TerminalNode
	Greater() antlr.TerminalNode
	TemplateArgs() ITemplateArgsContext
	LeftSquare() antlr.TerminalNode
	RightSquare() antlr.TerminalNode
	TypeName() ITypeNameContext
	Expression() IExpressionContext

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	baseType ITypeNameContext
	name     IIdContext
	args     ITemplateArgsContext
	size     IExpressionContext
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RiddleParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiddleParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) GetBaseType() ITypeNameContext { return s.baseType }

func (s *TypeNameContext) GetName() IIdContext { return s.name }

func (s *TypeNameContext) GetArgs() ITemplateArgsContext { return s.args }

func (s *TypeNameContext) GetSize() IExpressionContext { return s.size }

func (s *TypeNameContext) SetBaseType(v ITypeNameContext) { s.baseType = v }

func (s *TypeNameContext) SetName(v IIdContext) { s.name = v }

func (s *TypeNameContext) SetArgs(v ITemplateArgsContext) { s.args = v }

func (s *TypeNameContext) SetSize(v IExpressionContext) { s.size = v }

func (s *TypeNameContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *TypeNameContext) Less() antlr.TerminalNode {
	return s.GetToken(RiddleParserLess, 0)
}

func (s *TypeNameContext) Greater() antlr.TerminalNode {
	return s.GetToken(RiddleParserGreater, 0)
}

func (s *TypeNameContext) TemplateArgs() ITemplateArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemplateArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemplateArgsContext)
}

func (s *TypeNameContext) LeftSquare() antlr.TerminalNode {
	return s.GetToken(RiddleParserLeftSquare, 0)
}

func (s *TypeNameContext) RightSquare() antlr.TerminalNode {
	return s.GetToken(RiddleParserRightSquare, 0)
}

func (s *TypeNameContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeNameContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiddleParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RiddleParserVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RiddleParser) TypeName() (localctx ITypeNameContext) {
	return p.typeName(0)
}

func (p *RiddleParser) typeName(_p int) (localctx ITypeNameContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeNameContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, RiddleParserRULE_typeName, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(458)

			var _x = p.Id()

			localctx.(*TypeNameContext).name = _x
		}

	case 2:
		{
			p.SetState(459)

			var _x = p.Id()

			localctx.(*TypeNameContext).name = _x
		}
		{
			p.SetState(460)
			p.Match(RiddleParserLess)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)

			var _x = p.TemplateArgs()

			localctx.(*TypeNameContext).args = _x
		}
		{
			p.SetState(462)
			p.Match(RiddleParserGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeNameContext(p, _parentctx, _parentState)
			localctx.(*TypeNameContext).baseType = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RiddleParserRULE_typeName)
			p.SetState(466)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(467)
				p.Match(RiddleParserLeftSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(468)

				var _x = p.expression(0)

				localctx.(*TypeNameContext).size = _x
			}
			{
				p.SetState(469)
				p.Match(RiddleParserRightSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *RiddleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExprPtrContext = nil
		if localctx != nil {
			t = localctx.(*ExprPtrContext)
		}
		return p.ExprPtr_Sempred(t, predIndex)

	case 20:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 29:
		var t *TypeNameContext = nil
		if localctx != nil {
			t = localctx.(*TypeNameContext)
		}
		return p.TypeName_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RiddleParser) ExprPtr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RiddleParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 36)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 35)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 34)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 33)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 30)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RiddleParser) TypeName_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 20:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
