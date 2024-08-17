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

```bash
 (for local env)
$ make migrate_database
```

### 5.1 create migration file

```bash
$ make create_migration MIGRATION_NAME="name migration"
```

### 6. Running the app

```bash
# development
$ go run cmd/main.go
```

Navigate to your [host](http://localhost:8000) to check the server is online.