# Module: `api`

This module is the backend of Hatechnolog website. It grants access to various endpoints, utilizing the predefined messages API (Based on Protobuf).

## Setting Up

The API module works with a database. This database has to have some predefined data initialzed in order for everything to work.

The initialization code can be found in the `initdb` sub-module (it's a standalone module, because we don't want it to be part of the API final binary).

This module expects a "models" package with the relevant predefined data. For instance, in order to create predefined roles, you can do something like this:

```go
memberRole := models.Role{
  Name:        "member",
  Description: "Community Member",
  Permissions: []models.Permission{
    karmaReadPermission,
    postReactPermission,
    commentReactPermission,
  },
}
editorRole := models.Role{
  Name:        "editor",
  Description: "Community Editor",
  Permissions: []models.Permission{
    postCreatePermission,
    postEditPermission,
  },
}
adminRole := models.Role{
  Name:        "admin",
  Description: "Community Admin",
  Permissions: []models.Permission{
    postDeletePermission,
    commentEditPermission,
    commentDeletePermission,
  },
}
roles := []*models.Role{
  &memberRole,
  &editorRole,
  &adminRole,
}
for _, role := range roles {
  if err := tx.FirstOrCreate(role, &models.Role{Name: role.Name}).Error; err != nil {
    return fmt.Errorf("default roles: %w", err)
  }
}
```

## Configuration File

```yaml
config:
  server:
    listenAddr: STRING
    listenPort: STRING
    key_path: STRING
    fullchain_path: STRING
    allowed_origin: STRING
  identity:
    salt: STRING
    secret: STRING
  api:
    client_token: STRING
    server_token: STRING
  auth:
    redirect_param: STRING
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
            - STRING
        github:
          redirect_url: STRING
          client_id: STRING
          client_secret: STRING
          scopes:
            - STRING
            - STRING
  consts:
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
