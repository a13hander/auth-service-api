include .env

DOCKER_COMPOSE_FLAGS=-f docker-compose.yml -f docker-compose.override.yml
MIGRATIONS_DIR=$(CURDIR)/migrations

DSN="host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable"

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/kisielk/errcheck@latest
	GOBIN=$(CURDIR)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(CURDIR)/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(CURDIR)/bin go install github.com/envoyproxy/protoc-gen-validate@v1
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

generate:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 \
	--proto_path vendor.protogen \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate --validate_out=lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
	api/auth_v1/service.proto

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		 mkdir -p vendor.protogen/validate && \
		 git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate && \
		 mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate && \
		 rm -rf vendor.protogen/protoc-gen-validate ; \
	fi