// Code generated from TuringParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // TuringParser

import "github.com/antlr4-go/antlr/v4"

// TuringParserListener is a complete listener for a parse tree produced by TuringParser.
type TuringParserListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterEstadoActual is called when entering the estadoActual production.
	EnterEstadoActual(c *EstadoActualContext)

	// EnterSimboloLectura is called when entering the simboloLectura production.
	EnterSimboloLectura(c *SimboloLecturaContext)

	// EnterSimboloEscritura is called when entering the simboloEscritura production.
	EnterSimboloEscritura(c *SimboloEscrituraContext)

	// EnterMovimiento is called when entering the movimiento production.
	EnterMovimiento(c *MovimientoContext)

	// EnterDireccion is called when entering the direccion production.
	EnterDireccion(c *DireccionContext)

	// EnterSiguienteEstado is called when entering the siguienteEstado production.
	EnterSiguienteEstado(c *SiguienteEstadoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitEstadoActual is called when exiting the estadoActual production.
	ExitEstadoActual(c *EstadoActualContext)

	// ExitSimboloLectura is called when exiting the simboloLectura production.
	ExitSimboloLectura(c *SimboloLecturaContext)

	// ExitSimboloEscritura is called when exiting the simboloEscritura production.
	ExitSimboloEscritura(c *SimboloEscrituraContext)

	// ExitMovimiento is called when exiting the movimiento production.
	ExitMovimiento(c *MovimientoContext)

	// ExitDireccion is called when exiting the direccion production.
	ExitDireccion(c *DireccionContext)

	// ExitSiguienteEstado is called when exiting the siguienteEstado production.
	ExitSiguienteEstado(c *SiguienteEstadoContext)
}
