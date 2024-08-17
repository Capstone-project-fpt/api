# name app
APP_NAME = server
DB_HOST = localhost
DB_PORT = 5432
DB_NAME = postgres
SSL_MODE = disable

run:
	go run ./cmd/${APP_NAME}/

wire:
	cd internal/wire && wire

migrate_database:
	migrate -database postgres://postgres:@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE} -path migrations up

migrate_database_down:
	migrate -database postgres://postgres:@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE} -path migrations down 1

create_migration:
	migrate create -ext sql -dir migrations -seq MIGRAION_NAME