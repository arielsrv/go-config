package env_test

import (
	"github.com/arielsrv/go-config/env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	defer t.Cleanup(func() {
		os.Clearenv()
	})

	t.Setenv("key", "value")
	actual := env.Get("key")
	assert.Equal(t, "value", actual)
}
