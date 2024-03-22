services:
	docker compose up -d

dev:
	air

lint:
	golangci-lint run

