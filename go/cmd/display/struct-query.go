package display

import (
	"github.com/alterejoe/generate/sqlc-go-helper/cmd/parse"
	"github.com/dave/dst"
)

type TemplateSqlcquery_QStruct struct {
	Name   string
	Fields []*dst.Field
	parse.StandardData
}

func (sqqs *TemplateSqlcquery_QStruct) GetTypespec() *dst.TypeSpec {
	return &dst.TypeSpec{
		Name: dst.NewIdent(sqqs.Name),
		Type: &dst.StructType{
			Fields: &dst.FieldList{
				List: sqqs.GetStructFields(),
			},
		},
	}
}

func (sqqs *TemplateSqlcquery_QStruct) GetStructFields() []*dst.Field {
	return sqqs.Fields
}

func (sqqs *TemplateSqlcquery_QStruct) GetStandardData() parse.StandardData {
	return sqqs.StandardData
}
