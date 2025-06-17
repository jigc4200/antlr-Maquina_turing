package traductorturing

import (
	"fmt"
	"turing-machine-go/parser" // o el path correcto de tu módulo
	// Importar el paquete antlr
)

type Traductor struct {
	*parser.BaseTuringParserVisitor
	Tuplas []string
}

func NewTraductor() *Traductor {
	return &Traductor{
		BaseTuringParserVisitor: &parser.BaseTuringParserVisitor{},
		Tuplas:                  []string{}, // ← SOLUCIÓN AQUÍ
	}
}

func (v *Traductor) VisitInstruccion(ctx *parser.InstruccionContext) interface{} {
	fmt.Printf("Visitando instruccion: %s\n", ctx.GetText())
	estado := extraerEstado(ctx.EstadoActual().GetText())                 // ej. en_el_estado_a → a
	leer := traducirSimbolo(ctx.SimboloLectura().GetText())               // ej. si_se_lee_un_cero → 0
	escribir := traducirSimbolo(ctx.SimboloEscritura().GetText())         // ej. se_escribe_un_uno → 1
	mov := traducirMovimiento(ctx.Movimiento().GetText())                 // ej. se_mueve_hacia_la_derecha → d
	siguiente := traducirSiguienteEstado(ctx.SiguienteEstado().GetText()) // ej. continua_en_el_estado_b → b

	// Generar quíntupla LT
	tupla := fmt.Sprintf("%s %s %s %s %s", estado, leer, escribir, mov, siguiente)
	v.Tuplas = append(v.Tuplas, tupla)
	return true // En la tarea 2 se recolectarán para archivo
}

func (v *Traductor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	instrucciones := ctx.AllInstruccion()
	fmt.Printf("Visitando Programa. Número de instrucciones encontradas: %d\n", len(instrucciones))
	for i, inst := range instrucciones {
		fmt.Printf("Visitando instrucción %d: Tipo %T, Texto: %s\n", i, inst, inst.GetText())
		// Call VisitInstruccion directly for each instruction context
		if instruccionCtx, ok := inst.(*parser.InstruccionContext); ok {
			v.VisitInstruccion(instruccionCtx)
		} else {
			fmt.Printf("Error: Contexto de instrucción inesperado: %T\n", inst)
		}
	}
	return nil
}

func extraerEstado(texto string) string {
	r := []rune(texto)
	return string(r[len(r)-1]) // último carácter
}

func traducirSimbolo(texto string) string {
	switch {
	case texto == "si_se_lee_un_blanco", texto == "se_escribe_un_blanco":
		return "_"
	case texto == "si_se_lee_un_cero", texto == "se_escribe_un_cero":
		return "0"
	case texto == "si_se_lee_un_uno", texto == "se_escribe_un_uno":
		return "1"
	default:
		return "?" // error
	}
}

func traducirMovimiento(texto string) string {
	switch texto {
	case "se_mueve_hacia_la_izquierda":
		return "i"
	case "se_mueve_hacia_la_derecha":
		return "d"
	case "no_se_mueve":
		return "n"
	default:
		return "?" // error
	}
}

func traducirSiguienteEstado(texto string) string {
	if texto == "termina" {
		return "z" // estado final
	}
	return extraerEstado(texto)
}
