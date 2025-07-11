package conversions

import (
	"fmt"
	"strings"

	data "github.com/alterejoe/generate/sqlc-go-helper/cmd/display"
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

// types: struct and function
// outputs: query and display

// toType enum

func getSqlcStructFields(v *dst.FuncDecl) dst.FieldList {
	args := v.Type.Params.List
	switch v := len(args); v {
	case 1:
		return dst.FieldList{
			List: []*dst.Field{},
		}
	case 2:
		if strings.Contains(fmt.Sprint(args[1].Type), "Params") {
			return dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("Params")},
						Type:  dst.NewIdent(fmt.Sprint("db.", args[1].Type)),
					},
				},
			}
		} else {
			return dst.FieldList{
				List: []*dst.Field{
					{
						Names: args[1].Names,
						Type:  dst.NewIdent(fmt.Sprint(args[1].Type)),
					},
				}}
		}
	default:
		panic("Incorrect number of args")
	}
}

func getSqlcQueryReturns(v *dst.FuncDecl) []*dst.Field {
	for i, f := range v.Type.Results.List {
		fmt.Println(i, f)
	}
	return v.Type.Params.List
	// args := v.Type.Results.List
	// switch v := len(args); v {
	// case 1:
	// 	return []*dst.Field{
	// 		{
	// 			Names: args[1].Names,
	// 			Type:  dst.NewIdent(fmt.Sprint(args[1].Type)),
	// 		},
	// 	}
	// default:
	// 	panic("Incorrect number of args")
}

func FuncdeclToStruct(v *dst.FuncDecl, t string) interfaces.Struct {
	sd := &data.StandardData{
		Name: GetFuncName(v),
	}

	params := getSqlcStructFields(v)
	switch t {
	case "query":
		return &data.StructData_Query{
			StructData: data.StructData{
				Fields:       &params,
				StandardData: *sd,
			},
		}
	default:
		// panic("Incorrect string type for structType(t string)")
		fmt.Println("Incorrect string type for structType(t string)")
		return nil
	}
}

func GetFuncName(v *dst.FuncDecl) string {
	return v.Name.Name
}

func FuncdeclToFunction(v *dst.FuncDecl, t string) interfaces.Function {
	// sd := &data.StandardData{
	// 	Name: GetFuncName(v),
	// }

	getSqlcStructFields(v)
	switch t {
	case "query":
		return &data.Sqlcquery_QueryFunction{}
	default:
		panic("Incorrect string type for structType(t string)")
	}
}
