package conversions

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/data"
	dstto "github.com/alterejoe/generate/sqlc-go-helper/cmd/dst-to"
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

// types: struct and function
// outputs: query and display

// toType enum

func FuncToStruct(v *dst.FuncDecl, t string) interfaces.Struct {
	sd := &data.StandardData{
		Name: GetFuncName(v),
	}

	genTo := &dstto.FuncTo{FuncDecl: v}
	st, err := genTo.ToFunctionType()
	if err != nil {
		panic(err)
	}

	switch t {
	case "display":
		return &data.StructData_Display{
			StandardData: *sd,
			Params:       st.Fields,
		}
	case "query":
		return &data.StructData_Query{
			StandardData: *sd,
			Params:       st.Fields,
		}
	default:
		panic("Incorrect string type for structType(t string)")
	}
}

func GetFuncName(v *dst.FuncDecl) string {
	return v.Name.Name
}
