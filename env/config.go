package env

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Filename    string
	Folder      string
	Environment string
	Scope       string
}

func New() *Config {
	c := new(Config)
	c.Filename = "config.yaml"
	c.Folder = "config"

	return c
}

var config *Config

func init() {
	config = New()
}

func SetConfigPath(folder string) {
	if !IsEmptyString(folder) {
		config.Folder = folder
	}
}

func SetConfigFile(filename string) {
	if !IsEmptyString(filename) {
		config.Filename = filename
	}
}

func Reset() {
	os.Clearenv()
	config = New()
}

// Load
// SCOPE LOCAL (local) or REMOTE (remote)
// ENV DEV|UAT|REMOTE.
func Load() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	root := findRoot(wd)

	propertiesPath := fmt.Sprintf("%s/%s", root, config.Folder)
	var compositeConfig []string

	env := GetEnv()
	scope := GetScope()

	envConfig := fmt.Sprintf("%s/%s/%s.%s", propertiesPath, scope, env, config.Filename)
	if PathExists(envConfig) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", envConfig))
		compositeConfig = append(compositeConfig, envConfig)
	}

	scopeConfig := fmt.Sprintf("%s/%s/%s", propertiesPath, scope, config.Filename)
	if PathExists(scopeConfig) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	sharedConfig := fmt.Sprintf("%s/%s", propertiesPath, config.Filename)
	if PathExists(fmt.Sprintf("%s/%s", propertiesPath, config.Filename)) {
		slog.Info(fmt.Sprintf("go-config: append %s ...", sharedConfig))
		compositeConfig = append(compositeConfig, sharedConfig)
	}

	err = godotenv.Load(compositeConfig...)
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("ENV: %s, SCOPE: %s", env, scope))

	return nil
}

func findRoot(path string) string {
	if PathExists(filepath.Join(path, "go.mod")) {
		return path
	}

	parent := filepath.Dir(path)
	return findRoot(parent)
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
