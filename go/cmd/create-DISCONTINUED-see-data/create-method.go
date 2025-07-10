package create

//
// import (
// 	"go/token"
//
// 	"github.com/dave/dst"
// )
//
// func CreateQueryMethod(props QueryFunc) []dst.Decl {
// 	// Define the receiver
// 	receiver := &dst.Field{
// 		Names: []*dst.Ident{dst.NewIdent(props.GetAbbv())},
// 		Type:  dst.NewIdent(props.GetName()),
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
// 	results := props.GetReturns()
// 	// Define the function body
// 	r := []dst.Expr{
// 		dst.NewIdent(props.GetLowerName()),
// 	}
// 	r2 := []dst.Expr{
//
// 		dst.NewIdent(props.GetLowerName()),
// 	}
// 	props.GetQueryAddErr(&r, true)
// 	props.GetQueryAddErr(&r2, false)
//
// 	args := []dst.Expr{
// 		dst.NewIdent("r"),
// 	}
// 	props.GetQueryArgs(&args)
//
// 	body := &dst.BlockStmt{
// 		List: []dst.Stmt{
// 			&dst.AssignStmt{
// 				Lhs: r,
// 				Tok: token.DEFINE,
// 				Rhs: []dst.Expr{
// 					&dst.CallExpr{
// 						Fun: &dst.SelectorExpr{
// 							X:   dst.NewIdent("query"),
// 							Sel: dst.NewIdent(props.GetName()),
// 						},
// 						Args: args,
// 					},
// 				},
// 			},
// 			&dst.ReturnStmt{
// 				Results: r2,
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
