# name app
APP_NAME = server
# DB Local
# DB_HOST = localhost
# DB_PORT = 5433
# DB_NAME = postgres
# DB_PASSWORD = postgres
# DB_USERNAME = postgres
# SSL_MODE = disable

# DB Dev
DB_HOST = aws-0-ap-southeast-1.pooler.supabase.com
DB_PORT = 6543
DB_NAME = postgres
DB_PASSWORD = postgres123!
DB_USERNAME = postgres.igpctwhkfhikvrlgmkuu
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
	
swagger:
	swag init -d ./cmd/server,./internal/controller,./internal/dto,./pkg/response