package parse

import (
	"fmt"
	"go/token"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/data"
	dstto "github.com/alterejoe/generate/sqlc-go-helper/cmd/dst-to"
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/generators"
	"github.com/dave/dst"
)

func ParseModels(n dst.Node) []dst.Decl {
	switch v := n.(type) {
	case *dst.GenDecl:
		if v.Tok != token.TYPE {
			return []dst.Decl{}
		}

		genTo := dstto.GenTo{GenDecl: v}
		st, err := genTo.ToStructType()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fields := st.Fields
		decls := []dst.Decl{}
		for _, f := range fields.List {
			if funcdata := data.FieldToDisplayFunction(f.Names[0].Name, v); funcdata != nil {
				funcgen := generators.FunctionGenerate(funcdata)
				decls = append(decls, funcgen)
			}
		}
		return decls
	default:
		return nil
	}
}
