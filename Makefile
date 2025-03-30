include .env


DOCKER_COMPOSE := docker-compose


start-db:
	@echo "Starting database..."
	$(DOCKER_COMPOSE) --env-file ./.env up -d postgres

stop-db:
	@echo "Stopping and delete database..."
	$(DOCKER_COMPOSE) rm -sf postgres

start-local:
	@echo "Starting app..."
	go run .\cmd\main.go


#Запустить приложение(с Docker)
start-docker:
	@echo "Starting containers..."
	$(DOCKER_COMPOSE) up --build -d

# Остановить контейнеры
stop-docker:
	@echo "Stopping containers..."
	$(DOCKER_COMPOSE) down

# Запуск всех тестов
test:
	@echo "Running tests..."
	go test -v ./tests/...

# Применить миграции (
migrate:
	@echo "Applying migrations..."
	migrate -path ./migrations -database "postgres://user:qwerty@localhost:5432/postgres?sslmode=disable" up

# Откатить миграции
migrate-down:
	@echo "Rolling back migrations..."
	migrate -path ./migrations -database "postgres://user:qwerty@localhost:5432/postgres?sslmode=disable" down