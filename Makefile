# name app
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

wire:
	cd internal/wire && wire