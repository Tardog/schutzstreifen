project = schutzstreifen
buffalo_compose = docker-compose -p $(project) -f docker-compose.yml
buffalo_run = $(buffalo_compose) run --rm build
buffalo_env ?= development

install: create-db migrate compile-css

create-db:
	$(buffalo_run) pop create -e $(buffalo_env)

reset-db:
	$(buffalo_run) pop reset -e $(buffalo_env)

drop-db:
	$(buffalo_run) pop drop -e $(buffalo_env)

migrate:
	$(buffalo_run) pop migrate -e $(buffalo_env)

start:
	$(buffalo_compose) up app -d

start-dev:
	$(buffalo_compose) up app

stop:
	$(buffalo_compose) down

restart: stop start

build-containers:
	$(buffalo_compose) -f docker-compose.yml build --no-cache

test:
	-$(buffalo_run) test

css:
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css

cli:
	$(buffalo_compose) exec dev bash
