# Module: `api`

This module is the backend "face" to the frontend (the website). It grants access to various endpoints that can be interacted with through a websocket, utilizing the predefined messages API.

## Configuration File

```yaml
config:
  server:
    listenAddr: STRING
    listenPort: INTEGER
    key_path: STRING | null
    fullchain_path: STRING | null
  identity:
    salt: STRING
    secret: STRING
  api_token:
    client: STRING
    server: STRING
  auth:
    redirect_host: STRING
    session_cookies:
      session:
        name: STRING
        expiry_time_sec: INTEGER
        path: STRING
        domain: STRING
        secure: BOOL
        http_only: BOOL
      session_flag:
        name: STRING
        expiry_time_sec: INTEGER
        path: STRING
        domain: STRING
        secure: BOOL
        http_only: BOOL
    oauth2:
      state_cookie:
        name: STRING
        expiry_time_sec: INTEGER
        path: STRING
        domain: STRING
        secure: BOOL
        http_only: BOOL
      config:
        discord:
          redirect_url: STRING
          client_id: STRING
          client_secret: STRING
          scopes:
            - STRING
        github:
          redirect_url: STRING
          client_id: STRING
          client_secret: STRING
          scopes:
            - STRING
      hooks:
        discord:
          guild_id: STRING
          supporter_role_id: STRING
          vip_role_id: STRING
  database:
    host: STRING
    port: INTEGER
    user: STRING
    password: STRING
    database: STRING
    sslmode: STRING

```
