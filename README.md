# Configuration

```shell
$ go get github.com/arielsrv/go-config
```

```go
// config/config.yaml by default
err := env.Load()
if err != nil {
    log.Fatal(err)
}

log.Print(env.Get("message"))
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
    2023/08/15 13:47:16 INFO go-config: append /Users/ariel.pineiro/projects/iskaypetcom/go-config/config/remote/dev.config.yaml ...
    2023/08/15 13:47:16 INFO go-config: append /Users/ariel.pineiro/projects/iskaypetcom/go-config/config/remote/config.yaml ...
    2023/08/15 13:47:16 INFO go-config: append /Users/ariel.pineiro/projects/iskaypetcom/go-config/config/config.yaml ...
    2023/08/15 13:47:16 INFO ENV: dev, SCOPE: prod
    ```
