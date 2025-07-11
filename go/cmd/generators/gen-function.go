package generators

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
	"github.com/dave/dst"
)

// think of this as a function for the interface
// the generator provides the necessary information and
// standardizes it to the expected data
// then i can use functions to get the data i want
//
// by generating the final output here I can take advantage of the internal
// expected features of the Function
//
// This standardizes the function, and allows me to create alternates
type Gen_FunctionGenerator struct {
	Gen *dst.GenDecl
	interfaces.Function
}

func (sd *Gen_FunctionGenerator) Generate() *dst.FuncDecl {
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
