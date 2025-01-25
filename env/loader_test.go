package env

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	Reset()
	err := Load()

	require.NoError(t, err)
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

	require.NoError(t, err)
	assert.True(t, IsLocal())
}

func TestLoad_CustomConfig_By_Param(t *testing.T) {
	Reset()
	SetConfig(&Config{
		File: "config.yaml",
		Path: "config",
	})
	err := Load()

	require.NoError(t, err)
	assert.True(t, IsLocal())
}

func TestLoad_CustomConfig_Err(t *testing.T) {
	Reset()
	SetConfigPath("config")
	SetConfigFile("invalid.yaml")
	err := Load()

	require.Error(t, err)
}

func TestLoad_Env(t *testing.T) {
	Reset()
	t.Setenv("ENV", "dev")

	err := Load()
	require.NoError(t, err)
	assert.False(t, IsLocal())
}

func TestFindRoot(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	actual := findRoot(wd, "go.mod")
	if !strings.HasSuffix(actual, fmt.Sprintf("%c%s", os.PathSeparator, "go-config")) {
		t.Logf("%s", fmt.Sprintf("go-config: findRoot, go.mod not found  %s", actual))
		t.Fail()
	}
}

func TestFindRoot_Empty(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	actual := findRoot(wd, "invalid")
	assert.Empty(t, actual)
}
