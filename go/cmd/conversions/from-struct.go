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

func GenToStruct(v *dst.GenDecl, t string) interfaces.Struct {
	sd := &data.StandardData{
		Name: GetGenName(v),
	}

	genTo := &dstto.GenTo{GenDecl: v}
	st, err := genTo.ToStructType()
	if err != nil {
		panic(err)
	}

	switch t {
	case "display":
		return &data.StructData_Display{
			StructData: data.StructData{
				Params:       st.Fields,
				StandardData: *sd,
			},
		}
	case "query":
		return &data.StructData_Query{
			StructData: data.StructData{
				Params:       st.Fields,
				StandardData: *sd,
			},
		}
	default:
		panic("Incorrect string type for structType(t string)")
	}
}

func GetGenName(v *dst.GenDecl) string {
	for _, spec := range v.Specs {
		switch s := spec.(type) {
		case *dst.TypeSpec:

			return s.Name.Name
		}
	}
	return ""
}
