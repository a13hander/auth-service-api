# auth-service-api

## Зависимости проекта
- docker
- [docker compose plugin](https://docs.docker.com/compose/install/linux/)

## Установка проекта
выполните команды в терминале:
- `cp .env.dist .env`
- `cp docker-compose.override.yml.dist docker-compose.override.yml`

заполните файл `.env` актуальными значениями
заполните файл `docker-compose.override.yml` актуальными значениями

## Запуск проекта
выполните команды в терминале:
- `make env-up`

## Остановка проекта
выполните команды в терминале
- `make env-down`

## Дополнительно
Дополнительные команды:
- `make env-status` статус контейнеров окружения
