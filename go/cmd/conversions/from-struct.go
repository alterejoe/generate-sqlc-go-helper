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

func ToStruct(v *dst.GenDecl, t string) interfaces.Struct {
	sd := &data.StandardData{
		Name: GetName(v),
	}

	genTo := &dstto.GenTo{GenDecl: v}
	st, err := genTo.ToStructType()
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

func GetName(v *dst.GenDecl) string {
	for _, spec := range v.Specs {
		switch s := spec.(type) {
		case *dst.TypeSpec:

			return s.Name.Name
		}
	}
	return ""
}
