package main

import (
	"github.com/arielsrv/go-config/env"
	"log"
	"os"
)

func init() {
	err := os.Setenv("ENV", "dev")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// config/config.yaml by default
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(env.Get("message"))
}
