package data

import (
	"log/slog"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/helper"
	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it nasier to sort and delegate data to its respective parser
type GenToDisplayInterfaceProps struct {
	Name       string
	Gendecl    *dst.GenDecl
	TypeSpec   *dst.TypeSpec
	StructSpec *dst.StructType
	Logger     *slog.Logger
}

func GenToDisplayInterface(props *GenToDisplayInterfaceProps) *Gendecl_toDisplayInterface {
	fd_ts := &Gendecl_toDisplayInterface{
		Gendecl:    props.Gendecl,
		TypeSpec:   props.TypeSpec,
		StructSpec: props.StructSpec,
		StandardData: &StandardData{
			Name:   props.Name,
			Logger: props.Logger,
		},
	}
	return fd_ts
}

type Gendecl_toDisplayInterface struct {
	Gendecl    *dst.GenDecl
	TypeSpec   *dst.TypeSpec
	StructSpec *dst.StructType
	*StandardData
}

func (gdtdi *Gendecl_toDisplayInterface) GetStructFields() []*dst.Field {
	return gdtdi.StructSpec.Fields.List
}

func (gdtdi *Gendecl_toDisplayInterface) GetFuncitonFields() []*dst.Field {
	var fields []*dst.Field
	for _, v := range gdtdi.GetStructFields() {
		// gdtdi.GetLogger().Info("GetFuncitonFields", slog.Int("index", i), slog.String("value", ))
		if len(v.Names) == 0 {
			continue
		}
		t := &dst.Field{
			Names: []*dst.Ident{dst.NewIdent("Get" + v.Names[0].Name)},
			Type: &dst.FuncType{
				Params: &dst.FieldList{}, // no params
				Results: &dst.FieldList{
					List: []*dst.Field{
						{
							Type: dst.NewIdent(helper.ToStandardReturnType(&v.Type)),
						},
					},
				},
			},
		}
		fields = append(fields, t)
		if helper.CheckGenText(v.Type) {
			ttext := &dst.Field{
				Names: []*dst.Ident{dst.NewIdent("Get" + v.Names[0].Name + "Text")},
				Type: &dst.FuncType{
					Params: &dst.FieldList{}, // no params
					Results: &dst.FieldList{
						List: []*dst.Field{{
							Type: dst.NewIdent("string"),
						}},
					},
				},
			}
			fields = append(fields, ttext)
		}
	}
	return fields
}

func (gdtdi *Gendecl_toDisplayInterface) GetTypeSpec() *dst.TypeSpec {
	// return gdtdi.TypeSpec
	return &dst.TypeSpec{
		Name: dst.NewIdent(gdtdi.GetName()),
		Type: &dst.InterfaceType{
			Methods: &dst.FieldList{
				List: gdtdi.GetFuncitonFields(),
			},
		},
	}
}
