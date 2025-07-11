package generators

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

func FunctionGenerate(f interfaces.DisplayFunction) *dst.FuncDecl {
	return &dst.FuncDecl{
		Recv: f.GetReceiver(),
		Name: dst.NewIdent(f.GetFunctionName()),
		Type: &dst.FuncType{
			Params:  f.GetFunctionParams(),
			Results: f.GetResults(),
		},
		Body: f.GetBody(),
	}
}
