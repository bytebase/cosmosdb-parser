// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CosmosDBParser.
type CosmosDBParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CosmosDBParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select_specification.
	VisitSelect_specification(ctx *Select_specificationContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_specification.
	VisitFrom_specification(ctx *From_specificationContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_source.
	VisitFrom_source(ctx *From_sourceContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#container_expression.
	VisitContainer_expression(ctx *Container_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#container_name.
	VisitContainer_name(ctx *Container_nameContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#object_property_list.
	VisitObject_property_list(ctx *Object_property_listContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#object_property.
	VisitObject_property(ctx *Object_propertyContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#property_alias.
	VisitProperty_alias(ctx *Property_aliasContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#scalar_expression.
	VisitScalar_expression(ctx *Scalar_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#property_name.
	VisitProperty_name(ctx *Property_nameContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#array_index.
	VisitArray_index(ctx *Array_indexContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#input_alias.
	VisitInput_alias(ctx *Input_aliasContext) interface{}
}
