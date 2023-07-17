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
- However, modules are deployed as Docker containers, thus pleaase make sure the final output of your build system is a docker image, tagged like: `hatechnolog/modules/[MODULE]-[OS]:[MAJOR].[MINOR]`.

## Deploying Modules

- Current deploy system is designed to work with Docker Registry & SSH server.
- It is assumed that you have a docker registry available to push the images, and a SSH server to deply them.
- Use the `deploy.py` script next to this file.
- The configuration file (`deploy.yaml`) indicates what exactly to deploy, this might be a format of one:
```yaml
config:
  registry:
    url: STRING
    name: STRING
    user: STRING
    password: STRING
  modules:
    query:
      metadata:
        majorVersion: INTEGER
        minorVersion: INTERGER
      ssh:
        host: STRING
        port: INTEGER
        user: STRING
        password: STRING | null
        identityFile: STRING | null
        identityFilePassword: STRING | null
```
- This config will deploy `query` of the specified version to the specified SSH server, by pushing the relevant Docker registry.
- To add more modules, simply add another section under `modules`.
- Each deploy process creates a temporary tag to the selected inage in order to push it to the remote Docker registry. This tag should be removed automatically.