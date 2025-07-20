.PHONY: all help ui up clean logs run-api run-ui db-up db-down

UI_DIR := ui
API_DIR := cmd/main.go

all: help

help:
	@echo "BareCMS Development Commands"
	@echo ""
	@echo "  up       - Start the development environment"
	@echo "  ui       - Build UI (frontend)"
	@echo "  clean    - Stop and cleanup containers"
	@echo "  logs     - Show container logs"
	@echo "  help     - Show this help message"
	@echo ""

ui:
	cd $(UI_DIR) && npm install && npm run build

up:
	docker compose up

clean:
	docker compose down -v

logs:
	docker compose logs -f

db:
	@echo "Starting PostgreSQL database container..."
	docker compose up postgres

db-down:
	@echo "Stopping PostgreSQL database container..."
	docker compose stop postgres && docker compose rm -f postgres

api:
	@echo "Starting Go API server locally..."
	go run $(API_DIR)

dev:
	@echo "Starting UI in development mode with hot-reload..."
	cd $(UI_DIR) && npm run dev