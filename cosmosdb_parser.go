// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
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

type CosmosDBParser struct {
	*antlr.BaseParser
}

var CosmosDBParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cosmosdbparserParserInit() {
	staticData := &CosmosDBParserParserStaticData
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
		"root", "select", "select_clause", "select_specification", "from_clause",
		"from_specification", "from_source", "container_expression", "container_name",
		"object_property_list", "object_property", "property_alias", "scalar_expression",
		"property_name", "array_index", "input_alias",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 109, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3,
		44, 8, 3, 1, 3, 3, 3, 47, 8, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 3, 7, 58, 8, 7, 1, 7, 3, 7, 61, 8, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 5, 9, 68, 8, 9, 10, 9, 12, 9, 71, 9, 9, 1, 10, 1, 10, 3, 10,
		75, 8, 10, 1, 10, 3, 10, 78, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 87, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 5, 12, 98, 8, 12, 10, 12, 12, 12, 101, 9, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 1, 24, 16, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 0, 102, 0, 32, 1, 0, 0,
		0, 2, 35, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 46, 1, 0, 0, 0, 8, 48, 1, 0,
		0, 0, 10, 51, 1, 0, 0, 0, 12, 53, 1, 0, 0, 0, 14, 55, 1, 0, 0, 0, 16, 62,
		1, 0, 0, 0, 18, 64, 1, 0, 0, 0, 20, 72, 1, 0, 0, 0, 22, 79, 1, 0, 0, 0,
		24, 86, 1, 0, 0, 0, 26, 102, 1, 0, 0, 0, 28, 104, 1, 0, 0, 0, 30, 106,
		1, 0, 0, 0, 32, 33, 3, 2, 1, 0, 33, 34, 5, 0, 0, 1, 34, 1, 1, 0, 0, 0,
		35, 36, 3, 4, 2, 0, 36, 37, 3, 8, 4, 0, 37, 3, 1, 0, 0, 0, 38, 39, 5, 3,
		0, 0, 39, 40, 3, 6, 3, 0, 40, 5, 1, 0, 0, 0, 41, 47, 5, 1, 0, 0, 42, 44,
		5, 5, 0, 0, 43, 42, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0,
		45, 47, 3, 18, 9, 0, 46, 41, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 47, 7, 1,
		0, 0, 0, 48, 49, 5, 4, 0, 0, 49, 50, 3, 10, 5, 0, 50, 9, 1, 0, 0, 0, 51,
		52, 3, 12, 6, 0, 52, 11, 1, 0, 0, 0, 53, 54, 3, 14, 7, 0, 54, 13, 1, 0,
		0, 0, 55, 60, 3, 16, 8, 0, 56, 58, 5, 2, 0, 0, 57, 56, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 5, 11, 0, 0, 60, 57, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 15, 1, 0, 0, 0, 62, 63, 5, 11, 0, 0, 63,
		17, 1, 0, 0, 0, 64, 69, 3, 20, 10, 0, 65, 66, 5, 9, 0, 0, 66, 68, 3, 20,
		10, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69,
		70, 1, 0, 0, 0, 70, 19, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 77, 3, 24,
		12, 0, 73, 75, 5, 2, 0, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 78, 3, 22, 11, 0, 77, 74, 1, 0, 0, 0, 77, 78, 1, 0,
		0, 0, 78, 21, 1, 0, 0, 0, 79, 80, 5, 11, 0, 0, 80, 23, 1, 0, 0, 0, 81,
		82, 6, 12, -1, 0, 82, 87, 3, 30, 15, 0, 83, 84, 3, 28, 14, 0, 84, 85, 5,
		7, 0, 0, 85, 87, 1, 0, 0, 0, 86, 81, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 87,
		99, 1, 0, 0, 0, 88, 89, 10, 3, 0, 0, 89, 90, 5, 10, 0, 0, 90, 98, 3, 26,
		13, 0, 91, 92, 10, 2, 0, 0, 92, 93, 5, 6, 0, 0, 93, 94, 5, 8, 0, 0, 94,
		95, 3, 26, 13, 0, 95, 96, 5, 8, 0, 0, 96, 98, 1, 0, 0, 0, 97, 88, 1, 0,
		0, 0, 97, 91, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 25, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 103, 5,
		11, 0, 0, 103, 27, 1, 0, 0, 0, 104, 105, 5, 13, 0, 0, 105, 29, 1, 0, 0,
		0, 106, 107, 5, 11, 0, 0, 107, 31, 1, 0, 0, 0, 10, 43, 46, 57, 60, 69,
		74, 77, 86, 97, 99,
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

// CosmosDBParserInit initializes any static state used to implement CosmosDBParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCosmosDBParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CosmosDBParserInit() {
	staticData := &CosmosDBParserParserStaticData
	staticData.once.Do(cosmosdbparserParserInit)
}

// NewCosmosDBParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCosmosDBParser(input antlr.TokenStream) *CosmosDBParser {
	CosmosDBParserInit()
	this := new(CosmosDBParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CosmosDBParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CosmosDBParser.g4"

	return this
}

// CosmosDBParser tokens.
const (
	CosmosDBParserEOF                 = antlr.TokenEOF
	CosmosDBParserMULTIPLY_OPERATOR   = 1
	CosmosDBParserAS_SYMBOL           = 2
	CosmosDBParserSELECT_SYMBOL       = 3
	CosmosDBParserFROM_SYMBOL         = 4
	CosmosDBParserDISTINCT_SYMBOL     = 5
	CosmosDBParserLS_BRACKET_SYMBOL   = 6
	CosmosDBParserRS_BRACKET_SYMBOL   = 7
	CosmosDBParserDOUBLE_QUOTE_SYMBOL = 8
	CosmosDBParserCOMMA_SYMBOL        = 9
	CosmosDBParserDOT_SYMBOL          = 10
	CosmosDBParserIDENTIFIER          = 11
	CosmosDBParserWHITESPACE          = 12
	CosmosDBParserDECIMAL_DIGITS      = 13
)

// CosmosDBParser rules.
const (
	CosmosDBParserRULE_root                 = 0
	CosmosDBParserRULE_select               = 1
	CosmosDBParserRULE_select_clause        = 2
	CosmosDBParserRULE_select_specification = 3
	CosmosDBParserRULE_from_clause          = 4
	CosmosDBParserRULE_from_specification   = 5
	CosmosDBParserRULE_from_source          = 6
	CosmosDBParserRULE_container_expression = 7
	CosmosDBParserRULE_container_name       = 8
	CosmosDBParserRULE_object_property_list = 9
	CosmosDBParserRULE_object_property      = 10
	CosmosDBParserRULE_property_alias       = 11
	CosmosDBParserRULE_scalar_expression    = 12
	CosmosDBParserRULE_property_name        = 13
	CosmosDBParserRULE_array_index          = 14
	CosmosDBParserRULE_input_alias          = 15
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_() ISelectContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Select_() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CosmosDBParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Select_()
	}
	{
		p.SetState(33)
		p.Match(CosmosDBParserEOF)
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

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_clause() ISelect_clauseContext
	From_clause() IFrom_clauseContext

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select
	return p
}

func InitEmptySelectContext(p *SelectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) Select_clause() ISelect_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_clauseContext)
}

func (s *SelectContext) From_clause() IFrom_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_clauseContext)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_() (localctx ISelectContext) {
	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CosmosDBParserRULE_select)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Select_clause()
	}
	{
		p.SetState(36)
		p.From_clause()
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

// ISelect_clauseContext is an interface to support dynamic dispatch.
type ISelect_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT_SYMBOL() antlr.TerminalNode
	Select_specification() ISelect_specificationContext

	// IsSelect_clauseContext differentiates from other interfaces.
	IsSelect_clauseContext()
}

type Select_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_clauseContext() *Select_clauseContext {
	var p = new(Select_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_clause
	return p
}

func InitEmptySelect_clauseContext(p *Select_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_clause
}

func (*Select_clauseContext) IsSelect_clauseContext() {}

func NewSelect_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_clauseContext {
	var p = new(Select_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select_clause

	return p
}

func (s *Select_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_clauseContext) SELECT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserSELECT_SYMBOL, 0)
}

func (s *Select_clauseContext) Select_specification() ISelect_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_specificationContext)
}

func (s *Select_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect_clause(s)
	}
}

func (s *Select_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect_clause(s)
	}
}

