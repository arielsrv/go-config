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
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	root := findRoot(wd, "go.mod")

	configPath := fmt.Sprintf("%s/%s", root, config.Path)
	var compositeConfig []string

	env := GetEnv()
	scope := GetScope()

	// ../config/remote/test.config.yaml
	envConfig := fmt.Sprintf("%s/%s/%s.%s", configPath, scope, env, config.File)
	if pathExists(envConfig) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", envConfig))
		compositeConfig = append(compositeConfig, envConfig)
	}

	// ../config/remote/config.yaml
	scopeConfig := fmt.Sprintf("%s/%s/%s", configPath, scope, config.File)
	if pathExists(scopeConfig) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	// ../config/config.yaml
	sharedConfig := fmt.Sprintf("%s/%s", configPath, config.File)
	if pathExists(fmt.Sprintf("%s/%s", configPath, config.File)) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", sharedConfig))
		compositeConfig = append(compositeConfig, sharedConfig)
	}

	err = godotenv.Load(compositeConfig...)
	if err != nil {
		return err
	}

	config.Logger.Debug(fmt.Sprintf("ENV: %s, SCOPE: %s", env, scope))

	return nil
}

// findRoot Find go.mod recursively.
func findRoot(wd string, target string) string {
	if pathExists(filepath.Join(wd, target)) {
		return wd
	}

	parent := filepath.Dir(wd)
	if parent == wd {
		// Reached the filesystem root without finding "go.mod"
		return ""
	}
	return findRoot(parent, target)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("go-config: %s", err))
		return false
	}
	return true
}
