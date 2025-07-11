package generators

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

type Gen_StructGenerator struct {
	Gen *dst.GenDecl
	interfaces.Struct
}

func (sqs *Gen_StructGenerator) Generate() *dst.GenDecl {
	return &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			sqs.GetTypeSpec(),
		},
	}
}
