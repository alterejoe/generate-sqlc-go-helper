package data

import (
	"github.com/dave/dst"
)

type StructData struct {
	Fields *dst.FieldList
	StandardData
}

type StructData_Query struct {
	StructData
}

func (sd *StructData_Query) GetStructParams() []*dst.Field {
	var out []*dst.Field
	params := sd.Fields.List
	for _, param := range params {
		f := param.Names[0]
		t := param.Type
		out = append(out, &dst.Field{
			Names: []*dst.Ident{f},
			Type:  t,
		})
	}
	return out
}

type StructData_Display struct {
	StructData
}

func (sd *StructData_Display) GetStructParams() []*dst.Field {

	return []*dst.Field{}
}
