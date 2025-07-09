package main

import (
	"go/token"

	"github.com/dave/dst"
)

func CreateStruct(props *ParseProps) []dst.Decl {
	var fields []*dst.Field
	if props.StructParams() != nil {
		fields = append(fields, props.StructParams())
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
