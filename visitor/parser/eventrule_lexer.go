// Code generated from EventRule.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 150,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 5, 18, 111,
	10, 18, 3, 18, 6, 18, 114, 10, 18, 13, 18, 14, 18, 115, 3, 18, 3, 18, 6,
	18, 120, 10, 18, 13, 18, 14, 18, 121, 5, 18, 124, 10, 18, 3, 19, 6, 19,
	127, 10, 19, 13, 19, 14, 19, 128, 3, 20, 3, 20, 3, 20, 7, 20, 134, 10,
	20, 12, 20, 14, 20, 137, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 149, 10, 22, 3, 135, 2, 23, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 2, 3, 2, 6, 3, 2, 47, 47, 3, 2, 50, 59, 7, 2, 47, 48, 50, 59, 67, 92,
	97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 156, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 45,
	3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2, 2, 9, 53, 3, 2, 2, 2, 11,
	56, 3, 2, 2, 2, 13, 60, 3, 2, 2, 2, 15, 62, 3, 2, 2, 2, 17, 65, 3, 2, 2,
	2, 19, 67, 3, 2, 2, 2, 21, 69, 3, 2, 2, 2, 23, 72, 3, 2, 2, 2, 25, 75,
	3, 2, 2, 2, 27, 84, 3, 2, 2, 2, 29, 97, 3, 2, 2, 2, 31, 100, 3, 2, 2, 2,
	33, 107, 3, 2, 2, 2, 35, 110, 3, 2, 2, 2, 37, 126, 3, 2, 2, 2, 39, 130,
	3, 2, 2, 2, 41, 140, 3, 2, 2, 2, 43, 148, 3, 2, 2, 2, 45, 46, 7, 42, 2,
	2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48, 6, 3, 2, 2, 2, 49, 50, 7,
	99, 2, 2, 50, 51, 7, 112, 2, 2, 51, 52, 7, 102, 2, 2, 52, 8, 3, 2, 2, 2,
	53, 54, 7, 113, 2, 2, 54, 55, 7, 116, 2, 2, 55, 10, 3, 2, 2, 2, 56, 57,
	7, 112, 2, 2, 57, 58, 7, 113, 2, 2, 58, 59, 7, 118, 2, 2, 59, 12, 3, 2,
	2, 2, 60, 61, 7, 63, 2, 2, 61, 14, 3, 2, 2, 2, 62, 63, 7, 35, 2, 2, 63,
	64, 7, 63, 2, 2, 64, 16, 3, 2, 2, 2, 65, 66, 7, 64, 2, 2, 66, 18, 3, 2,
	2, 2, 67, 68, 7, 62, 2, 2, 68, 20, 3, 2, 2, 2, 69, 70, 7, 64, 2, 2, 70,
	71, 7, 63, 2, 2, 71, 22, 3, 2, 2, 2, 72, 73, 7, 62, 2, 2, 73, 74, 7, 63,
	2, 2, 74, 24, 3, 2, 2, 2, 75, 76, 7, 101, 2, 2, 76, 77, 7, 113, 2, 2, 77,
	78, 7, 112, 2, 2, 78, 79, 7, 118, 2, 2, 79, 80, 7, 99, 2, 2, 80, 81, 7,
	107, 2, 2, 81, 82, 7, 112, 2, 2, 82, 83, 7, 117, 2, 2, 83, 26, 3, 2, 2,
	2, 84, 85, 7, 112, 2, 2, 85, 86, 7, 113, 2, 2, 86, 87, 7, 118, 2, 2, 87,
	88, 7, 34, 2, 2, 88, 89, 7, 101, 2, 2, 89, 90, 7, 113, 2, 2, 90, 91, 7,
	112, 2, 2, 91, 92, 7, 118, 2, 2, 92, 93, 7, 99, 2, 2, 93, 94, 7, 107, 2,
	2, 94, 95, 7, 112, 2, 2, 95, 96, 7, 117, 2, 2, 96, 28, 3, 2, 2, 2, 97,
	98, 7, 107, 2, 2, 98, 99, 7, 112, 2, 2, 99, 30, 3, 2, 2, 2, 100, 101, 7,
	112, 2, 2, 101, 102, 7, 113, 2, 2, 102, 103, 7, 118, 2, 2, 103, 104, 7,
	34, 2, 2, 104, 105, 7, 107, 2, 2, 105, 106, 7, 112, 2, 2, 106, 32, 3, 2,
	2, 2, 107, 108, 7, 46, 2, 2, 108, 34, 3, 2, 2, 2, 109, 111, 9, 2, 2, 2,
	110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2, 2, 112,
	114, 9, 3, 2, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113,
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 123, 3, 2, 2, 2, 117, 119, 7, 48,
	2, 2, 118, 120, 9, 3, 2, 2, 119, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2,
	121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 124, 3, 2, 2, 2, 123,
	117, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 36, 3, 2, 2, 2, 125, 127, 9,
	4, 2, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 126, 3, 2, 2,
	2, 128, 129, 3, 2, 2, 2, 129, 38, 3, 2, 2, 2, 130, 135, 7, 36, 2, 2, 131,
	134, 5, 43, 22, 2, 132, 134, 11, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132,
	3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 135, 133, 3, 2,
	2, 2, 136, 138, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 139, 7, 36, 2, 2,
	139, 40, 3, 2, 2, 2, 140, 141, 9, 5, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143,
	8, 21, 2, 2, 143, 42, 3, 2, 2, 2, 144, 145, 7, 94, 2, 2, 145, 149, 7, 36,
	2, 2, 146, 147, 7, 94, 2, 2, 147, 149, 7, 94, 2, 2, 148, 144, 3, 2, 2,
	2, 148, 146, 3, 2, 2, 2, 149, 44, 3, 2, 2, 2, 11, 2, 110, 115, 121, 123,
	128, 133, 135, 148, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'and'", "'or'", "'not'", "'='", "'!='", "'>'", "'<'",
	"'>='", "'<='", "'contains'", "'not contains'", "'in'", "'not in'", "','",
}

