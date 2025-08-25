package inspectors

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/dave/dst"
)

type StructParams struct {
	StructName string
	FieldList  *dst.FieldList
}

func Struct(n dst.Node, d deps.Deps, output func(StructParams, deps.Deps) []dst.Decl) []dst.Decl {
	switch v := n.(type) {
	// this is a struct
	case *dst.GenDecl:
		if v.Tok != token.TYPE {
			return []dst.Decl{}
		}

		if len(v.Specs) != 1 {
			return []dst.Decl{}
		}

		ts, ok := v.Specs[0].(*dst.TypeSpec)
		if !ok {
			return []dst.Decl{}
		}

		st, ok := ts.Type.(*dst.StructType)
		if !ok {
			return []dst.Decl{}
		}

		params := StructParams{
			StructName: ts.Name.Name,
			FieldList:  st.Fields,
		}

		// d.Logger.Info("found struct", "name", ts.Name.Name, "params", params)
		return output(params, d)

	default:
		return nil
	}
}
