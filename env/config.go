package env

import (
	"log/slog"
	"os"
)

type Config struct {
	Logger *slog.Logger
	File   string
	Path   string
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

// SetConfigPath "config" from root by default.
func SetConfigPath(configPath string) {
	if !IsEmptyString(configPath) {
		config.Path = configPath
	}
}

// SetConfigFile "config.yaml" by default.
func SetConfigFile(configFile string) {
	if !IsEmptyString(configFile) {
		config.File = configFile
	}
}

// SetLogger text warn logger by default.
func SetLogger(logger *slog.Logger) {
	if logger != nil {
		config.Logger = logger
	}
}

// SetConfig set custom config.
func SetConfig(cfg *Config) {
	if cfg != nil {
		SetConfigFile(cfg.File)
		SetConfigPath(cfg.Path)
	}
}
