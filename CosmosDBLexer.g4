lexer grammar CosmosDBLexer;

options {
	caseInsensitive = true;
}

fragment A: [a];
fragment B: [b];
fragment C: [c];
fragment D: [d];
fragment E: [e];
fragment F: [f];
fragment G: [g];
fragment H: [h];
fragment I: [i];
fragment J: [j];
fragment K: [k];
fragment L: [l];
fragment M: [m];
fragment N: [n];
fragment O: [o];
fragment P: [p];
fragment Q: [q];
fragment R: [r];
fragment S: [s];
fragment T: [t];
fragment U: [u];
fragment V: [v];
fragment W: [w];
fragment X: [x];
fragment Y: [y];
fragment Z: [z];

MULTIPLY_OPERATOR: '*';

AS_SYMBOL: 'AS';
SELECT_SYMBOL: 'SELECT';
FROM_SYMBOL: 'FROM';
DISTINCT_SYMBOL: 'DISTINCT';

LS_BRACKET_SYMBOL: '[';
RS_BRACKET_SYMBOL: ']';
DOUBLE_QUOTE_SYMBOL: '"';
COMMA_SYMBOL: ',';
DOT_SYMBOL: '.';
/* Identifiers */
IDENTIFIER: [a-z] [a-z_0-9]*;

// White space handling
WHITESPACE:
	[ \t\f\r\n] -> channel(HIDDEN); // Ignore whitespaces.

fragment DECIMAL_DIGIT: [0-9];
DECIMAL_DIGITS: DECIMAL_DIGIT+;