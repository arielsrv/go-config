package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/arielsrv/go-config/env"
)

func main() {
	// config/config.yaml by default
	// native text logger | warn level by default
	env.SetConfigPath("config")
	env.SetConfigFile("config.yaml")
	env.SetLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("VAR: %s", env.Get("message"))
	log.Printf("MY-VAR1: %s", env.Get("my-var1"))
	log.Printf("MY_VAR2: %s", env.Get("MY_VAR2"))
	log.Printf("ENV: %s", env.GetEnv())
	log.Printf("SCOPE: %s", env.GetScope())
	log.Printf("NOT FOUND: %s", env.Get("MISSING"))

	for env.IsRemote() {
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}
