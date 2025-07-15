package parse

import (
	"fmt"
	"go/token"
	"log/slog"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/data"
	dstto "github.com/alterejoe/generate/sqlc-go-helper/cmd/dst-to"
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/generators"
	"github.com/dave/dst"
)

func ParseInterfaces(n dst.Node, logger *slog.Logger) []dst.Decl {
	switch v := n.(type) {
	case *dst.GenDecl:
		if v.Tok != token.TYPE {
			return []dst.Decl{}
		}

		genTo := dstto.GenTo{GenDecl: v}
		ts, err := genTo.ToTypeSpec()
		if err != nil {
			// fmt.Println(err)
			// panic(err)
		}
		st, err := genTo.ToStructType()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		decls := []dst.Decl{}

		props := data.GenToDisplayInterfaceProps{
			Name:       ts.Name.String(),
			Gendecl:    v,
			TypeSpec:   ts,
			StructSpec: st,
			Logger:     logger,
		}
		if structdata := data.GenToDisplayInterface(&props); structdata != nil {
			structgen := generators.StructGenerate(structdata)
			decls = append(decls, structgen)
		}
		return decls
	default:
		return nil
	}
}
