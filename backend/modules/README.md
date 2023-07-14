# Adding a New Module

Generally, module is a pretty standalone unit, but there are some guidelines when creating a new one:
- You must of course provide your own `main.go` file under a `main` package.
- Use the `entrypoint.Entrypoint()` function to make sure all the modules use the same basic command line arguments & logger. It will panic if it cannot find the config file.
```go
var config Config
configPath, logger := entrypoint.Entrypoint(&config)
```
- Ensure that commonly used code is exported to the library, allowing all modules to utilize it :)