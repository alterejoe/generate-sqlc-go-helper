package data

import (
	"log/slog"

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
		Gendecl:  props.Gendecl,
		TypeSpec: props.TypeSpec,
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

func (gdtdi *Gendecl_toDisplayInterface) GetTypeSpec() *dst.TypeSpec {
	return gdtdi.TypeSpec
}
