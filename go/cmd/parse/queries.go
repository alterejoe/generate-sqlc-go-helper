package parse

import (
	"log/slog"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/data"
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/generators"
	"github.com/dave/dst"
)

func ParseQueries(n dst.Node, logger *slog.Logger) []dst.Decl {
	switch v := n.(type) {
	case *dst.FuncDecl:
		//output:
		// one struct
		structdata := data.FuncToQueryStruct(v)
		funcdata := data.FuncToQueryFunction(v)
		structgen := generators.StructGenerate(structdata)
		funcgen := generators.FunctionGenerate(funcdata)
		return []dst.Decl{
			structgen,
			funcgen,
		}
	default:
		return nil
	}
}
