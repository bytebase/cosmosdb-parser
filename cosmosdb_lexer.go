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
		"", "'*'", "'AS'", "'SELECT'", "'FROM'", "'DISTINCT'", "'['", "']'",
		"'\"'", "','", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL",
		"DISTINCT_SYMBOL", "LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL", "DOUBLE_QUOTE_SYMBOL",
		"COMMA_SYMBOL", "DOT_SYMBOL", "IDENTIFIER", "WHITESPACE", "DECIMAL_DIGITS",
	}
	staticData.RuleNames = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "MULTIPLY_OPERATOR",
		"AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL", "DISTINCT_SYMBOL", "LS_BRACKET_SYMBOL",
		"RS_BRACKET_SYMBOL", "DOUBLE_QUOTE_SYMBOL", "COMMA_SYMBOL", "DOT_SYMBOL",
		"IDENTIFIER", "WHITESPACE", "DECIMAL_DIGIT", "DECIMAL_DIGITS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 187, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		5, 36, 172, 8, 36, 10, 36, 12, 36, 175, 9, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 39, 4, 39, 184, 8, 39, 11, 39, 12, 39, 185, 0, 0,
		40, 1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0, 21,
		0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 0,
		43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 1, 55, 2, 57, 3, 59, 4, 61, 5, 63,
		6, 65, 7, 67, 8, 69, 9, 71, 10, 73, 11, 75, 12, 77, 0, 79, 13, 1, 0, 30,
		2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0,
		68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0,
		71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0,
		74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0,
		77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0,
		80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0,
		83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0,
		86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0,
		89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 2, 0, 65, 90, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 12, 13, 32, 32, 1, 0, 48,
		57, 161, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 81, 1, 0, 0, 0, 3, 83, 1, 0, 0,
		0, 5, 85, 1, 0, 0, 0, 7, 87, 1, 0, 0, 0, 9, 89, 1, 0, 0, 0, 11, 91, 1,
		0, 0, 0, 13, 93, 1, 0, 0, 0, 15, 95, 1, 0, 0, 0, 17, 97, 1, 0, 0, 0, 19,
		99, 1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23, 103, 1, 0, 0, 0, 25, 105, 1, 0,
		0, 0, 27, 107, 1, 0, 0, 0, 29, 109, 1, 0, 0, 0, 31, 111, 1, 0, 0, 0, 33,
		113, 1, 0, 0, 0, 35, 115, 1, 0, 0, 0, 37, 117, 1, 0, 0, 0, 39, 119, 1,
		0, 0, 0, 41, 121, 1, 0, 0, 0, 43, 123, 1, 0, 0, 0, 45, 125, 1, 0, 0, 0,
		47, 127, 1, 0, 0, 0, 49, 129, 1, 0, 0, 0, 51, 131, 1, 0, 0, 0, 53, 133,
		1, 0, 0, 0, 55, 135, 1, 0, 0, 0, 57, 138, 1, 0, 0, 0, 59, 145, 1, 0, 0,
		0, 61, 150, 1, 0, 0, 0, 63, 159, 1, 0, 0, 0, 65, 161, 1, 0, 0, 0, 67, 163,
		1, 0, 0, 0, 69, 165, 1, 0, 0, 0, 71, 167, 1, 0, 0, 0, 73, 169, 1, 0, 0,
		0, 75, 176, 1, 0, 0, 0, 77, 180, 1, 0, 0, 0, 79, 183, 1, 0, 0, 0, 81, 82,
		7, 0, 0, 0, 82, 2, 1, 0, 0, 0, 83, 84, 7, 1, 0, 0, 84, 4, 1, 0, 0, 0, 85,
		86, 7, 2, 0, 0, 86, 6, 1, 0, 0, 0, 87, 88, 7, 3, 0, 0, 88, 8, 1, 0, 0,
		0, 89, 90, 7, 4, 0, 0, 90, 10, 1, 0, 0, 0, 91, 92, 7, 5, 0, 0, 92, 12,
		1, 0, 0, 0, 93, 94, 7, 6, 0, 0, 94, 14, 1, 0, 0, 0, 95, 96, 7, 7, 0, 0,
		96, 16, 1, 0, 0, 0, 97, 98, 7, 8, 0, 0, 98, 18, 1, 0, 0, 0, 99, 100, 7,
		9, 0, 0, 100, 20, 1, 0, 0, 0, 101, 102, 7, 10, 0, 0, 102, 22, 1, 0, 0,
		0, 103, 104, 7, 11, 0, 0, 104, 24, 1, 0, 0, 0, 105, 106, 7, 12, 0, 0, 106,
		26, 1, 0, 0, 0, 107, 108, 7, 13, 0, 0, 108, 28, 1, 0, 0, 0, 109, 110, 7,
		14, 0, 0, 110, 30, 1, 0, 0, 0, 111, 112, 7, 15, 0, 0, 112, 32, 1, 0, 0,
		0, 113, 114, 7, 16, 0, 0, 114, 34, 1, 0, 0, 0, 115, 116, 7, 17, 0, 0, 116,
		36, 1, 0, 0, 0, 117, 118, 7, 18, 0, 0, 118, 38, 1, 0, 0, 0, 119, 120, 7,
		19, 0, 0, 120, 40, 1, 0, 0, 0, 121, 122, 7, 20, 0, 0, 122, 42, 1, 0, 0,
		0, 123, 124, 7, 21, 0, 0, 124, 44, 1, 0, 0, 0, 125, 126, 7, 22, 0, 0, 126,
		46, 1, 0, 0, 0, 127, 128, 7, 23, 0, 0, 128, 48, 1, 0, 0, 0, 129, 130, 7,
		24, 0, 0, 130, 50, 1, 0, 0, 0, 131, 132, 7, 25, 0, 0, 132, 52, 1, 0, 0,
		0, 133, 134, 5, 42, 0, 0, 134, 54, 1, 0, 0, 0, 135, 136, 7, 0, 0, 0, 136,
		137, 7, 18, 0, 0, 137, 56, 1, 0, 0, 0, 138, 139, 7, 18, 0, 0, 139, 140,
		7, 4, 0, 0, 140, 141, 7, 11, 0, 0, 141, 142, 7, 4, 0, 0, 142, 143, 7, 2,
		0, 0, 143, 144, 7, 19, 0, 0, 144, 58, 1, 0, 0, 0, 145, 146, 7, 5, 0, 0,
		146, 147, 7, 17, 0, 0, 147, 148, 7, 14, 0, 0, 148, 149, 7, 12, 0, 0, 149,
		60, 1, 0, 0, 0, 150, 151, 7, 3, 0, 0, 151, 152, 7, 8, 0, 0, 152, 153, 7,
		18, 0, 0, 153, 154, 7, 19, 0, 0, 154, 155, 7, 8, 0, 0, 155, 156, 7, 13,
		0, 0, 156, 157, 7, 2, 0, 0, 157, 158, 7, 19, 0, 0, 158, 62, 1, 0, 0, 0,
		159, 160, 5, 91, 0, 0, 160, 64, 1, 0, 0, 0, 161, 162, 5, 93, 0, 0, 162,
		66, 1, 0, 0, 0, 163, 164, 5, 34, 0, 0, 164, 68, 1, 0, 0, 0, 165, 166, 5,
		44, 0, 0, 166, 70, 1, 0, 0, 0, 167, 168, 5, 46, 0, 0, 168, 72, 1, 0, 0,
		0, 169, 173, 7, 26, 0, 0, 170, 172, 7, 27, 0, 0, 171, 170, 1, 0, 0, 0,
		172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		74, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 7, 28, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 179, 6, 37, 0, 0, 179, 76, 1, 0, 0, 0, 180, 181, 7, 29,
		0, 0, 181, 78, 1, 0, 0, 0, 182, 184, 3, 77, 38, 0, 183, 182, 1, 0, 0, 0,
		184, 185, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		80, 1, 0, 0, 0, 3, 0, 173, 185, 1, 0, 1, 0,
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
	CosmosDBLexerMULTIPLY_OPERATOR   = 1
	CosmosDBLexerAS_SYMBOL           = 2
	CosmosDBLexerSELECT_SYMBOL       = 3
	CosmosDBLexerFROM_SYMBOL         = 4
	CosmosDBLexerDISTINCT_SYMBOL     = 5
	CosmosDBLexerLS_BRACKET_SYMBOL   = 6
	CosmosDBLexerRS_BRACKET_SYMBOL   = 7
	CosmosDBLexerDOUBLE_QUOTE_SYMBOL = 8
	CosmosDBLexerCOMMA_SYMBOL        = 9
	CosmosDBLexerDOT_SYMBOL          = 10
	CosmosDBLexerIDENTIFIER          = 11
	CosmosDBLexerWHITESPACE          = 12
	CosmosDBLexerDECIMAL_DIGITS      = 13
)
