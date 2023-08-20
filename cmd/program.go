package main

import (
	"log"
	"os"

	"github.com/arielsrv/go-config/env"
)

func init() {
	err := os.Setenv("ENV", "dev")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// config/config.yaml by default
	env.SetConfigPath("config")
	env.SetConfigFile("config.yaml")
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("VAR: %s", os.Getenv("message"))
	log.Printf("ENV: %s", env.GetEnv())
	log.Printf("SCOPE: %s", env.GetScope())
	log.Printf("NOT FOUND: %s", env.Get("MISSING"))
}
