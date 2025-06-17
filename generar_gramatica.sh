#!/bin/bash

# Ruta al JAR de ANTLR
antlr='java -jar /home/jigc4200/antlr-4.13.1-complete.jar'

# Generar lexer y parser con visitor en lenguaje Go se guardan en la carpeta parser
$antlr -Dlanguage=Go TuringLexer.g4 -o parser    
$antlr -Dlanguage=Go TuringParser.g4 -visitor -o parser
