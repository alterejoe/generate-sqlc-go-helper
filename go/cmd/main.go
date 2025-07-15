package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/parse"
	"github.com/dave/dst"
	"github.com/golang-cz/devslog"
	"github.com/joho/godotenv"
)

type Paths struct {
	Dir    string
	WinDir string
}

func main() {
	// source

	slogOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	opts := &devslog.Options{
		HandlerOptions:    slogOpts,
		MaxSlicePrintSize: 10,
		SortKeys:          true,
		NewLineAfterLog:   true,
		StringerFormatter: true,
		NoColor:           true,
	}
	handler := devslog.NewHandler(os.Stdout, opts)
	logger := slog.New(handler)

	godotenv.Load()
	paths := Paths{
		Dir: os.Getenv("SQLCDIR"),
	}

	var dbqueries []dst.Decl
	var querys_displayfunctions []dst.Decl
	var model_displayfunctions []dst.Decl
	var sqlc_queryies []string

	_ = filepath.WalkDir(paths.Dir, func(path string, d fs.DirEntry, err error) error {

		if strings.HasSuffix(path, "models.go") {
			m := runner(path, logger, parse.ParseModels)
			s := strRunner(path, logger, parse.ParseModelsSqlc)
			model_displayfunctions = append(model_displayfunctions, m...)
			sqlc_queryies = append(sqlc_queryies, s...)
		} else if strings.HasSuffix(path, ".sql.go") {
			qm := runner(path, logger, parse.ParseModels)
			querys_displayfunctions = append(querys_displayfunctions, qm...)
		}
		if strings.HasSuffix(path, ".sql.go") {
			q := runner(path, logger, parse.ParseQueries)
			// will need to parse structs related to queries as well
			// for display purposes
			// each struct ex: GerFirstAccountRow related to a query
			// need the display functions to be attached not just the original models.go
			dbqueries = append(dbqueries, q...)
		}
		return nil
	})

	for i, q := range sqlc_queryies {
		fmt.Println(i, q)
	}
	// displayFile(queries)
	// displayFile(models)
	//
	db_imports := []string{
		"context",
		"github.com/alterejoe/budget/db",
	}
	db_queries_file := dst.File{
		Name:  dst.NewIdent("queries"),
		Decls: dbqueries,
	}
	model_df_imports := []string{
		"fmt",
		"time",
		"github.com/google/uuid",
	}
	model_df_file := dst.File{
		Name:  dst.NewIdent("db"),
		Decls: model_displayfunctions,
	}

	query_df_file := dst.File{
		Name:  dst.NewIdent("db"),
		Decls: querys_displayfunctions,
	}

	// sqlc_queryies_file := dst.File{
	// 	Name:  dst.NewIdent("queries"),
	// 	Decls: sqlc_queryies,
	// }

	addImports(&db_queries_file, db_imports)
	addImports(&model_df_file, model_df_imports)
	addImports(&query_df_file, model_df_imports)
	// addImports(&sqlc_queryies_file, model_df_imports)
	writeToFile(&db_queries_file, "../../../budget/web-budget/web/internal/queries/generated.go")
	writeToFile(&query_df_file, "../../../budget/web-budget/web/db/generated-queries.go")
	writeToFile(&model_df_file, "../../../budget/web-budget/web/db/generated-models.go")
	// writeToFile(&sqlc_queryies_file, "../../../budget/web-budget/web/internal/sqlc/generated.go")
	// writeToFile(&dst.File{Name: dst.NewIdent("models"), Decls: models}, "models.go")
}
