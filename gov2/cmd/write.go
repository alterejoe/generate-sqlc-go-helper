package main

import (
	"fmt"
	"go/token"
	"os"
	"os/exec"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type WriteFile struct {
	Imports     []string
	FileName    string
	PackageName string
	Out         []dst.Decl
}

func writeToFile(wf WriteFile) {
	file := dst.File{
		Name:  dst.NewIdent("interfaces"),
		Decls: wf.Out,
	}

	addImports(&file, wf.Imports)

	// Create the file
	f, err := os.Create(wf.PackageName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// Write to the file
	if err := decorator.Fprint(f, &file); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Format the file using gofmt
	cmd := exec.Command("gofmt", "-w", wf.FileName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error formatting file:", err)
	}
}

func addImports(file *dst.File, imports []string) {
	for _, imp := range imports {
		importSpec := &dst.ImportSpec{
			Path: &dst.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%q", imp),
			},
		}
		importDecl := &dst.GenDecl{
			Tok: token.IMPORT,
			Specs: []dst.Spec{
				importSpec,
			},
		}
		file.Decls = append([]dst.Decl{importDecl}, file.Decls...)
	}
}
