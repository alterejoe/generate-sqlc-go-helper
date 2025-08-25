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
	cfg := GetEnv()

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
		Dir: cfg.DbProjectPath,
	}

	var dbqueries []dst.Decl
	var querys_displayfunctions []dst.Decl
	var model_displayfunctions []dst.Decl
	var model_displayinterfaces []dst.Decl

	_ = filepath.WalkDir(paths.Dir, func(path string, d fs.DirEntry, err error) error {

		if strings.HasSuffix(path, "models.go") {
			m := dst_inspect_file(path, logger, parse.ParseModels)
			i := dst_inspect_file(path, logger, parse.ParseInterfaces)

			model_displayfunctions = append(model_displayfunctions, m...)
			model_displayinterfaces = append(model_displayinterfaces, i...)

		} else if strings.HasSuffix(path, ".sql.go") {
			qm := dst_inspect_file(path, logger, parse.ParseModels)

			querys_displayfunctions = append(querys_displayfunctions, qm...)
		}
		if strings.HasSuffix(path, ".sql.go") {
			q := dst_inspect_file(path, logger, parse.ParseQueries)
			i := dst_inspect_file(path, logger, parse.ParseInterfaces)

			model_displayinterfaces = append(model_displayinterfaces, i...)
			dbqueries = append(dbqueries, q...)
		}
		return nil
	})

	db_imports := []string{
		"context",
		cfg.DbProjectUrl,
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

	model_interface_imports := []string{
		"github.com/google/uuid",
		"time",
	}
	model_interface_file := dst.File{
		Name:  dst.NewIdent("interfaces"),
		Decls: model_displayinterfaces,
	}

	addImports(&db_queries_file, db_imports)
	addImports(&model_df_file, model_df_imports)
	addImports(&query_df_file, model_df_imports)
	addImports(&model_interface_file, model_interface_imports)

	// writeToFile(&db_queries_file, fmt.Sprint(writeto, "/internal/queries/generated.go"))
	// writeToFile(&model_df_file, fmt.Sprint(writeto, "/db/generated-queries.go"))
	// writeToFile(&model_interface_file, fmt.Sprint(writeto, "/interfaces/generated.go"))
	writeToFile(&db_queries_file, cfg.DbQueryParamOut)
	writeToFile(&model_df_file, cfg.DbInterfaceAdaptersOut)
	writeToFile(&query_df_file, cfg.DbInterfaceAdaptersOut)

	logger.Info("Done", "writeto", cfg.DbQueryParamOut)
	fmt.Println("Done")
}
