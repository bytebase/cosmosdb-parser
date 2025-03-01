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
		"", "'*'", "'SELECT'", "'FROM'",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "SELECT_SYMBOL", "FROM_SYMBOL", "IDENTIFIER",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "MULTIPLY_OPERATOR",
		"SELECT_SYMBOL", "FROM_SYMBOL", "IDENTIFIER", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 140, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 5, 29, 132, 8, 29, 10, 29, 12, 29, 135, 9, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 0, 0, 31, 1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0,
		15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35,
		0, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 1, 55, 2,
		57, 3, 59, 4, 61, 5, 1, 0, 29, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98,
		98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101,
		2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104,
		2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107,
		2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110,
		2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113,
		2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116,
		2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119,
		2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122,
		2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9,
		10, 12, 13, 32, 32, 114, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 0, 0, 0, 3, 65,
		1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 69, 1, 0, 0, 0, 9, 71, 1, 0, 0, 0, 11,
		73, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 77, 1, 0, 0, 0, 17, 79, 1, 0, 0,
		0, 19, 81, 1, 0, 0, 0, 21, 83, 1, 0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 87,
		1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 91, 1, 0, 0, 0, 31, 93, 1, 0, 0, 0,
		33, 95, 1, 0, 0, 0, 35, 97, 1, 0, 0, 0, 37, 99, 1, 0, 0, 0, 39, 101, 1,
		0, 0, 0, 41, 103, 1, 0, 0, 0, 43, 105, 1, 0, 0, 0, 45, 107, 1, 0, 0, 0,
		47, 109, 1, 0, 0, 0, 49, 111, 1, 0, 0, 0, 51, 113, 1, 0, 0, 0, 53, 115,
		1, 0, 0, 0, 55, 117, 1, 0, 0, 0, 57, 124, 1, 0, 0, 0, 59, 129, 1, 0, 0,
		0, 61, 136, 1, 0, 0, 0, 63, 64, 7, 0, 0, 0, 64, 2, 1, 0, 0, 0, 65, 66,
		7, 1, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 7, 2, 0, 0, 68, 6, 1, 0, 0, 0, 69,
		70, 7, 3, 0, 0, 70, 8, 1, 0, 0, 0, 71, 72, 7, 4, 0, 0, 72, 10, 1, 0, 0,
		0, 73, 74, 7, 5, 0, 0, 74, 12, 1, 0, 0, 0, 75, 76, 7, 6, 0, 0, 76, 14,
		1, 0, 0, 0, 77, 78, 7, 7, 0, 0, 78, 16, 1, 0, 0, 0, 79, 80, 7, 8, 0, 0,
		80, 18, 1, 0, 0, 0, 81, 82, 7, 9, 0, 0, 82, 20, 1, 0, 0, 0, 83, 84, 7,
		10, 0, 0, 84, 22, 1, 0, 0, 0, 85, 86, 7, 11, 0, 0, 86, 24, 1, 0, 0, 0,
		87, 88, 7, 12, 0, 0, 88, 26, 1, 0, 0, 0, 89, 90, 7, 13, 0, 0, 90, 28, 1,
		0, 0, 0, 91, 92, 7, 14, 0, 0, 92, 30, 1, 0, 0, 0, 93, 94, 7, 15, 0, 0,
		94, 32, 1, 0, 0, 0, 95, 96, 7, 16, 0, 0, 96, 34, 1, 0, 0, 0, 97, 98, 7,
		17, 0, 0, 98, 36, 1, 0, 0, 0, 99, 100, 7, 18, 0, 0, 100, 38, 1, 0, 0, 0,
		101, 102, 7, 19, 0, 0, 102, 40, 1, 0, 0, 0, 103, 104, 7, 20, 0, 0, 104,
		42, 1, 0, 0, 0, 105, 106, 7, 21, 0, 0, 106, 44, 1, 0, 0, 0, 107, 108, 7,
		22, 0, 0, 108, 46, 1, 0, 0, 0, 109, 110, 7, 23, 0, 0, 110, 48, 1, 0, 0,
		0, 111, 112, 7, 24, 0, 0, 112, 50, 1, 0, 0, 0, 113, 114, 7, 25, 0, 0, 114,
		52, 1, 0, 0, 0, 115, 116, 5, 42, 0, 0, 116, 54, 1, 0, 0, 0, 117, 118, 7,
		18, 0, 0, 118, 119, 7, 4, 0, 0, 119, 120, 7, 11, 0, 0, 120, 121, 7, 4,
		0, 0, 121, 122, 7, 2, 0, 0, 122, 123, 7, 19, 0, 0, 123, 56, 1, 0, 0, 0,
		124, 125, 7, 5, 0, 0, 125, 126, 7, 17, 0, 0, 126, 127, 7, 14, 0, 0, 127,
		128, 7, 12, 0, 0, 128, 58, 1, 0, 0, 0, 129, 133, 7, 26, 0, 0, 130, 132,
		7, 27, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 60, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0,
		136, 137, 7, 28, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 6, 30, 0, 0, 139,
		62, 1, 0, 0, 0, 2, 0, 133, 1, 0, 1, 0,
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
	CosmosDBLexerMULTIPLY_OPERATOR = 1
	CosmosDBLexerSELECT_SYMBOL     = 2
	CosmosDBLexerFROM_SYMBOL       = 3
	CosmosDBLexerIDENTIFIER        = 4
	CosmosDBLexerWHITESPACE        = 5
)
