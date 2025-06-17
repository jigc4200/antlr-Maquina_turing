// Code generated from TuringParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // TuringParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TuringParser.
type TuringParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TuringParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by TuringParser#instruccion.
	VisitInstruccion(ctx *InstruccionContext) interface{}

	// Visit a parse tree produced by TuringParser#estadoActual.
	VisitEstadoActual(ctx *EstadoActualContext) interface{}

	// Visit a parse tree produced by TuringParser#simboloLectura.
	VisitSimboloLectura(ctx *SimboloLecturaContext) interface{}

	// Visit a parse tree produced by TuringParser#simboloEscritura.
	VisitSimboloEscritura(ctx *SimboloEscrituraContext) interface{}

	// Visit a parse tree produced by TuringParser#movimiento.
	VisitMovimiento(ctx *MovimientoContext) interface{}

	// Visit a parse tree produced by TuringParser#direccion.
	VisitDireccion(ctx *DireccionContext) interface{}

	// Visit a parse tree produced by TuringParser#siguienteEstado.
	VisitSiguienteEstado(ctx *SiguienteEstadoContext) interface{}
}
