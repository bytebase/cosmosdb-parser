// Code generated from CosmosDBLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CosmosDBLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CosmosDBLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cosmosdblexerLexerInit() {
	staticData := &CosmosDBLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'*'", "'AS'", "'SELECT'", "'FROM'", "'DISTINCT'", "'UNDEFINED'",
		"'NULL'", "'FALSE'", "'TRUE'", "'NOT'", "'UDF'", "'WHERE'", "'AND'",
		"'OR'", "'@'", "'{'", "'}'", "'['", "']'", "'('", "')'", "'''", "'\"'",
		"','", "'.'", "'?'", "':'", "'+'", "'-'", "'~'", "'/'", "'%'", "'&'",
		"'|'", "'||'", "'^'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL",
		"DISTINCT_SYMBOL", "UNDEFINED_SYMBOL", "NULL_SYMBOL", "FALSE_SYMBOL",
		"TRUE_SYMBOL", "NOT_SYMBOL", "UDF_SYMBOL", "WHERE_SYMBOL", "AND_SYMBOL",
		"OR_SYMBOL", "AT_SYMBOL", "LC_BRACKET_SYMBOL", "RC_BRACKET_SYMBOL",
		"LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL", "LR_BRACKET_SYMBOL", "RR_BRACKET_SYMBOL",
		"SINGLE_QUOTE_SYMBOL", "DOUBLE_QUOTE_SYMBOL", "COMMA_SYMBOL", "DOT_SYMBOL",
		"QUESTION_MARK_SYMBOL", "COLON_SYMBOL", "PLUS_SYMBOL", "MINUS_SYMBOL",
		"BIT_NOT_SYMBOL", "DIVIDE_SYMBOL", "MODULO_SYMBOL", "BIT_AND_SYMBOL",
		"BIT_OR_SYMBOL", "DOUBLE_BAR_SYMBOL", "BIT_XOR_SYMBOL", "EQUAL_SYMBOL",
		"IDENTIFIER", "WHITESPACE", "DECIMAL", "REAL", "FLOAT", "HEXADECIMAL",
		"STRING_LITERAL",
	}
	staticData.RuleNames = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "MULTIPLY_OPERATOR",
		"AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL", "DISTINCT_SYMBOL", "UNDEFINED_SYMBOL",
		"NULL_SYMBOL", "FALSE_SYMBOL", "TRUE_SYMBOL", "NOT_SYMBOL", "UDF_SYMBOL",
		"WHERE_SYMBOL", "AND_SYMBOL", "OR_SYMBOL", "AT_SYMBOL", "LC_BRACKET_SYMBOL",
		"RC_BRACKET_SYMBOL", "LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL", "LR_BRACKET_SYMBOL",
		"RR_BRACKET_SYMBOL", "SINGLE_QUOTE_SYMBOL", "DOUBLE_QUOTE_SYMBOL", "COMMA_SYMBOL",
		"DOT_SYMBOL", "QUESTION_MARK_SYMBOL", "COLON_SYMBOL", "PLUS_SYMBOL",
		"MINUS_SYMBOL", "BIT_NOT_SYMBOL", "DIVIDE_SYMBOL", "MODULO_SYMBOL",
		"BIT_AND_SYMBOL", "BIT_OR_SYMBOL", "DOUBLE_BAR_SYMBOL", "BIT_XOR_SYMBOL",
		"EQUAL_SYMBOL", "IDENTIFIER", "WHITESPACE", "DEC_DIGIT", "DEC_DOT_DEC",
		"DECIMAL", "REAL", "FLOAT", "HEX_DIGIT", "HEXADECIMAL", "FullWidthLetter",
		"ESCAPE_SEQUENCE", "STRING_CHAR", "STRING_LITERAL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 431, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54,
		1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1,
		59, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 5, 63,
		328, 8, 63, 10, 63, 12, 63, 331, 9, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1,
		65, 1, 65, 1, 66, 4, 66, 340, 8, 66, 11, 66, 12, 66, 341, 1, 66, 1, 66,
		4, 66, 346, 8, 66, 11, 66, 12, 66, 347, 1, 66, 4, 66, 351, 8, 66, 11, 66,
		12, 66, 352, 1, 66, 1, 66, 1, 66, 1, 66, 4, 66, 359, 8, 66, 11, 66, 12,
		66, 360, 3, 66, 363, 8, 66, 1, 67, 4, 67, 366, 8, 67, 11, 67, 12, 67, 367,
		1, 68, 1, 68, 3, 68, 372, 8, 68, 1, 68, 1, 68, 3, 68, 376, 8, 68, 1, 68,
		4, 68, 379, 8, 68, 11, 68, 12, 68, 380, 1, 69, 1, 69, 1, 70, 1, 70, 1,
		71, 1, 71, 1, 71, 4, 71, 390, 8, 71, 11, 71, 12, 71, 391, 1, 72, 1, 72,
		1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 3,
		73, 406, 8, 73, 1, 74, 1, 74, 3, 74, 410, 8, 74, 1, 75, 1, 75, 5, 75, 414,
		8, 75, 10, 75, 12, 75, 417, 9, 75, 1, 75, 1, 75, 1, 75, 1, 75, 5, 75, 423,
		8, 75, 10, 75, 12, 75, 426, 9, 75, 1, 75, 1, 75, 3, 75, 430, 8, 75, 0,
		0, 76, 1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0,
		21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41,
		0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 1, 55, 2, 57, 3, 59, 4, 61, 5,
		63, 6, 65, 7, 67, 8, 69, 9, 71, 10, 73, 11, 75, 12, 77, 13, 79, 14, 81,
		15, 83, 16, 85, 17, 87, 18, 89, 19, 91, 20, 93, 21, 95, 22, 97, 23, 99,
		24, 101, 25, 103, 26, 105, 27, 107, 28, 109, 29, 111, 30, 113, 31, 115,
		32, 117, 33, 119, 34, 121, 35, 123, 36, 125, 37, 127, 38, 129, 39, 131,
		0, 133, 0, 135, 40, 137, 41, 139, 42, 141, 0, 143, 43, 145, 0, 147, 0,
		149, 0, 151, 44, 1, 0, 35, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98,
		2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2,
		0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2,
		0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2,
		0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2,
		0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2,
		0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2,
		0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2,
		0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 2,
		0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10,
		12, 13, 32, 32, 1, 0, 48, 57, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70,
		97, 102, 10, 0, 192, 214, 216, 246, 248, 8191, 11264, 12287, 12352, 12687,
		13056, 13183, 13312, 16383, 19968, 55295, 63744, 64255, 65280, 65520, 14,
		0, 34, 34, 39, 39, 47, 47, 66, 66, 70, 70, 78, 78, 82, 82, 84, 84, 92,
		92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 5, 0, 10, 10, 13, 13,
		34, 34, 39, 39, 92, 92, 415, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1,
		0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1,
		0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0,
		125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 135, 1, 0,
		0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 151,
		1, 0, 0, 0, 1, 153, 1, 0, 0, 0, 3, 155, 1, 0, 0, 0, 5, 157, 1, 0, 0, 0,
		7, 159, 1, 0, 0, 0, 9, 161, 1, 0, 0, 0, 11, 163, 1, 0, 0, 0, 13, 165, 1,
		0, 0, 0, 15, 167, 1, 0, 0, 0, 17, 169, 1, 0, 0, 0, 19, 171, 1, 0, 0, 0,
		21, 173, 1, 0, 0, 0, 23, 175, 1, 0, 0, 0, 25, 177, 1, 0, 0, 0, 27, 179,
		1, 0, 0, 0, 29, 181, 1, 0, 0, 0, 31, 183, 1, 0, 0, 0, 33, 185, 1, 0, 0,
		0, 35, 187, 1, 0, 0, 0, 37, 189, 1, 0, 0, 0, 39, 191, 1, 0, 0, 0, 41, 193,
		1, 0, 0, 0, 43, 195, 1, 0, 0, 0, 45, 197, 1, 0, 0, 0, 47, 199, 1, 0, 0,
		0, 49, 201, 1, 0, 0, 0, 51, 203, 1, 0, 0, 0, 53, 205, 1, 0, 0, 0, 55, 207,
		1, 0, 0, 0, 57, 210, 1, 0, 0, 0, 59, 217, 1, 0, 0, 0, 61, 222, 1, 0, 0,
		0, 63, 231, 1, 0, 0, 0, 65, 241, 1, 0, 0, 0, 67, 246, 1, 0, 0, 0, 69, 252,
		1, 0, 0, 0, 71, 257, 1, 0, 0, 0, 73, 261, 1, 0, 0, 0, 75, 265, 1, 0, 0,
		0, 77, 271, 1, 0, 0, 0, 79, 275, 1, 0, 0, 0, 81, 278, 1, 0, 0, 0, 83, 280,
		1, 0, 0, 0, 85, 282, 1, 0, 0, 0, 87, 284, 1, 0, 0, 0, 89, 286, 1, 0, 0,
		0, 91, 288, 1, 0, 0, 0, 93, 290, 1, 0, 0, 0, 95, 292, 1, 0, 0, 0, 97, 294,
		1, 0, 0, 0, 99, 296, 1, 0, 0, 0, 101, 298, 1, 0, 0, 0, 103, 300, 1, 0,
		0, 0, 105, 302, 1, 0, 0, 0, 107, 304, 1, 0, 0, 0, 109, 306, 1, 0, 0, 0,
		111, 308, 1, 0, 0, 0, 113, 310, 1, 0, 0, 0, 115, 312, 1, 0, 0, 0, 117,
		314, 1, 0, 0, 0, 119, 316, 1, 0, 0, 0, 121, 318, 1, 0, 0, 0, 123, 321,
		1, 0, 0, 0, 125, 323, 1, 0, 0, 0, 127, 325, 1, 0, 0, 0, 129, 332, 1, 0,
		0, 0, 131, 336, 1, 0, 0, 0, 133, 362, 1, 0, 0, 0, 135, 365, 1, 0, 0, 0,
		137, 371, 1, 0, 0, 0, 139, 382, 1, 0, 0, 0, 141, 384, 1, 0, 0, 0, 143,
		386, 1, 0, 0, 0, 145, 393, 1, 0, 0, 0, 147, 405, 1, 0, 0, 0, 149, 409,
		1, 0, 0, 0, 151, 429, 1, 0, 0, 0, 153, 154, 7, 0, 0, 0, 154, 2, 1, 0, 0,
		0, 155, 156, 7, 1, 0, 0, 156, 4, 1, 0, 0, 0, 157, 158, 7, 2, 0, 0, 158,
		6, 1, 0, 0, 0, 159, 160, 7, 3, 0, 0, 160, 8, 1, 0, 0, 0, 161, 162, 7, 4,
		0, 0, 162, 10, 1, 0, 0, 0, 163, 164, 7, 5, 0, 0, 164, 12, 1, 0, 0, 0, 165,
		166, 7, 6, 0, 0, 166, 14, 1, 0, 0, 0, 167, 168, 7, 7, 0, 0, 168, 16, 1,
		0, 0, 0, 169, 170, 7, 8, 0, 0, 170, 18, 1, 0, 0, 0, 171, 172, 7, 9, 0,
		0, 172, 20, 1, 0, 0, 0, 173, 174, 7, 10, 0, 0, 174, 22, 1, 0, 0, 0, 175,
		176, 7, 11, 0, 0, 176, 24, 1, 0, 0, 0, 177, 178, 7, 12, 0, 0, 178, 26,
		1, 0, 0, 0, 179, 180, 7, 13, 0, 0, 180, 28, 1, 0, 0, 0, 181, 182, 7, 14,
		0, 0, 182, 30, 1, 0, 0, 0, 183, 184, 7, 15, 0, 0, 184, 32, 1, 0, 0, 0,
		185, 186, 7, 16, 0, 0, 186, 34, 1, 0, 0, 0, 187, 188, 7, 17, 0, 0, 188,
		36, 1, 0, 0, 0, 189, 190, 7, 18, 0, 0, 190, 38, 1, 0, 0, 0, 191, 192, 7,
		19, 0, 0, 192, 40, 1, 0, 0, 0, 193, 194, 7, 20, 0, 0, 194, 42, 1, 0, 0,
		0, 195, 196, 7, 21, 0, 0, 196, 44, 1, 0, 0, 0, 197, 198, 7, 22, 0, 0, 198,
		46, 1, 0, 0, 0, 199, 200, 7, 23, 0, 0, 200, 48, 1, 0, 0, 0, 201, 202, 7,
		24, 0, 0, 202, 50, 1, 0, 0, 0, 203, 204, 7, 25, 0, 0, 204, 52, 1, 0, 0,
		0, 205, 206, 5, 42, 0, 0, 206, 54, 1, 0, 0, 0, 207, 208, 7, 0, 0, 0, 208,
		209, 7, 18, 0, 0, 209, 56, 1, 0, 0, 0, 210, 211, 7, 18, 0, 0, 211, 212,
		7, 4, 0, 0, 212, 213, 7, 11, 0, 0, 213, 214, 7, 4, 0, 0, 214, 215, 7, 2,
		0, 0, 215, 216, 7, 19, 0, 0, 216, 58, 1, 0, 0, 0, 217, 218, 7, 5, 0, 0,
		218, 219, 7, 17, 0, 0, 219, 220, 7, 14, 0, 0, 220, 221, 7, 12, 0, 0, 221,
		60, 1, 0, 0, 0, 222, 223, 7, 3, 0, 0, 223, 224, 7, 8, 0, 0, 224, 225, 7,
		18, 0, 0, 225, 226, 7, 19, 0, 0, 226, 227, 7, 8, 0, 0, 227, 228, 7, 13,
		0, 0, 228, 229, 7, 2, 0, 0, 229, 230, 7, 19, 0, 0, 230, 62, 1, 0, 0, 0,
		231, 232, 7, 20, 0, 0, 232, 233, 7, 13, 0, 0, 233, 234, 7, 3, 0, 0, 234,
		235, 7, 4, 0, 0, 235, 236, 7, 5, 0, 0, 236, 237, 7, 8, 0, 0, 237, 238,
		7, 13, 0, 0, 238, 239, 7, 4, 0, 0, 239, 240, 7, 3, 0, 0, 240, 64, 1, 0,
		0, 0, 241, 242, 7, 13, 0, 0, 242, 243, 7, 20, 0, 0, 243, 244, 7, 11, 0,
		0, 244, 245, 7, 11, 0, 0, 245, 66, 1, 0, 0, 0, 246, 247, 7, 5, 0, 0, 247,
		248, 7, 0, 0, 0, 248, 249, 7, 11, 0, 0, 249, 250, 7, 18, 0, 0, 250, 251,
		7, 4, 0, 0, 251, 68, 1, 0, 0, 0, 252, 253, 7, 19, 0, 0, 253, 254, 7, 17,
		0, 0, 254, 255, 7, 20, 0, 0, 255, 256, 7, 4, 0, 0, 256, 70, 1, 0, 0, 0,
		257, 258, 7, 13, 0, 0, 258, 259, 7, 14, 0, 0, 259, 260, 7, 19, 0, 0, 260,
		72, 1, 0, 0, 0, 261, 262, 7, 20, 0, 0, 262, 263, 7, 3, 0, 0, 263, 264,
		7, 5, 0, 0, 264, 74, 1, 0, 0, 0, 265, 266, 7, 22, 0, 0, 266, 267, 7, 7,
		0, 0, 267, 268, 7, 4, 0, 0, 268, 269, 7, 17, 0, 0, 269, 270, 7, 4, 0, 0,
		270, 76, 1, 0, 0, 0, 271, 272, 7, 0, 0, 0, 272, 273, 7, 13, 0, 0, 273,
		274, 7, 3, 0, 0, 274, 78, 1, 0, 0, 0, 275, 276, 7, 14, 0, 0, 276, 277,
		7, 17, 0, 0, 277, 80, 1, 0, 0, 0, 278, 279, 5, 64, 0, 0, 279, 82, 1, 0,
		0, 0, 280, 281, 5, 123, 0, 0, 281, 84, 1, 0, 0, 0, 282, 283, 5, 125, 0,
		0, 283, 86, 1, 0, 0, 0, 284, 285, 5, 91, 0, 0, 285, 88, 1, 0, 0, 0, 286,
		287, 5, 93, 0, 0, 287, 90, 1, 0, 0, 0, 288, 289, 5, 40, 0, 0, 289, 92,
		1, 0, 0, 0, 290, 291, 5, 41, 0, 0, 291, 94, 1, 0, 0, 0, 292, 293, 5, 39,
		0, 0, 293, 96, 1, 0, 0, 0, 294, 295, 5, 34, 0, 0, 295, 98, 1, 0, 0, 0,
		296, 297, 5, 44, 0, 0, 297, 100, 1, 0, 0, 0, 298, 299, 5, 46, 0, 0, 299,
		102, 1, 0, 0, 0, 300, 301, 5, 63, 0, 0, 301, 104, 1, 0, 0, 0, 302, 303,
		5, 58, 0, 0, 303, 106, 1, 0, 0, 0, 304, 305, 5, 43, 0, 0, 305, 108, 1,
		0, 0, 0, 306, 307, 5, 45, 0, 0, 307, 110, 1, 0, 0, 0, 308, 309, 5, 126,
		0, 0, 309, 112, 1, 0, 0, 0, 310, 311, 5, 47, 0, 0, 311, 114, 1, 0, 0, 0,
		312, 313, 5, 37, 0, 0, 313, 116, 1, 0, 0, 0, 314, 315, 5, 38, 0, 0, 315,
		118, 1, 0, 0, 0, 316, 317, 5, 124, 0, 0, 317, 120, 1, 0, 0, 0, 318, 319,
		5, 124, 0, 0, 319, 320, 5, 124, 0, 0, 320, 122, 1, 0, 0, 0, 321, 322, 5,
		94, 0, 0, 322, 124, 1, 0, 0, 0, 323, 324, 5, 61, 0, 0, 324, 126, 1, 0,
		0, 0, 325, 329, 7, 26, 0, 0, 326, 328, 7, 27, 0, 0, 327, 326, 1, 0, 0,
		0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330,
		128, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 332, 333, 7, 28, 0, 0, 333, 334,
		1, 0, 0, 0, 334, 335, 6, 64, 0, 0, 335, 130, 1, 0, 0, 0, 336, 337, 7, 29,
		0, 0, 337, 132, 1, 0, 0, 0, 338, 340, 3, 131, 65, 0, 339, 338, 1, 0, 0,
		0, 340, 341, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342,
		343, 1, 0, 0, 0, 343, 345, 5, 46, 0, 0, 344, 346, 3, 131, 65, 0, 345, 344,
		1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 347, 348, 1, 0,
		0, 0, 348, 363, 1, 0, 0, 0, 349, 351, 3, 131, 65, 0, 350, 349, 1, 0, 0,
		0, 351, 352, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353,
		354, 1, 0, 0, 0, 354, 355, 5, 46, 0, 0, 355, 363, 1, 0, 0, 0, 356, 358,
		5, 46, 0, 0, 357, 359, 3, 131, 65, 0, 358, 357, 1, 0, 0, 0, 359, 360, 1,
		0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 363, 1, 0, 0,
		0, 362, 339, 1, 0, 0, 0, 362, 350, 1, 0, 0, 0, 362, 356, 1, 0, 0, 0, 363,
		134, 1, 0, 0, 0, 364, 366, 3, 131, 65, 0, 365, 364, 1, 0, 0, 0, 366, 367,
		1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 136, 1, 0,
		0, 0, 369, 372, 3, 135, 67, 0, 370, 372, 3, 133, 66, 0, 371, 369, 1, 0,
		0, 0, 371, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 375, 7, 4, 0, 0,
		374, 376, 7, 30, 0, 0, 375, 374, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376,
		378, 1, 0, 0, 0, 377, 379, 3, 131, 65, 0, 378, 377, 1, 0, 0, 0, 379, 380,
		1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 138, 1, 0,
		0, 0, 382, 383, 3, 133, 66, 0, 383, 140, 1, 0, 0, 0, 384, 385, 7, 31, 0,
		0, 385, 142, 1, 0, 0, 0, 386, 387, 5, 48, 0, 0, 387, 389, 7, 23, 0, 0,
		388, 390, 3, 141, 70, 0, 389, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391,
		389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 144, 1, 0, 0, 0, 393, 394,
		7, 32, 0, 0, 394, 146, 1, 0, 0, 0, 395, 396, 5, 92, 0, 0, 396, 406, 7,
		33, 0, 0, 397, 398, 5, 92, 0, 0, 398, 399, 7, 20, 0, 0, 399, 400, 1, 0,
		0, 0, 400, 401, 3, 141, 70, 0, 401, 402, 3, 141, 70, 0, 402, 403, 3, 141,
		70, 0, 403, 404, 3, 141, 70, 0, 404, 406, 1, 0, 0, 0, 405, 395, 1, 0, 0,
		0, 405, 397, 1, 0, 0, 0, 406, 148, 1, 0, 0, 0, 407, 410, 3, 147, 73, 0,
		408, 410, 8, 34, 0, 0, 409, 407, 1, 0, 0, 0, 409, 408, 1, 0, 0, 0, 410,
		150, 1, 0, 0, 0, 411, 415, 3, 95, 47, 0, 412, 414, 3, 149, 74, 0, 413,
		412, 1, 0, 0, 0, 414, 417, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 415, 416,
		1, 0, 0, 0, 416, 418, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 418, 419, 3, 95,
		47, 0, 419, 430, 1, 0, 0, 0, 420, 424, 3, 97, 48, 0, 421, 423, 3, 149,
		74, 0, 422, 421, 1, 0, 0, 0, 423, 426, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0,
		424, 425, 1, 0, 0, 0, 425, 427, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 427,
		428, 3, 97, 48, 0, 428, 430, 1, 0, 0, 0, 429, 411, 1, 0, 0, 0, 429, 420,
		1, 0, 0, 0, 430, 152, 1, 0, 0, 0, 17, 0, 329, 341, 347, 352, 360, 362,
		367, 371, 375, 380, 391, 405, 409, 415, 424, 429, 1, 0, 1, 0,
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

