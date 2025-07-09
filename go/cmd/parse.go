package main

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

// stdout

func parse_models(n dst.Node) []dst.Decl {
	switch n.(type) {
	case *dst.GenDecl:
		// s := CreateStruct(v)
		// return []dst.Decl{s}
	default:
		return nil
	}
	return nil
}

func addToDecls(add *[]dst.Decl, decls *[]dst.Decl) {
	for _, decl := range *decls {
		*add = append(*add, decl)
	}
}

type ParseProps struct {
	Name    string
	Params  []*dst.Field
	Results *dst.FieldList
}

func (qmp *ParseProps) GetLowerName() string {
	return strings.ToLower(qmp.Name)
}

func (qmp *ParseProps) GetAbbv() string {
	abbv := ""
	for _, c := range qmp.Name {
		if rune('A') <= c && c <= rune('Z') {
			abbv += string(c)
		}
	}
	s := strings.ToLower(abbv)
	return s
}

func (qmp *ParseProps) GetName() string {
	return qmp.Name
}

func (qmp *ParseProps) StructParams() *dst.Field {
	if len(qmp.Params) > 1 {
		t := qmp.Params[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return &dst.Field{
				Names: qmp.Params[1].Names,
				Type:  dst.NewIdent(fmt.Sprint("*", t)),
			}
		} else {
			field := &dst.Field{
				Names: []*dst.Ident{dst.NewIdent("Params")},
				Type:  dst.NewIdent(fmt.Sprint("*db.", t)),
			}
			return field
		}
	}
	return nil
}

func (qmp *ParseProps) GetSecondArg() string {
	if len(qmp.Params) > 1 {
		t := qmp.Params[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return fmt.Sprint(qmp.Params[1].Names[0])
		} else {
			return "Params"
		}
	}
	return ""
}

func (qmp *ParseProps) QueryArgs(params *[]dst.Expr) {
	if len(qmp.Params) > 1 {
		*params = append(*params, dst.NewIdent(fmt.Sprintf("*%s.%s", qmp.GetAbbv(), qmp.GetSecondArg())))
	}
}

func (qmp *ParseProps) Returns() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *ParseProps) QueryAddErr(results *[]dst.Expr, function bool) {
	if len(qmp.Results.List) > 1 {
		*results = append(*results, dst.NewIdent("err"))
	} else if !function {
		*results = append(*results, dst.NewIdent("nil"))
	}
}

func retrieveProps(n *dst.FuncDecl) *ParseProps {
	return &ParseProps{
		Name:    n.Name.Name,
		Params:  n.Type.Params.List,
		Results: n.Type.Results,
	}
}

func parse_queries(n dst.Node) []dst.Decl {
	var decls []dst.Decl
	switch v := n.(type) {
	case *dst.FuncDecl:
		// if is struct
		props := retrieveProps(v)
		s := CreateStruct(props)
		m := CreateQueryMethod(props)

		addToDecls(&decls, &s)
		addToDecls(&decls, &m)
		return decls
	default:
		return nil
	}
}
