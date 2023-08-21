package main

import (
	"log"
	"os"

	"github.com/arielsrv/go-config/env"

	_ "github.com/arielsrv/go-config/autoload"
)

func main() {
	log.Printf("VAR: %s", os.Getenv("message"))
	log.Printf("VAR: %s", env.Get("message"))
}
