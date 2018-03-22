// Code generated from spec/RCL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package rcl // RCL
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRCLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRCLVisitor) VisitRcl(ctx *RclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitValBool(ctx *ValBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitValInt(ctx *ValIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitValDouble(ctx *ValDoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitValString(ctx *ValStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitValObject(ctx *ValObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCLVisitor) VisitValArray(ctx *ValArrayContext) interface{} {
	return v.VisitChildren(ctx)
}
