package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDir = "golang_project"

// InitializeEnvironment selects and loads an environment file using the "GO_APP_ENV" environment key.
// If it fails program exits with -1 status.
func InitializeEnvironment() {
	env, exists := os.LookupEnv("GO_APP_ENV")
	if !exists {
		log.Fatal("environment variable is not set")
		os.Exit(-1)
	}

	re := regexp.MustCompile(`^(.*` + projectDir + `)`)
	cwd, _ := os.Getwd()
	rootPath := string(re.Find([]byte(cwd)))

	if env == "prod" {
		loadEnv(rootPath, ".env")
	} else {
		loadEnv(rootPath, ".env."+env)
	}
}

// loadEnv loads the speficied environment file.
// If it fails program exits with -1 status.
func loadEnv(rootPath string, env string) {
	err := godotenv.Load(rootPath + "/" + env)
	if err != nil {
		log.Fatal("environment variable is not set" + env)
		os.Exit(-1)
	}
}
