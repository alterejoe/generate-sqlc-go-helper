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
		ts, err := genTo.ToTypeSpec()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		st, err := genTo.ToStructType()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fields := st.Fields.List
		decls := []dst.Decl{}
		for _, f := range fields {
			props := &data.GenToDisplayFunctionProps{
				Name:    ts.Name.String(),
				Field:   f.Names[0].String(),
				Gendecl: v,
			}
			if funcdata := data.GenToDisplayFunction(props); funcdata != nil {
				funcgen := generators.FunctionGenerate(funcdata)
				decls = append(decls, funcgen)
			}
		}
		return decls
	default:
		return nil
	}
}
