package main

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/conversions"
	data "github.com/alterejoe/generate/sqlc-go-helper/cmd/display"
	"github.com/dave/dst"
)

const QUERY = "query"
const DISPLAY = "display"

func parse_models(n dst.Node) []dst.Decl {
	/*
		This parses models.go file to get the table data.
		It will then generate display pointer methods that
		return as strings for html templates.

		Types will be limited so we can generate string conversions.
	*/
	switch v := n.(type) {
	case *dst.GenDecl:
		if v.Tok != token.TYPE {
			return nil
		}

		// structdata := conversions.GendeclToStruct(v, DISPLAY)

		// //model to display function
		funcgen := &data.FunctionGenerator{
			Function: &data.Modelgo_DisplayFunction{},
		}
		return []dst.Decl{
			funcgen.Generate(),
		}
	default:
		return nil
	}
}

func parse_queries(n dst.Node) []dst.Decl {
	// var decls []dst.Decl
	switch v := n.(type) {
	case *dst.FuncDecl:
		structdata := conversions.FuncdeclToStruct(v, QUERY)
		funcdata := conversions.FuncdeclToFunction(v, QUERY)

		structgen := &data.StructGenerator{
			Sqlcquery_QueryStruct: data.Sqlcquery_QueryStruct{
				Name:   structdata.GetName(),
				Fields: structdata.GetStructParams(),
			},
		}

		funcgen := &data.FunctionGenerator{

			Function: funcdata,
		}
		return []dst.Decl{
			structgen.Generate(),
			funcgen.Generate(),
		}
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
