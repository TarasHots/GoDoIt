DOCKER_BUILD_VARS := COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1
DOCKER_BUILD := ${DOCKER_BUILD_VARS} docker build

COMPOSE := $(DOCKER_BUILD_VARS) docker-compose

build:
	${COMPOSE} pull --ignore-pull-failures --include-deps
	${COMPOSE} build

start: build
	${COMPOSE} up -d

test:
	 docker build -f Dockerfile.multistage -t docker-go-do-it --progress plain --no-cache --target run-test-stage .

stop:
	${COMPOSE} down

destroy: stop
	${COMPOSE} rm --force --stop -v 
