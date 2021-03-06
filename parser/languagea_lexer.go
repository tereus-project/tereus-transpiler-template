// Code generated from LanguageA.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type LanguageALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var languagealexerLexerStaticData struct {
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

func languagealexerLexerInit() {
	staticData := &languagealexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.symbolicNames = []string{
		"", "Equal", "Add", "Subtract", "Multiply", "Divide", "Identifier",
		"FloatConstant", "IntegerConstant", "Whitespace", "Newline",
	}
	staticData.ruleNames = []string{
		"Equal", "Add", "Subtract", "Multiply", "Divide", "Identifier", "FloatConstant",
		"IntegerConstant", "Whitespace", "Newline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 72, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5,
		5, 34, 8, 5, 10, 5, 12, 5, 37, 9, 5, 1, 6, 4, 6, 40, 8, 6, 11, 6, 12, 6,
		41, 1, 6, 1, 6, 4, 6, 46, 8, 6, 11, 6, 12, 6, 47, 1, 7, 1, 7, 5, 7, 52,
		8, 7, 10, 7, 12, 7, 55, 9, 7, 1, 8, 4, 8, 58, 8, 8, 11, 8, 12, 8, 59, 1,
		8, 1, 8, 1, 9, 1, 9, 3, 9, 66, 8, 9, 1, 9, 3, 9, 69, 8, 9, 1, 9, 1, 9,
		0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 9, 9, 32, 32, 78, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25,
		1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 31, 1, 0, 0, 0, 13,
		39, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 57, 1, 0, 0, 0, 19, 68, 1, 0, 0,
		0, 21, 22, 5, 61, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 43, 0, 0, 24, 4,
		1, 0, 0, 0, 25, 26, 5, 45, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 42, 0, 0,
		28, 8, 1, 0, 0, 0, 29, 30, 5, 47, 0, 0, 30, 10, 1, 0, 0, 0, 31, 35, 7,
		0, 0, 0, 32, 34, 7, 1, 0, 0, 33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35,
		33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 12, 1, 0, 0, 0, 37, 35, 1, 0, 0,
		0, 38, 40, 7, 2, 0, 0, 39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39,
		1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 45, 5, 46, 0, 0,
		44, 46, 7, 2, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45, 1,
		0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 14, 1, 0, 0, 0, 49, 53, 7, 3, 0, 0, 50,
		52, 7, 2, 0, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0,
		0, 53, 54, 1, 0, 0, 0, 54, 16, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 58,
		7, 4, 0, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 6, 8, 0, 0, 62, 18, 1,
		0, 0, 0, 63, 65, 5, 13, 0, 0, 64, 66, 5, 10, 0, 0, 65, 64, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 69, 5, 10, 0, 0, 68, 63, 1,
		0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 6, 9, 0, 0, 71,
		20, 1, 0, 0, 0, 8, 0, 35, 41, 47, 53, 59, 65, 68, 1, 6, 0, 0,
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

// LanguageALexerInit initializes any static state used to implement LanguageALexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLanguageALexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LanguageALexerInit() {
	staticData := &languagealexerLexerStaticData
	staticData.once.Do(languagealexerLexerInit)
}

// NewLanguageALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLanguageALexer(input antlr.CharStream) *LanguageALexer {
	LanguageALexerInit()
	l := new(LanguageALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &languagealexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "LanguageA.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LanguageALexer tokens.
const (
	LanguageALexerEqual           = 1
	LanguageALexerAdd             = 2
	LanguageALexerSubtract        = 3
	LanguageALexerMultiply        = 4
	LanguageALexerDivide          = 5
	LanguageALexerIdentifier      = 6
	LanguageALexerFloatConstant   = 7
	LanguageALexerIntegerConstant = 8
	LanguageALexerWhitespace      = 9
	LanguageALexerNewline         = 10
)
