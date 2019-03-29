project=schutzstreifen
buffalo_exec=docker-compose -p $(project) -f docker-compose.build.yml run --rm build buffalo

install:
	$(buffalo_exec) pop create -a
	$(buffalo_exec) pop migrate
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css

reset-db:
	$(buffalo_exec) pop reset -a

drop-db:
	$(buffalo_exec) pop drop -a

migrate:
	$(buffalo_exec) pop migrate

start:
	docker-compose -p $(project) -f docker-compose.yml up -d

stop:
	docker-compose -p $(project) -f docker-compose.yml down

restart: stop start

build-containers:
	docker-compose -p $(project) -f docker-compose.yml build --no-cache

test:
	-$(buffalo_exec) test

css:
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css

cli:
	docker-compose -p $(project) -f docker-compose.yml exec dev bash
