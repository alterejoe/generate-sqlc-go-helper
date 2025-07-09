package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type Paths struct {
	Dir string
}

func main() {
	paths := Paths{
		Dir: "/home/altjoe/Documents/projects/budget/web-budget/web/db",
	}
	var queries []dst.Decl
	var models []dst.Decl

	_ = filepath.WalkDir(paths.Dir, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, "models.go") {
			m := runner(path, parse_models)
			models = append(models, m...)
		} else if strings.HasSuffix(path, ".sql.go") {
			q := runner(path, parse_queries)
			queries = append(queries, q...)
		}
		return nil
	})
	queriesfile := &dst.File{
		Name:  dst.NewIdent("queries"),
		Decls: queries,
	}
	// output to std
	if err := decorator.Fprint(os.Stdout, queriesfile); err != nil {
		fmt.Println(err)
		return
	}
	modelsfile := &dst.File{
		Name:  dst.NewIdent("models"),
		Decls: models,
	}

	if err := decorator.Fprint(os.Stdout, modelsfile); err != nil {
		fmt.Println(err)
		return
	}
}

func runner(path string, inspector func(dst.Node) []dst.Decl) []dst.Decl {
	fset := token.NewFileSet()
	f, err := decorator.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	var decls []dst.Decl
	dst.Inspect(f, func(n dst.Node) bool {
		d := inspector(n)
		decls = append(decls, d...)
		return true
	})

	return decls
}