func (s *Select_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_clause() (localctx ISelect_clauseContext) {
	localctx = NewSelect_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CosmosDBParserRULE_select_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(CosmosDBParserSELECT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Select_specification()
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

// ISelect_specificationContext is an interface to support dynamic dispatch.
type ISelect_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLY_OPERATOR() antlr.TerminalNode
	Object_property_list() IObject_property_listContext
	DISTINCT_SYMBOL() antlr.TerminalNode

	// IsSelect_specificationContext differentiates from other interfaces.
	IsSelect_specificationContext()
}

type Select_specificationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_specificationContext() *Select_specificationContext {
	var p = new(Select_specificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_specification
	return p
}

func InitEmptySelect_specificationContext(p *Select_specificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_specification
}

func (*Select_specificationContext) IsSelect_specificationContext() {}

func NewSelect_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_specificationContext {
	var p = new(Select_specificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select_specification

	return p
}

func (s *Select_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_specificationContext) MULTIPLY_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMULTIPLY_OPERATOR, 0)
}

func (s *Select_specificationContext) Object_property_list() IObject_property_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_property_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_property_listContext)
}

func (s *Select_specificationContext) DISTINCT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDISTINCT_SYMBOL, 0)
}

func (s *Select_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect_specification(s)
	}
}

func (s *Select_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect_specification(s)
	}
}

func (s *Select_specificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect_specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_specification() (localctx ISelect_specificationContext) {
	localctx = NewSelect_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CosmosDBParserRULE_select_specification)
	var _la int

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserMULTIPLY_OPERATOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(CosmosDBParserMULTIPLY_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CosmosDBParserDISTINCT_SYMBOL, CosmosDBParserIDENTIFIER, CosmosDBParserDECIMAL_DIGITS:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserDISTINCT_SYMBOL {
			{
				p.SetState(42)
				p.Match(CosmosDBParserDISTINCT_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(45)
			p.Object_property_list()
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

// IFrom_clauseContext is an interface to support dynamic dispatch.
type IFrom_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM_SYMBOL() antlr.TerminalNode
	From_specification() IFrom_specificationContext

	// IsFrom_clauseContext differentiates from other interfaces.
	IsFrom_clauseContext()
}

type From_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_clauseContext() *From_clauseContext {
	var p = new(From_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_clause
	return p
}

func InitEmptyFrom_clauseContext(p *From_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_clause
}

func (*From_clauseContext) IsFrom_clauseContext() {}

func NewFrom_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_clauseContext {
	var p = new(From_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_clause

	return p
}

func (s *From_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *From_clauseContext) FROM_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserFROM_SYMBOL, 0)
}

func (s *From_clauseContext) From_specification() IFrom_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_specificationContext)
}

func (s *From_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_clause(s)
	}
}

func (s *From_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_clause(s)
	}
}

func (s *From_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_clause() (localctx IFrom_clauseContext) {
	localctx = NewFrom_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CosmosDBParserRULE_from_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(CosmosDBParserFROM_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.From_specification()
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

// IFrom_specificationContext is an interface to support dynamic dispatch.
type IFrom_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	From_source() IFrom_sourceContext

	// IsFrom_specificationContext differentiates from other interfaces.
	IsFrom_specificationContext()
}

type From_specificationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_specificationContext() *From_specificationContext {
	var p = new(From_specificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_specification
	return p
}

func InitEmptyFrom_specificationContext(p *From_specificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_specification
}

func (*From_specificationContext) IsFrom_specificationContext() {}

func NewFrom_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_specificationContext {
	var p = new(From_specificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_specification

	return p
}

func (s *From_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *From_specificationContext) From_source() IFrom_sourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_sourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_sourceContext)
}

func (s *From_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_specification(s)
	}
}

func (s *From_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_specification(s)
	}
}

func (s *From_specificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_specification() (localctx IFrom_specificationContext) {
	localctx = NewFrom_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CosmosDBParserRULE_from_specification)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.From_source()
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

// IFrom_sourceContext is an interface to support dynamic dispatch.
type IFrom_sourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_expression() IContainer_expressionContext

	// IsFrom_sourceContext differentiates from other interfaces.
	IsFrom_sourceContext()
}

