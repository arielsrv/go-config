package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var cfg = viper.New()

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
	var compositeConfig []*viper.Viper

	env := GetEnv()
	scope := GetScope()

	// ../config/config.yaml
	sharedConfig := filepath.Join(configPath, config.File)
	if pathExists(sharedConfig) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", sharedConfig))
		conf := viper.New()
		conf.SetConfigFile(sharedConfig)
		compositeConfig = append(compositeConfig, conf)
	}

	// ../config/remote/config.yaml
	scopeConfig := filepath.Join(configPath, scope, config.File)
	if pathExists(scopeConfig) {
		config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", scopeConfig))
		conf := viper.New()
		conf.SetConfigFile(scopeConfig)
		compositeConfig = append(compositeConfig, conf)
	}

	if !IsLocal() {
		// ../config/remote/test.config.yaml
		envConfig := filepath.Join(configPath, scope, fmt.Sprintf("%s.%s", env, config.File))
		if pathExists(envConfig) {
			config.Logger.Debug(fmt.Sprintf("go-config: append %s ...", envConfig))
			conf := viper.New()
			conf.SetConfigFile(envConfig)
			compositeConfig = append(compositeConfig, conf)
		}
	}

	if len(compositeConfig) == 0 {
		return errors.New("no config files found")
	}

	for i := 0; i < len(compositeConfig); i++ {
		c := compositeConfig[i]
		err = c.ReadInConfig()
		if err != nil {
			return err
		}
		err = cfg.MergeConfigMap(c.AllSettings())
		if err != nil {
			return err
		}
	}

	config.Logger.Info(fmt.Sprintf("ENV: %s, SCOPE: %s", env, scope))

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
		return StringEmpty
	}
	return findRoot(parent, target)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
