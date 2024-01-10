package env_test

import (
	"os"
	"testing"

	"github.com/arielsrv/go-config/env"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	defer t.Cleanup(func() {
		os.Clearenv()
	})

	t.Setenv("key", "value")
	actual := env.Get("key")
	assert.Equal(t, "value", actual)
}

func TestIsLocal(t *testing.T) {
	actual := env.IsLocal()
	assert.True(t, actual)
}

func TestIsRemote(t *testing.T) {
	defer t.Cleanup(func() {
		os.Clearenv()
	})

	t.Setenv("ENV", "staging")
	actual := env.IsRemote()
	assert.True(t, actual)
}

func TestIsEmptyString(t *testing.T) {
	actual := env.IsEmptyString("")
	assert.True(t, actual)

	actual = env.IsEmptyString(" ")
	assert.True(t, actual)
}

func TestGetInt(t *testing.T) {
	defer t.Cleanup(func() {
		os.Clearenv()
	})

	t.Setenv("key", "1000")
	actual := env.GetInt("key", 0)
	assert.Equal(t, 1000, actual)
}

func TestGetInt_Default(t *testing.T) {
	defer t.Cleanup(func() {
		os.Clearenv()
	})

	actual := env.GetInt("key", 1000)
	assert.Equal(t, 1000, actual)
}
