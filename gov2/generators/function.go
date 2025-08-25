package generators

import (
	"fmt"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/inspectors"
	"github.com/dave/dst"
)

// Recv *FieldList // receiver (methods); or nil (functions)
// Name *Ident     // function/method name
// Type *FuncType  // function signature: type and value parameters, results, and position of "func" keyword
// Body *BlockStmt // function body; or nil for external (non-Go) function
// Decs FuncDeclDecorations
func Function(p inspectors.StructParams, deps deps.Deps) []dst.Decl {
	var funcs []dst.Decl
	for key := range p.FieldList.List {
		funcdecl := &dst.FuncDecl{
			Recv: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent(deps.Environment.DbModuleNameOut)},
						Type:  &dst.Ident{Name: "DBAdapter"}, // <-- you must define what db is
					},
				},
			},
			Name: dst.NewIdent(fmt.Sprintf("Get%s", key)),
			Type: &dst.FuncType{
				Params: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent(deps.Environment.DbModuleNameOut)},
							Type:  &dst.Ident{Name: "string"}, // example type
						},
					},
				},
				Results: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("err")},
							Type:  &dst.Ident{Name: "error"}, // required
						},
					},
				},
			},
		}
		funcs = append(funcs, funcdecl)
	}
	return funcs
}
