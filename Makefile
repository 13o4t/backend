MAKEFILE_PATH = $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_NAME = $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))

SERVICE = 

CHECK_FORMAT = text

COMPLEXITY_COEF = 10
COMPLEXITY_DIR = $(shell ls -d */ | grep -v vendor)

ENV_FILE = init/.env
COMPOSE_FILE = deploy/docker-compose.yml

DEV_ENV_FILE = init/.dev.env
DEV_COMPOSE_FILE = deploy/docker-compose.dev.yml

# generate code
.PHONY: openapi
openapi:
	@./scripts/openapi.sh $(SERVICE) api/openapi internal/myapp/ports ports

.PHONY: protobuf
protobuf:
	@./scripts/protobuf.sh $(SERVICE) api/protobuf internal/pkg/genprotobuf

# static check
.PHONY: staticcheck
staticcheck:
	-staticcheck -checks all -f $(CHECK_FORMAT) ./...

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: complexity
complexity:
	@gocyclo -over $(COMPLEXITY_COEF) $(COMPLEXITY_DIR)

# test
.PHONY: test
test:
	@go test ./...

# docker compose
.PHONY: compose-build
compose-build:
	@docker compose --env-file=$(ENV_FILE) -f=$(COMPOSE_FILE) -p=$(PROJECT_NAME) build

.PHONY: compose-up
compose-up:
	@docker compose --env-file=$(ENV_FILE) -f=$(COMPOSE_FILE) -p=$(PROJECT_NAME) up

.PHONY: compose-down
compose-down:
	@docker compose --env-file=$(ENV_FILE) -f=$(COMPOSE_FILE) -p=$(PROJECT_NAME) down --remove-orphans

# develop
.PHONY: run-dev
run-dev:
	@docker compose --env-file=$(DEV_ENV_FILE) -f=$(DEV_COMPOSE_FILE) -p=$(PROJECT_NAME)_dev up

.PHONY: clean-dev
clean-dev:
	@docker compose --env-file=$(DEV_ENV_FILE) -f=$(DEV_COMPOSE_FILE) -p=$(PROJECT_NAME)_dev down --remove-orphans

# clean
.PHONY: clean
clean: compose-down
