# Configuration

```shell
go get github.com/arielsrv/go-config
```

example

```go
package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/arielsrv/go-config/env"
)

func init() {
	err := os.Setenv("ENV", "dev")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// config/config.yaml by default
	// native text logger | warn level by default
	env.SetConfigPath("config")
	env.SetConfigFile("config.yaml")
	env.SetLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("VAR: %s", os.Getenv("message"))
	log.Printf("ENV: %s", env.GetEnv())
	log.Printf("SCOPE: %s", env.GetScope())
	log.Printf("NOT FOUND: %s", env.Get("MISSING"))
}
```

Environment configuration is based on **Archaius Config**, you should use a similar folder
structure.
*SCOPE* env variable in remote environment is required

```
├── config
│	├── config.yaml 				(shared config)
│	├── local
│	│	└── config.yaml             (for local development)
│	└── remote                      (for remote environment)
│		├── config.yaml             (base shared config by env)
│		├── dev.config.yaml
│		├── prod.config.yaml
│		└── {env}.config.yaml
```

The SDK provides a simple configuration hierarchy

* config/config.properties (shared config)
* config/{scope}/config.properties (override shared config by environment)
* config/{scope}/{env}.config.properties (override env and shared config by scope)

Example *test.pets-api.internal.com*

```
└── config
    ├── config.yml                              3th (third)
    └── local
        └── config.yml                          ignored
    └── remote
        └── config.yml (base config)            2nd (second)
        └── test.config.yml (base config)       1st (first)
    ```

    * 1st (first)   remote/test.config.yml
    * 2nd (second)  remote/config.yml
    * 3th (third)   config.yml
```
```
INFO go-config: append ../my-app/config/remote/test.config.yaml ...
INFO go-config: append ../my-app/config/remote/config.yaml ...
INFO go-config: append ../my-app/config/config.yaml ...
INFO go-config: ENV: test, SCOPE: remote
ENV: test
SCOPE: remote
WARN go-config: config with name SOME.CONFIG not found
```
	
