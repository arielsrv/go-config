package autoload

/*
	You can just read the config/config.yaml file on import just by doing

		import _ "github.com/arielsrv/go-config/autoload"

	And Bob's your mother's brother
*/

import (
	"github.com/arielsrv/go-config/env"
)

func init() {
	err := env.Load()
	if err != nil {
		return
	}
}
