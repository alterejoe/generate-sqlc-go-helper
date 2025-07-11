package data

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it easier to sort and delegate data to its respective parser
type GenToDisplayFunctionProps struct {
	Name    string
	Field   string
	Gendecl *dst.GenDecl
}

func GenToDisplayFunction(props *GenToDisplayFunctionProps) *Gendecl_toDisplayFunction {
	fd_ts := &Gendecl_toDisplayFunction{
		Gendecl: props.Gendecl,
		Field:   props.Field,
		StandardData: &StandardData{
			Name: props.Name,
		},
	}
	return fd_ts
}

type Gendecl_toDisplayFunction struct {
	Gendecl *dst.GenDecl
	Field   string
	*StandardData
}

func (qmp *Gendecl_toDisplayFunction) GetParams() []*dst.Field {
	return []*dst.Field{}
}

func (qmp *Gendecl_toDisplayFunction) GetReturns() []*dst.Field {
	return []*dst.Field{}
}

func (qmp *Gendecl_toDisplayFunction) GetQueryResults() []dst.Expr {
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

func (qmp *Gendecl_toDisplayFunction) GetReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent(qmp.GetAbbv())},
				Type:  dst.NewIdent(fmt.Sprint("*" + qmp.GetName())),
			},
		},
	}
}

func (qmp *Gendecl_toDisplayFunction) GetFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{},
	}
}

func (qmp *Gendecl_toDisplayFunction) GetResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *Gendecl_toDisplayFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: []dst.Stmt{ // this is basically the function body
			// other stuff can go here
			&dst.ReturnStmt{ // this is a return statement to the body
				Results: qmp.GetQueryResults(),
			},
		},
	}
}

func (qmp *Gendecl_toDisplayFunction) GetQueryArgs() []dst.Expr {
	switch len(qmp.GetParams()) {
	case 2:
		secondarg := qmp.GetParams()[1]
		if strings.Contains(fmt.Sprint(secondarg.Type), "Params") {
			return []dst.Expr{
				dst.NewIdent("r"),
				dst.NewIdent(fmt.Sprint(qmp.GetAbbv(), ".Params")),
			}
		} else {
			return []dst.Expr{
				dst.NewIdent("r"),
				dst.NewIdent(fmt.Sprint(qmp.GetAbbv(), ".", secondarg.Names[0].Name)),
			}
		}
	default:
		return []dst.Expr{
			dst.NewIdent("r"),
		}
	}

}

func (qmp *Gendecl_toDisplayFunction) GetFunctionName() string {
	return "Get" + qmp.Field
}
