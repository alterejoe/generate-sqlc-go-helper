package main

import (
	"github.com/dave/dst"
)

func CreateQueryMethod(props *ParseProps) []dst.Decl {
	return nil
}

// func CreateQueryMethod(dcl *dst.FuncDecl) []dst.Decl {
// 	props := ParseProps{
// 		Name:   dcl.Name.Name,
// 		Params: dcl.Type.Params.List,
// 	}
// 	// Define the receiver
// 	receiver := &dst.Field{
// 		Names: []*dst.Ident{dst.NewIdent(props.GetAbbv())},
// 		Type:  &dst.StarExpr{X: dst.NewIdent(props.GetName())},
// 	}
//
// 	// Define the parameters
// 	params := &dst.FieldList{
// 		List: []*dst.Field{
// 			{
// 				Names: []*dst.Ident{dst.NewIdent("query")},
// 				Type:  &dst.StarExpr{X: dst.NewIdent("db.Queries")},
// 			},
// 			{
// 				Names: []*dst.Ident{dst.NewIdent("r")},
// 				Type:  dst.NewIdent("context.Context"),
// 			},
// 		},
// 	}
//
// 	// Define the return types
// 	results := &dst.FieldList{
// 		List: []*dst.Field{
// 			{
// 				Type: dst.NewIdent("any"),
// 			},
// 			{
// 				Type: dst.NewIdent("error"),
// 			},
// 		},
// 	}
//
// 	// Define the function body
// 	body := &dst.BlockStmt{
// 		List: []dst.Stmt{
// 			&dst.AssignStmt{
// 				Lhs: []dst.Expr{
// 					dst.NewIdent(props.GetLowerName()),
// 					dst.NewIdent("err"),
// 				},
// 				Tok: token.DEFINE,
// 				Rhs: []dst.Expr{
// 					&dst.CallExpr{
// 						Fun: &dst.SelectorExpr{
// 							X:   dst.NewIdent("query"),
// 							Sel: dst.NewIdent(props.GetName()),
// 						},
// 						Args: []dst.Expr{
// 							dst.NewIdent("r"),
// 						},
// 					},
// 				},
// 			},
// 			&dst.ReturnStmt{
// 				Results: []dst.Expr{
// 					dst.NewIdent(props.GetLowerName()),
// 					dst.NewIdent("err"),
// 				},
// 			},
// 		},
// 	}
//
// 	// Define the function declaration
// 	funcDecl := &dst.FuncDecl{
// 		Recv: &dst.FieldList{
// 			List: []*dst.Field{receiver},
// 		},
// 		Name: dst.NewIdent("Query"),
// 		Type: &dst.FuncType{
// 			Params:  params,
// 			Results: results,
// 		},
// 		Body: body,
// 	}
//
// 	return []dst.Decl{funcDecl}
// }
