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

type FunctionGenerator struct {
	ModelDisplayFunction
}

func (sd *FunctionGenerator) Generate() *dst.FuncDecl {
	return &dst.FuncDecl{
		Recv: sd.GetReceiver(),
		Name: dst.NewIdent("Query"),
		Type: &dst.FuncType{
			Params:  sd.GetParams(),
			Results: sd.GetResults(),
		},
		Body: sd.GetBody(),
	}
}

type ModelDisplayFunction struct{}

// Using the current GetReceiver function
// The 'Reciever' is the (qmp *ModelDisplayFunction)
func (qmp *ModelDisplayFunction) GetReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent("ABBV")},
				Type:  dst.NewIdent("FUNCTION_NAME"),
			},
		},
	}
}

func (qmp *ModelDisplayFunction) GetParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("r")},
			{Type: dst.NewIdent("LOWER FUNCTION NAME")},
		},
	}
}

// The 'QueryResults' would be a return of a Query function within the generated code
func (qmp *ModelDisplayFunction) GetResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("THIS_IS_WHERE_LOWER_NAME_GOES")},
			{Type: dst.NewIdent("err")},
		},
	}
}

func (qmp *ModelDisplayFunction) GetQueryResults() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("r"),
		dst.NewIdent("THIS_IS_WHERE_PARAMETERS_'COULD'_GO"),
	}
}

func (qmp *ModelDisplayFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{ // this is basically the function body
			qmp.GenerateQuery(), // this is a function call within the body
			qmp.GetBodyReturn(),
		},
	}
}
func (qmp *ModelDisplayFunction) GetArguments() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("r"),
		dst.NewIdent("THIS_IS_WHERE_PARAMETERS_'COULD'_GO"),
	}
}

const QUERY_PACKAGE = "query"

func (qmp *ModelDisplayFunction) GetQueryCall() *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   dst.NewIdent(QUERY_PACKAGE),
		Sel: dst.NewIdent("SQLC_FUNCTION_NAME"),
	}
}

func (qmp *ModelDisplayFunction) GenerateQuery() *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs: qmp.GetQueryResults(),
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun:  qmp.GetQueryCall(),
				Args: qmp.GetArguments(),
			},
		},
	}
}

func (qmp *ModelDisplayFunction) GetBodyReturn() *dst.ReturnStmt {
	return &dst.ReturnStmt{ // this is a return statement to the body
		Results: qmp.GetQueryResults(),
	}
}
