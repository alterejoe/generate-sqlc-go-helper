package deps

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DbProjectPath          string
	DbProjectUrl           string
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

	cfg := Env{
		DbProjectPath:          os.Getenv("IN_DB_PROJECT_PATH"),
		DbProjectUrl:           os.Getenv("IN_DB_PROJECT_URL"),
		DbQueryParamOut:        os.Getenv("OUT_DB_QUERY_PARAM_OUT"),
		DbInterfaceAdaptersOut: os.Getenv("OUT_DB_INTERFACE_ADAPTERS_OUT"),
		DbInterfacesOut:        os.Getenv("OUT_DB_INTERFACES_OUT"),
	}
	fmt.Println("Loaded .env file")
	return cfg
}
