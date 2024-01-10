package autoload

/*
	You can just read the config/config.yaml file on import just by doing

		import _ "gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config/autoload"

	And Bob's your mother's brother
*/

import (
	"gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config/env"
)

func init() {
	err := env.Load()
	if err != nil {
		return
	}
}
