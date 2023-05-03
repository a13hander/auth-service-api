DOCKER_COMPOSE_FLAGS=-f docker-compose.yml -f docker-compose.override.yml

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

env-up:
	docker compose ${DOCKER_COMPOSE_FLAGS} up -d

env-down:
	docker compose ${DOCKER_COMPOSE_FLAGS} down

env-status:
	docker compose ${DOCKER_COMPOSE_FLAGS} ps
