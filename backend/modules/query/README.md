# Module: `query`

This module has the responsibility of retrieving pertinent data from different services, which is then displayed on the "Hatechnolog" platforms.

## Configuration File

```yaml
config:
  main:
    queryIntervalSec: INTEGER
    database:
      host: STRING
      port: INTEGER
      user: STRING
      password: STRING
      database: STRING
      sslmode: STRING
  discord:
    botToken: STRING
    targetGuilds:
    - name: STRING
      newMemberPeriodDays: INTEGER
      avgJoinTimePeriodDays: INTEGER
  youtube:
    apiKey: STRING
    targetChannels:
    - name: STRING
```

## Adding a New Service Checklist

- Create a folder under `services` with your service names.
- Create the following files:
    - `config.go` - The services config, then add it to the main config under `main.go`. Also, update your local `config.yaml` file, and the config documentation in this file.
    - `consts.go` - General service consts.
    - `models.go` - All the GORM models you will be using in this service. Make sure to expose an `AutoMigrate` function so these models will actually get created.
    - `worker.go` - This is the code which will run periodically. Please use the same function names as declared on the other modules :)
    - `[SERVICE].go` - Actual service definition.
- In `main.go`:
    - Add your service's `AutoMigrate` function to the main `AutoMigrate` function.
    - Add your worker & config to the workers definition list.
    - Again, make sure you updated the main config with your config.
