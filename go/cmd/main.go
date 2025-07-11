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

		if strings.HasSuffix(path, "models.go") {
			m := runner(path, parse.ParseModels)
			models = append(models, m...)
		} else if strings.HasSuffix(path, ".sql.go") {
			q := runner(path, parse.ParseQueries)
			queries = append(queries, q...)
		}
		return nil
	})

	// displayFile(queries)
	displayFile(models)

	imports := []string{
		"context",
		"github.com/alterejoe/budget/db",
		"github.com/alterejoe/budget/web/internal/queries"}
	queriesfile := dst.File{
		Name:  dst.NewIdent("queries"),
		Decls: queries,
	}
	addImports(&queriesfile, imports)
	// writeToFile(&queriesfile, "../../../budget/web-budget/web/internal/queries/generated.go")
	// writeToFile(&dst.File{Name: dst.NewIdent("models"), Decls: models}, "models.go")
}
