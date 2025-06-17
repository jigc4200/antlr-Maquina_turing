lexer grammar TuringLexer;

EN_EL_ESTADO: 'en_el_estado_';
SI_SE_LEE_UN: 'si_se_lee_un_';
SE_ESCRIBE_UN: 'se_escribe_un_';
SE_MUEVE_HACIA_LA: 'se_mueve_hacia_la_';
NO_SE_MUEVE: 'no_se_mueve';
CONTINUA_EN_EL_ESTADO: 'continua_en_el_estado_';
TERMINA: 'termina';

BLANCO: 'blanco';
CERO: 'cero';
UNO: 'uno';
IZQUIERDA: 'izquierda';
DERECHA: 'derecha';

COMA: ',';

ESTADO: [a-z] | [1-9][0-9]*;

WS: [ \t\r\n]+ -> skip;
