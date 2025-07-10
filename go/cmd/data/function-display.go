package data

import (
	"go/token"

	"github.com/dave/dst"
)

//	type Function interface {
//		Standard
//		GetBody() *dst.BlockStmt
//		GetFunctionType() *dst.FuncType
//		GetReceiver() *dst.FieldList
//	}

type FunctionData_Display struct {
	StandardData
}

func (qmp *FunctionData_Display) GetReturns() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("string")},
		},
	}
}

func (qmp *FunctionData_Display) GetReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("THIS_IS_WHERE_DB_MODELS_GO")},
		},
	}
}

func (qmp *FunctionData_Display) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{
			&dst.AssignStmt{
				Lhs: dst.NewIdent("LOWER_NAME"),
				Tok: token.DEFINE,
				Rhs: []dst.Expr{
					&dst.CallExpr{
						Fun: &dst.SelectorExpr{
							X:   dst.NewIdent("Query"),
							Sel: dst.NewIdent("SQLC_FUNCTION_NAME"),
						},
						Args: dst.NewIdent("PROBABLY_NEED_NEW_INTERFACE_POINTER_FUNCTION"),
					},
				},
			},
			&dst.ReturnStmt{
				Results: r2,
			},
		},
	}
}
