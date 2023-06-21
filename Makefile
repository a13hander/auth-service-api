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
	GOBIN=$(CURDIR)/bin go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(CURDIR)/bin go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
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

run:
	go run cmd/server/main.go

generate:
	mkdir -p pkg/swagger
	make generate-auth-api
	make generate-access-api
	statik -src=pkg/swagger

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path api/auth_v1 \
	--proto_path vendor.protogen \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate --validate_out=lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
	--grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
  --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
  --openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	api/auth_v1/auth.proto

generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 \
	--proto_path vendor.protogen \
	--go_out=pkg/access_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/access_v1/access.proto

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
  		mkdir -p vendor.protogen/validate && \
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate && \
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate && \
		rm -rf vendor.protogen/protoc-gen-validate ; \
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
		rm -rf vendor.protogen/openapiv2 ;\
	fi
