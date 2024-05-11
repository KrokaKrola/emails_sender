package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Environment struct {
	env string
}

func NewEnvironment() *Environment {
	return &Environment{}
}

func (e *Environment) Load() error {
	err := e.getAppEnv()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error getting current working directory %v", err)
	}

	envPath := filepath.Join(cwd, fmt.Sprintf(".env.%s", e.env))

	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return err
}

func (e *Environment) getAppEnv() error {
	env := flag.String("app_env", "", "app_env variable")

	flag.Parse()

	if *env == "" {
		return fmt.Errorf("app_env flag is not specified")
	}

	e.env = *env

	return nil
}
