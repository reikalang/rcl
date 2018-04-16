// Code generated from spec/RCL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // RCL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RCLParser.
type RCLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RCLParser#rcl.
	VisitRcl(ctx *RclContext) interface{}

	// Visit a parse tree produced by RCLParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by RCLParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by RCLParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by RCLParser#ValBool.
	VisitValBool(ctx *ValBoolContext) interface{}

	// Visit a parse tree produced by RCLParser#ValInt.
	VisitValInt(ctx *ValIntContext) interface{}

	// Visit a parse tree produced by RCLParser#ValDouble.
	VisitValDouble(ctx *ValDoubleContext) interface{}

	// Visit a parse tree produced by RCLParser#ValString.
	VisitValString(ctx *ValStringContext) interface{}

	// Visit a parse tree produced by RCLParser#ValObject.
	VisitValObject(ctx *ValObjectContext) interface{}

	// Visit a parse tree produced by RCLParser#ValArray.
	VisitValArray(ctx *ValArrayContext) interface{}
}
