# Hatechnolog

Website & all other services Hatechnolog has to offer.

## Directory Tree
- `api`: RPC-messages that define the frontend-backend relations (`protobuf`).
- `backend`: Backend modules (`golang` based, scripts are in `python`).
- `frontend`: `hatechnolog.com` frontend (`NextJS`, `typescript` based).
- `setup_api_messages.sh` - Compiles all the messages from the API directory, both for the backend modules, and for the frontend.
- `deploy.py & deploylib` - Docker registry-based eployment helpers.

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
