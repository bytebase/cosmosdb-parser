parser grammar CosmosDBParser;

options {
	tokenVocab = CosmosDBLexer;
}

root: select EOF;

select: select_clause from_clause;

select_clause: SELECT_SYMBOL select_specification;

select_specification:
	MULTIPLY_OPERATOR
	| DISTINCT_SYMBOL? object_property_list;

from_clause: FROM_SYMBOL from_specification;

from_specification: from_source;

from_source: container_expression;

container_expression: container_name (AS_SYMBOL? IDENTIFIER)?;

container_name: IDENTIFIER;

object_property_list:
	object_property (COMMA_SYMBOL object_property)*;

object_property: scalar_expression (AS_SYMBOL? property_alias)?;

property_alias: IDENTIFIER;

// scalar_expression: https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/query/scalar-expressions
scalar_expression:
	input_alias
	| scalar_expression DOT_SYMBOL property_name
	| scalar_expression LS_BRACKET_SYMBOL (
		(DOUBLE_QUOTE_SYMBOL property_name DOUBLE_QUOTE_SYMBOL)
		| (array_index)
	) RS_BRACKET_SYMBOL;

property_name: IDENTIFIER;

array_index: DECIMAL_DIGITS;

input_alias: IDENTIFIER;