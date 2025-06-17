// Code generated from TuringParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // TuringParser

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

type TuringParser struct {
	*antlr.BaseParser
}

var TuringParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func turingparserParserInit() {
	staticData := &TuringParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'en_el_estado_'", "'si_se_lee_un_'", "'se_escribe_un_'", "'se_mueve_hacia_la_'",
		"'no_se_mueve'", "'continua_en_el_estado_'", "'termina'", "'blanco'",
		"'cero'", "'uno'", "'izquierda'", "'derecha'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "EN_EL_ESTADO", "SI_SE_LEE_UN", "SE_ESCRIBE_UN", "SE_MUEVE_HACIA_LA",
		"NO_SE_MUEVE", "CONTINUA_EN_EL_ESTADO", "TERMINA", "BLANCO", "CERO",
		"UNO", "IZQUIERDA", "DERECHA", "COMA", "ESTADO", "WS",
	}
	staticData.RuleNames = []string{
		"programa", "instruccion", "estadoActual", "simboloLectura", "simboloEscritura",
		"movimiento", "direccion", "siguienteEstado",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 55, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 18, 8, 0, 11, 0, 12,
		0, 19, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 3, 5, 46, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 53, 8, 7, 1,
		7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 2, 1, 0, 8, 10, 1, 0, 11, 12,
		49, 0, 17, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 36, 1,
		0, 0, 0, 8, 39, 1, 0, 0, 0, 10, 45, 1, 0, 0, 0, 12, 47, 1, 0, 0, 0, 14,
		52, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0,
		0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22,
		5, 0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 24, 3, 4, 2, 0, 24, 25, 5, 13, 0, 0,
		25, 26, 3, 6, 3, 0, 26, 27, 5, 13, 0, 0, 27, 28, 3, 8, 4, 0, 28, 29, 5,
		13, 0, 0, 29, 30, 3, 10, 5, 0, 30, 31, 5, 13, 0, 0, 31, 32, 3, 14, 7, 0,
		32, 3, 1, 0, 0, 0, 33, 34, 5, 1, 0, 0, 34, 35, 5, 14, 0, 0, 35, 5, 1, 0,
		0, 0, 36, 37, 5, 2, 0, 0, 37, 38, 7, 0, 0, 0, 38, 7, 1, 0, 0, 0, 39, 40,
		5, 3, 0, 0, 40, 41, 7, 0, 0, 0, 41, 9, 1, 0, 0, 0, 42, 43, 5, 4, 0, 0,
		43, 46, 3, 12, 6, 0, 44, 46, 5, 5, 0, 0, 45, 42, 1, 0, 0, 0, 45, 44, 1,
		0, 0, 0, 46, 11, 1, 0, 0, 0, 47, 48, 7, 1, 0, 0, 48, 13, 1, 0, 0, 0, 49,
		50, 5, 6, 0, 0, 50, 53, 5, 14, 0, 0, 51, 53, 5, 7, 0, 0, 52, 49, 1, 0,
		0, 0, 52, 51, 1, 0, 0, 0, 53, 15, 1, 0, 0, 0, 3, 19, 45, 52,
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

// TuringParserInit initializes any static state used to implement TuringParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTuringParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TuringParserInit() {
	staticData := &TuringParserParserStaticData
	staticData.once.Do(turingparserParserInit)
}

// NewTuringParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTuringParser(input antlr.TokenStream) *TuringParser {
	TuringParserInit()
	this := new(TuringParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TuringParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TuringParser.g4"

	return this
}

// TuringParser tokens.
const (
	TuringParserEOF                   = antlr.TokenEOF
	TuringParserEN_EL_ESTADO          = 1
	TuringParserSI_SE_LEE_UN          = 2
	TuringParserSE_ESCRIBE_UN         = 3
	TuringParserSE_MUEVE_HACIA_LA     = 4
	TuringParserNO_SE_MUEVE           = 5
	TuringParserCONTINUA_EN_EL_ESTADO = 6
	TuringParserTERMINA               = 7
	TuringParserBLANCO                = 8
	TuringParserCERO                  = 9
	TuringParserUNO                   = 10
	TuringParserIZQUIERDA             = 11
	TuringParserDERECHA               = 12
	TuringParserCOMA                  = 13
	TuringParserESTADO                = 14
	TuringParserWS                    = 15
)

