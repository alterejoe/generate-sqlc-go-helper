package data

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

type StandardData struct {
	Name string
}

func (sd *StandardData) GetAbbv() string {
	abbv := ""
	for _, c := range sd.Name {
		if rune('A') <= c && c <= rune('Z') {
			abbv += string(c)
		}
	}
	s := strings.ToLower(abbv)
	return s
}

type StructData_Query struct {
	Params *dst.FieldList
	StandardData
}

func (sd *StructData_Query) GetStructParams() []*dst.Field {
	if len(sd.Params.List) > 1 {
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
	return nil
}

type StructData_Display struct {
	Params *dst.FieldList
	StandardData
}

func (sd *StructData_Display) GetStructParams() []*dst.Field {
	return []*dst.Field{}
}
