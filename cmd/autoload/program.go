package main

import (
	"log"
	"os"

	_ "gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config/autoload"
	"gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config/env"
)

func main() {
	log.Printf("VAR: %s", os.Getenv("message"))
	log.Printf("VAR: %s", env.Get("message"))
}