type From_sourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_sourceContext() *From_sourceContext {
	var p = new(From_sourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_source
	return p
}

func InitEmptyFrom_sourceContext(p *From_sourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_source
}

func (*From_sourceContext) IsFrom_sourceContext() {}

func NewFrom_sourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_sourceContext {
	var p = new(From_sourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_source

	return p
}

func (s *From_sourceContext) GetParser() antlr.Parser { return s.parser }

func (s *From_sourceContext) Container_expression() IContainer_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_expressionContext)
}

func (s *From_sourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_sourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_sourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_source(s)
	}
}

func (s *From_sourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_source(s)
	}
}

func (s *From_sourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_source(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_source() (localctx IFrom_sourceContext) {
	localctx = NewFrom_sourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CosmosDBParserRULE_from_source)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Container_expression()
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

// IContainer_expressionContext is an interface to support dynamic dispatch.
type IContainer_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_name() IContainer_nameContext
	IDENTIFIER() antlr.TerminalNode
	AS_SYMBOL() antlr.TerminalNode

	// IsContainer_expressionContext differentiates from other interfaces.
	IsContainer_expressionContext()
}

type Container_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_expressionContext() *Container_expressionContext {
	var p = new(Container_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_expression
	return p
}

func InitEmptyContainer_expressionContext(p *Container_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_expression
}

func (*Container_expressionContext) IsContainer_expressionContext() {}

func NewContainer_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_expressionContext {
	var p = new(Container_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_container_expression

	return p
}

func (s *Container_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_expressionContext) Container_name() IContainer_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_nameContext)
}

func (s *Container_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Container_expressionContext) AS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAS_SYMBOL, 0)
}

func (s *Container_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterContainer_expression(s)
	}
}

func (s *Container_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitContainer_expression(s)
	}
}

func (s *Container_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitContainer_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Container_expression() (localctx IContainer_expressionContext) {
	localctx = NewContainer_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CosmosDBParserRULE_container_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Container_name()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserAS_SYMBOL || _la == CosmosDBParserIDENTIFIER {
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserAS_SYMBOL {
			{
				p.SetState(56)
				p.Match(CosmosDBParserAS_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(59)
			p.Match(CosmosDBParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IContainer_nameContext is an interface to support dynamic dispatch.
type IContainer_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsContainer_nameContext differentiates from other interfaces.
	IsContainer_nameContext()
}

type Container_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_nameContext() *Container_nameContext {
	var p = new(Container_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_name
	return p
}

func InitEmptyContainer_nameContext(p *Container_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_name
}

func (*Container_nameContext) IsContainer_nameContext() {}

func NewContainer_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_nameContext {
	var p = new(Container_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_container_name

	return p
}

func (s *Container_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Container_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterContainer_name(s)
	}
}

func (s *Container_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitContainer_name(s)
	}
}

func (s *Container_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitContainer_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Container_name() (localctx IContainer_nameContext) {
	localctx = NewContainer_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CosmosDBParserRULE_container_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(CosmosDBParserIDENTIFIER)
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

// IObject_property_listContext is an interface to support dynamic dispatch.
type IObject_property_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllObject_property() []IObject_propertyContext
	Object_property(i int) IObject_propertyContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsObject_property_listContext differentiates from other interfaces.
	IsObject_property_listContext()
}

type Object_property_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_property_listContext() *Object_property_listContext {
	var p = new(Object_property_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property_list
	return p
}

func InitEmptyObject_property_listContext(p *Object_property_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property_list
}

func (*Object_property_listContext) IsObject_property_listContext() {}

func NewObject_property_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_property_listContext {
	var p = new(Object_property_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_property_list

	return p
}

func (s *Object_property_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_property_listContext) AllObject_property() []IObject_propertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_propertyContext); ok {
			len++
		}
	}

	tst := make([]IObject_propertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_propertyContext); ok {
			tst[i] = t.(IObject_propertyContext)
			i++
		}
	}

	return tst
}

func (s *Object_property_listContext) Object_property(i int) IObject_propertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_propertyContext); ok {
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

	return t.(IObject_propertyContext)
}

func (s *Object_property_listContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Object_property_listContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Object_property_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_property_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_property_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_property_list(s)
	}
}

func (s *Object_property_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_property_list(s)
	}
}

func (s *Object_property_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_property_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_property_list() (localctx IObject_property_listContext) {
	localctx = NewObject_property_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CosmosDBParserRULE_object_property_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Object_property()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(65)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Object_property()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IObject_propertyContext is an interface to support dynamic dispatch.
type IObject_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Scalar_expression() IScalar_expressionContext
	Property_alias() IProperty_aliasContext
	AS_SYMBOL() antlr.TerminalNode

	// IsObject_propertyContext differentiates from other interfaces.
	IsObject_propertyContext()
}

type Object_propertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_propertyContext() *Object_propertyContext {
	var p = new(Object_propertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property
	return p
}

func InitEmptyObject_propertyContext(p *Object_propertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property
}

func (*Object_propertyContext) IsObject_propertyContext() {}

func NewObject_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_propertyContext {
	var p = new(Object_propertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_property

	return p
}

func (s *Object_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_propertyContext) Scalar_expression() IScalar_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expressionContext)
}

func (s *Object_propertyContext) Property_alias() IProperty_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_aliasContext)
}

func (s *Object_propertyContext) AS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAS_SYMBOL, 0)
}

