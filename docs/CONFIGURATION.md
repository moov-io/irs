# IRS
**[Purpose](README.md)** | **Configuration** | **[Running](RUNNING.md)**

---

## Configuration
Custom configuration for this application may be specified via an environment variable `APP_CONFIG` to a configuration file that will be merged with the default configuration file.

- [Default Configuration](../configs/config.default.yml)
- [Config Source Code](../pkg/service/model_config.go)
- Full Configuration
  ```yaml
  IRS:

    # Service configurations
    Servers:

      # Public service configuration
      Public:
        Bind:
          # Address and port to listen on.
          Address: ":8200"

      # Health/Admin service configuration.
      Admin:
        Bind:
          # Address and port to listen on.
          Address: ":8201"

    # All database configuration is done here. Only one connector can be configured.
    Database:

      # Database name to use for selected connector.
      DatabaseName: "identity"

      # MySql configuration
      MySQL:
        Address: tcp(mysqlidentity:3306)
        User: identity
        Password: identity

      # OR uses the sqllite db
      SQLLite:
        Path: ":memory:"
  ```

---
**[Next - Running](RUNNING.md)**
