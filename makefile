DB_CONNECTION_STRING=postgres://postgres:example@localhost:5432/postgres?sslmode=disable

services:
	docker compose up -d

services-down:
	docker compose down

dev:
	air

migrate:
	goose -dir db/migrations postgres $(DB_CONNECTION_STRING) up

migrate-down:
	goose -dir db/migrations postgres $(DB_CONNECTION_STRING) down

seed:
	go run cmd/seed/main.go

lint:
	golangci-lint run

