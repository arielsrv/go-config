package env_test

import (
	"github.com/arielsrv/go-config/env"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	os.Clearenv()
	err := env.Load()

	assert.NoError(t, err)
	assert.True(t, env.IsLocal())
}

func TestLoad_CustomConfig(t *testing.T) {
	os.Clearenv()
	err := env.Load(env.Config{
		Filename: "config.yaml",
		Folder:   "config",
	})

	assert.NoError(t, err)
	assert.True(t, env.IsLocal())
}

func TestLoad_CustomConfig_Err(t *testing.T) {
	os.Clearenv()
	err := env.Load(env.Config{
		Filename: "invalid.yaml",
		Folder:   "config",
	})

	assert.Error(t, err)
}

func TestLoad_Env(t *testing.T) {
	os.Clearenv()
	t.Setenv("ENV", "dev")

	err := env.Load()
	assert.NoError(t, err)
	assert.True(t, !env.IsLocal())
}

func TestLoad_Env_Override(t *testing.T) {
	os.Clearenv()
	t.Setenv("ENV", "dev")

	err := env.Load()
	assert.NoError(t, err)
	assert.True(t, !env.IsLocal())
	assert.Equal(t, "env-override", os.Getenv("app.name"))
}

func TestLoad_Scope_Override(t *testing.T) {
	os.Clearenv()
	t.Setenv("ENV", "dev")

	err := env.Load()
	assert.NoError(t, err)
	assert.True(t, !env.IsLocal())
	assert.Equal(t, "scope-override", os.Getenv("key"))
}
