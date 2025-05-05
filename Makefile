# Переменные для подключения к БД
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?=
DB_NAME ?= persons
SSL_MODE ?= disable

MIGRATIONS_DIR ?= ./migrations

DSN := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

MIGRATE := migrate

.PHONY: migrate-up migrate-down migrate-force migrate-create migrate-version

# Применить все миграции
migrate-up:
	@echo "Applying all up migrations..."
	@$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DSN)" up

# Откатить все миграции
migrate-down:
	@echo "Applying all down migrations..."
	@$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DSN)" down

migrate-create:
	@echo "Creating new migration..."
	@$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

run:
	go run ./cmd/rest_api main.go