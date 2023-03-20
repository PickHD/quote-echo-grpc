.PHONY: test

# Load .env configurations
ifneq (,$(wildcard ./.env))
    include .env
    export
    MIGRATIONS_SOURCE = file://db/migrations
	DATABASE_URL = postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
endif

# Run GRPC server
grpc:
	go run cmd/main.go

# Run golang-migrate up
migrate-up:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" up

# Run golang-migrate down
migrate-down:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" down -all

# Run golang-migrate create
migrate-create:
	migrate create -ext sql -dir ./db/migrations $(filter-out $@,$(MAKECMDGOALS))

# Run code generator
gen:
	go generate ./...
	buf generate

# Run unit tests with coverage info
test:
	go test ./... -cover

# Run linters
lint:
	golangci-lint run
	buf lint

# Run code formatter
fmt:
	gofmt -w -s .
	find . -name '*.proto' | xargs clang-format -i