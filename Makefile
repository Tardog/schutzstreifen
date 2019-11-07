project = schutzstreifen
docker = docker-compose -p $(project) -f docker-compose.yml
build_run = $(docker) run --rm build
buffalo_env ?= development
build_arguments ?=
task ?=

setup: build-containers install-plugins reset-db migrate css

install-plugins:
	$(build_run) plugins install

create-db:
	$(build_run) pop create -e $(buffalo_env)

reset-db:
	$(build_run) pop reset -e $(buffalo_env)

drop-db:
	$(build_run) pop drop -e $(buffalo_env)

migrate:
	$(build_run) pop migrate -e $(buffalo_env)

start:
	$(docker) up -d app

start-dev:
	$(docker) up app

stop:
	$(docker) down

restart: stop start

build-containers:
	$(docker) -f docker-compose.yml build $(build_arguments)

test:
	-$(build_run) test

css:
	$(docker) run --rm --entrypoint sassc build -t compressed public/assets/scss/application.scss public/assets/application.css

cli: start
	$(docker) exec app bash

task: start
	$(docker) exec app buffalo task $(task)
