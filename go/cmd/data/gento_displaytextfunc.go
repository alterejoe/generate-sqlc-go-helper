package data

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it easier to sort and delegate data to its respective parser

func GenToDisplayTextFunction(props *GenToDisplayFunctionProps) *Gendecl_toDisplayTextFunction {
	fd_ts := &Gendecl_toDisplayTextFunction{
		Gendecl:  props.Gendecl,
		Field:    props.Field,
		TypeSpec: props.TypeSpec,
		StandardData: &StandardData{
			Name:   props.Name,
			Logger: props.Logger,
		},
	}
	return fd_ts
}

type Gendecl_toDisplayTextFunction struct {
	Gendecl  *dst.GenDecl
	Field    *dst.Field
	TypeSpec *dst.TypeSpec
	*StandardData
}

// dst.NewIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + ".Time"),
func (qmp *Gendecl_toDisplayTextFunction) StandardIdent(param string) *dst.Ident {
	return dst.NewIdent(param)
}
func (qmp *Gendecl_toDisplayTextFunction) StandardPrintIdent(param string) *dst.Ident {
	return dst.NewIdent("fmt.Sprint(" + param + ")")
}

func (qmp *Gendecl_toDisplayTextFunction) StandardPgIdent(param string) *dst.Ident {
	return dst.NewIdent(qmp.StandardPgParam(param))
}

func (qmp *Gendecl_toDisplayTextFunction) StandardPgPrintIdent(param string) *dst.Ident {
	return dst.NewIdent(qmp.StandardPgPrintParam(param))
}

func (qmp *Gendecl_toDisplayTextFunction) StandardPgParam(param string) string {
	return fmt.Sprint(qmp.GetAbbv(), ".", qmp.GetFieldName(), ".", param)
}

func (qmp *Gendecl_toDisplayTextFunction) StandardPgPrintParam(param string) string {
	return fmt.Sprint("fmt.Sprint(", qmp.GetAbbv(), ".", qmp.GetFieldName(), ".", param, ")")
}

func (qmp *Gendecl_toDisplayTextFunction) GetParams() []*dst.Field {
	return []*dst.Field{}
}

func (qmp *Gendecl_toDisplayTextFunction) GetReturns() []*dst.Field {
	return []*dst.Field{}
}

func (qmp *Gendecl_toDisplayTextFunction) GetConversionResults() []dst.Expr {
	return []dst.Expr{
		dst.NewIdent("nil"),
	}
}

func (qmp *Gendecl_toDisplayTextFunction) GetGenerateReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent(qmp.GetAbbv())},
				Type:  dst.NewIdent(fmt.Sprint("*" + qmp.GetName())),
			},
		},
	}
}

func (qmp *Gendecl_toDisplayTextFunction) GetGenerateFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{},
	}
}

func (qmp *Gendecl_toDisplayTextFunction) GetGenerateResults() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("string")},
		},
	}
}

func (qmp *Gendecl_toDisplayTextFunction) GetTypeConversion() []dst.Stmt {
	var conversion dst.Stmt
	var results []dst.Stmt
	nilResults := &dst.ReturnStmt{
		Results: []dst.Expr{
			dst.NewIdent("\"\""),
		},
	}
	switch v := fmt.Sprint(qmp.Field.Type); v {
	case "bool", "int64", "int32", "string":
		conversion = &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.StandardPrintIdent(qmp.GetAbbv() + "." + qmp.Field.Names[0].Name),
			},
		}
		results = []dst.Stmt{
			conversion,
		}
	case "&{pgtype Text {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardPgIdent("Valid"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							dst.NewIdent(qmp.StandardPgParam("String")),
						},
					},
				},
			},
		}
		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{pgtype Float8 {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardPgIdent("Valid"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							qmp.StandardPgPrintIdent("Float64"),
						},
					},
				},
			},
		}
		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{pgtype Int4 {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardPgIdent("Valid"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							qmp.StandardPgPrintIdent("Int32"),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{pgtype Bool {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardPgIdent("Valid"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							qmp.StandardPgPrintIdent("Bool"),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{pgtype Timestamp {{None [] [] None} []}}", "&{pgtype Timestamptz {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardPgIdent("Valid"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							qmp.StandardPgPrintIdent("Time"),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}

	///  come back here
	case "&{<nil> bool {{None [] [] None} [] []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + " != nil"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							dst.NewIdent("string(" + qmp.GetAbbv() + "." + qmp.GetFieldName() + ")"),
						},
					},
				},
			},
		}
	case "&{<nil> byte {{None [] [] None} [] []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + " != nil"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							dst.NewIdent("string(" + qmp.GetAbbv() + "." + qmp.GetFieldName() + ")"),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}

	case "&{<nil> string {{None [] [] None} [] []}}", "&{<nil> float64 {{None [] [] None} [] []}}", "&{<nil> int32 {{None [] [] None} [] []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + " != nil"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							dst.NewIdent("fmt.Sprint(" + qmp.GetAbbv() + "." + qmp.GetFieldName() + ")"),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{pgtype UUID {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.StandardPgIdent("Valid"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							qmp.StandardPgPrintIdent("Bytes"),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	default:
		qmp.GetLogger().Error("Unknown Type", "Type", fmt.Sprint(qmp.Field.Type))
		results = []dst.Stmt{
			nilResults,
		}
	}

	return results
}

func (qmp *Gendecl_toDisplayTextFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: qmp.GetTypeConversion(),
	}
}

func (qmp *Gendecl_toDisplayTextFunction) GetQueryArgs() []dst.Expr {
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

func (qmp *Gendecl_toDisplayTextFunction) GetFieldName() string {
	return qmp.Field.Names[0].Name
}

func (qmp *Gendecl_toDisplayTextFunction) GetGenerateFunctionName() string {
	return "Get" + qmp.GetFieldName() + "Text"
}
