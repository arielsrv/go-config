package main

import (
	"log"
	"os"

	_ "github.com/arielsrv/go-config/autoload"
	"github.com/arielsrv/go-config/env"
)

func main() {
	log.Printf("VAR: %s", os.Getenv("message"))
	log.Printf("VAR: %s", env.Get("message"))
}
