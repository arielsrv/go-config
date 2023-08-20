package env

import (
	"fmt"
	"os"
	"strings"
)

type Scope int

const (
	LOCAL Scope = iota
	REMOTE
)

func (s Scope) String() string {
	return []string{
		"local",
		"remote",
	}[s]
}

// GetEnv
// Get env variable from the System.
// Example for test.golang-template.internal.com is test.
func GetEnv() string {
	env := os.Getenv("ENV")
	return strings.ToLower(env)
}

// GetScope
// * Get scope name from the System.
// Priority order is as follows:
// * 1. It looks in "app.scope" system property.
// * 2. If empty, it looks in ENV system env variable
// *	2.1 If empty, it is a local scope
// *	2.2 If not empty and starts with "test," it is a test scope
// *	2.3 Otherwise, it is a "prod" environment.
func GetScope() string {
	scope := os.Getenv("app.scope")
	if !IsEmptyString(scope) {
		return scope
	}

	env := os.Getenv("ENV")
	if IsEmptyString(env) {
		return LOCAL.String()
	}

	return REMOTE.String()
}

func IsEmptyString(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func IsLocal() bool {
	return LOCAL.String() == GetScope()
}

func IsRemote() bool {
	return !IsLocal()
}

func Get(key string) string {
	value := os.Getenv(key)
	if IsEmptyString(value) {
		config.Logger.Warn(fmt.Sprintf("go-config: config with name %s not found", key))
	}
	return value
}
