LOCAL_BIN:=$(CURDIR)/bin
PROJECT = github.com/sparhokm/go-course-ms-auth

init: local-env-build docker-down docker-pull docker-build docker-up  \
	  wait-db db-migrations-up
up: docker-up
down: docker-down
restart: down up

local-env-build:
	chmod 777 ./docker/common/env-init.sh
	./docker/common/env-init.sh ./.env ./docker/.env ./docker/.local.env ./.env.local

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-pull:
	docker-compose pull

docker-build:
	docker-compose build --pull

wait-db:
	docker-compose run --rm --no-deps migrator wait-for-it db:5432 -t 30

db-migrations-create:
	docker-compose run --rm --no-deps migrator goose -dir migrations create $(filter-out $@,$(MAKECMDGOALS)) sql

db-migrations-status:
	docker-compose run --rm --no-deps migrator goose -dir migrations status

db-migrations-up:
	docker-compose run --rm --no-deps migrator goose -dir migrations up -v

db-migrations-down:
	docker-compose run --rm --no-deps migrator goose -dir migrations down -v

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/vektra/mockery/v2@v2.42.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/grpc

fmt:
	go fmt ./...

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

mockery:
	$(LOCAL_BIN)/mockery

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

test:
	go test ./... | grep -v "no test files"

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode=count -coverpkg=$(PROJECT)/... -count=5
	grep -v 'mocks\|config\|/pkg/user_v1\|/cmd\|/internal/app' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	rm coverage.out

test-coverage-ci:
	go test ./... -coverprofile=coverage.tmp.out -covermode=atomic -coverpkg=$(PROJECT)/... -race -count=5
	grep -v 'mocks\|config\|/pkg/user_v1\|/cmd\|/internal/app' coverage.tmp.out  > coverage.out
	rm coverage.tmp.out