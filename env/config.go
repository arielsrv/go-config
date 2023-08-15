package env

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Filename string
	Folder   string
}

// Load
// SCOPE LOCAL (local) or REMOTE (remote)
// ENV DEV|UAT|REMOTE.
func Load(config ...Config) error {
	environment := GetEnv()
	scope := GetScope()
	settings := &Config{
		Filename: "config.yaml",
		Folder:   "config",
	}

	if len(config) > 0 {
		settings = &config[0]
	}

	root, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		if isProjectRoot(root) {
			break
		}
		root = filepath.Dir(root)
	}

	propertiesPath := fmt.Sprintf("%s/%s", root, settings.Folder)
	var compositeConfig []string

	envConfig := fmt.Sprintf("%s/%s/%s.%s", propertiesPath, scope, environment, settings.Filename)
	if found(envConfig) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", envConfig))
		compositeConfig = append(compositeConfig, envConfig)
	}

	scopeConfig := fmt.Sprintf("%s/%s/%s", propertiesPath, scope, settings.Filename)
	if found(scopeConfig) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	sharedConfig := fmt.Sprintf("%s/%s", propertiesPath, settings.Filename)
	if found(fmt.Sprintf("%s/%s", propertiesPath, settings.Filename)) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", sharedConfig))
		compositeConfig = append(compositeConfig, sharedConfig)
	}

	err = godotenv.Load(compositeConfig...)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("ENV: %s, SCOPE: %s", environment, scope))

	return nil
}

func isProjectRoot(dir string) bool {
	return found(filepath.Join(dir, "go.mod"))
}

func found(path string) bool {
	_, err := os.Stat(path) // ignore
	return err == nil
}
