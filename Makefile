include .env

DOCKER_COMPOSE_FLAGS=-f docker-compose.yml -f docker-compose.override.yml
MIGRATIONS_DIR=$(CURDIR)/migrations

DSN="host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable"

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/kisielk/errcheck@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(CURDIR)/bin

env-up:
	docker compose ${DOCKER_COMPOSE_FLAGS} up -d

env-down:
	docker compose ${DOCKER_COMPOSE_FLAGS} down

env-status:
	docker compose ${DOCKER_COMPOSE_FLAGS} ps

migrate-up:
	goose -dir ${MIGRATIONS_DIR} postgres ${DSN} up

migrate-down:
	goose -dir ${MIGRATIONS_DIR} postgres ${DSN} down

migrate-status:
	goose -dir ${MIGRATIONS_DIR} postgres ${DSN} status

lint:
	bin/golangci-lint run ./...