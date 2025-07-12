package generators

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

func FunctionGenerate(f interfaces.DisplayFunction) *dst.FuncDecl {
	return &dst.FuncDecl{
		Recv: f.GetGenerateReceiver(),
		Name: dst.NewIdent(f.GetGenerateFunctionName()),
		Type: &dst.FuncType{
			Params:  f.GetGenerateFunctionParams(),
			Results: f.GetGenerateResults(),
		},
		Body: f.GetBody(),
	}
}
