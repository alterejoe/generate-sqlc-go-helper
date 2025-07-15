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
		if interfacedata := data.GenToDisplayInterface(&props); interfacedata != nil {
			interfacegen := generators.InterfaceGenerate(interfacedata)
			decls = append(decls, interfacegen)
		}
		return decls
	default:
		return nil
	}
}