// CosmosDBLexerInit initializes any static state used to implement CosmosDBLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCosmosDBLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CosmosDBLexerInit() {
	staticData := &CosmosDBLexerLexerStaticData
	staticData.once.Do(cosmosdblexerLexerInit)
}

// NewCosmosDBLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCosmosDBLexer(input antlr.CharStream) *CosmosDBLexer {
	CosmosDBLexerInit()
	l := new(CosmosDBLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CosmosDBLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CosmosDBLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CosmosDBLexer tokens.
const (
	CosmosDBLexerMULTIPLY_OPERATOR    = 1
	CosmosDBLexerAS_SYMBOL            = 2
	CosmosDBLexerSELECT_SYMBOL        = 3
	CosmosDBLexerFROM_SYMBOL          = 4
	CosmosDBLexerDISTINCT_SYMBOL      = 5
	CosmosDBLexerUNDEFINED_SYMBOL     = 6
	CosmosDBLexerNULL_SYMBOL          = 7
	CosmosDBLexerFALSE_SYMBOL         = 8
	CosmosDBLexerTRUE_SYMBOL          = 9
	CosmosDBLexerNOT_SYMBOL           = 10
	CosmosDBLexerUDF_SYMBOL           = 11
	CosmosDBLexerWHERE_SYMBOL         = 12
	CosmosDBLexerAND_SYMBOL           = 13
	CosmosDBLexerOR_SYMBOL            = 14
	CosmosDBLexerAT_SYMBOL            = 15
	CosmosDBLexerLC_BRACKET_SYMBOL    = 16
	CosmosDBLexerRC_BRACKET_SYMBOL    = 17
	CosmosDBLexerLS_BRACKET_SYMBOL    = 18
	CosmosDBLexerRS_BRACKET_SYMBOL    = 19
	CosmosDBLexerLR_BRACKET_SYMBOL    = 20
	CosmosDBLexerRR_BRACKET_SYMBOL    = 21
	CosmosDBLexerSINGLE_QUOTE_SYMBOL  = 22
	CosmosDBLexerDOUBLE_QUOTE_SYMBOL  = 23
	CosmosDBLexerCOMMA_SYMBOL         = 24
	CosmosDBLexerDOT_SYMBOL           = 25
	CosmosDBLexerQUESTION_MARK_SYMBOL = 26
	CosmosDBLexerCOLON_SYMBOL         = 27
	CosmosDBLexerPLUS_SYMBOL          = 28
	CosmosDBLexerMINUS_SYMBOL         = 29
	CosmosDBLexerBIT_NOT_SYMBOL       = 30
	CosmosDBLexerDIVIDE_SYMBOL        = 31
	CosmosDBLexerMODULO_SYMBOL        = 32
	CosmosDBLexerBIT_AND_SYMBOL       = 33
	CosmosDBLexerBIT_OR_SYMBOL        = 34
	CosmosDBLexerDOUBLE_BAR_SYMBOL    = 35
	CosmosDBLexerBIT_XOR_SYMBOL       = 36
	CosmosDBLexerEQUAL_SYMBOL         = 37
	CosmosDBLexerIDENTIFIER           = 38
	CosmosDBLexerWHITESPACE           = 39
	CosmosDBLexerDECIMAL              = 40
	CosmosDBLexerREAL                 = 41
	CosmosDBLexerFLOAT                = 42
	CosmosDBLexerHEXADECIMAL          = 43
	CosmosDBLexerSTRING_LITERAL       = 44
)
