package main

import (
	"go/token"

	"github.com/dave/dst"
)

func combineFields(add []*dst.Field, decls []*dst.Field) {
	for _, decl := range decls {
		add = append(add, decl)
	}
}

func CreateQueryStruct(props FromStruct) []dst.Decl {
	var fields []*dst.Field
	if props.GetStructParams() != nil {
		// fields = append(fields, props.GetStructParams()...)
		// fields = append(fields, props.GetStructParams())
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
