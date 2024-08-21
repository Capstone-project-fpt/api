# name app
APP_NAME = server
DB_HOST = localhost
DB_PORT = 5433
DB_NAME = postgres
DB_PASSWORD = postgres
DB_USERNAME = postgres
SSL_MODE = disable

run:
	go run ./cmd/${APP_NAME}/

wire:
	cd internal/wire && wire

migrate_up:
	migrate -database postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE} -path database/migrations up

migrate_down:
	migrate -database postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE} -path database/migrations down 1

create_migration:
	migrate create -ext sql -dir database/migrations -seq $(MIGRATION_NAME)

sqlc:
	sqlc generate