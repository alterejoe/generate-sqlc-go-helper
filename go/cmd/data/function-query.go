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

type FunctionData_Query struct {
	StandardData
}

func (qmp *FunctionData_Query) GetReturns() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *FunctionData_Query) GetReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("ABBV")},
			{Type: dst.NewIdent("SQLC_FUNCTION_NAME")},
		},
	}
}

func (qmp *FunctionData_Query) GetBody() *dst.BlockStmt {
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

func (qmp *FunctionData_Query) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{}
}
