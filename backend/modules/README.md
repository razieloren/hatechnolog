# Adding a New Module

Generally, module is a pretty standalone unit, but there are some guidelines when creating a new one:
- You must of course provide your own `main.go` file under a `main` package.
- Use the `entrypoint.Entrypoint()` function to make sure all the modules use the same basic command line arguments & logger. It will panic if it cannot find the config file.
```go
var config Config
configPath, logger := entrypoint.Entrypoint(&config)
```
- Ensure that commonly used code is exported to the library (`x`), allowing all modules to utilize it :)

## Configuration Files

- For the proper functioning of a module, a configuration file is required. Each module provides documentation on its configuration structure in the corresponding folder's "README.md" file.
- The configuration parser utilizes the `YAML` format.
- By default, modules will search for their configuration file under the `./config.yaml` path.
    - A custom path can be provided with the `-c` flag.
## Building Modules

- Each module contains its own buildsystem.
- However, modules are deployed as Docker containers, thus please make sure the final output of your build system is a docker image, tagged like: `hatechnolog/modules/[MODULE]-[OS]:[MAJOR].[MINOR]`.