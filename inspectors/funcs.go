package inspectors

import (
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/dave/dst"
)

type FuncParams struct {
	FuncName string
	Args     *dst.FieldList
	Results  *dst.FieldList
}

func Funcs(n dst.Node, d deps.Deps, output func(FuncParams, deps.Deps) []dst.Decl) []dst.Decl {
	switch node := n.(type) {
	case *dst.FuncDecl:

		params := FuncParams{
			FuncName: node.Name.Name,
			Args:     node.Type.Params,
			Results:  node.Type.Results,
		}

		return output(params, d)

	default:
		return nil
	}
}
