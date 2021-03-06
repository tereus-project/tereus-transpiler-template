// Code generated from LanguageA.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LanguageA

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

type LanguageAParser struct {
	*antlr.BaseParser
}

var languageaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func languageaParserInit() {
	staticData := &languageaParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.symbolicNames = []string{
		"", "Equal", "Add", "Subtract", "Multiply", "Divide", "Identifier",
		"FloatConstant", "IntegerConstant", "Whitespace", "Newline",
	}
	staticData.ruleNames = []string{
		"translation", "topLevelStatement", "variableDeclaration", "expression",
		"binaryOperator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 36, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 4, 0, 12, 8, 0, 11, 0, 12, 0, 13, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 29, 8, 3, 10, 3,
		12, 3, 32, 9, 3, 1, 4, 1, 4, 1, 4, 0, 1, 6, 5, 0, 2, 4, 6, 8, 0, 2, 1,
		0, 7, 8, 1, 0, 2, 5, 32, 0, 11, 1, 0, 0, 0, 2, 15, 1, 0, 0, 0, 4, 17, 1,
		0, 0, 0, 6, 21, 1, 0, 0, 0, 8, 33, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11,
		10, 1, 0, 0, 0, 12, 13, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0,
		0, 14, 1, 1, 0, 0, 0, 15, 16, 3, 4, 2, 0, 16, 3, 1, 0, 0, 0, 17, 18, 5,
		6, 0, 0, 18, 19, 5, 1, 0, 0, 19, 20, 3, 6, 3, 0, 20, 5, 1, 0, 0, 0, 21,
		22, 6, 3, -1, 0, 22, 23, 7, 0, 0, 0, 23, 30, 1, 0, 0, 0, 24, 25, 10, 1,
		0, 0, 25, 26, 3, 8, 4, 0, 26, 27, 3, 6, 3, 2, 27, 29, 1, 0, 0, 0, 28, 24,
		1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0,
		31, 7, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 7, 1, 0, 0, 34, 9, 1, 0,
		0, 0, 2, 13, 30,
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

// LanguageAParserInit initializes any static state used to implement LanguageAParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLanguageAParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LanguageAParserInit() {
	staticData := &languageaParserStaticData
	staticData.once.Do(languageaParserInit)
}

// NewLanguageAParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLanguageAParser(input antlr.TokenStream) *LanguageAParser {
	LanguageAParserInit()
	this := new(LanguageAParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &languageaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "LanguageA.g4"

	return this
}

// LanguageAParser tokens.
const (
	LanguageAParserEOF             = antlr.TokenEOF
	LanguageAParserEqual           = 1
	LanguageAParserAdd             = 2
	LanguageAParserSubtract        = 3
	LanguageAParserMultiply        = 4
	LanguageAParserDivide          = 5
	LanguageAParserIdentifier      = 6
	LanguageAParserFloatConstant   = 7
	LanguageAParserIntegerConstant = 8
	LanguageAParserWhitespace      = 9
	LanguageAParserNewline         = 10
)

// LanguageAParser rules.
const (
	LanguageAParserRULE_translation         = 0
	LanguageAParserRULE_topLevelStatement   = 1
	LanguageAParserRULE_variableDeclaration = 2
	LanguageAParserRULE_expression          = 3
	LanguageAParserRULE_binaryOperator      = 4
)

// ITranslationContext is an interface to support dynamic dispatch.
type ITranslationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslationContext differentiates from other interfaces.
	IsTranslationContext()
}

type TranslationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslationContext() *TranslationContext {
	var p = new(TranslationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageAParserRULE_translation
	return p
}

func (*TranslationContext) IsTranslationContext() {}

func NewTranslationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslationContext {
	var p = new(TranslationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageAParserRULE_translation

	return p
}

func (s *TranslationContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslationContext) AllTopLevelStatement() []ITopLevelStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelStatementContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelStatementContext); ok {
			tst[i] = t.(ITopLevelStatementContext)
			i++
		}
	}

	return tst
}

