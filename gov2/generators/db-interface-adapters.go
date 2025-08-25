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
func DbInterfaceAdapters(p inspectors.Params, deps deps.Deps) []dst.Decl {
	var i []dst.Decl
	for key, field := range p.Fields {
		f := dst.Field{
			Names: []*dst.Ident{dst.NewIdent(fmt.Sprintf("Get%s", key))},
			Type: &dst.FuncType{
				Params: &dst.FieldList{
					List: []*dst.Field{},
				},
				Results: &dst.FieldList{
					List: []*dst.Field{
						field,
					},
				},
			},
		}

		i = append(i, &f)
	}

	gen := &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name: dst.NewIdent("DbInterface"),
				Type: &dst.InterfaceType{
					Methods: &dst.FieldList{
						List: i,
					},
				},
			},
		},
	}
	i = append(i, gen)
	return i
}