// TuringParser rules.
const (
	TuringParserRULE_programa         = 0
	TuringParserRULE_instruccion      = 1
	TuringParserRULE_estadoActual     = 2
	TuringParserRULE_simboloLectura   = 3
	TuringParserRULE_simboloEscritura = 4
	TuringParserRULE_movimiento       = 5
	TuringParserRULE_direccion        = 6
	TuringParserRULE_siguienteEstado  = 7
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(TuringParserEOF, 0)
}

func (s *ProgramaContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramaContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
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

	return t.(IInstruccionContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TuringParserRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TuringParserEN_EL_ESTADO {
		{
			p.SetState(16)
			p.Instruccion()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(TuringParserEOF)
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

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EstadoActual() IEstadoActualContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode
	SimboloLectura() ISimboloLecturaContext
	SimboloEscritura() ISimboloEscrituraContext
	Movimiento() IMovimientoContext
	SiguienteEstado() ISiguienteEstadoContext

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_instruccion
	return p
}

func InitEmptyInstruccionContext(p *InstruccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_instruccion
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) EstadoActual() IEstadoActualContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstadoActualContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstadoActualContext)
}

func (s *InstruccionContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(TuringParserCOMA)
}

func (s *InstruccionContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserCOMA, i)
}

func (s *InstruccionContext) SimboloLectura() ISimboloLecturaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimboloLecturaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimboloLecturaContext)
}

func (s *InstruccionContext) SimboloEscritura() ISimboloEscrituraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimboloEscrituraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimboloEscrituraContext)
}

func (s *InstruccionContext) Movimiento() IMovimientoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMovimientoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMovimientoContext)
}

