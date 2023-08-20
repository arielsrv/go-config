package env

import (
	"log/slog"
	"os"
)

type Config struct {
	File   string
	Path   string
	Logger *slog.Logger
}

var config *Config

func init() {
	config = New()
}

func New() *Config {
	c := new(Config)
	c.File = "config.yaml"
	c.Path = "config"
	c.Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	return c
}

func SetConfigPath(configPath string) {
	if !IsEmptyString(configPath) {
		config.Path = configPath
	}
}

func SetConfigFile(configFile string) {
	if !IsEmptyString(configFile) {
		config.File = configFile
	}
}

func SetLogger(logger *slog.Logger) {
	if logger != nil {
		config.Logger = logger
	}
}
