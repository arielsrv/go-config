package env

import (
	"fmt"
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
	if !strings.HasSuffix(actual, "/go-config") {
		t.Logf(fmt.Sprintf("not root found %s", actual))
		t.Fail()
	}
}

func TestFindRoot_Empty(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	actual := findRoot(wd, "invalid")
	assert.Empty(t, actual)
}
