DOCKER_BUILD_VARS := COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1
DOCKER_BUILD := ${DOCKER_BUILD_VARS} docker build

COMPOSE := $(DOCKER_BUILD_VARS) docker-compose

build:
	${COMPOSE} pull --ignore-pull-failures --include-deps
	${COMPOSE} build

start:
	${COMPOSE} up -d

bash:
	${COMPOSE} run app bash
