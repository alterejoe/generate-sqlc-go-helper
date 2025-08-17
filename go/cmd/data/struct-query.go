package data

import (
	"fmt"
	"strings"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/helper"
	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it easier to sort and delegate data to its respective parser
func FuncToQueryStruct(f *dst.FuncDecl) *Funcdecl_toQueryStruct {
	fd_ts := &Funcdecl_toQueryStruct{
		Funcdecl: f,
		StandardData: &StandardData{
			Name: f.Name.String(),
		},
	}
	return fd_ts
}

type Funcdecl_toQueryStruct struct {
	Funcdecl *dst.FuncDecl
	*StandardData
}

func (fdts *Funcdecl_toQueryStruct) GetTypeSpec() *dst.TypeSpec {
	return &dst.TypeSpec{
		Name: dst.NewIdent(fdts.GetName()),
		Type: &dst.StructType{
			Fields: &dst.FieldList{
				List: fdts.GetStructFields(),
			},
		},
	}
}
func (sqqs *Funcdecl_toQueryStruct) GetStructFields() []*dst.Field {
	// fmt.Println(sqqs.Funcdecl)
	switch len(sqqs.Funcdecl.Type.Params.List) {
	case 2:
		second := sqqs.Funcdecl.Type.Params.List[1]
		if strings.Contains(fmt.Sprint(second.Type), "Params") {
			return []*dst.Field{
				{
					Names: []*dst.Ident{dst.NewIdent("Params")},
					Type:  dst.NewIdent(fmt.Sprint("db.", second.Type)),
				},
			}
		} else {
			propercasename := strings.Title(second.Names[0].Name)
			return []*dst.Field{
				{
					Names: []*dst.Ident{dst.NewIdent(propercasename)},
					Type:  dst.NewIdent(helper.ToPgtype(&second.Type)),
				},
			}
		}
	default:
		return []*dst.Field{}
	}
}
