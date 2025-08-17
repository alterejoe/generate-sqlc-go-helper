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
	writeto := os.Getenv("WRITETO")
	sqlcdir := writeto + "/db/"
	paths := Paths{
		Dir: sqlcdir,
	}

	var dbqueries []dst.Decl
	var querys_displayfunctions []dst.Decl
	var model_displayfunctions []dst.Decl
	var sqlc_queryies []string
	var model_displayinterfaces []dst.Decl

	_ = filepath.WalkDir(paths.Dir, func(path string, d fs.DirEntry, err error) error {

		if strings.HasSuffix(path, "models.go") {
			m := runner(path, logger, parse.ParseModels)
			s := strRunner(path, logger, parse.ParseModelsSqlc)
			i := runner(path, logger, parse.ParseInterfaces)
			model_displayfunctions = append(model_displayfunctions, m...)
			sqlc_queryies = append(sqlc_queryies, s...)
			model_displayinterfaces = append(model_displayinterfaces, i...)
		} else if strings.HasSuffix(path, ".sql.go") {
			qm := runner(path, logger, parse.ParseModels)
			querys_displayfunctions = append(querys_displayfunctions, qm...)
		}
		if strings.HasSuffix(path, ".sql.go") {
			q := runner(path, logger, parse.ParseQueries)
			// will need to parse structs related to queries as well
			i := runner(path, logger, parse.ParseInterfaces)
			model_displayinterfaces = append(model_displayinterfaces, i...)
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
		fmt.Sprint("github.com/alterejoe/", os.Getenv("PROJECT"), "/web/db"),
		// "github.com/alterejoe/order/db",
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

	/*this is the operation of the query as it relates to the Query interface
	  type DatabaseQueryParams interface {
	  	Query(query *model.Queries, r context.Context) (any, error)
	  	GetParams() any
	  }*/
	writeToFile(&db_queries_file, fmt.Sprint(writeto, "/internal/queries/generated.go"))

	/*This is to generate the parameter and text parameter functions to add onto sqlc generated models.
	This will make it easier to get the data after database fetch.
	*/
	writeToFile(&model_df_file, fmt.Sprint(writeto, "/db/generated-queries.go"))

	/*not sure what this is needed to do*/
	// writeToFile(&query_df_file, fmt.Sprint(writeto, "/db/generated-models.go"))

	/*This generates the interfaces for each model and param/text param functions.
	This makes it easier to get the data transfered from the endpoint (database fetch) to the html template.
	*/
	writeToFile(&model_interface_file, fmt.Sprint(writeto, "/interfaces/generated.go"))
}
