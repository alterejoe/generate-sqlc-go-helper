package data

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

type StructData struct {
	Params *dst.FieldList
	StandardData
}

type StructData_Query struct {
	StructData
}

func (sd *StructData_Query) GetStructParams() []*dst.Field {

	t := sd.Params.List[1].Type
	if !strings.Contains(fmt.Sprint(t), "Params") {
		return []*dst.Field{
			{
				Names: sd.Params.List[1].Names,
				Type:  dst.NewIdent(fmt.Sprint("*", t)),
			},
		}
	} else {
		return []*dst.Field{
			{
				Names: []*dst.Ident{dst.NewIdent("Params")},
				Type:  dst.NewIdent(fmt.Sprint("*db.", t)),
			},
		}
	}
}

type StructData_Display struct {
	StructData
}

func (sd *StructData_Display) GetStructParams() []*dst.Field {
	fmt.Printf("sd.Params: %v\n", sd.Params)
	return []*dst.Field{}
}
