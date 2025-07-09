package conversions

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/data"
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

// types: struct and function
// outputs: query and display

// toType enum

func FromStruct(v *dst.GenDecl, t string) interfaces.Struct {
	_ = &data.StandardData{
		Name: GetName(v),
	}

	switch t {
	case "display":
		return &data.StructData_Display{}
	case "query":
		return &data.StructData_Query{}
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
