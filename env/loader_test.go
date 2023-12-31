package env

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	Reset()
	err := Load()

	assert.NoError(t, err)
	assert.True(t, IsLocal())
}

func TestLoad_CustomConfig(t *testing.T) {
	Reset()
	SetConfigPath("config")
	SetConfigFile("config.yaml")
	SetLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))
	err := Load()

	assert.NoError(t, err)
	assert.True(t, IsLocal())
}

func TestLoad_CustomConfig_By_Param(t *testing.T) {
	Reset()
	SetConfig(&Config{
		File: "config.yaml",
		Path: "config",
	})
	err := Load()

	assert.NoError(t, err)
	assert.True(t, IsLocal())
}

func TestLoad_CustomConfig_Err(t *testing.T) {
	Reset()
	SetConfigPath("config")
	SetConfigFile("invalid.yaml")
	err := Load()

	assert.Error(t, err)
}

func TestLoad_Env(t *testing.T) {
	Reset()
	t.Setenv("ENV", "dev")

	err := Load()
	assert.NoError(t, err)
	assert.True(t, !IsLocal())
}

func TestLoad_Env_Override(t *testing.T) {
	Reset()
	t.Setenv("ENV", "dev")

	err := Load()
	assert.NoError(t, err)
	assert.True(t, !IsLocal())
	assert.Equal(t, "env-override", os.Getenv("app.name"))
}

func TestLoad_Msg_Override(t *testing.T) {
	Reset()
	t.Setenv("ENV", "dev")

	err := Load()
	assert.NoError(t, err)
	assert.True(t, !IsLocal())
	assert.Equal(t, "remote-override", os.Getenv("message"))
}

func TestFindRoot(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	actual := findRoot(wd, "go.mod")
	if !strings.HasSuffix(actual, fmt.Sprintf("%c%s", os.PathSeparator, "go-config")) {
		t.Logf(fmt.Sprintf("go-config: findRoot, go.mod not found  %s", actual))
		t.Fail()
	}
}

func TestFindRoot_Empty(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	actual := findRoot(wd, "invalid")
	assert.Empty(t, actual)
}
