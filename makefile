services:
	docker compose up -d

services-logs:
	docker compose logs -f

services-down:
	docker compose down

build-styles:
	npx tailwindcss -i ./styles/styles.css -o ./public/styles.css

dev:
	go run main.go

build-templates:
	templ generate
