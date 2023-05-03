# auth-service-api

## Зависимости проекта
- docker
- [docker compose plugin](https://docs.docker.com/compose/install/linux/)

## Установка проекта
выполните команды в терминале:
- `make install-deps`
- `cp .env.dist .env`
- `cp docker-compose.override.yml.dist docker-compose.override.yml`

заполните файлы `.env` и `docker-compose.override.yml` актуальными значениями
- `make migrate-up`

## Запуск проекта
выполните команды в терминале:
- `make env-up`

## Остановка проекта
выполните команды в терминале
- `make env-down`

## Дополнительно
Дополнительные команды:
- `make env-status` статус контейнеров окружения
