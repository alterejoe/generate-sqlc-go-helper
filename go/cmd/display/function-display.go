package display

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/parse"
	"github.com/dave/dst"
)

//	type Function interface {
//		Standard
//		GetBody() *dst.BlockStmt
//		GetFunctionType() *dst.FuncTypeGetBody
//		GetReceiver() *dst.FieldList
//	}

type TemplateModelgo_Function struct {
	parse.StandardData
}

// Using the current GetReceiver function
// The 'Reciever' is the (qmp *Modelgo_DisplayFunction)
func (qmp *TemplateModelgo_Function) GetReceiver() *dst.FieldList {
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

func (qmp *TemplateModelgo_Function) GetFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("r")},
			{Type: dst.NewIdent("FUNCTION_NAME")},
		},
	}
}

// The 'QueryResults' would be a return of a Query function within the generated code
func (qmp *TemplateModelgo_Function) GetResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("THIS_IS_WHERE_LOWER_NAME_GOES")},
			{Type: dst.NewIdent("err")},
		},
	}
}

func (qmp *TemplateModelgo_Function) GetQueryResults() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("r"),
		dst.NewIdent("THIS_IS_WHERE_PARAMETERS_'COULD'_GO"),
	}
}

func (qmp *TemplateModelgo_Function) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{ // this is basically the function body
			qmp.GenerateQuery(), // this is a function call within the body
			qmp.GetBodyReturn(),
		},
	}
}
func (qmp *TemplateModelgo_Function) GetQueryArguments() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("r"),
		dst.NewIdent("THIS_IS_WHERE_PARAMETERS_'COULD'_GO"),
	}
}

const QUERY_PACKAGE = "query"

func (qmp *TemplateModelgo_Function) GetQueryCall() *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(QUERY_PACKAGE),
		Sel: dst.NewIdent("SQLC_FUNCTION_NAME"),
	}
}

func (qmp *TemplateModelgo_Function) GenerateQuery() *dst.AssignStmt {
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

func (qmp *TemplateModelgo_Function) GetBodyReturn() *dst.ReturnStmt {
	return &dst.ReturnStmt{ // this is a return statement to the body
		Results: qmp.GetQueryResults(),
	}
}
