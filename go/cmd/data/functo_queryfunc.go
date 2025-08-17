package data

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it easier to sort and delegate data to its respective parser
func FuncToQueryFunction(f *dst.FuncDecl) *Funcdecl_toQueryFunction {
	fd_ts := &Funcdecl_toQueryFunction{
		Funcdecl: f,
		StandardData: &StandardData{
			Name: f.Name.String(),
		},
	}
	return fd_ts
}

type Funcdecl_toQueryFunction struct {
	Funcdecl *dst.FuncDecl
	*StandardData
}

func (qmp *Funcdecl_toQueryFunction) GetParams() []*dst.Field {
	return qmp.Funcdecl.Type.Params.List
}

func (qmp *Funcdecl_toQueryFunction) GetReturns() []*dst.Field {
	return qmp.Funcdecl.Type.Results.List
}

func (qmp *Funcdecl_toQueryFunction) GetGenerateFunctionName() string {
	return "Query"
}

func (qmp *Funcdecl_toQueryFunction) GetQueryResults() []dst.Expr {
	switch len(qmp.GetReturns()) {
	case 2:

		return []dst.Expr{
			dst.NewIdent("results"),
			dst.NewIdent("err"),
		}
	default:

		return []dst.Expr{
			dst.NewIdent("err"),
		}
	}
}

func (qmp *Funcdecl_toQueryFunction) GetGenerateReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent(qmp.GetAbbv())},
				Type:  dst.NewIdent(fmt.Sprint(qmp.GetName())),
			},
		},
	}
}

func (qmp *Funcdecl_toQueryFunction) GetGenerateFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent("query")},
				Type:  dst.NewIdent("*db.Queries"),
			},
			{
				Names: []*dst.Ident{dst.NewIdent("r")},
				Type:  dst.NewIdent("context.Context"),
			},
		},
	}
}

func (qmp *Funcdecl_toQueryFunction) GetGenerateResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *Funcdecl_toQueryFunction) GetFunctionReturn() []dst.Expr {

	switch len(qmp.GetReturns()) {
	case 2:

		return []dst.Expr{
			dst.NewIdent("results"),
			dst.NewIdent("err"),
		}
	default:

		return []dst.Expr{
			dst.NewIdent("nil"),
			dst.NewIdent("err"),
		}
	}
}

func (qmp *Funcdecl_toQueryFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{ // this is basically the function body
			// other stuff can go here
			qmp.GenerateQuery(), // this is a function call within the body
			&dst.ReturnStmt{ // this is a return statement to the body
				Results: qmp.GetFunctionReturn(),
			},
		},
	}
}

func (qmp *Funcdecl_toQueryFunction) GetQueryArgs() []dst.Expr {
	switch len(qmp.GetParams()) {
	case 2:
		secondarg := qmp.GetParams()[1]
		if strings.Contains(fmt.Sprint(secondarg.Type), "Params") {
			return []dst.Expr{
				dst.NewIdent("r"),
				dst.NewIdent(fmt.Sprint(qmp.GetAbbv(), ".Params")),
			}
		} else {

			propercasename := strings.Title(secondarg.Names[0].Name)
			return []dst.Expr{
				dst.NewIdent("r"),
				dst.NewIdent(fmt.Sprint(qmp.GetAbbv(), ".", propercasename)),
			}
		}
	default:
		return []dst.Expr{
			dst.NewIdent("r"),
		}
	}

}

// /// internal database query
func (qmp *Funcdecl_toQueryFunction) GenerateQuery() *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs: qmp.GetQueryResults(),
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   dst.NewIdent(QUERY_PACKAGE),
					Sel: dst.NewIdent(qmp.GetName()),
				},
				Args: qmp.GetQueryArgs(),
			},
		},
	}
}
