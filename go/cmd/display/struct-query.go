package data

import "github.com/dave/dst"

type Sqlcquery_QueryStruct struct {
	Name   string
	Fields []*dst.Field
	StandardData
}

func (sqqs *Sqlcquery_QueryStruct) GetTypespec() *dst.TypeSpec {
	return &dst.TypeSpec{
		Name: dst.NewIdent(sqqs.Name),
		Type: &dst.StructType{
			Fields: &dst.FieldList{
				List: sqqs.GetStructFields(),
			},
		},
	}
}

func (sqqs *Sqlcquery_QueryStruct) GetStructFields() []*dst.Field {
	return sqqs.Fields
}

func (sqqs *Sqlcquery_QueryStruct) GetStandardData() StandardData {
	return sqqs.StandardData
}
