DB_CONNECTION_STRING=postgres://postgres:example@localhost:5432/postgres?sslmode=disable

services:
	docker compose up -d

dev:
	air

migrate:
	goose -dir db/migrations postgres $(DB_CONNECTION_STRING) up

lint:
	golangci-lint run

