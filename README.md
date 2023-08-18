![coverage](https://img.shields.io/badge/Coverage-98.1%25-brightgreen)

> This package provides a config files based on archaius netflix with some features

# Configuration

```shell
$ go get github.com/arielsrv/go-config
```

example

```go
package main

import (
	"fmt"
	"github.com/arielsrv/go-config/env"
	"log"
	"os"
)

func init() {
	err := os.Setenv("ENV", "dev")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// config/config.yaml by default
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("VAR: %s", env.Get("message"))
	log.Printf("ENV: %s", env.GetEnv())
	log.Printf("SCOPE: %s", env.GetScope())
}
```

Environment configuration is based on **Archaius Config**, you should use a similar folder
structure.
*SCOPE* env variable in remote environment is required

```
└── config
    ├── config.yml (shared config)
    └── local
        └── config.yml (for local development)
    └── remote (for remote environment)
        └── config.yml (base config)
        └── {environment}.config.yml (base config)
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
        └── config.yml                          <ignored>
    └── remote
        └── config.yml (base config)            2nd (second)
        └── test.config.yml (base config)       1st (first)
    ```

    * 1st (first)   remote/test.config.yml
    * 2nd (second)  remote/config.yml
    * 3th (third)   config.yml
```
```
DEBUG go-config: append /go-config/config/remote/dev.config.yaml ... 
DEBUG go-config: append /go-config/config/remote/config.yaml ... 
DEBUG go-config: append /go-config/config/config.yaml ... 
DEBUG ENV: dev, SCOPE: remote 
INFO  VAR: remote-override 
INFO  ENV: dev 
INFO  SCOPE: remote
```
