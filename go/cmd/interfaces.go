package main

import "github.com/dave/dst"

// this is so that the inteface will be verbose
// these will be used to convert from *dst.GenDecl
type FromFunc interface {
	QueryFunc
}

// this is so that the inteface will be verbose
// these will be used to convert from *dst.GenDecl
type FromStruct interface {
	QueryStruct
}

type GenStandard interface {
	GetName() string
	GetLowerName() string
	GetReturns() *dst.FieldList
}

type FuncStandard interface {
	GetAbbv() string
	GetSecondArg() string
	GetQueryArgs(params *[]dst.Expr)
	GetQueryAddErr(results *[]dst.Expr, function bool)
}

type StructStandard interface {
	StructParams() *dst.Field
}

type QueryStruct interface {
	GenStandard
	StructStandard
}

type QueryFunc interface {
	GenStandard
	FuncStandard
}
