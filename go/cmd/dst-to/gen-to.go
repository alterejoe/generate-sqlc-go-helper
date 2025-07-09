package dstto

import (
	"fmt"

	"github.com/dave/dst"
)

type GenTo struct{ *dst.GenDecl }

func (gt *GenTo) ToTypeSpec() (*dst.TypeSpec, error) {
	if len(gt.Specs) == 0 {
		return nil, fmt.Errorf("len(gt.Specs) == 0")
	}
	ts, ok := gt.Specs[0].(*dst.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("gt.Specs[0].(*dst.TypeSpec)")
	}

	return ts, nil
}

func (gt *GenTo) ToStructType() (*dst.StructType, error) {
	ts, err := gt.ToTypeSpec()
	if err != nil {
		return nil, err
	}

	st, ok := ts.Type.(*dst.StructType)
	if !ok {
		return nil, fmt.Errorf("ts.Type.(*dst.StructType)")
	}
	return st, nil
}
