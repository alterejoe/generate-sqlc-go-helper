package data

import (
	"fmt"
	"go/token"
	"log/slog"
	"strings"

	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it easier to sort and delegate data to its respective parser
type GenToDisplayFunctionProps struct {
	Name       string
	Field      *dst.Field
	Gendecl    *dst.GenDecl
	TypeSpec   *dst.TypeSpec
	StructSpec *dst.StructType
	Logger     *slog.Logger
}

func GenToDisplayFunction(props *GenToDisplayFunctionProps) *Gendecl_toDisplayFunction {
	fd_ts := &Gendecl_toDisplayFunction{
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

type Gendecl_toDisplayFunction struct {
	Gendecl  *dst.GenDecl
	Field    *dst.Field
	TypeSpec *dst.TypeSpec
	*StandardData
}

// // dst.NewIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + ".Time"),
func (qmp *Gendecl_toDisplayFunction) ParamIdent(param string) *dst.Ident {
	return dst.NewIdent(param)
}

func (qmp *Gendecl_toDisplayFunction) PreidentInt(param string) string {
	return fmt.Sprint("int(", param, ")")
}

func (qmp *Gendecl_toDisplayFunction) PreidentPgfield(param string) string {
	return fmt.Sprint(qmp.GetAbbv(), ".", param)
}

func (qmp *Gendecl_toDisplayFunction) PreidentPgsubparam(param string) string {
	return fmt.Sprint(qmp.GetAbbv(), ".", qmp.GetFieldName(), ".", param)
}

func (qmp *Gendecl_toDisplayFunction) Ident(param string) *dst.Ident {
	return dst.NewIdent(fmt.Sprint(param))
}

//
// func (qmp *Gendecl_toDisplayFunction) PgSubparam(param string) string {
// 	return fmt.Sprint(qmp.GetAbbv(), ".", qmp.GetFieldName(), ".", param)
// }
//
// func (qmp *Gendecl_toDisplayFunction) PgSubparamIdentInt(param string) *dst.Ident {
// 	return dst.NewIdent("int(" + qmp.PgSubparam(param) + ")")
// }
//
// func (qmp *Gendecl_toDisplayFunction) PgSubparamIdent(param string) *dst.Ident {
// 	return dst.NewIdent(qmp.PgSubparam(param))
// }
//
// func (qmp *Gendecl_toDisplayFunction) PgParam() string {
// 	return fmt.Sprint(qmp.GetAbbv(), ".", qmp.GetFieldName())
// }
//
// func (qmp *Gendecl_toDisplayFunction) PgParamIdent() *dst.Ident {
// 	return dst.NewIdent(qmp.PgParam())
// }
//
// func (qmp *Gendecl_toDisplayFunction) PgParamIdent() *dst.Ident {
// 	return dst.NewIdent(qmp.PgParam())
// }

func (qmp *Gendecl_toDisplayFunction) GetParams() []*dst.Field {
	return []*dst.Field{}
}

func (qmp *Gendecl_toDisplayFunction) GetReturns() []*dst.Field {

	return []*dst.Field{
		qmp.Field,
	}
}

func (qmp *Gendecl_toDisplayFunction) GetGenerateReceiver() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent(qmp.GetAbbv())},
				Type:  dst.NewIdent(fmt.Sprint("*" + qmp.GetName())),
			},
		},
	}
}

func (qmp *Gendecl_toDisplayFunction) GetGenerateFunctionParams() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{},
	}
}

func (qmp *Gendecl_toDisplayFunction) GetGenerateResults() *dst.FieldList {
	switch fmt.Sprint(qmp.Field.Type) {
	case "string":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("string")},
			},
		}
	case "int64":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("int")},
			},
		}
	case "int32":

		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("int")},
			},
		}
	case "bool":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("bool")},
			},
		}
	case "&{pgtype Bool {{None [] [] None} []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("bool")},
			},
		}
	case "&{pgtype Int4 {{None [] [] None} []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("int")},
			},
		}
	case "&{pgtype Float8 {{None [] [] None} []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("float64")},
			},
		}
	case "&{pgtype Text {{None [] [] None} []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("string")},
			},
		}
	case "&{pgtype Timestamp {{None [] [] None} []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("time.Time")},
			},
		}
	case "&{pgtype Timestamptz {{None [] [] None} []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("time.Time")},
			},
		}
	case "&{<nil> byte {{None [] [] None} [] []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("[]byte")},
			},
		}
	case "&{<nil> string {{None [] [] None} [] []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("[]string")},
			},
		}
	case "&{<nil> float64 {{None [] [] None} [] []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("[]float64")},
			},
		}
	case "&{<nil> int32 {{None [] [] None} [] []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("[]int")},
			},
		}
	case "&{<nil> bool {{None [] [] None} [] []}}":
		return &dst.FieldList{
			List: []*dst.Field{
				{Type: dst.NewIdent("[]bool")},
			},
		}
	default:
		fmt.Println("Wrong type", qmp.Field.Type)
		return nil
	}
}