func (s *Object_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_property(s)
	}
}

func (s *Object_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_property(s)
	}
}

func (s *Object_propertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_property(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_property() (localctx IObject_propertyContext) {
	localctx = NewObject_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CosmosDBParserRULE_object_property)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.scalar_expression(0)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserAS_SYMBOL || _la == CosmosDBParserIDENTIFIER {
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserAS_SYMBOL {
			{
				p.SetState(73)
				p.Match(CosmosDBParserAS_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(76)
			p.Property_alias()
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

// IProperty_aliasContext is an interface to support dynamic dispatch.
type IProperty_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsProperty_aliasContext differentiates from other interfaces.
	IsProperty_aliasContext()
}

type Property_aliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_aliasContext() *Property_aliasContext {
	var p = new(Property_aliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_alias
	return p
}

func InitEmptyProperty_aliasContext(p *Property_aliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_alias
}

func (*Property_aliasContext) IsProperty_aliasContext() {}

func NewProperty_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_aliasContext {
	var p = new(Property_aliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_property_alias

	return p
}

func (s *Property_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Property_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterProperty_alias(s)
	}
}

func (s *Property_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitProperty_alias(s)
	}
}

func (s *Property_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitProperty_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Property_alias() (localctx IProperty_aliasContext) {
	localctx = NewProperty_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CosmosDBParserRULE_property_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(CosmosDBParserIDENTIFIER)
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

// IScalar_expressionContext is an interface to support dynamic dispatch.
type IScalar_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Input_alias() IInput_aliasContext
	Array_index() IArray_indexContext
	RS_BRACKET_SYMBOL() antlr.TerminalNode
	Scalar_expression() IScalar_expressionContext
	DOT_SYMBOL() antlr.TerminalNode
	Property_name() IProperty_nameContext
	LS_BRACKET_SYMBOL() antlr.TerminalNode
	AllDOUBLE_QUOTE_SYMBOL() []antlr.TerminalNode
	DOUBLE_QUOTE_SYMBOL(i int) antlr.TerminalNode

	// IsScalar_expressionContext differentiates from other interfaces.
	IsScalar_expressionContext()
}

type Scalar_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalar_expressionContext() *Scalar_expressionContext {
	var p = new(Scalar_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_expression
	return p
}

func InitEmptyScalar_expressionContext(p *Scalar_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_expression
}

func (*Scalar_expressionContext) IsScalar_expressionContext() {}

func NewScalar_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scalar_expressionContext {
	var p = new(Scalar_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_scalar_expression

	return p
}

func (s *Scalar_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Scalar_expressionContext) Input_alias() IInput_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_aliasContext)
}

func (s *Scalar_expressionContext) Array_index() IArray_indexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_indexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_indexContext)
}

func (s *Scalar_expressionContext) RS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRS_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expressionContext) Scalar_expression() IScalar_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expressionContext)
}

func (s *Scalar_expressionContext) DOT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOT_SYMBOL, 0)
}

func (s *Scalar_expressionContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *Scalar_expressionContext) LS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLS_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expressionContext) AllDOUBLE_QUOTE_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserDOUBLE_QUOTE_SYMBOL)
}

