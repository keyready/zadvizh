# Переменные
COMPOSE_FILE := d-c.prod.yml
DOCKER_COMPOSE := docker-compose -f $(COMPOSE_FILE)

.PHONY: build
build: ## Сборка приложения
	@echo "[+] Сборка приложения..."
	$(DOCKER_COMPOSE) build --no-cache

.PHONY: up
up: ## Запуск приложения
	@echo "[+] Запуск сервисов..."
	$(DOCKER_COMPOSE) up -d

.PHONY: down
down: ## Остановка приложения
	@echo "[-] Остановка приложения..."
	$(DOCKER_COMPOSE) down

.PHONY: logs_and_ps
logs: ## Просмотр логов
	@echo "[+] Просмотр логов и состояния контейнеров"
	$(DOCKER_COMPOSE) ps && $(DOCKER_COMPOSE) logs -f

.PHONY: clean
clean: ## Удаление приложения
	@echo "[-] Удаление приложения..."
	$(DOCKER_COMPOSE) down -v