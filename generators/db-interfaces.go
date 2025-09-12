package generators

import (
	"fmt"
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/inspectors"
	"github.com/dave/dst"
)

// DbInterfaceTemplate generates an interface like:
//
//	type DbInterface interface {
//	    GetUser(id string) (User, error)
//	    SaveUser(u User) error
//	}
func DbInterfaces(p inspectors.StructParams, deps deps.Deps) []dst.Decl {
	var decls []dst.Decl

	// allocate FieldList properly
	fl := &dst.FieldList{
		List: []*dst.Field{},
	}

	for key, field := range p.FieldList.List {
		method := &dst.Field{
			Names: []*dst.Ident{dst.NewIdent(fmt.Sprintf("Get%s", key))},
			Type: &dst.FuncType{
				Params: &dst.FieldList{List: []*dst.Field{}}, // no params
				Results: &dst.FieldList{
					List: []*dst.Field{
						{Type: field.Type},
					},
				},
			},
		}
		fl.List = append(fl.List, method)
	}

	gen := &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name: dst.NewIdent("DbInterface"),
				Type: &dst.InterfaceType{
					Methods: fl,
				},
			},
		},
	}

	decls = append(decls, gen)
	return decls
}
