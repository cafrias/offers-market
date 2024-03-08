services:
	docker compose up -d

services-logs:
	docker compose logs -f

services-down:
	docker compose down

build-styles:
	npx tailwindcss -i ./styles/styles.css -o ./public/styles.css

dev:
	air

build-proto:
	protoc -I=api --go_out=api --go_opt=paths=source_relative api/api.proto

build-templates:
	templ generate
