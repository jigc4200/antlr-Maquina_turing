package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/antlr4-go/antlr/v4"

	"turing-machine-go/parser"
	traductorturing "turing-machine-go/traductor-turing"
)

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{Errors: make([]string, 0)}
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("Error de sintaxis en línea %d:%d %s", line, column, msg))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func comprobar(archivo string, path string) bool {
	inf, error := os.ReadFile(path + archivo)
	check(error)

	fmt.Println("Ruta verificada:", path+archivo)

	if inf == nil {
		fmt.Println("Error de Lectura de Archivo de Prueba ")
		return false
	}

	fmt.Println("Prueba de Lectura realizada con éxito")
	return true
}

func lecturaD(archivo string, path string, wg *sync.WaitGroup) {
	var lineasOriginales []string
	defer wg.Done()

	fullPath := path + archivo
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Mostrar contenido
	scanner := bufio.NewScanner(file)
	var contenido string
	for scanner.Scan() {
		line := scanner.Text()
		contenido += line + "\n"
		lineasOriginales = append(lineasOriginales, line) // Capturar líneas originales
	}
	// Eliminar el último salto de línea si existe
	if len(contenido) > 0 && contenido[len(contenido)-1] == '\n' {
		contenido = contenido[:len(contenido)-1]
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Procesar con ANTLR
	input := antlr.NewInputStream(contenido)
	lexer := parser.NewTuringLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTuringParser(tokens)
	// Configurar el parser
	p.RemoveErrorListeners()
	errorListener := NewCustomErrorListener()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true

	tree := p.Programa()

	// Validar si se encontraron errores de sintaxis
	if len(errorListener.Errors) > 0 {
		fmt.Println("Se encontraron errores de sintaxis:")
		for _, msg := range errorListener.Errors {
			fmt.Println(msg)
		}
		fmt.Println("Abortando.")
		return
	}

	// Validar si el árbol de sintaxis fue generado correctamente
	if tree == nil {
		fmt.Println("El árbol de sintaxis es nil. Abortando.")
		return
	}
	// Visitar el árbol
	visitor := traductorturing.NewTraductor()

	fmt.Printf("Tipo del árbol de sintaxis: %T\n", tree)
	// Asegurarse de que el árbol es del tipo esperado antes de llamar a VisitPrograma
	if programaTree, ok := tree.(*parser.ProgramaContext); ok {
		visitor.VisitPrograma(programaTree)
	} else {
		fmt.Printf("Error: El árbol de sintaxis no es de tipo *parser.ProgramaContext, es %T\n", tree)
	}
	fmt.Println("Tuplas generadas:")
	for _, t := range visitor.Tuplas {
		fmt.Println(t)
	}

	// Escribir en archivo
	outputPath := path + "output.txt"
	fmt.Printf("Intentando escribir %d tuplas en %s\n", len(visitor.Tuplas), outputPath)

	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creando archivo de salida:", err)
		return
	}
	defer f.Close()

	for i, tupla := range visitor.Tuplas {
		_, err := f.WriteString(tupla + "\n")
		if err != nil {
			fmt.Printf("Error escribiendo tupla %d: %v\n", i, err)
			return
		}
	}

	htmlPath := path + "output.html"
	htmlFile, err := os.Create(htmlPath)
	if err != nil {
		fmt.Println("Error creando archivo HTML:", err)
		return
	}
	defer htmlFile.Close()

	htmlFile.WriteString("<!DOCTYPE html>\n<html lang=\"es\">\n<head>\n<meta charset=\"UTF-8\">\n<title>Tabla MT</title>\n<style>table {border-collapse: collapse;} td, th {border: 1px solid #aaa; padding: 8px;} th {background-color: #ddd;}</style>\n</head>\n<body>\n")
	htmlFile.WriteString("<h2>Traducción de Instrucciones a Máquina de Turing</h2>\n")
	htmlFile.WriteString("<table>\n<tr><th>Instrucción (Lenguaje Natural)</th><th>Quíntupla (MT)</th></tr>\n")

	for i := range visitor.Tuplas {
		if i < len(lineasOriginales) {
			htmlFile.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", lineasOriginales[i], visitor.Tuplas[i]))
		} else {
			htmlFile.WriteString(fmt.Sprintf("<tr><td>Instrucción %d</td><td>%s</td></tr>\n", i+1, visitor.Tuplas[i]))
		}
	}

	htmlFile.WriteString("</table>\n</body>\n</html>")

	fmt.Println("Archivo HTML generado:", htmlPath)
	fmt.Println("Escritura de tuplas completada.")

	fmt.Println("Traducción ANTLR completada. Archivo generado:", outputPath)
}

func main() {
	const archivo = "test.txt"
	const path = "/home/jigc4200/Documents/antlr-Maquina_turing/"

	fmt.Println(" Turing Machine Simulator - Version 1.0")

	if !comprobar(archivo, path) {
		fmt.Println("Abortando.")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go lecturaD(archivo, path, &wg)
	wg.Wait()
}