func (s *Scalar_expressionContext) DOUBLE_QUOTE_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOUBLE_QUOTE_SYMBOL, i)
}

func (s *Scalar_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scalar_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scalar_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterScalar_expression(s)
	}
}

func (s *Scalar_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitScalar_expression(s)
	}
}

func (s *Scalar_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitScalar_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Scalar_expression() (localctx IScalar_expressionContext) {
	return p.scalar_expression(0)
}

func (p *CosmosDBParser) scalar_expression(_p int) (localctx IScalar_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewScalar_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalar_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, CosmosDBParserRULE_scalar_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserIDENTIFIER:
		{
			p.SetState(82)
			p.Input_alias()
		}

	case CosmosDBParserDECIMAL_DIGITS:
		{
			p.SetState(83)
			p.Array_index()
		}
		{
			p.SetState(84)
			p.Match(CosmosDBParserRS_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(89)
					p.Match(CosmosDBParserDOT_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(90)
					p.Property_name()
				}

			case 2:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(92)
					p.Match(CosmosDBParserLS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(93)
					p.Match(CosmosDBParserDOUBLE_QUOTE_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)
					p.Property_name()
				}
				{
					p.SetState(95)
					p.Match(CosmosDBParserDOUBLE_QUOTE_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

// IProperty_nameContext is an interface to support dynamic dispatch.
type IProperty_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsProperty_nameContext differentiates from other interfaces.
	IsProperty_nameContext()
}

type Property_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_nameContext() *Property_nameContext {
	var p = new(Property_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_name
	return p
}

func InitEmptyProperty_nameContext(p *Property_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_name
}

func (*Property_nameContext) IsProperty_nameContext() {}

func NewProperty_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_nameContext {
	var p = new(Property_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_property_name

	return p
}

func (s *Property_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Property_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterProperty_name(s)
	}
}

func (s *Property_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitProperty_name(s)
	}
}

func (s *Property_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitProperty_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Property_name() (localctx IProperty_nameContext) {
	localctx = NewProperty_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CosmosDBParserRULE_property_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(CosmosDBParserIDENTIFIER)
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

// IArray_indexContext is an interface to support dynamic dispatch.
type IArray_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_DIGITS() antlr.TerminalNode

	// IsArray_indexContext differentiates from other interfaces.
	IsArray_indexContext()
}

type Array_indexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_indexContext() *Array_indexContext {
	var p = new(Array_indexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_array_index
	return p
}

func InitEmptyArray_indexContext(p *Array_indexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_array_index
}

func (*Array_indexContext) IsArray_indexContext() {}

func NewArray_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_indexContext {
	var p = new(Array_indexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_array_index

	return p
}

func (s *Array_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_indexContext) DECIMAL_DIGITS() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDECIMAL_DIGITS, 0)
}

func (s *Array_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterArray_index(s)
	}
}

func (s *Array_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitArray_index(s)
	}
}

func (s *Array_indexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitArray_index(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Array_index() (localctx IArray_indexContext) {
	localctx = NewArray_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CosmosDBParserRULE_array_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(CosmosDBParserDECIMAL_DIGITS)
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

// IInput_aliasContext is an interface to support dynamic dispatch.
type IInput_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsInput_aliasContext differentiates from other interfaces.
	IsInput_aliasContext()
}

type Input_aliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInput_aliasContext() *Input_aliasContext {
	var p = new(Input_aliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_input_alias
	return p
}

func InitEmptyInput_aliasContext(p *Input_aliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_input_alias
}

func (*Input_aliasContext) IsInput_aliasContext() {}

func NewInput_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_aliasContext {
	var p = new(Input_aliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_input_alias

	return p
}

func (s *Input_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Input_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterInput_alias(s)
	}
}

func (s *Input_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitInput_alias(s)
	}
}

func (s *Input_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitInput_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Input_alias() (localctx IInput_aliasContext) {
	localctx = NewInput_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CosmosDBParserRULE_input_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(CosmosDBParserIDENTIFIER)
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

func (p *CosmosDBParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *Scalar_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Scalar_expressionContext)
		}
		return p.Scalar_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CosmosDBParser) Scalar_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
