package generators

import (
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/inspectors"
	"github.com/dave/dst"
)

func DbInterfaceAdapters(p inspectors.StructParams, deps deps.Deps) []dst.Decl {
	var decls []dst.Decl

	gen := &dst.FuncDecl{}

	decls = append(decls, gen)
	return decls
}
