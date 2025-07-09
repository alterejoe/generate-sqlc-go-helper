package create

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

func combineFields(add []*dst.Field, decls []*dst.Field) {
	for _, decl := range decls {
		add = append(add, decl)
	}
}

func QueryStruct(props interfaces.Struct) []dst.Decl {
	var fields []*dst.Field
	if props.GetStructParams() != nil {
		combineFields(fields, props.GetStructParams())
	}

	newstruct := &dst.TypeSpec{
		Name: dst.NewIdent(props.GetName()),
		Type: &dst.StructType{
			Fields: &dst.FieldList{
				List: fields,
			},
		},
	}
	returnstruct := &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			newstruct,
		},
	}
	return []dst.Decl{returnstruct}
}
