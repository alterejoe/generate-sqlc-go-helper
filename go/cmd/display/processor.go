package data

import (
	"fmt"

	"github.com/dave/dst"
)

type Processor struct {
}

func (p *Processor) ProcessFunc(n *dst.FuncDecl) {
	fmt.Printf("\"processing func\": %v\n", "processing func")
}

func (p *Processor) ProcessStruct(n *dst.GenDecl) {
	fmt.Printf("\"processing struct\": %v\n", "processing struct")
}
