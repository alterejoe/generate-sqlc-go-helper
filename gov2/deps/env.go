package deps

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	DbProjectPathIn string
	DbProjectUrlIn  string

	DbModulePathOut    string
	DbModuleNameOut    string
	DbModuleImportsOut []string

	DbQueryParamOut        string
	DbInterfaceAdaptersOut string
	DbInterfacesOut        string
}

func GetEnv() Env {
	fmt.Println("Loading .env file...")
	envFile := flag.String("env-file", ".env", "Path to .env file")
	// flag.Parse()

	if *envFile == "" {
		panic("missing --env-file, must provide one")
	}

	// use the flag value here
	if err := godotenv.Load(*envFile); err != nil {
		fmt.Println("Error loading .env file:", err)
		panic(err)
	}

	imports := strings.Split(os.Getenv("OUT_DB_MODULE_IMPORTS"), ",")

	cfg := Env{
		DbProjectPathIn: os.Getenv("IN_DB_PROJECT_PATH"),
		DbProjectUrlIn:  os.Getenv("IN_DB_PROJECT_URL"),

		DbModulePathOut:    os.Getenv("OUT_DB_MODULE_PATH"),
		DbModuleNameOut:    os.Getenv("OUT_DB_MODULE_NAME"),
		DbModuleImportsOut: imports,

		DbQueryParamOut:        os.Getenv("OUT_DB_QUERY_PARAM_OUT"),
		DbInterfaceAdaptersOut: os.Getenv("OUT_DB_INTERFACE_ADAPTERS_OUT"),
		DbInterfacesOut:        os.Getenv("OUT_DB_INTERFACES_OUT"),
	}
	fmt.Println("Loaded .env file")
	return cfg
}
