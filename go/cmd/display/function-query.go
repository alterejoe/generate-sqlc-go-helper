package display

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/parse"
	"github.com/dave/dst"
)

type TemplateSqlcquery_QFunction struct {
	Func   *dst.FuncDecl
	Params []*dst.Field
	*parse.StandardData
}

func (qmp *TemplateSqlcquery_QFunction) GetParams() dst.Expr {
	secondArg := qmp.Params[1]
	if strings.Contains(fmt.Sprint(secondArg.Type), "Params") {
		return dst.NewIdent(fmt.Sprint(qmp.GetAbbv(), ".", qmp.GetName()))
	} else {
		return dst.NewIdent(fmt.Sprint(qmp.GetAbbv(), ".", fmt.Sprint(secondArg.Names[0].Name)))
	}

}

// Using the current GetReceiver function
// The 'Reciever' is the (qmp *ModelQueryFunction)
func (qmp *TemplateSqlcquery_QFunction) GetReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent(qmp.GetAbbv())},
				Type:  dst.NewIdent(fmt.Sprint(qmp.GetName())),
			},
		},
	}
}

func (qmp *TemplateSqlcquery_QFunction) GetFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Names: []*dst.Ident{dst.NewIdent("query")},
				Type: dst.NewIdent("*db.Queries")},
			{Names: []*dst.Ident{dst.NewIdent("r")},
				Type: dst.NewIdent("context.Context")},
		},
	}
}

// The 'QueryResults' would be a return of a Query function within the generated code
func (qmp *TemplateSqlcquery_QFunction) GetResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *TemplateSqlcquery_QFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{ // this is basically the function body
			qmp.GenerateQuery(), // this is a function call within the body
			qmp.GetBodyReturn(),
		},
	}
}

func (qmp *TemplateSqlcquery_QFunction) GetQueryArguments() []dst.Expr {
	// args := qmp.Params
	switch len(qmp.Params) {
	case 2:
		return []dst.Expr{
			dst.NewIdent("r"),
			qmp.GetParams(),
		}
	default:
		return []dst.Expr{
			dst.NewIdent("r"),
		}
	}
}

func (qmp *TemplateSqlcquery_QFunction) GetQueryResults() []dst.Expr {
	args := qmp.Params
	switch len(args) {
	case 1:
		return []dst.Expr{
			dst.NewIdent(qmp.GetLowerName()),
			dst.NewIdent("err"),
		}
	default:
		return []dst.Expr{
			dst.NewIdent("err"),
		}
	}
}
func (qmp *TemplateSqlcquery_QFunction) GetQueryCall() *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(QUERY_PACKAGE),
		Sel: dst.NewIdent(qmp.GetName()),
	}
}

func (qmp *TemplateSqlcquery_QFunction) GenerateQuery() *dst.AssignStmt {
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

func (qmp *TemplateSqlcquery_QFunction) GetBodyReturn() *dst.ReturnStmt {
	return &dst.ReturnStmt{ // this is a return statement to the body
		Results: qmp.GetQueryResults(),
	}
}
