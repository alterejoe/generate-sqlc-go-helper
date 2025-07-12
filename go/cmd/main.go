package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/parse"
	"github.com/dave/dst"
	"github.com/joho/godotenv"
)

type Paths struct {
	Dir    string
	WinDir string
}

func main() {
	// source
	godotenv.Load()
	paths := Paths{
		Dir: os.Getenv("SQLCDIR"),
	}

	var queries []dst.Decl
	var models []dst.Decl

	_ = filepath.WalkDir(paths.Dir, func(path string, d fs.DirEntry, err error) error {

		if strings.HasSuffix(path, "models.go") || strings.HasSuffix(path, ".sql.go") {
			m := runner(path, parse.ParseModels)
			models = append(models, m...)
		} else if strings.HasSuffix(path, ".sql.go") {
			q := runner(path, parse.ParseQueries)
			// will need to parse structs related to queries as well
			// for display purposes
			// each struct ex: GerFirstAccountRow related to a query
			// need the display functions to be attached not just the original models.go
			queries = append(queries, q...)
		}
		return nil
	})

	// displayFile(queries)
	// displayFile(models)
	//
	queriesimports := []string{
		"context",
		"github.com/alterejoe/budget/db",
	}
	queriesfile := dst.File{
		Name:  dst.NewIdent("queries"),
		Decls: queries,
	}
	modelsimports := []string{
		"fmt",
		"time",
	}
	modelsfile := dst.File{
		Name:  dst.NewIdent("db"),
		Decls: models,
	}
	addImports(&queriesfile, queriesimports)
	addImports(&modelsfile, modelsimports)
	writeToFile(&queriesfile, "../../../budget/web-budget/web/internal/queries/generated.go")
	writeToFile(&modelsfile, "../../../budget/web-budget/web/db/generated.go")
	// writeToFile(&dst.File{Name: dst.NewIdent("models"), Decls: models}, "models.go")
}
