# Module: `api`

This module is the backend "face" to the frontend (the website). It grants access to various endpoints that can be interacted with through a websocket, utilizing the predefined messages API.

## Configuration File

```yaml
config:
  main:
    server:
      listenAddr: STRING
      listenPort: INT
      tls:
        key_path: STRING | null
        certificate_path: STRING | null
    database:
      host: STRING
      port: INT
      user: STRING
      password: STRING
      database: STRING
      sslmode: STRING
  endpoints:
    stats:
      latest_stats_push_interval_sec: INT
```

## Adding a New Endpoint Checklist

- Make sure to add all the proper messages under the root `api` directory, and also adding a new endpoint field in the endpoints intent enum.
- Create a folder under `endpoint` with your endpoint names.
- Write your code (preferably under files named: `endpoint.go` for main API functions, `handlers.go` for message request-response, and `helpers.go` for util functions).
- Register the handler in `router.go` with your new endpoint intent.
- If this endpoint needs a config, you can also edit the router's config struct. Make sure to update the documentation accordingly.

## Building

- Inspect the `Makefile`.
- The build process is made of 2 main stages:
    - Creating the `query` binary.
    - Building a local Docker image that runs `query`.
- The build system expects a `config.yaml` to be found in the root directory of the module (next to this file).
- The buils system creates 2 binaries & 2 docker images, one for `darwin-arm64` (Used for local tests) and one for `linux-amd64` (Used for operational purposes).
- You may build only for one platform using `make *-image`.
- Anyhow, in order to test the module, please use the docker container and not the binary directly since this is how it is going to be executed operationally.
- In order to clean some artficats, you may user: `make clean-binaries`, `make clean-images`, and `make clean` - Which is cleaning everything.
