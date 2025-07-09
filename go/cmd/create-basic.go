package main

import (
	"go/token"

	"github.com/dave/dst"
)

//
// func CreateStruct(dcl *dst.FuncDecl) []dst.Decl {
// 	structname := dcl.Name.Name
// 	newstruct := &dst.TypeSpec{
// 		Name: dst.NewIdent(structname),
// 		Type: &dst.StructType{
// 			Fields: &dst.FieldList{},
// 		},
// 	}
//
// 	returnstruct := &dst.GenDecl{
// 		Tok: token.TYPE,
// 		Specs: []dst.Spec{
// 			newstruct,
// 		},
// 	}
// 	return []dst.Decl{returnstruct}
// }

func CreateStruct(props *ParseProps) []dst.Decl {
	var fields []*dst.Field
	if props.ExtraStructParam() != nil {
		fields = append(fields, props.ExtraStructParam())
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
