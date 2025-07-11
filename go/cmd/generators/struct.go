package generators

import (
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

func StructGenerate(s interfaces.DisplayStruct) *dst.GenDecl {
	return &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			s.GetTypeSpec(),
		},
	}
}
