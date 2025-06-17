parser grammar TuringParser;

options { tokenVocab=TuringLexer; }

programa: instruccion+ EOF;

instruccion: estadoActual COMA simboloLectura COMA simboloEscritura COMA movimiento COMA siguienteEstado;

estadoActual: EN_EL_ESTADO ESTADO;
simboloLectura: SI_SE_LEE_UN (BLANCO | CERO | UNO);
simboloEscritura: SE_ESCRIBE_UN (BLANCO | CERO | UNO);
movimiento: SE_MUEVE_HACIA_LA direccion | NO_SE_MUEVE;
direccion: IZQUIERDA | DERECHA;
siguienteEstado: CONTINUA_EN_EL_ESTADO ESTADO | TERMINA;
