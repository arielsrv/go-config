package env_test

import (
	"os"
	"strings"
	"testing"

	"github.com/arielsrv/go-config/env"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	env.Reset()
	err := env.Load()

	assert.NoError(t, err)
	assert.True(t, env.IsLocal())
}

func TestLoad_CustomConfig(t *testing.T) {
	env.Reset()
	env.SetConfigPath("config")
	env.SetConfigFile("config.yaml")
	err := env.Load()

	assert.NoError(t, err)
	assert.True(t, env.IsLocal())
}

func TestLoad_CustomConfig_Err(t *testing.T) {
	env.Reset()
	env.SetConfigPath("config")
	env.SetConfigFile("invalid.yaml")
	err := env.Load()

	assert.Error(t, err)
}

func TestLoad_Env(t *testing.T) {
	env.Reset()
	t.Setenv("ENV", "dev")

	err := env.Load()
	assert.NoError(t, err)
	assert.True(t, !env.IsLocal())
}

func TestLoad_Env_Override(t *testing.T) {
	env.Reset()
	t.Setenv("ENV", "dev")

	err := env.Load()
	assert.NoError(t, err)
	assert.True(t, !env.IsLocal())
	assert.Equal(t, "env-override", os.Getenv("app.name"))
}

func TestLoad_Msg_Override(t *testing.T) {
	env.Reset()
	t.Setenv("ENV", "dev")

	err := env.Load()
	assert.NoError(t, err)
	assert.True(t, !env.IsLocal())
	assert.Equal(t, "remote-override", os.Getenv("message"))
}

func TestFindRoot(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	actual := env.FindRoot(wd, "go.mod")
	assert.True(t, strings.HasSuffix(actual, "/go-config"))
}

func TestFindRoot_Empty(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	actual := env.FindRoot(wd, "invalid")
	assert.Empty(t, actual)
}
