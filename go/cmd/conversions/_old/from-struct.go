package conversions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

type QueryStruct struct {
	Name   string
	Params *dst.FieldList
}

func (qmp *QueryStruct) Initialize() {
	qmp.Params = Params
}

func (qmp *QueryStruct) GetStructParams() []*dst.Field {
	if len(qmp.Params.List) > 1 {
		t := qmp.Params.List[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return []*dst.Field{
				{
					Names: qmp.Params.List[1].Names,
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
	return nil
}

type DisplayStruct struct {
	Name   string
	Params *dst.FieldList
}

func (qmp *DisplayStruct) GetStructParams() []*dst.Field {
	if len(qmp.Params.List) > 1 {
		t := qmp.Params.List[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return []*dst.Field{
				{
					Names: qmp.Params.List[1].Names,
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
	return nil
}
