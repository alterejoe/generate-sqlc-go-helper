package main

import (
	"strings"

	"github.com/dave/dst"
)

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
	Name   string
	Params []*dst.Field
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

func (qmp *ParseProps) ExtraStructParam() *dst.Field {
	if len(qmp.Params) > 1 {
		t := qmp.Params[1].Type

		field := &dst.Field{
			Names: []*dst.Ident{dst.NewIdent("Params")},
			Type:  qmp.Params[1].Type,
		}
		return field
	}
	return nil
}

func retrieveProps(n *dst.FuncDecl) *ParseProps {
	return &ParseProps{
		Name:   n.Name.Name,
		Params: n.Type.Params.List,
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
