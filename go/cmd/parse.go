package main

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/conversions"
	"github.com/dave/dst"
)

func parse_models(n dst.Node) []dst.Decl {
	switch v := n.(type) {
	case *dst.GenDecl:
		if v.Tok != token.TYPE {
			return nil
		}
		conv_gendecl := &conversions.FromGenDecl{
			Input: v,
		}
		// fmt.Println(conv_gendecl.GetName(), "Name")
		// fmt.Println(conv_gendecl.GetLowerName(), "Lower")
		// fmt.Println(conv_gendecl.GetReturns(), "Returnsn")
		s := CreateQueryStruct(conv_gendecl)
		return s
	default:
		return nil
	}
}

func parse_queries(n dst.Node) []dst.Decl {
	var decls []dst.Decl
	switch v := n.(type) {
	case *dst.FuncDecl:
		conv_func := &conversions.FromFuncDecl{
			Input: v,
		}

		s := CreateQueryStruct(conv_func)
		m := CreateQueryMethod(conv_func)

		CombineDecls(&decls, &s)
		CombineDecls(&decls, &m)
		return decls
	default:
		return nil
	}
}

func CombineDecls(add *[]dst.Decl, decls *[]dst.Decl) {
	for _, decl := range *decls {
		*add = append(*add, decl)
	}
}

// func retrieveFuncProps(n *dst.FuncDecl) *ParseFuncProps {
// 	return &ParseFuncProps{
// 		ParseStructProps: ParseStructProps{
// 			Name:   n.Name.Name,
// 			Params: n.Type.Params,
// 		},
// 		Results: n.Type.Results,
// 	}
// }
//
// func retrieveStructProps(n *dst.GenDecl) *ParseStructProps {
// 	typeSpec, ok := n.Specs[0].(*dst.TypeSpec)
// 	if !ok {
// 		return nil
// 	}
// 	structSpec := typeSpec.Type.(*dst.StructType)
// 	return &ParseStructProps{
// 		Name:   typeSpec.Name.Name,
// 		Params: structSpec.Fields,
// 	}
// }
