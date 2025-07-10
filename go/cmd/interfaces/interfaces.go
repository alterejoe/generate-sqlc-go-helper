package interfaces

import "github.com/dave/dst"

// type GenStandard interface {
// 	GetName() string
// 	GetLowerName() string
// 	GetReturns() *dst.FieldList
// }
//
// type FuncStandard interface {
// 	GetAbbv() string
// 	GetSecondArg() string
// 	GetQueryArgs(params *[]dst.Expr)
// 	GetQueryAddErr(results *[]dst.Expr, function bool)
// }
//
// type StructStandard interface {
// 	GetStructParams() []*dst.Field
// }
//
// type QueryStruct interface {
// 	GenStandard
// 	StructStandard
// }
//
// type QueryFunc interface {
// 	GenStandard
// 	FuncStandard
// }

type Standard interface {
	GetName() string
	GetLowerName() string
	GetAbbv() string
}

type Struct interface {
	Standard
	GetStructParams() []*dst.Field
}

type Function interface {
	Standard
	GetBody() *dst.BlockStmt
	GetFunctionType() *dst.FuncType
	GetReceiver() *dst.FieldList
	GetResults() *dst.FieldList
}
