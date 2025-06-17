// Code generated from TuringParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // TuringParser

import "github.com/antlr4-go/antlr/v4"

type BaseTuringParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTuringParserVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitInstruccion(ctx *InstruccionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitEstadoActual(ctx *EstadoActualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitSimboloLectura(ctx *SimboloLecturaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitSimboloEscritura(ctx *SimboloEscrituraContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitMovimiento(ctx *MovimientoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitDireccion(ctx *DireccionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTuringParserVisitor) VisitSiguienteEstado(ctx *SiguienteEstadoContext) interface{} {
	return v.VisitChildren(ctx)
}
