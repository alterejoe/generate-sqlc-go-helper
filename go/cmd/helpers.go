package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"os/exec"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func writeToFile(file *dst.File, filename string) {
	// Create the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// Write to the file
	if err := decorator.Fprint(f, file); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Format the file using gofmt
	cmd := exec.Command("gofmt", "-w", filename)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error formatting file:", err)
	}
}

func runner(path string, logger *slog.Logger, inspector func(dst.Node, *slog.Logger) []dst.Decl) []dst.Decl {
	fset := token.NewFileSet()
	f, err := decorator.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	var decls []dst.Decl
	dst.Inspect(f, func(n dst.Node) bool {
		d := inspector(n, logger)
		decls = append(decls, d...)
		return true
	})

	return decls
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
