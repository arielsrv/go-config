package main

import (
	"fmt"
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

	log.Println(fmt.Sprintf("VAR: %s", env.Get("message")))
	log.Println(fmt.Sprintf("ENV: %s", env.GetEnv()))
	log.Println(fmt.Sprintf("SCOPE: %s", env.GetScope()))
}
