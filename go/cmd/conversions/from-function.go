package conversions

import (
	"fmt"

	"github.com/dave/dst"
)

// types: struct and function
// outputs: query and display

// toType enum

func FromFunction(v *dst.FuncDecl) {
	fmt.Println("Function")
}
