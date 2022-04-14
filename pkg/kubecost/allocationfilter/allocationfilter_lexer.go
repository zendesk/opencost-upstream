// Code generated from allocationfilter.g4 by ANTLR 4.10. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type allocationfilterLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var allocationfilterlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func allocationfilterlexerLexerInit() {
	staticData := &allocationfilterlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "'AND'", "')'", "'OR'", "'NOT'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "COMP", "CF1", "CF2", "CV", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "COMP", "CF1", "CF2", "CV",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 103, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 59, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 73, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 90, 8, 7, 1, 8,
		4, 8, 93, 8, 8, 11, 8, 12, 8, 94, 1, 9, 4, 9, 98, 8, 9, 11, 9, 12, 9, 99,
		1, 9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 1, 0, 2, 2, 0, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 109, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3,
		23, 1, 0, 0, 0, 5, 27, 1, 0, 0, 0, 7, 29, 1, 0, 0, 0, 9, 32, 1, 0, 0, 0,
		11, 58, 1, 0, 0, 0, 13, 72, 1, 0, 0, 0, 15, 89, 1, 0, 0, 0, 17, 92, 1,
		0, 0, 0, 19, 97, 1, 0, 0, 0, 21, 22, 5, 40, 0, 0, 22, 2, 1, 0, 0, 0, 23,
		24, 5, 65, 0, 0, 24, 25, 5, 78, 0, 0, 25, 26, 5, 68, 0, 0, 26, 4, 1, 0,
		0, 0, 27, 28, 5, 41, 0, 0, 28, 6, 1, 0, 0, 0, 29, 30, 5, 79, 0, 0, 30,
		31, 5, 82, 0, 0, 31, 8, 1, 0, 0, 0, 32, 33, 5, 78, 0, 0, 33, 34, 5, 79,
		0, 0, 34, 35, 5, 84, 0, 0, 35, 10, 1, 0, 0, 0, 36, 37, 3, 15, 7, 0, 37,
		38, 5, 58, 0, 0, 38, 39, 3, 17, 8, 0, 39, 40, 5, 61, 0, 0, 40, 41, 3, 17,
		8, 0, 41, 42, 5, 42, 0, 0, 42, 59, 1, 0, 0, 0, 43, 44, 3, 15, 7, 0, 44,
		45, 5, 58, 0, 0, 45, 46, 3, 17, 8, 0, 46, 47, 5, 61, 0, 0, 47, 48, 3, 17,
		8, 0, 48, 59, 1, 0, 0, 0, 49, 50, 3, 13, 6, 0, 50, 51, 5, 58, 0, 0, 51,
		52, 3, 17, 8, 0, 52, 53, 5, 42, 0, 0, 53, 59, 1, 0, 0, 0, 54, 55, 3, 13,
		6, 0, 55, 56, 5, 58, 0, 0, 56, 57, 3, 17, 8, 0, 57, 59, 1, 0, 0, 0, 58,
		36, 1, 0, 0, 0, 58, 43, 1, 0, 0, 0, 58, 49, 1, 0, 0, 0, 58, 54, 1, 0, 0,
		0, 59, 12, 1, 0, 0, 0, 60, 61, 5, 110, 0, 0, 61, 62, 5, 97, 0, 0, 62, 63,
		5, 109, 0, 0, 63, 64, 5, 101, 0, 0, 64, 65, 5, 115, 0, 0, 65, 66, 5, 112,
		0, 0, 66, 67, 5, 97, 0, 0, 67, 68, 5, 99, 0, 0, 68, 73, 5, 101, 0, 0, 69,
		70, 5, 112, 0, 0, 70, 71, 5, 111, 0, 0, 71, 73, 5, 100, 0, 0, 72, 60, 1,
		0, 0, 0, 72, 69, 1, 0, 0, 0, 73, 14, 1, 0, 0, 0, 74, 75, 5, 108, 0, 0,
		75, 76, 5, 97, 0, 0, 76, 77, 5, 98, 0, 0, 77, 78, 5, 101, 0, 0, 78, 90,
		5, 108, 0, 0, 79, 80, 5, 97, 0, 0, 80, 81, 5, 110, 0, 0, 81, 82, 5, 110,
		0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 97, 0, 0,
		85, 86, 5, 116, 0, 0, 86, 87, 5, 105, 0, 0, 87, 88, 5, 111, 0, 0, 88, 90,
		5, 110, 0, 0, 89, 74, 1, 0, 0, 0, 89, 79, 1, 0, 0, 0, 90, 16, 1, 0, 0,
		0, 91, 93, 7, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92,
		1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 18, 1, 0, 0, 0, 96, 98, 7, 1, 0, 0,
		97, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1,
		0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 6, 9, 0, 0, 102, 20, 1, 0, 0,
		0, 6, 0, 58, 72, 89, 94, 99, 1, 6, 0, 0,
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

// allocationfilterLexerInit initializes any static state used to implement allocationfilterLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewallocationfilterLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AllocationfilterLexerInit() {
	staticData := &allocationfilterlexerLexerStaticData
	staticData.once.Do(allocationfilterlexerLexerInit)
}

// NewallocationfilterLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewallocationfilterLexer(input antlr.CharStream) *allocationfilterLexer {
	AllocationfilterLexerInit()
	l := new(allocationfilterLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &allocationfilterlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "allocationfilter.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// allocationfilterLexer tokens.
const (
	allocationfilterLexerT__0 = 1
	allocationfilterLexerT__1 = 2
	allocationfilterLexerT__2 = 3
	allocationfilterLexerT__3 = 4
	allocationfilterLexerT__4 = 5
	allocationfilterLexerCOMP = 6
	allocationfilterLexerCF1  = 7
	allocationfilterLexerCF2  = 8
	allocationfilterLexerCV   = 9
	allocationfilterLexerWS   = 10
)
