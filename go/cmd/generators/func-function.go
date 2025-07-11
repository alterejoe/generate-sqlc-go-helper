package generators

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

type Func_FunctionGenerator struct {
	Func *dst.FuncDecl
	interfaces.Function
}

func (sd *Func_FunctionGenerator) Generate() *dst.FuncDecl {
	return &dst.FuncDecl{
		Recv: sd.GetReceiver(),
		Name: dst.NewIdent("Query"),
		Type: &dst.FuncType{
			Params:  sd.GetFunctionParams(),
			Results: sd.GetResults(),
		},
		Body: sd.GetBody(),
	}
}
