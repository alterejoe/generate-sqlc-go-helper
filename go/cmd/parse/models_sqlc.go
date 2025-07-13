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

func ParseModelsSqlc(n dst.Node, logger *slog.Logger) []string {
	switch v := n.(type) {
	case *dst.GenDecl:
		if v.Tok != token.TYPE {
			return []string{}
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
		var strs []string
		props := &data.GenToSqlcQueryProps{
			Name:       ts.Name.String(),
			Gendecl:    v,
			TypeSpec:   ts,
			StructSpec: st,
			Logger:     logger,
		}
		if sqlcdata := data.GenToSqlcQuery(props); sqlcdata != nil {
			sqlcquery := generators.SqlcSelectGenerate(sqlcdata)
			strs = append(strs, sqlcquery)
		}
		return strs
	default:
		return nil
	}
}
