package inspectors

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/dave/dst"
)

type Params struct {
	StructName string
	Fields     map[string]dst.Expr
}

func Struct(n dst.Node, d deps.Deps, output func(Params, deps.Deps) []dst.Decl) []dst.Decl {
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

		fields := make(map[string]dst.Expr)
		if len(st.Fields.List) > 0 {
			for _, field := range st.Fields.List {
				if len(field.Names) > 0 {
					fields[field.Names[0].Name] = field.Type
				}
			}
		}

		params := Params{
			StructName: ts.Name.Name,
			Fields:     fields,
		}

		// d.Logger.Info("found struct", "name", ts.Name.Name, "params", params)
		return output(params, d)

	default:
		return nil
	}
}
