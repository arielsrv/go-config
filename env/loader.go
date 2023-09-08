package env

import (
	"fmt"
	"log"
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

	configPath := filepath.Join(root, config.Path)
	log.Println(configPath)
	var compositeConfig []string

	env := GetEnv()
	scope := GetScope()

	// ../config/remote/test.config.yaml
	envConfig := filepath.Join(configPath, scope, fmt.Sprintf("%s.%s", env, config.File))
	if pathExists(envConfig) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", envConfig))
		compositeConfig = append(compositeConfig, envConfig)
	}

	// ../config/remote/config.yaml
	scopeConfig := filepath.Join(configPath, scope, config.File)
	if pathExists(scopeConfig) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	// ../config/config.yaml
	sharedConfig := filepath.Join(configPath, config.File)
	if pathExists(sharedConfig) {
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
func findRoot(wd, target string) string {
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
	return err == nil
}
