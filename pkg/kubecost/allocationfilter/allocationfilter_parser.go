// Code generated from allocationfilter.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // allocationfilter

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type allocationfilterParser struct {
	*antlr.BaseParser
}

var allocationfilterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func allocationfilterParserInit() {
	staticData := &allocationfilterParserStaticData
	staticData.literalNames = []string{
		"", "'('", "'AND'", "')'", "'OR'", "'NOT'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "COMP", "CF1", "CF2", "CV", "WS",
	}
	staticData.ruleNames = []string{
		"filter",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 26, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 3, 0, 24, 8, 0, 1, 0, 0, 0, 1, 0, 0, 0, 29, 0, 23, 1, 0, 0, 0,
		2, 3, 5, 1, 0, 0, 3, 4, 3, 0, 0, 0, 4, 5, 5, 2, 0, 0, 5, 6, 3, 0, 0, 0,
		6, 7, 5, 3, 0, 0, 7, 24, 1, 0, 0, 0, 8, 9, 5, 1, 0, 0, 9, 10, 3, 0, 0,
		0, 10, 11, 5, 4, 0, 0, 11, 12, 3, 0, 0, 0, 12, 13, 5, 3, 0, 0, 13, 24,
		1, 0, 0, 0, 14, 15, 5, 6, 0, 0, 15, 16, 5, 2, 0, 0, 16, 24, 5, 6, 0, 0,
		17, 18, 5, 6, 0, 0, 18, 19, 5, 4, 0, 0, 19, 24, 5, 6, 0, 0, 20, 21, 5,
		5, 0, 0, 21, 24, 5, 6, 0, 0, 22, 24, 5, 6, 0, 0, 23, 2, 1, 0, 0, 0, 23,
		8, 1, 0, 0, 0, 23, 14, 1, 0, 0, 0, 23, 17, 1, 0, 0, 0, 23, 20, 1, 0, 0,
		0, 23, 22, 1, 0, 0, 0, 24, 1, 1, 0, 0, 0, 1, 23,
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

// allocationfilterParserInit initializes any static state used to implement allocationfilterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewallocationfilterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AllocationfilterParserInit() {
	staticData := &allocationfilterParserStaticData
	staticData.once.Do(allocationfilterParserInit)
}

// NewallocationfilterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewallocationfilterParser(input antlr.TokenStream) *allocationfilterParser {
	AllocationfilterParserInit()
	this := new(allocationfilterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &allocationfilterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "allocationfilter.g4"

	return this
}

// allocationfilterParser tokens.
const (
	allocationfilterParserEOF  = antlr.TokenEOF
	allocationfilterParserT__0 = 1
	allocationfilterParserT__1 = 2
	allocationfilterParserT__2 = 3
	allocationfilterParserT__3 = 4
	allocationfilterParserT__4 = 5
	allocationfilterParserCOMP = 6
	allocationfilterParserCF1  = 7
	allocationfilterParserCF2  = 8
	allocationfilterParserCV   = 9
	allocationfilterParserWS   = 10
)

// allocationfilterParserRULE_filter is the allocationfilterParser rule.
const allocationfilterParserRULE_filter = 0

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = allocationfilterParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = allocationfilterParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) AllFilter() []IFilterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterContext); ok {
			len++
		}
	}

	tst := make([]IFilterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterContext); ok {
			tst[i] = t.(IFilterContext)
			i++
		}
	}

	return tst
}

func (s *FilterContext) Filter(i int) IFilterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
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

	return t.(IFilterContext)
}

func (s *FilterContext) AllCOMP() []antlr.TerminalNode {
	return s.GetTokens(allocationfilterParserCOMP)
}

func (s *FilterContext) COMP(i int) antlr.TerminalNode {
	return s.GetToken(allocationfilterParserCOMP, i)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(allocationfilterListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(allocationfilterListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *allocationfilterParser) Filter() (localctx IFilterContext) {
	this := p
	_ = this

	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, allocationfilterParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(2)
			p.Match(allocationfilterParserT__0)
		}
		{
			p.SetState(3)
			p.Filter()
		}
		{
			p.SetState(4)
			p.Match(allocationfilterParserT__1)
		}
		{
			p.SetState(5)
			p.Filter()
		}
		{
			p.SetState(6)
			p.Match(allocationfilterParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(8)
			p.Match(allocationfilterParserT__0)
		}
		{
			p.SetState(9)
			p.Filter()
		}
		{
			p.SetState(10)
			p.Match(allocationfilterParserT__3)
		}
		{
			p.SetState(11)
			p.Filter()
		}
		{
			p.SetState(12)
			p.Match(allocationfilterParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(14)
			p.Match(allocationfilterParserCOMP)
		}
		{
			p.SetState(15)
			p.Match(allocationfilterParserT__1)
		}
		{
			p.SetState(16)
			p.Match(allocationfilterParserCOMP)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(17)
			p.Match(allocationfilterParserCOMP)
		}
		{
			p.SetState(18)
			p.Match(allocationfilterParserT__3)
		}
		{
			p.SetState(19)
			p.Match(allocationfilterParserCOMP)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(20)
			p.Match(allocationfilterParserT__4)
		}
		{
			p.SetState(21)
			p.Match(allocationfilterParserCOMP)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(22)
			p.Match(allocationfilterParserCOMP)
		}

	}

	return localctx
}
