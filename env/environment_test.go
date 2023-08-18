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
