package dstto

import (
	"github.com/dave/dst"
)

type FuncTo struct{ *dst.FuncDecl }

func (f *FuncTo) ToFunctionType() (*dst.FuncType, error) {
	return f.FuncDecl.Type.(*dst.FuncType), nil
}
