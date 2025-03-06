// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// CosmosDBParserListener is a complete listener for a parse tree produced by CosmosDBParser.
type CosmosDBParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterSelect_specification is called when entering the select_specification production.
	EnterSelect_specification(c *Select_specificationContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterFrom_specification is called when entering the from_specification production.
	EnterFrom_specification(c *From_specificationContext)

	// EnterFrom_source is called when entering the from_source production.
	EnterFrom_source(c *From_sourceContext)

	// EnterContainer_expression is called when entering the container_expression production.
	EnterContainer_expression(c *Container_expressionContext)

	// EnterContainer_name is called when entering the container_name production.
	EnterContainer_name(c *Container_nameContext)

	// EnterObject_property_list is called when entering the object_property_list production.
	EnterObject_property_list(c *Object_property_listContext)

	// EnterObject_property is called when entering the object_property production.
	EnterObject_property(c *Object_propertyContext)

	// EnterProperty_alias is called when entering the property_alias production.
	EnterProperty_alias(c *Property_aliasContext)

	// EnterScalar_expression is called when entering the scalar_expression production.
	EnterScalar_expression(c *Scalar_expressionContext)

	// EnterProperty_name is called when entering the property_name production.
	EnterProperty_name(c *Property_nameContext)

	// EnterArray_index is called when entering the array_index production.
	EnterArray_index(c *Array_indexContext)

	// EnterInput_alias is called when entering the input_alias production.
	EnterInput_alias(c *Input_aliasContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitSelect_specification is called when exiting the select_specification production.
	ExitSelect_specification(c *Select_specificationContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitFrom_specification is called when exiting the from_specification production.
	ExitFrom_specification(c *From_specificationContext)

	// ExitFrom_source is called when exiting the from_source production.
	ExitFrom_source(c *From_sourceContext)

	// ExitContainer_expression is called when exiting the container_expression production.
	ExitContainer_expression(c *Container_expressionContext)

	// ExitContainer_name is called when exiting the container_name production.
	ExitContainer_name(c *Container_nameContext)

	// ExitObject_property_list is called when exiting the object_property_list production.
	ExitObject_property_list(c *Object_property_listContext)

	// ExitObject_property is called when exiting the object_property production.
	ExitObject_property(c *Object_propertyContext)

	// ExitProperty_alias is called when exiting the property_alias production.
	ExitProperty_alias(c *Property_aliasContext)

	// ExitScalar_expression is called when exiting the scalar_expression production.
	ExitScalar_expression(c *Scalar_expressionContext)

	// ExitProperty_name is called when exiting the property_name production.
	ExitProperty_name(c *Property_nameContext)

	// ExitArray_index is called when exiting the array_index production.
	ExitArray_index(c *Array_indexContext)

	// ExitInput_alias is called when exiting the input_alias production.
	ExitInput_alias(c *Input_aliasContext)
}
