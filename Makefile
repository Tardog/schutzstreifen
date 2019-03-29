project = schutzstreifen
buffalo_exec = docker-compose -p $(project) -f docker-compose.build.yml run --rm build buffalo
buffalo_env ?= development

install: create-db migrate compile-css

create-db:
	$(buffalo_exec) pop create -e $(buffalo_env)

reset-db:
	$(buffalo_exec) pop reset -e $(buffalo_env)

drop-db:
	$(buffalo_exec) pop drop -e $(buffalo_env)

migrate:
	$(buffalo_exec) pop migrate -e $(buffalo_env)

start:
	docker-compose -p $(project) -f docker-compose.yml up -d

stop:
	docker-compose -p $(project) -f docker-compose.yml down

restart: stop start

build-containers:
	docker-compose -p $(project) -f docker-compose.yml build --no-cache

test:
	-$(buffalo_exec) test

compile-css:
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css

cli:
	docker-compose -p $(project) -f docker-compose.yml exec dev bash
