package generators

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

type Func_StructGenerator struct {
	Funcs *dst.FuncDecl
	interfaces.Struct
}

func (sqs *Func_StructGenerator) Generate() *dst.GenDecl {
	return &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			sqs.GetTypeSpec(),
		},
	}
}
