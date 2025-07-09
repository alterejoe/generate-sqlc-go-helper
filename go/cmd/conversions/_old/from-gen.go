package conversions

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
)

type FromGenDecl struct {
	Name   string
	Input  *dst.GenDecl
	Params *dst.FieldList
}

func (fgd *FromGenDecl) GetName() string {
	for _, spec := range fgd.Input.Specs {
		switch s := spec.(type) {
		case *dst.TypeSpec:

			fgd.Name = s.Name.Name
			return s.Name.Name
		}
	}
	return ""
}

func (qmp *FromGenDecl) GetLowerName() string {
	return strings.ToLower(qmp.Name)
}

func (qmp *FromGenDecl) GetReturns() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *FromGenDecl) GetStructParams() []*dst.Field {
	if qmp.Input.Tok != token.TYPE {
		return nil
	}

	structDecl := qmp.Input
	typeSpec, ok := structDecl.Specs[0].(*dst.TypeSpec)
	if !ok {
		return nil
	}
	structSpec := typeSpec.Type.(*dst.StructType)
	qmp.Params = structSpec.Fields
	if len(qmp.Params.List) > 1 {
		t := qmp.Params.List[1].Type

		if !strings.Contains(fmt.Sprint(t), "Params") {
			return []*dst.Field{{
				Names: qmp.Params.List[1].Names,
				Type:  dst.NewIdent(fmt.Sprint("*", t)),
			}}
		} else {

			return []*dst.Field{
				{
					Names: []*dst.Ident{dst.NewIdent("Params")},
					Type:  dst.NewIdent(fmt.Sprint("*db.", t)),
				},
			}
		}
	}
	return nil
}
