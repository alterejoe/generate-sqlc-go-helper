package main

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/deps"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/generators"
	"github.com/alterejoe/generate/sqlc-go-helper/gov2/inspectors"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func main() {
	gd := deps.GetDeps()

	var finaloutput Output
	_ = filepath.WalkDir(gd.Environment.DbProjectPathIn, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		output := delegator(path, gd)

		finaloutput.Dbqueryparam_out = append(finaloutput.Dbqueryparam_out, output.Dbqueryparam_out...)
		finaloutput.Dbinterfaceadapters_out = append(finaloutput.Dbinterfaceadapters_out, output.Dbinterfaceadapters_out...)
		finaloutput.DbInterfaces_out = append(finaloutput.DbInterfaces_out, output.DbInterfaces_out...)

		return nil
	})

	imports := []string{
		"context",
		gd.Environment.DbModulePathOut,
	}
	imports = append(imports, gd.Environment.DbModuleImportsOut...)

	wf_dbqueryparam_out := WriteFile{
		Interface:   "sqlcqueries",
		Imports:     imports,
		FileName:    gd.Environment.DbQueryParamOut,
		PackageName: gd.Environment.DbQueryParamOut,
		Out:         finaloutput.Dbqueryparam_out,
	}
	// wf_dbinterfaceadapters_out := WriteFile{
	// 	Imports:     []string{"database/sql"},
	// 	FileName:    gd.Environment.DbInterfaceAdaptersOut,
	// 	PackageName: gd.Environment.DbInterfaceAdaptersOut,
	// 	Out:         finaloutput.Dbinterfaceadapters_out,
	// }
	// wf_displayinterfaces_out := WriteFile{
	// 	Imports:     []string{"database/sql"},
	// 	FileName:    gd.Environment.DbInterfacesOut,
	// 	PackageName: gd.Environment.DbInterfacesOut,
	// 	Out:         finaloutput.DbInterfaces_out,
	// }

	gd.Logger.Info("Writing files...")
	if len(wf_dbqueryparam_out.Out) > 0 {
		writeToFile(wf_dbqueryparam_out)
	} else {
		gd.Logger.Info("No files to write")
	}
	// if len(wf_dbinterfaceadapters_out.Out) > 0 {
	// 	writeToFile(wf_dbinterfaceadapters_out)
	// }
	// if len(wf_displayinterfaces_out.Out) > 0 {
	// 	writeToFile(wf_displayinterfaces_out)
	// }
}

func Inspector(node dst.Node, deps deps.Deps, output func(inspectors.StructParams, deps.Deps) []dst.Decl) []dst.Decl {
	return nil
}

func Generator(p inspectors.StructParams, deps deps.Deps) []dst.Decl {
	return nil
}

type Output struct {
	Dbqueryparam_out        []dst.Decl
	Dbinterfaceadapters_out []dst.Decl
	DbInterfaces_out        []dst.Decl
}

func delegator(path string, deps deps.Deps) *Output {
	var dbqueryparam_out []dst.Decl
	var dbinterfaceadapters_out []dst.Decl
	var dbinterfaces_out []dst.Decl

	if strings.HasSuffix(path, ".go") {
		if strings.EqualFold(path, "db.go") {
			return &Output{}
		}
		dbinterfaces_out = append(dbinterfaces_out, dst_inspect_file[inspectors.StructParams](path, deps, generators.DbInterfaces, inspectors.Struct)...)
	}
	if strings.HasSuffix(path, ".sql.go") {
		dbqueryparam_out = append(dbqueryparam_out, dst_inspect_file[inspectors.FuncParams](path, deps, generators.DbQueryParams_PointerFunctions, inspectors.Funcs)...)
		// dbqueryparam_out = append(dbqueryparam_out, dst_inspect_file[inspectors.StructParams](path, deps, generators.DbQueryParams_PointerFunctions, inspectors.Funcs)...)
		// dbinterfaceadapters_out = append(dbinterfaceadapters_out, dst_inspect_file(path, deps, generators.DbInterfaceAdapters, inspectors.Struct)...)
	}

	return &Output{
		Dbqueryparam_out:        dbqueryparam_out,
		Dbinterfaceadapters_out: dbinterfaceadapters_out,
		DbInterfaces_out:        dbinterfaces_out,
	}
}

//	func dst_inspect_file(path string, d deps.Deps, output func(inspectors.StructParams, deps.Deps) []dst.Decl, inspector func(dst.Node, deps.Deps, func(inspectors.StructParams, deps.Deps) []dst.Decl) []dst.Decl) []dst.Decl {
//		fset := token.NewFileSet()
//		f, err := decorator.ParseFile(fset, path, nil, parser.AllErrors)
//		if err != nil {
//			panic(err)
//		}
//
//		// basically this will iterate over every node in the file and the Inspector
//		// will parse each individual node
//		var decls []dst.Decl
//		dst.Inspect(f, func(n dst.Node) bool {
//			d := inspector(n, d, output)
//			decls = append(decls, d...)
//			return true
//		})
//
//		return decls
//	}
func dst_inspect_file[T any](path string, d deps.Deps, output func(T, deps.Deps) []dst.Decl, inspector func(dst.Node, deps.Deps, func(T, deps.Deps) []dst.Decl) []dst.Decl) []dst.Decl {
	fset := token.NewFileSet()
	f, err := decorator.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	var decls []dst.Decl
	dst.Inspect(f, func(n dst.Node) bool {
		d := inspector(n, d, output)
		decls = append(decls, d...)
		return true
	})
	return decls
}
