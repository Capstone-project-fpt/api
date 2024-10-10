# Capstone Backend

### 1. Prerequisites:

- Golang

### 2. Install dependencies:

```bash
$ go mod download
```

### 3. Create file `config.yaml` from `config.yaml.example`

```bash
  cp config.yaml.example config.yaml
```

### 4. Database Migration Guide

Config DB_HOST, DB_HOST, DB_NAME, DB_PASSWORD, DB_USERNAME and SSL_MODE before run migration

```bash
// Up
$ make migrate_up

// Down
$ make migrate_down
```

### 5.1 create migration file

```bash
$ make create_migration MIGRATION_NAME="NameOfMigration"
```

### 6. Generate Swagger docs

```bash
$ make swagger
```

### 7. Running the app

```bash
# development
$ make run
```

#### Optional: Live reload for Go apps in development

1. Install [air](https://github.com/air-verse/air)

If air not command found fix with
   Export path in `~/.zshrc`:

  ```bash
  export PATH=$PATH:$(go env GOPATH)/bin
  ```

   Update path environment:

  ```bash
  source ~/.zshrc
  ```

2. Run with bash:

  ```bash
  air
  ```

  Happy coding!


