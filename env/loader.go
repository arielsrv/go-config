package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Reset only for test.
func Reset() {
	os.Clearenv()
	config = New()
}

// Load
// SCOPE LOCAL (local) or REMOTE (remote)
// ENV DEV|UAT|REMOTE.
func Load() error {
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

	propertiesPath := fmt.Sprintf("%s/%s", root, config.Path)
	var compositeConfig []string

	env := GetEnv()
	scope := GetScope()

	envConfig := fmt.Sprintf("%s/%s/%s.%s", propertiesPath, scope, env, config.File)
	if PathExists(envConfig) {
		config.Logger.Info(fmt.Sprintf("go-config: append %s ...", envConfig))
		compositeConfig = append(compositeConfig, envConfig)
	}

	scopeConfig := fmt.Sprintf("%s/%s/%s", propertiesPath, scope, config.File)
	if PathExists(scopeConfig) {
		config.Logger.Info(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	sharedConfig := fmt.Sprintf("%s/%s", propertiesPath, config.File)
	if PathExists(fmt.Sprintf("%s/%s", propertiesPath, config.File)) {
		config.Logger.Info(fmt.Sprintf("go-config: append %s ...", sharedConfig))
		compositeConfig = append(compositeConfig, sharedConfig)
	}

	err = godotenv.Load(compositeConfig...)
	if err != nil {
		return err
	}

	config.Logger.Info(fmt.Sprintf("ENV: %s, SCOPE: %s", env, scope))

	return nil
}

func PathExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("go-config: %s", err))
		return false
	}
	config.Logger.Debug(fmt.Sprintf("go-config: path %s, fileInfo: %s", path, fileInfo.Name()))
	return true
}
