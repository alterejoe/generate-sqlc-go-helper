package data

import (
	"fmt"
	"log/slog"

	"github.com/dave/dst"
)

// / come back to this if you get lost again
// using this factory style function we can pass parameters to children
// this makes it easier to sort and delegate data to its respective parser
type GenToSqlcQueryProps struct {
	Name       string
	Fields     []*dst.Field
	Gendecl    *dst.GenDecl
	TypeSpec   *dst.TypeSpec
	StructSpec *dst.StructType
	Logger     *slog.Logger
}

func GenToSqlcQuery(props *GenToSqlcQueryProps) *Gendecl_toSqlcQuery {
	fd_ts := &Gendecl_toSqlcQuery{
		Fields:   props.Fields,
		Gendecl:  props.Gendecl,
		TypeSpec: props.TypeSpec,
		StandardData: &StandardData{
			Name:   props.Name,
			Logger: props.Logger,
		},
	}
	return fd_ts
}

type Gendecl_toSqlcQuery struct {
	Fields   []*dst.Field
	Gendecl  *dst.GenDecl
	TypeSpec *dst.TypeSpec
	*StandardData
}

// // dst.NewIdent(qmp.GetAbbv() + "." + qmp.GetFieldName() + ".Time"),
func (qmp *Gendecl_toSqlcQuery) ParamIdent(param string) *dst.Ident {
	return dst.NewIdent(param)
}

func (qmp *Gendecl_toSqlcQuery) PreidentInt(param string) string {
	return fmt.Sprint("int(", param, ")")
}

func (qmp *Gendecl_toSqlcQuery) PreidentPgfield(param string) string {
	return fmt.Sprint(qmp.GetAbbv(), ".", param)
}

func (qmp *Gendecl_toSqlcQuery) Ident(param string) *dst.Ident {
	return dst.NewIdent(fmt.Sprint(param))
}

// f.GetSelect(),
// f.GetFrom(),
// f.GetWhere(),
// f.GetOrderBy(),
// f.GetLimit())
func (qmp *Gendecl_toSqlcQuery) GetIdentifier() string {
	return "identifier"
}

func (qmp *Gendecl_toSqlcQuery) GetQuery() string {
	return "SELECT select, columns, here"
}

func (qmp *Gendecl_toSqlcQuery) GetFrom() string {
	return "some.stupid.table"
}

func (qmp *Gendecl_toSqlcQuery) GetWhere() string {
	return "where the=fuck"
}

func (qmp *Gendecl_toSqlcQuery) GetOrderBy() string {
	return "order by yourmom"
}

func (qmp *Gendecl_toSqlcQuery) GetLimit() string {
	return "limit a-lot"
}

func (qmp *Gendecl_toSqlcQuery) GetReturns() string {
	return ":one"
}