func (s *TranslationContext) TopLevelStatement(i int) ITopLevelStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelStatementContext); ok {
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

	return t.(ITopLevelStatementContext)
}

func (s *TranslationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageAVisitor:
		return t.VisitTranslation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageAParser) Translation() (localctx ITranslationContext) {
	this := p
	_ = this

	localctx = NewTranslationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LanguageAParserRULE_translation)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LanguageAParserIdentifier {
		{
			p.SetState(10)
			p.TopLevelStatement()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopLevelStatementContext is an interface to support dynamic dispatch.
type ITopLevelStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelStatementContext differentiates from other interfaces.
	IsTopLevelStatementContext()
}

type TopLevelStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelStatementContext() *TopLevelStatementContext {
	var p = new(TopLevelStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageAParserRULE_topLevelStatement
	return p
}

func (*TopLevelStatementContext) IsTopLevelStatementContext() {}

func NewTopLevelStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelStatementContext {
	var p = new(TopLevelStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageAParserRULE_topLevelStatement

	return p
}

func (s *TopLevelStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *TopLevelStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageAVisitor:
		return t.VisitTopLevelStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageAParser) TopLevelStatement() (localctx ITopLevelStatementContext) {
	this := p
	_ = this

	localctx = NewTopLevelStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LanguageAParserRULE_topLevelStatement)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.VariableDeclaration()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageAParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageAParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LanguageAParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Equal() antlr.TerminalNode {
	return s.GetToken(LanguageAParserEqual, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
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

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageAVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageAParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	this := p
	_ = this

	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LanguageAParserRULE_variableDeclaration)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(LanguageAParserIdentifier)
	}
	{
		p.SetState(18)
		p.Match(LanguageAParserEqual)
	}
	{
		p.SetState(19)
		p.expression(0)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageAParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageAParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionBinaryContext struct {
	*ExpressionContext
}

func NewExpressionBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBinaryContext {
	var p = new(ExpressionBinaryContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBinaryContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionBinaryContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionBinaryContext) BinaryOperator() IBinaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *ExpressionBinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageAVisitor:
		return t.VisitExpressionBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionConstantContext struct {
	*ExpressionContext
}

func NewExpressionConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionConstantContext {
	var p = new(ExpressionConstantContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionConstantContext) FloatConstant() antlr.TerminalNode {
	return s.GetToken(LanguageAParserFloatConstant, 0)
}

func (s *ExpressionConstantContext) IntegerConstant() antlr.TerminalNode {
	return s.GetToken(LanguageAParserIntegerConstant, 0)
}

func (s *ExpressionConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageAVisitor:
		return t.VisitExpressionConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageAParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LanguageAParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, LanguageAParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewExpressionConstantContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(22)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LanguageAParserFloatConstant || _la == LanguageAParserIntegerConstant) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LanguageAParserRULE_expression)
			p.SetState(24)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(25)
				p.BinaryOperator()
			}
			{
				p.SetState(26)
				p.expression(2)
			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageAParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageAParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) Add() antlr.TerminalNode {
	return s.GetToken(LanguageAParserAdd, 0)
}

func (s *BinaryOperatorContext) Subtract() antlr.TerminalNode {
	return s.GetToken(LanguageAParserSubtract, 0)
}

func (s *BinaryOperatorContext) Multiply() antlr.TerminalNode {
	return s.GetToken(LanguageAParserMultiply, 0)
}

func (s *BinaryOperatorContext) Divide() antlr.TerminalNode {
	return s.GetToken(LanguageAParserDivide, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageAVisitor:
		return t.VisitBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageAParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	this := p
	_ = this

	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LanguageAParserRULE_binaryOperator)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageAParserAdd)|(1<<LanguageAParserSubtract)|(1<<LanguageAParserMultiply)|(1<<LanguageAParserDivide))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *LanguageAParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LanguageAParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
