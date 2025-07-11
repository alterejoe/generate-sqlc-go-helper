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

type Modelgo_DisplayFunction struct {
	StandardData
	Processor
}

// Using the current GetReceiver function
// The 'Reciever' is the (qmp *Modelgo_DisplayFunction)
func (qmp *Modelgo_DisplayFunction) GetReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				// Names: []*dst.Ident{dst.NewIdent("ABBV")},
				Names: []*dst.Ident{dst.NewIdent("r")},

				// Type: dst.NewIdent("FUNCTION_NAME"),
				Type: dst.NewIdent(qmp.GetName()),
			},
		},
	}
}

func (qmp *Modelgo_DisplayFunction) GetFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("r")},
			{Type: dst.NewIdent("FUNCTION_NAME")},
		},
	}
}

// The 'QueryResults' would be a return of a Query function within the generated code
func (qmp *Modelgo_DisplayFunction) GetResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("THIS_IS_WHERE_LOWER_NAME_GOES")},
			{Type: dst.NewIdent("err")},
		},
	}
}

func (qmp *Modelgo_DisplayFunction) GetQueryResults() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("r"),
		dst.NewIdent("THIS_IS_WHERE_PARAMETERS_'COULD'_GO"),
	}
}

func (qmp *Modelgo_DisplayFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{ // this is basically the function body
			qmp.GenerateQuery(), // this is a function call within the body
			qmp.GetBodyReturn(),
		},
	}
}
func (qmp *Modelgo_DisplayFunction) GetQueryArguments() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("r"),
		dst.NewIdent("THIS_IS_WHERE_PARAMETERS_'COULD'_GO"),
	}
}

const QUERY_PACKAGE = "query"

func (qmp *Modelgo_DisplayFunction) GetQueryCall() *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(QUERY_PACKAGE),
		Sel: dst.NewIdent("SQLC_FUNCTION_NAME"),
	}
}

func (qmp *Modelgo_DisplayFunction) GenerateQuery() *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs: qmp.GetQueryResults(),
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun:  qmp.GetQueryCall(),
				Args: qmp.GetQueryArguments(),
			},
		},
	}
}

func (qmp *Modelgo_DisplayFunction) GetBodyReturn() *dst.ReturnStmt {
	return &dst.ReturnStmt{ // this is a return statement to the body
		Results: qmp.GetQueryResults(),
	}
}
