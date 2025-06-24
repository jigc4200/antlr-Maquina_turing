# Simulador/Traductor de Máquina de Turing con ANTLR y Go

Este proyecto es un simulador y traductor de instrucciones de Máquina de Turing, implementado utilizando ANTLR4 para el análisis sintáctico y léxico, y Go para la lógica de la aplicación. Permite a los usuarios definir instrucciones de Máquina de Turing en un lenguaje natural simplificado y las traduce a quíntuplas (el formato estándar de la Máquina de Turing), generando una salida en texto plano y un archivo HTML para una visualización más amigable.

## Características

*   **Análisis de Lenguaje Natural**: Procesa instrucciones de Máquina de Turing escritas en un formato legible por humanos.
*   **Generación de Quíntuplas**: Convierte las instrucciones a su representación formal de quíntuplas (estado_actual, símbolo_leído, nuevo_símbolo, movimiento, nuevo_estado).
*   **Validación de Sintaxis**: Utiliza ANTLR para validar que las instrucciones sigan la gramática definida.
*   **Salida en Texto y HTML**: Genera un archivo de texto (`output.txt`) con las quíntuplas y un archivo HTML (`output.html`) que muestra las instrucciones originales junto a sus quíntuplas correspondientes en una tabla.
*   **Fácil de Extender**: La gramática ANTLR puede ser modificada para soportar nuevas instrucciones o variaciones.

## Requisitos Previos

Antes de empezar, asegúrate de tener instalado lo siguiente:

*   **Go**: Versión 1.16 o superior. Puedes descargarlo desde [golang.org](https://golang.org/dl/).
*   **Java Development Kit (JDK)**: Necesario para ejecutar ANTLR. Puedes descargarlo desde el sitio oficial de Oracle o usar una alternativa de código abierto como OpenJDK.
*   **ANTLR 4 JAR**: El archivo `antlr-4.13.1-complete.jar` (o la versión más reciente). Puedes descargarlo desde la página de lanzamientos de ANTLR en GitHub: [ANTLR Releases](https://github.com/antlr/antlr4/releases). Asegúrate de colocar este archivo en una ubicación accesible, por ejemplo, en tu directorio de usuario o en el directorio raíz del proyecto. Para este proyecto, se asume que está en `/home/jigc4200/antlr-4.13.1-complete.jar`.

## Instalación y Configuración

Sigue estos pasos para configurar y ejecutar el proyecto en tu máquina local:

1.  **Clonar el Repositorio (si aplica)**:
    Si este código proviene de un repositorio, clónalo:
    ```bash
    git clone [URL_DEL_REPOSITORIO]
    cd antlr-Maquina_turing
    ```
    Si ya tienes los archivos, asegúrate de estar en el directorio raíz del proyecto (`antlr-Maquina_turing`).

2.  **Inicializar el Módulo Go**:
    Abre una terminal en el directorio raíz del proyecto y ejecuta:
    ```bash
    go mod init turing-machine-go
    ```
    Si ya existe un archivo `go.mod`, este paso puede no ser necesario.

3.  **Descargar Dependencias de Go**:
    Obtén el paquete de ANTLR para Go:
    ```bash
    go get github.com/antlr/antlr4/runtime/Go/antlr
    ```

4.  **Generar el Lexer y Parser de ANTLR**:
    Este proyecto utiliza gramáticas ANTLR (`TuringLexer.g4` y `TuringParser.g4`) para procesar las instrucciones. Necesitas generar los archivos Go correspondientes.

    *   **Asegúrate de que el archivo `generar_gramatica.sh` sea ejecutable**:
        ```bash
        chmod +x generar_gramatica.sh
        ```

    *   **Ejecuta el script para generar los archivos del parser**:
        ```bash
        ./generar_gramatica.sh
        ```
        Este script utilizará el JAR de ANTLR para generar los archivos Go en el directorio `parser/`. Asegúrate de que la ruta al JAR de ANTLR en el script (`antlr='java -jar /home/jigc4200/antlr-4.13.1-complete.jar'`) sea correcta para tu sistema.

## Uso

Para ejecutar el traductor de Máquina de Turing:

1.  **Prepara tu archivo de entrada**:
    El programa espera un archivo llamado `test.txt` en el directorio raíz del proyecto (`/home/jigc4200/Documents/antlr-Maquina_turing/`). Este archivo debe contener las instrucciones de la Máquina de Turing que deseas traducir, una instrucción por línea.

    **Ejemplo de `test.txt`:**
    ```
    estado q0, lee 0, escribe 1, mueve derecha, nuevo estado q1
    estado q1, lee 1, escribe 0, mueve izquierda, nuevo estado q0
    estado q1, lee blanco, escribe blanco, mueve alto, nuevo estado q_final
    ```

2.  **Ejecuta el programa Go**:
    Desde el directorio raíz del proyecto, ejecuta:
    ```bash
    go run main.go
    ```

3.  **Verifica la Salida**:
    Después de la ejecución, se generarán dos archivos en el mismo directorio:
    *   `output.txt`: Contiene las quíntuplas generadas, una por línea.
    *   `output.html`: Un archivo HTML que muestra una tabla con las instrucciones originales y sus quíntuplas traducidas. Puedes abrir este archivo en cualquier navegador web para una visualización clara.

## Estructura del Proyecto

*   `main.go`: El punto de entrada principal de la aplicación. Contiene la lógica para leer el archivo de entrada, invocar el lexer y parser de ANTLR, procesar el árbol de sintaxis y generar los archivos de salida.
*   `TuringLexer.g4`: Define el lexer para la gramática de la Máquina de Turing.
*   `TuringParser.g4`: Define el parser para la gramática de la Máquina de Turing.
*   `generar_gramatica.sh`: Script de shell para compilar las gramáticas ANTLR y generar los archivos Go del lexer y parser.
*   `parser/`: Directorio que contiene los archivos Go generados por ANTLR (lexer, parser, listener, visitor).
*   `traductor-turing/traductor.go`: Contiene la implementación del `visitor` de ANTLR que recorre el árbol de sintaxis y realiza la traducción de las instrucciones a quíntuplas.
*   `test.txt`: Archivo de entrada de ejemplo con instrucciones de Máquina de Turing.
*   `output.txt`: Archivo de salida generado con las quíntuplas.
*   `output.html`: Archivo HTML generado para la visualización de las instrucciones y quíntuplas.
*   `DETALLES_TECNICOS.TXT`: Archivo con notas técnicas y pasos iniciales para el desarrollo del proyecto.

## Contribución

Si deseas contribuir a este proyecto, por favor, sigue los siguientes pasos:

1.  Haz un "fork" del repositorio.
2.  Crea una nueva rama (`git checkout -b feature/nueva-caracteristica`).
3.  Realiza tus cambios y haz "commit" de ellos (`git commit -am 'Agrega nueva característica'`).
4.  Sube tus cambios a la rama (`git push origin feature/nueva-caracteristica`).
5.  Abre un "Pull Request".