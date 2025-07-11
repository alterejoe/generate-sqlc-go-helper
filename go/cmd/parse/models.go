package parse

import (
	"github.com/dave/dst"
)

func ParseModels(n dst.Node) []dst.Decl {
	switch n.(type) {
	case *dst.GenDecl:
		//output:
		// one function
		return []dst.Decl{}
	default:
		return nil
	}
}
