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
		"'NULL'", "'FALSE'", "'TRUE'", "'NOT'", "'UDF'", "'WHERE'", "'@'", "'{'",
		"'}'", "'['", "']'", "'('", "')'", "'''", "'\"'", "','", "'.'", "'?'",
		"':'", "'+'", "'-'", "'~'", "'/'", "'%'", "'&'", "'|'", "'||'", "'^'",
		"'='",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL",
		"DISTINCT_SYMBOL", "UNDEFINED_SYMBOL", "NULL_SYMBOL", "FALSE_SYMBOL",
		"TRUE_SYMBOL", "NOT_SYMBOL", "UDF_SYMBOL", "WHERE_SYMBOL", "AT_SYMBOL",
		"LC_BRACKET_SYMBOL", "RC_BRACKET_SYMBOL", "LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL",
		"LR_BRACKET_SYMBOL", "RR_BRACKET_SYMBOL", "SINGLE_QUOTE_SYMBOL", "DOUBLE_QUOTE_SYMBOL",
		"COMMA_SYMBOL", "DOT_SYMBOL", "QUESTION_MARK_SYMBOL", "COLON_SYMBOL",
		"PLUS_SYMBOL", "MINUS_SYMBOL", "BIT_NOT_SYMBOL", "DIVIDE_SYMBOL", "MODULO_SYMBOL",
		"BIT_AND_SYMBOL", "BIT_OR_SYMBOL", "OR_SYMBOL", "BIT_XOR_SYMBOL", "EUQAL_SYMBOL",
		"IDENTIFIER", "WHITESPACE", "DECIMAL", "REAL", "FLOAT", "HEXADECIMAL",
		"STRING_LITERAL",
	}
	staticData.RuleNames = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "MULTIPLY_OPERATOR",
		"AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL", "DISTINCT_SYMBOL", "UNDEFINED_SYMBOL",
		"NULL_SYMBOL", "FALSE_SYMBOL", "TRUE_SYMBOL", "NOT_SYMBOL", "UDF_SYMBOL",
		"WHERE_SYMBOL", "AT_SYMBOL", "LC_BRACKET_SYMBOL", "RC_BRACKET_SYMBOL",
		"LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL", "LR_BRACKET_SYMBOL", "RR_BRACKET_SYMBOL",
		"SINGLE_QUOTE_SYMBOL", "DOUBLE_QUOTE_SYMBOL", "COMMA_SYMBOL", "DOT_SYMBOL",
		"QUESTION_MARK_SYMBOL", "COLON_SYMBOL", "PLUS_SYMBOL", "MINUS_SYMBOL",
		"BIT_NOT_SYMBOL", "DIVIDE_SYMBOL", "MODULO_SYMBOL", "BIT_AND_SYMBOL",
		"BIT_OR_SYMBOL", "OR_SYMBOL", "BIT_XOR_SYMBOL", "EUQAL_SYMBOL", "IDENTIFIER",
		"WHITESPACE", "DEC_DIGIT", "DEC_DOT_DEC", "DECIMAL", "REAL", "FLOAT",
		"HEX_DIGIT", "HEXADECIMAL", "FullWidthLetter", "ESCAPE_SEQUENCE", "STRING_CHAR",
		"STRING_LITERAL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 420, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		73, 7, 73, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1,
		58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 5, 61, 317,
		8, 61, 10, 61, 12, 61, 320, 9, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1,
		63, 1, 64, 4, 64, 329, 8, 64, 11, 64, 12, 64, 330, 1, 64, 1, 64, 4, 64,
		335, 8, 64, 11, 64, 12, 64, 336, 1, 64, 4, 64, 340, 8, 64, 11, 64, 12,
		64, 341, 1, 64, 1, 64, 1, 64, 1, 64, 4, 64, 348, 8, 64, 11, 64, 12, 64,
		349, 3, 64, 352, 8, 64, 1, 65, 4, 65, 355, 8, 65, 11, 65, 12, 65, 356,
		1, 66, 1, 66, 3, 66, 361, 8, 66, 1, 66, 1, 66, 3, 66, 365, 8, 66, 1, 66,
		4, 66, 368, 8, 66, 11, 66, 12, 66, 369, 1, 67, 1, 67, 1, 68, 1, 68, 1,
		69, 1, 69, 1, 69, 4, 69, 379, 8, 69, 11, 69, 12, 69, 380, 1, 70, 1, 70,
		1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 3,
		71, 395, 8, 71, 1, 72, 1, 72, 3, 72, 399, 8, 72, 1, 73, 1, 73, 5, 73, 403,
		8, 73, 10, 73, 12, 73, 406, 9, 73, 1, 73, 1, 73, 1, 73, 1, 73, 5, 73, 412,
		8, 73, 10, 73, 12, 73, 415, 9, 73, 1, 73, 1, 73, 3, 73, 419, 8, 73, 0,
		0, 74, 1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0,
		21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41,
		0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 1, 55, 2, 57, 3, 59, 4, 61, 5,
		63, 6, 65, 7, 67, 8, 69, 9, 71, 10, 73, 11, 75, 12, 77, 13, 79, 14, 81,
		15, 83, 16, 85, 17, 87, 18, 89, 19, 91, 20, 93, 21, 95, 22, 97, 23, 99,
		24, 101, 25, 103, 26, 105, 27, 107, 28, 109, 29, 111, 30, 113, 31, 115,
		32, 117, 33, 119, 34, 121, 35, 123, 36, 125, 37, 127, 0, 129, 0, 131, 38,
		133, 39, 135, 40, 137, 0, 139, 41, 141, 0, 143, 0, 145, 0, 147, 42, 1,
		0, 35, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99,
		2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102,
		2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105,
		2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108,
		2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111,
		2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114,
		2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117,
		2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120,
		2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 2, 0, 65, 90, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 12, 13, 32, 32, 1,
		0, 48, 57, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70, 97, 102, 10, 0,
		192, 214, 216, 246, 248, 8191, 11264, 12287, 12352, 12687, 13056, 13183,
		13312, 16383, 19968, 55295, 63744, 64255, 65280, 65520, 14, 0, 34, 34,
		39, 39, 47, 47, 66, 66, 70, 70, 78, 78, 82, 82, 84, 84, 92, 92, 98, 98,
		102, 102, 110, 110, 114, 114, 116, 116, 5, 0, 10, 10, 13, 13, 34, 34, 39,
		39, 92, 92, 404, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0,
		0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0,
		0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111,
		1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0,
		0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1,
		0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0,
		139, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 1, 149, 1, 0, 0, 0, 3, 151, 1, 0,
		0, 0, 5, 153, 1, 0, 0, 0, 7, 155, 1, 0, 0, 0, 9, 157, 1, 0, 0, 0, 11, 159,
		1, 0, 0, 0, 13, 161, 1, 0, 0, 0, 15, 163, 1, 0, 0, 0, 17, 165, 1, 0, 0,
		0, 19, 167, 1, 0, 0, 0, 21, 169, 1, 0, 0, 0, 23, 171, 1, 0, 0, 0, 25, 173,
		1, 0, 0, 0, 27, 175, 1, 0, 0, 0, 29, 177, 1, 0, 0, 0, 31, 179, 1, 0, 0,
		0, 33, 181, 1, 0, 0, 0, 35, 183, 1, 0, 0, 0, 37, 185, 1, 0, 0, 0, 39, 187,
		1, 0, 0, 0, 41, 189, 1, 0, 0, 0, 43, 191, 1, 0, 0, 0, 45, 193, 1, 0, 0,
		0, 47, 195, 1, 0, 0, 0, 49, 197, 1, 0, 0, 0, 51, 199, 1, 0, 0, 0, 53, 201,
		1, 0, 0, 0, 55, 203, 1, 0, 0, 0, 57, 206, 1, 0, 0, 0, 59, 213, 1, 0, 0,
		0, 61, 218, 1, 0, 0, 0, 63, 227, 1, 0, 0, 0, 65, 237, 1, 0, 0, 0, 67, 242,
		1, 0, 0, 0, 69, 248, 1, 0, 0, 0, 71, 253, 1, 0, 0, 0, 73, 257, 1, 0, 0,
		0, 75, 261, 1, 0, 0, 0, 77, 267, 1, 0, 0, 0, 79, 269, 1, 0, 0, 0, 81, 271,
		1, 0, 0, 0, 83, 273, 1, 0, 0, 0, 85, 275, 1, 0, 0, 0, 87, 277, 1, 0, 0,
		0, 89, 279, 1, 0, 0, 0, 91, 281, 1, 0, 0, 0, 93, 283, 1, 0, 0, 0, 95, 285,
		1, 0, 0, 0, 97, 287, 1, 0, 0, 0, 99, 289, 1, 0, 0, 0, 101, 291, 1, 0, 0,
		0, 103, 293, 1, 0, 0, 0, 105, 295, 1, 0, 0, 0, 107, 297, 1, 0, 0, 0, 109,
		299, 1, 0, 0, 0, 111, 301, 1, 0, 0, 0, 113, 303, 1, 0, 0, 0, 115, 305,
		1, 0, 0, 0, 117, 307, 1, 0, 0, 0, 119, 310, 1, 0, 0, 0, 121, 312, 1, 0,
		0, 0, 123, 314, 1, 0, 0, 0, 125, 321, 1, 0, 0, 0, 127, 325, 1, 0, 0, 0,
		129, 351, 1, 0, 0, 0, 131, 354, 1, 0, 0, 0, 133, 360, 1, 0, 0, 0, 135,
		371, 1, 0, 0, 0, 137, 373, 1, 0, 0, 0, 139, 375, 1, 0, 0, 0, 141, 382,
		1, 0, 0, 0, 143, 394, 1, 0, 0, 0, 145, 398, 1, 0, 0, 0, 147, 418, 1, 0,
		0, 0, 149, 150, 7, 0, 0, 0, 150, 2, 1, 0, 0, 0, 151, 152, 7, 1, 0, 0, 152,
		4, 1, 0, 0, 0, 153, 154, 7, 2, 0, 0, 154, 6, 1, 0, 0, 0, 155, 156, 7, 3,
		0, 0, 156, 8, 1, 0, 0, 0, 157, 158, 7, 4, 0, 0, 158, 10, 1, 0, 0, 0, 159,
		160, 7, 5, 0, 0, 160, 12, 1, 0, 0, 0, 161, 162, 7, 6, 0, 0, 162, 14, 1,
		0, 0, 0, 163, 164, 7, 7, 0, 0, 164, 16, 1, 0, 0, 0, 165, 166, 7, 8, 0,
		0, 166, 18, 1, 0, 0, 0, 167, 168, 7, 9, 0, 0, 168, 20, 1, 0, 0, 0, 169,
		170, 7, 10, 0, 0, 170, 22, 1, 0, 0, 0, 171, 172, 7, 11, 0, 0, 172, 24,
		1, 0, 0, 0, 173, 174, 7, 12, 0, 0, 174, 26, 1, 0, 0, 0, 175, 176, 7, 13,
		0, 0, 176, 28, 1, 0, 0, 0, 177, 178, 7, 14, 0, 0, 178, 30, 1, 0, 0, 0,
		179, 180, 7, 15, 0, 0, 180, 32, 1, 0, 0, 0, 181, 182, 7, 16, 0, 0, 182,
		34, 1, 0, 0, 0, 183, 184, 7, 17, 0, 0, 184, 36, 1, 0, 0, 0, 185, 186, 7,
		18, 0, 0, 186, 38, 1, 0, 0, 0, 187, 188, 7, 19, 0, 0, 188, 40, 1, 0, 0,
		0, 189, 190, 7, 20, 0, 0, 190, 42, 1, 0, 0, 0, 191, 192, 7, 21, 0, 0, 192,
		44, 1, 0, 0, 0, 193, 194, 7, 22, 0, 0, 194, 46, 1, 0, 0, 0, 195, 196, 7,
		23, 0, 0, 196, 48, 1, 0, 0, 0, 197, 198, 7, 24, 0, 0, 198, 50, 1, 0, 0,
		0, 199, 200, 7, 25, 0, 0, 200, 52, 1, 0, 0, 0, 201, 202, 5, 42, 0, 0, 202,
		54, 1, 0, 0, 0, 203, 204, 7, 0, 0, 0, 204, 205, 7, 18, 0, 0, 205, 56, 1,
		0, 0, 0, 206, 207, 7, 18, 0, 0, 207, 208, 7, 4, 0, 0, 208, 209, 7, 11,
		0, 0, 209, 210, 7, 4, 0, 0, 210, 211, 7, 2, 0, 0, 211, 212, 7, 19, 0, 0,
		212, 58, 1, 0, 0, 0, 213, 214, 7, 5, 0, 0, 214, 215, 7, 17, 0, 0, 215,
		216, 7, 14, 0, 0, 216, 217, 7, 12, 0, 0, 217, 60, 1, 0, 0, 0, 218, 219,
		7, 3, 0, 0, 219, 220, 7, 8, 0, 0, 220, 221, 7, 18, 0, 0, 221, 222, 7, 19,
		0, 0, 222, 223, 7, 8, 0, 0, 223, 224, 7, 13, 0, 0, 224, 225, 7, 2, 0, 0,
		225, 226, 7, 19, 0, 0, 226, 62, 1, 0, 0, 0, 227, 228, 7, 20, 0, 0, 228,
		229, 7, 13, 0, 0, 229, 230, 7, 3, 0, 0, 230, 231, 7, 4, 0, 0, 231, 232,
		7, 5, 0, 0, 232, 233, 7, 8, 0, 0, 233, 234, 7, 13, 0, 0, 234, 235, 7, 4,
		0, 0, 235, 236, 7, 3, 0, 0, 236, 64, 1, 0, 0, 0, 237, 238, 7, 13, 0, 0,
		238, 239, 7, 20, 0, 0, 239, 240, 7, 11, 0, 0, 240, 241, 7, 11, 0, 0, 241,
		66, 1, 0, 0, 0, 242, 243, 7, 5, 0, 0, 243, 244, 7, 0, 0, 0, 244, 245, 7,
		11, 0, 0, 245, 246, 7, 18, 0, 0, 246, 247, 7, 4, 0, 0, 247, 68, 1, 0, 0,
		0, 248, 249, 7, 19, 0, 0, 249, 250, 7, 17, 0, 0, 250, 251, 7, 20, 0, 0,
		251, 252, 7, 4, 0, 0, 252, 70, 1, 0, 0, 0, 253, 254, 7, 13, 0, 0, 254,
		255, 7, 14, 0, 0, 255, 256, 7, 19, 0, 0, 256, 72, 1, 0, 0, 0, 257, 258,
		7, 20, 0, 0, 258, 259, 7, 3, 0, 0, 259, 260, 7, 5, 0, 0, 260, 74, 1, 0,
		0, 0, 261, 262, 7, 22, 0, 0, 262, 263, 7, 7, 0, 0, 263, 264, 7, 4, 0, 0,
		264, 265, 7, 17, 0, 0, 265, 266, 7, 4, 0, 0, 266, 76, 1, 0, 0, 0, 267,
		268, 5, 64, 0, 0, 268, 78, 1, 0, 0, 0, 269, 270, 5, 123, 0, 0, 270, 80,
		1, 0, 0, 0, 271, 272, 5, 125, 0, 0, 272, 82, 1, 0, 0, 0, 273, 274, 5, 91,
		0, 0, 274, 84, 1, 0, 0, 0, 275, 276, 5, 93, 0, 0, 276, 86, 1, 0, 0, 0,
		277, 278, 5, 40, 0, 0, 278, 88, 1, 0, 0, 0, 279, 280, 5, 41, 0, 0, 280,
		90, 1, 0, 0, 0, 281, 282, 5, 39, 0, 0, 282, 92, 1, 0, 0, 0, 283, 284, 5,
		34, 0, 0, 284, 94, 1, 0, 0, 0, 285, 286, 5, 44, 0, 0, 286, 96, 1, 0, 0,
		0, 287, 288, 5, 46, 0, 0, 288, 98, 1, 0, 0, 0, 289, 290, 5, 63, 0, 0, 290,
		100, 1, 0, 0, 0, 291, 292, 5, 58, 0, 0, 292, 102, 1, 0, 0, 0, 293, 294,
		5, 43, 0, 0, 294, 104, 1, 0, 0, 0, 295, 296, 5, 45, 0, 0, 296, 106, 1,
		0, 0, 0, 297, 298, 5, 126, 0, 0, 298, 108, 1, 0, 0, 0, 299, 300, 5, 47,
		0, 0, 300, 110, 1, 0, 0, 0, 301, 302, 5, 37, 0, 0, 302, 112, 1, 0, 0, 0,
		303, 304, 5, 38, 0, 0, 304, 114, 1, 0, 0, 0, 305, 306, 5, 124, 0, 0, 306,
		116, 1, 0, 0, 0, 307, 308, 5, 124, 0, 0, 308, 309, 5, 124, 0, 0, 309, 118,
		1, 0, 0, 0, 310, 311, 5, 94, 0, 0, 311, 120, 1, 0, 0, 0, 312, 313, 5, 61,
		0, 0, 313, 122, 1, 0, 0, 0, 314, 318, 7, 26, 0, 0, 315, 317, 7, 27, 0,
		0, 316, 315, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318,
		319, 1, 0, 0, 0, 319, 124, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 322,
		7, 28, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 6, 62, 0, 0, 324, 126, 1,
		0, 0, 0, 325, 326, 7, 29, 0, 0, 326, 128, 1, 0, 0, 0, 327, 329, 3, 127,
		63, 0, 328, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0,
		330, 331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 334, 5, 46, 0, 0, 333,
		335, 3, 127, 63, 0, 334, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 352, 1, 0, 0, 0, 338, 340, 3, 127,
		63, 0, 339, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0,
		341, 342, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 5, 46, 0, 0, 344,
		352, 1, 0, 0, 0, 345, 347, 5, 46, 0, 0, 346, 348, 3, 127, 63, 0, 347, 346,
		1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0,
		0, 0, 350, 352, 1, 0, 0, 0, 351, 328, 1, 0, 0, 0, 351, 339, 1, 0, 0, 0,
		351, 345, 1, 0, 0, 0, 352, 130, 1, 0, 0, 0, 353, 355, 3, 127, 63, 0, 354,
		353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 356, 357,
		1, 0, 0, 0, 357, 132, 1, 0, 0, 0, 358, 361, 3, 131, 65, 0, 359, 361, 3,
		129, 64, 0, 360, 358, 1, 0, 0, 0, 360, 359, 1, 0, 0, 0, 361, 362, 1, 0,
		0, 0, 362, 364, 7, 4, 0, 0, 363, 365, 7, 30, 0, 0, 364, 363, 1, 0, 0, 0,
		364, 365, 1, 0, 0, 0, 365, 367, 1, 0, 0, 0, 366, 368, 3, 127, 63, 0, 367,
		366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370,
		1, 0, 0, 0, 370, 134, 1, 0, 0, 0, 371, 372, 3, 129, 64, 0, 372, 136, 1,
		0, 0, 0, 373, 374, 7, 31, 0, 0, 374, 138, 1, 0, 0, 0, 375, 376, 5, 48,
		0, 0, 376, 378, 7, 23, 0, 0, 377, 379, 3, 137, 68, 0, 378, 377, 1, 0, 0,
		0, 379, 380, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381,
		140, 1, 0, 0, 0, 382, 383, 7, 32, 0, 0, 383, 142, 1, 0, 0, 0, 384, 385,
		5, 92, 0, 0, 385, 395, 7, 33, 0, 0, 386, 387, 5, 92, 0, 0, 387, 388, 7,
		20, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390, 3, 137, 68, 0, 390, 391, 3, 137,
		68, 0, 391, 392, 3, 137, 68, 0, 392, 393, 3, 137, 68, 0, 393, 395, 1, 0,
		0, 0, 394, 384, 1, 0, 0, 0, 394, 386, 1, 0, 0, 0, 395, 144, 1, 0, 0, 0,
		396, 399, 3, 143, 71, 0, 397, 399, 8, 34, 0, 0, 398, 396, 1, 0, 0, 0, 398,
		397, 1, 0, 0, 0, 399, 146, 1, 0, 0, 0, 400, 404, 3, 91, 45, 0, 401, 403,
		3, 145, 72, 0, 402, 401, 1, 0, 0, 0, 403, 406, 1, 0, 0, 0, 404, 402, 1,
		0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 407, 1, 0, 0, 0, 406, 404, 1, 0, 0,
		0, 407, 408, 3, 91, 45, 0, 408, 419, 1, 0, 0, 0, 409, 413, 3, 93, 46, 0,
		410, 412, 3, 145, 72, 0, 411, 410, 1, 0, 0, 0, 412, 415, 1, 0, 0, 0, 413,
		411, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 416, 1, 0, 0, 0, 415, 413,
		1, 0, 0, 0, 416, 417, 3, 93, 46, 0, 417, 419, 1, 0, 0, 0, 418, 400, 1,
		0, 0, 0, 418, 409, 1, 0, 0, 0, 419, 148, 1, 0, 0, 0, 17, 0, 318, 330, 336,
		341, 349, 351, 356, 360, 364, 369, 380, 394, 398, 404, 413, 418, 1, 0,
		1, 0,
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
	CosmosDBLexerAT_SYMBOL            = 13
	CosmosDBLexerLC_BRACKET_SYMBOL    = 14
	CosmosDBLexerRC_BRACKET_SYMBOL    = 15
	CosmosDBLexerLS_BRACKET_SYMBOL    = 16
	CosmosDBLexerRS_BRACKET_SYMBOL    = 17
	CosmosDBLexerLR_BRACKET_SYMBOL    = 18
	CosmosDBLexerRR_BRACKET_SYMBOL    = 19
	CosmosDBLexerSINGLE_QUOTE_SYMBOL  = 20
	CosmosDBLexerDOUBLE_QUOTE_SYMBOL  = 21
	CosmosDBLexerCOMMA_SYMBOL         = 22
	CosmosDBLexerDOT_SYMBOL           = 23
	CosmosDBLexerQUESTION_MARK_SYMBOL = 24
	CosmosDBLexerCOLON_SYMBOL         = 25
	CosmosDBLexerPLUS_SYMBOL          = 26
	CosmosDBLexerMINUS_SYMBOL         = 27
	CosmosDBLexerBIT_NOT_SYMBOL       = 28
	CosmosDBLexerDIVIDE_SYMBOL        = 29
	CosmosDBLexerMODULO_SYMBOL        = 30
	CosmosDBLexerBIT_AND_SYMBOL       = 31
	CosmosDBLexerBIT_OR_SYMBOL        = 32
	CosmosDBLexerOR_SYMBOL            = 33
	CosmosDBLexerBIT_XOR_SYMBOL       = 34
	CosmosDBLexerEUQAL_SYMBOL         = 35
	CosmosDBLexerIDENTIFIER           = 36
	CosmosDBLexerWHITESPACE           = 37
	CosmosDBLexerDECIMAL              = 38
	CosmosDBLexerREAL                 = 39
	CosmosDBLexerFLOAT                = 40
	CosmosDBLexerHEXADECIMAL          = 41
	CosmosDBLexerSTRING_LITERAL       = 42
)
