// Code generated from TuringParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // TuringParser

import "github.com/antlr4-go/antlr/v4"

// BaseTuringParserListener is a complete listener for a parse tree produced by TuringParser.
type BaseTuringParserListener struct{}

var _ TuringParserListener = &BaseTuringParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTuringParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTuringParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTuringParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTuringParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseTuringParserListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseTuringParserListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseTuringParserListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseTuringParserListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterEstadoActual is called when production estadoActual is entered.
func (s *BaseTuringParserListener) EnterEstadoActual(ctx *EstadoActualContext) {}

// ExitEstadoActual is called when production estadoActual is exited.
func (s *BaseTuringParserListener) ExitEstadoActual(ctx *EstadoActualContext) {}

// EnterSimboloLectura is called when production simboloLectura is entered.
func (s *BaseTuringParserListener) EnterSimboloLectura(ctx *SimboloLecturaContext) {}

// ExitSimboloLectura is called when production simboloLectura is exited.
func (s *BaseTuringParserListener) ExitSimboloLectura(ctx *SimboloLecturaContext) {}

// EnterSimboloEscritura is called when production simboloEscritura is entered.
func (s *BaseTuringParserListener) EnterSimboloEscritura(ctx *SimboloEscrituraContext) {}

// ExitSimboloEscritura is called when production simboloEscritura is exited.
func (s *BaseTuringParserListener) ExitSimboloEscritura(ctx *SimboloEscrituraContext) {}

// EnterMovimiento is called when production movimiento is entered.
func (s *BaseTuringParserListener) EnterMovimiento(ctx *MovimientoContext) {}

// ExitMovimiento is called when production movimiento is exited.
func (s *BaseTuringParserListener) ExitMovimiento(ctx *MovimientoContext) {}

// EnterDireccion is called when production direccion is entered.
func (s *BaseTuringParserListener) EnterDireccion(ctx *DireccionContext) {}

// ExitDireccion is called when production direccion is exited.
func (s *BaseTuringParserListener) ExitDireccion(ctx *DireccionContext) {}

// EnterSiguienteEstado is called when production siguienteEstado is entered.
func (s *BaseTuringParserListener) EnterSiguienteEstado(ctx *SiguienteEstadoContext) {}

// ExitSiguienteEstado is called when production siguienteEstado is exited.
func (s *BaseTuringParserListener) ExitSiguienteEstado(ctx *SiguienteEstadoContext) {}
