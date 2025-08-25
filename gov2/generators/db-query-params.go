package generators

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/inspectors"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/tools"
	"github.com/dave/dst"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func DbQueryParams_PointerFunctions(p inspectors.FuncParams, deps deps.Deps) []dst.Decl {
	var decls []dst.Decl

	var structparams *dst.FieldList
	var queryparams []dst.Expr

	switch len(p.Args.List) {
	case 1:
		structparams = &dst.FieldList{
			List: []*dst.Field{},
		}
		queryparams = []dst.Expr{
			dst.NewIdent("r"),
		}
	case 2:
		passedparams := p.Args.List[1]
		c := cases.Title(language.English)
		name := c.String(passedparams.Names[0].Name)
		dbparamtype := &dst.SelectorExpr{
			X:   dst.NewIdent(deps.Environment.DbModuleNameOut),
			Sel: dst.NewIdent(fmt.Sprintf("%sParams", p.FuncName)),
		}

		if strings.EqualFold(name, "Arg") {
			structparams = &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("Params")},
						Type:  dbparamtype,
					},
				},
			}
			queryparams = []dst.Expr{
				dst.NewIdent("r"),
				dst.NewIdent(fmt.Sprintf("%s.Params", tools.GetAbbv(p.FuncName))),
			}
		} else {
			structparams = &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent(name)},
						Type:  passedparams.Type,
					},
				},
			}
			queryparams = []dst.Expr{
				dst.NewIdent("r"),
				dst.NewIdent(fmt.Sprintf("%s.%s", tools.GetAbbv(p.FuncName), name)),
			}
		}
	default:
		deps.Logger.Debug("No fields in struct", "struct", p.FuncName)
	}

	var queryreturns []dst.Expr
	var funcreturns []dst.Expr

	switch len(p.Results.List) {
	case 1:
		queryreturns = []dst.Expr{
			dst.NewIdent("err"),
		}

		funcreturns = []dst.Expr{
			dst.NewIdent("nil"),
			dst.NewIdent("err"),
		}

	case 2:
		queryreturns = []dst.Expr{
			dst.NewIdent("results"),
			dst.NewIdent("err"),
		}

		funcreturns = []dst.Expr{
			dst.NewIdent("results"),
			dst.NewIdent("err"),
		}

	default:
		deps.Logger.Debug("No fields in struct", "struct", p.FuncName)
	}

	gendecl := &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name: dst.NewIdent(p.FuncName),
				Type: &dst.StructType{
					Fields: structparams,
				},
			},
		},
	}

	// query pointer function
	funcdecl := &dst.FuncDecl{
		Recv: &dst.FieldList{
			List: []*dst.Field{
				{
					Names: []*dst.Ident{dst.NewIdent(tools.GetAbbv(p.FuncName))},
					Type:  &dst.Ident{Name: fmt.Sprintf("*%s", p.FuncName)},
				},
			},
		},
		Name: dst.NewIdent("Query"),
		Type: &dst.FuncType{
			Params: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("query")},
						Type:  &dst.Ident{Name: "*db.Queries"}, // example type
					},
					{
						Names: []*dst.Ident{dst.NewIdent("r")},
						Type:  &dst.Ident{Name: "context.Context"}, // example type
					},
				},
			},
			Results: &dst.FieldList{
				List: []*dst.Field{
					{
						Type: &dst.Ident{Name: "any"}, // required
					},
					{
						Type: &dst.Ident{Name: "error"}, // required
					},
				},
			},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				// results, err := query.SelectAllTransferCustomers(r)
				&dst.AssignStmt{
					Lhs: queryreturns,
					Tok: token.DEFINE,
					Rhs: []dst.Expr{
						&dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   dst.NewIdent("query"),
								Sel: dst.NewIdent(p.FuncName),
							},
							// Args: []dst.Expr{dst.NewIdent("r")},
							Args: queryparams,
						},
					},
				},
				// return results, err
				&dst.ReturnStmt{
					Results: funcreturns,
				},
			},
		},
	}

	decls = append(decls, gendecl)
	decls = append(decls, funcdecl)
	// }
	return decls
}
