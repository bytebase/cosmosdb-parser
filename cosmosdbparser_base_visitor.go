// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

type BaseCosmosDBParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCosmosDBParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitSelect_specification(ctx *Select_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitFrom_specification(ctx *From_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitFrom_source(ctx *From_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitContainer_expression(ctx *Container_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitContainer_name(ctx *Container_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitObject_property_list(ctx *Object_property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitObject_property(ctx *Object_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitProperty_alias(ctx *Property_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitScalar_expression(ctx *Scalar_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitProperty_name(ctx *Property_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitArray_index(ctx *Array_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitInput_alias(ctx *Input_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}
