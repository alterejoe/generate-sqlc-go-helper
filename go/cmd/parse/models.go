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

func GenText(t dst.Expr) bool {
	// switch v := fmt.Sprintf(t.Type), v {
	switch v := fmt.Sprintf("%s", t); v {
	case "&{pgtype Text {{None [] [] None} []}}", "string", "&{<nil> bool {{None [] [] None} [] []}}", "&{<nil> string {{None [] [] None} [] []}}":
		return false
	default:
		return true
	}
}

func ParseModels(n dst.Node, logger *slog.Logger) []dst.Decl {
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
		fields := st.Fields.List
		decls := []dst.Decl{}
		for _, f := range fields {
			props := &data.GenToDisplayFunctionProps{
				Name:       ts.Name.String(),
				Field:      f,
				Gendecl:    v,
				TypeSpec:   ts,
				StructSpec: st,
				Logger:     logger,
			}
			if funcdata := data.GenToDisplayFunction(props); funcdata != nil {
				funcgen := generators.FunctionGenerate(funcdata)
				decls = append(decls, funcgen)
			}
			if GenText(f.Type) {
				if fd := data.GenToDisplayTextFunction(props); fd != nil {
					f := generators.FunctionGenerate(fd)
					decls = append(decls, f)
				}
			}
		}
		return decls
	default:
		return nil
	}
}