func (s *InstruccionContext) SiguienteEstado() ISiguienteEstadoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISiguienteEstadoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISiguienteEstadoContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (s *InstruccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitInstruccion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TuringParserRULE_instruccion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.EstadoActual()
	}
	{
		p.SetState(24)
		p.Match(TuringParserCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.SimboloLectura()
	}
	{
		p.SetState(26)
		p.Match(TuringParserCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.SimboloEscritura()
	}
	{
		p.SetState(28)
		p.Match(TuringParserCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Movimiento()
	}
	{
		p.SetState(30)
		p.Match(TuringParserCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.SiguienteEstado()
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

// IEstadoActualContext is an interface to support dynamic dispatch.
type IEstadoActualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EN_EL_ESTADO() antlr.TerminalNode
	ESTADO() antlr.TerminalNode

	// IsEstadoActualContext differentiates from other interfaces.
	IsEstadoActualContext()
}

type EstadoActualContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEstadoActualContext() *EstadoActualContext {
	var p = new(EstadoActualContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_estadoActual
	return p
}

func InitEmptyEstadoActualContext(p *EstadoActualContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_estadoActual
}

func (*EstadoActualContext) IsEstadoActualContext() {}

func NewEstadoActualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EstadoActualContext {
	var p = new(EstadoActualContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_estadoActual

	return p
}

func (s *EstadoActualContext) GetParser() antlr.Parser { return s.parser }

func (s *EstadoActualContext) EN_EL_ESTADO() antlr.TerminalNode {
	return s.GetToken(TuringParserEN_EL_ESTADO, 0)
}

func (s *EstadoActualContext) ESTADO() antlr.TerminalNode {
	return s.GetToken(TuringParserESTADO, 0)
}

func (s *EstadoActualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EstadoActualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EstadoActualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterEstadoActual(s)
	}
}

func (s *EstadoActualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitEstadoActual(s)
	}
}

func (s *EstadoActualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitEstadoActual(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) EstadoActual() (localctx IEstadoActualContext) {
	localctx = NewEstadoActualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TuringParserRULE_estadoActual)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(TuringParserEN_EL_ESTADO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(TuringParserESTADO)
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

// ISimboloLecturaContext is an interface to support dynamic dispatch.
type ISimboloLecturaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SI_SE_LEE_UN() antlr.TerminalNode
	BLANCO() antlr.TerminalNode
	CERO() antlr.TerminalNode
	UNO() antlr.TerminalNode

	// IsSimboloLecturaContext differentiates from other interfaces.
	IsSimboloLecturaContext()
}

type SimboloLecturaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimboloLecturaContext() *SimboloLecturaContext {
	var p = new(SimboloLecturaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_simboloLectura
	return p
}

func InitEmptySimboloLecturaContext(p *SimboloLecturaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_simboloLectura
}

func (*SimboloLecturaContext) IsSimboloLecturaContext() {}

func NewSimboloLecturaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimboloLecturaContext {
	var p = new(SimboloLecturaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_simboloLectura

	return p
}

func (s *SimboloLecturaContext) GetParser() antlr.Parser { return s.parser }

func (s *SimboloLecturaContext) SI_SE_LEE_UN() antlr.TerminalNode {
	return s.GetToken(TuringParserSI_SE_LEE_UN, 0)
}

func (s *SimboloLecturaContext) BLANCO() antlr.TerminalNode {
	return s.GetToken(TuringParserBLANCO, 0)
}

func (s *SimboloLecturaContext) CERO() antlr.TerminalNode {
	return s.GetToken(TuringParserCERO, 0)
}

func (s *SimboloLecturaContext) UNO() antlr.TerminalNode {
	return s.GetToken(TuringParserUNO, 0)
}

func (s *SimboloLecturaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimboloLecturaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimboloLecturaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterSimboloLectura(s)
	}
}

func (s *SimboloLecturaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitSimboloLectura(s)
	}
}

func (s *SimboloLecturaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitSimboloLectura(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) SimboloLectura() (localctx ISimboloLecturaContext) {
	localctx = NewSimboloLecturaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TuringParserRULE_simboloLectura)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(TuringParserSI_SE_LEE_UN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1792) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ISimboloEscrituraContext is an interface to support dynamic dispatch.
type ISimboloEscrituraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SE_ESCRIBE_UN() antlr.TerminalNode
	BLANCO() antlr.TerminalNode
	CERO() antlr.TerminalNode
	UNO() antlr.TerminalNode

	// IsSimboloEscrituraContext differentiates from other interfaces.
	IsSimboloEscrituraContext()
}

type SimboloEscrituraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimboloEscrituraContext() *SimboloEscrituraContext {
	var p = new(SimboloEscrituraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_simboloEscritura
	return p
}

func InitEmptySimboloEscrituraContext(p *SimboloEscrituraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_simboloEscritura
}

func (*SimboloEscrituraContext) IsSimboloEscrituraContext() {}

func NewSimboloEscrituraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimboloEscrituraContext {
	var p = new(SimboloEscrituraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_simboloEscritura

	return p
}

func (s *SimboloEscrituraContext) GetParser() antlr.Parser { return s.parser }

func (s *SimboloEscrituraContext) SE_ESCRIBE_UN() antlr.TerminalNode {
	return s.GetToken(TuringParserSE_ESCRIBE_UN, 0)
}

func (s *SimboloEscrituraContext) BLANCO() antlr.TerminalNode {
	return s.GetToken(TuringParserBLANCO, 0)
}

func (s *SimboloEscrituraContext) CERO() antlr.TerminalNode {
	return s.GetToken(TuringParserCERO, 0)
}

func (s *SimboloEscrituraContext) UNO() antlr.TerminalNode {
	return s.GetToken(TuringParserUNO, 0)
}

func (s *SimboloEscrituraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimboloEscrituraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimboloEscrituraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterSimboloEscritura(s)
	}
}

func (s *SimboloEscrituraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitSimboloEscritura(s)
	}
}

func (s *SimboloEscrituraContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitSimboloEscritura(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) SimboloEscritura() (localctx ISimboloEscrituraContext) {
	localctx = NewSimboloEscrituraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TuringParserRULE_simboloEscritura)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(TuringParserSE_ESCRIBE_UN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1792) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IMovimientoContext is an interface to support dynamic dispatch.
type IMovimientoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SE_MUEVE_HACIA_LA() antlr.TerminalNode
	Direccion() IDireccionContext
	NO_SE_MUEVE() antlr.TerminalNode

	// IsMovimientoContext differentiates from other interfaces.
	IsMovimientoContext()
}

type MovimientoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMovimientoContext() *MovimientoContext {
	var p = new(MovimientoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_movimiento
	return p
}

func InitEmptyMovimientoContext(p *MovimientoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_movimiento
}

func (*MovimientoContext) IsMovimientoContext() {}

func NewMovimientoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MovimientoContext {
	var p = new(MovimientoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_movimiento

	return p
}

func (s *MovimientoContext) GetParser() antlr.Parser { return s.parser }

func (s *MovimientoContext) SE_MUEVE_HACIA_LA() antlr.TerminalNode {
	return s.GetToken(TuringParserSE_MUEVE_HACIA_LA, 0)
}

func (s *MovimientoContext) Direccion() IDireccionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDireccionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDireccionContext)
}

func (s *MovimientoContext) NO_SE_MUEVE() antlr.TerminalNode {
	return s.GetToken(TuringParserNO_SE_MUEVE, 0)
}

func (s *MovimientoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MovimientoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MovimientoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterMovimiento(s)
	}
}

func (s *MovimientoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitMovimiento(s)
	}
}

func (s *MovimientoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitMovimiento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Movimiento() (localctx IMovimientoContext) {
	localctx = NewMovimientoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TuringParserRULE_movimiento)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TuringParserSE_MUEVE_HACIA_LA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Match(TuringParserSE_MUEVE_HACIA_LA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Direccion()
		}

	case TuringParserNO_SE_MUEVE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(TuringParserNO_SE_MUEVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IDireccionContext is an interface to support dynamic dispatch.
type IDireccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IZQUIERDA() antlr.TerminalNode
	DERECHA() antlr.TerminalNode

	// IsDireccionContext differentiates from other interfaces.
	IsDireccionContext()
}

type DireccionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDireccionContext() *DireccionContext {
	var p = new(DireccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_direccion
	return p
}

func InitEmptyDireccionContext(p *DireccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_direccion
}

func (*DireccionContext) IsDireccionContext() {}

func NewDireccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DireccionContext {
	var p = new(DireccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_direccion

	return p
}

func (s *DireccionContext) GetParser() antlr.Parser { return s.parser }

func (s *DireccionContext) IZQUIERDA() antlr.TerminalNode {
	return s.GetToken(TuringParserIZQUIERDA, 0)
}

func (s *DireccionContext) DERECHA() antlr.TerminalNode {
	return s.GetToken(TuringParserDERECHA, 0)
}

func (s *DireccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DireccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DireccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterDireccion(s)
	}
}

func (s *DireccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitDireccion(s)
	}
}

func (s *DireccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitDireccion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Direccion() (localctx IDireccionContext) {
	localctx = NewDireccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TuringParserRULE_direccion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TuringParserIZQUIERDA || _la == TuringParserDERECHA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ISiguienteEstadoContext is an interface to support dynamic dispatch.
type ISiguienteEstadoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUA_EN_EL_ESTADO() antlr.TerminalNode
	ESTADO() antlr.TerminalNode
	TERMINA() antlr.TerminalNode

	// IsSiguienteEstadoContext differentiates from other interfaces.
	IsSiguienteEstadoContext()
}

type SiguienteEstadoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySiguienteEstadoContext() *SiguienteEstadoContext {
	var p = new(SiguienteEstadoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_siguienteEstado
	return p
}

func InitEmptySiguienteEstadoContext(p *SiguienteEstadoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TuringParserRULE_siguienteEstado
}

func (*SiguienteEstadoContext) IsSiguienteEstadoContext() {}

func NewSiguienteEstadoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SiguienteEstadoContext {
	var p = new(SiguienteEstadoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_siguienteEstado

	return p
}

func (s *SiguienteEstadoContext) GetParser() antlr.Parser { return s.parser }

func (s *SiguienteEstadoContext) CONTINUA_EN_EL_ESTADO() antlr.TerminalNode {
	return s.GetToken(TuringParserCONTINUA_EN_EL_ESTADO, 0)
}

func (s *SiguienteEstadoContext) ESTADO() antlr.TerminalNode {
	return s.GetToken(TuringParserESTADO, 0)
}

func (s *SiguienteEstadoContext) TERMINA() antlr.TerminalNode {
	return s.GetToken(TuringParserTERMINA, 0)
}

func (s *SiguienteEstadoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SiguienteEstadoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SiguienteEstadoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.EnterSiguienteEstado(s)
	}
}

func (s *SiguienteEstadoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringParserListener); ok {
		listenerT.ExitSiguienteEstado(s)
	}
}

func (s *SiguienteEstadoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringParserVisitor:
		return t.VisitSiguienteEstado(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) SiguienteEstado() (localctx ISiguienteEstadoContext) {
	localctx = NewSiguienteEstadoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TuringParserRULE_siguienteEstado)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TuringParserCONTINUA_EN_EL_ESTADO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Match(TuringParserCONTINUA_EN_EL_ESTADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(TuringParserESTADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TuringParserTERMINA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(TuringParserTERMINA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