var lexerSymbolicNames = []string{
	"", "", "", "AND", "OR", "NOT", "EQU", "NEQ", "GT", "LT", "GTE", "LTE",
	"CONTAINS", "NOTCONTAINS", "IN", "NOTIN", "COMMA", "NUMBER", "VAR", "STRING",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "AND", "OR", "NOT", "EQU", "NEQ", "GT", "LT", "GTE", "LTE",
	"CONTAINS", "NOTCONTAINS", "IN", "NOTIN", "COMMA", "NUMBER", "VAR", "STRING",
	"WHITESPACE", "ESC",
}

type EventRuleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEventRuleLexer(input antlr.CharStream) *EventRuleLexer {

	l := new(EventRuleLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "EventRule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EventRuleLexer tokens.
const (
	EventRuleLexerT__0        = 1
	EventRuleLexerT__1        = 2
	EventRuleLexerAND         = 3
	EventRuleLexerOR          = 4
	EventRuleLexerNOT         = 5
	EventRuleLexerEQU         = 6
	EventRuleLexerNEQ         = 7
	EventRuleLexerGT          = 8
	EventRuleLexerLT          = 9
	EventRuleLexerGTE         = 10
	EventRuleLexerLTE         = 11
	EventRuleLexerCONTAINS    = 12
	EventRuleLexerNOTCONTAINS = 13
	EventRuleLexerIN          = 14
	EventRuleLexerNOTIN       = 15
	EventRuleLexerCOMMA       = 16
	EventRuleLexerNUMBER      = 17
	EventRuleLexerVAR         = 18
	EventRuleLexerSTRING      = 19
	EventRuleLexerWHITESPACE  = 20
)
