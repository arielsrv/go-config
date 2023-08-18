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
	currentEnv := GetEnv()
	currentScope := GetScope()
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
		if PathExists(filepath.Join(root, "go.mod")) {
			break
		}
		root = filepath.Dir(root)
	}

	propertiesPath := fmt.Sprintf("%s/%s", root, settings.Folder)
	var compositeConfig []string

	envConfig := fmt.Sprintf("%s/%s/%s.%s", propertiesPath, currentScope, currentEnv, settings.Filename)
	if PathExists(envConfig) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", envConfig))
		compositeConfig = append(compositeConfig, envConfig)
	}

	scopeConfig := fmt.Sprintf("%s/%s/%s", propertiesPath, currentScope, settings.Filename)
	if PathExists(scopeConfig) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	sharedConfig := fmt.Sprintf("%s/%s", propertiesPath, settings.Filename)
	if PathExists(fmt.Sprintf("%s/%s", propertiesPath, settings.Filename)) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", sharedConfig))
		compositeConfig = append(compositeConfig, sharedConfig)
	}

	err = godotenv.Load(compositeConfig...)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("ENV: %s, SCOPE: %s", currentEnv, currentScope))

	return nil
}

func PathExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		slog.Error(fmt.Sprintf("go-config: %s", err))
		return false
	}
	slog.Debug(fmt.Sprintf("go-config: path %s, fileInfo: %s", path, fileInfo.Name()))
	return true
}