func (qmp *Gendecl_toDisplayFunction) GetTypeConversionReturn() *dst.ReturnStmt {
	switch fmt.Sprint(qmp.Field.Type) {
	case "string", "int64", "int32", "bool":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.Field.Type,
			},
		}
	case "&{pgtype Bool {{None [] [] None} []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("false"),
			},
		}
	case "&{pgtype Int4 {{None [] [] None} []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("-1"),
			},
		}
	case "&{pgtype Float8 {{None [] [] None} []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("-1"),
			},
		}
	case "&{pgtype Text {{None [] [] None} []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("\"\""),
			},
		}
	case "&{pgtype Timestamp {{None [] [] None} []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("time.Time{}"),
			},
		}
	case "&{pgtype Timestamptz {{None [] [] None} []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("time.Time{}"),
			},
		}
	case "&{<nil> byte {{None [] [] None} [] []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("[]byte{}"),
			},
		}
	case "&{<nil> string {{None [] [] None} [] []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("[]string{}"),
			},
		}
	case "&{<nil> float64 {{None [] [] None} [] []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("[]float64{}"),
			},
		}
	case "&{<nil> int32 {{None [] [] None} [] []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("[]int{}"),
			},
		}
	case "&{<nil> bool {{None [] [] None} [] []}}":
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("[]bool{}"),
			},
		}
	// case "&{<nil> string {{None [] [] None} [] []}}":
	// 	return &dst.ReturnStmt{
	// 		Results: []dst.Expr{
	// 			qmp.ParamIdent("[]string{}"),
	// 		},
	// 	}
	//
	default:
		fmt.Println("GetTypeConversionReturn: Type not found: ", qmp.Field.Type)
		return &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.ParamIdent("\"\""),
			},
		}
		// default:
		// 	return &dst.ReturnStmt{
		// 		Results: []dst.Expr{
		// 			qmp.Field.Type,
		// 		},
		// 	}
	}
}

func (qmp *Gendecl_toDisplayFunction) GetTypeConversion() []dst.Stmt {
	var conversion dst.Stmt
	var results []dst.Stmt

	nilResults := qmp.GetTypeConversionReturn()
	switch v := fmt.Sprint(qmp.Field.Type); v {
	case "bool", "string":
		conversion = &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.Ident(qmp.PreidentPgfield(qmp.GetFieldName())),
			},
		}
		results = []dst.Stmt{
			conversion,
		}
	case "int64", "int32":
		conversion = &dst.ReturnStmt{
			Results: []dst.Expr{
				qmp.Ident(qmp.PreidentInt(qmp.PreidentPgfield(qmp.GetFieldName()))),
			},
		}

		results = []dst.Stmt{
			conversion,
		}
	case "&{pgtype Text {{None [] [] None} []}}":
		conversion = &dst.IfStmt{
			// Cond: qmp.PgSubparamIdent("Valid"),
			Cond: qmp.Ident(qmp.PreidentPgsubparam("Valid")),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgSubparamIdent("String"),
							qmp.Ident(qmp.PreidentPgsubparam("String")),
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
			// Cond: qmp.PgSubparamIdent("Valid"),
			Cond: qmp.Ident(qmp.PreidentPgsubparam("Valid")),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgSubparamIdent("Float64")},
							qmp.Ident(qmp.PreidentPgsubparam("Float64")),
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
			// Cond: qmp.PgSubparamIdent("Valid"),
			Cond: qmp.Ident(qmp.PreidentPgsubparam("Valid")),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgSubparamIdentInt("Int32"),
							qmp.Ident(qmp.PreidentInt(qmp.PreidentPgsubparam("Int32"))),
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
			// Cond: qmp.PgSubparamIdent("Valid"),
			Cond: qmp.Ident(qmp.PreidentPgsubparam("Valid")),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgSubparamIdent("Bool"),
							qmp.Ident(qmp.PreidentPgsubparam("Bool")),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{<nil> byte {{None [] [] None} [] []}}":
		conversion = &dst.IfStmt{
			// Cond: qmp.ParamIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + " != nil"),
			Cond: qmp.Ident(qmp.PreidentPgfield(qmp.GetFieldName()) + " != nil"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgParamIdent(),
							qmp.Ident(qmp.PreidentPgfield(qmp.GetFieldName())),
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
			// Cond: qmp.PgSubparamIdent("Valid"),
			Cond: qmp.Ident(qmp.PreidentPgsubparam("Valid")),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgSubparamIdent("Time"),
							qmp.Ident(qmp.PreidentPgsubparam("Time")),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}
	case "&{<nil> int32 {{None [] [] None} [] []}}", "&{<nil> int64 {{None [] [] None} [] []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.ParamIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + " != nil"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					//var lowerFieldName returnType
					&dst.DeclStmt{
						Decl: &dst.GenDecl{
							Tok: token.VAR,
							Specs: []dst.Spec{
								&dst.ValueSpec{
									Names: []*dst.Ident{
										qmp.Ident(strings.ToLower(qmp.GetFieldName())),
									},
									Type: qmp.GetGenerateResults().List[0].Type,
								},
							},
						},
					},
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgParamIdent(),
							qmp.Ident(strings.ToLower(qmp.GetFieldName())),
						},
					},
				},
			},
		}

		results = []dst.Stmt{
			conversion,
			nilResults,
		}

	case "&{<nil> string {{None [] [] None} [] []}}", "&{<nil> float64 {{None [] [] None} [] []}}", "&{<nil> bool {{None [] [] None} [] []}}":
		conversion = &dst.IfStmt{
			Cond: qmp.ParamIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + " != nil"),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							// qmp.PgParamIdent(),
							qmp.Ident(qmp.PreidentPgfield(qmp.GetFieldName())),
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
		fmt.Println("Wrong -- ", qmp.Field.Type)
		panic("Type hasn't been implemented yet")
	}

	return results
}

func (qmp *Gendecl_toDisplayFunction) GetBody() *dst.BlockStmt {
	return &dst.BlockStmt{
		List: qmp.GetTypeConversion(),
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

func (qmp *Gendecl_toDisplayFunction) GetFieldName() string {
	return qmp.Field.Names[0].Name
}

func (qmp *Gendecl_toDisplayFunction) GetGenerateFunctionName() string {
	return "Get" + qmp.GetFieldName()
}
