PROJECT=schutzstreifen

install:
	buffalo db create -a
	buffalo db migrate up

start:
	docker-compose -p $(PROJECT) -f docker-compose.yml up -d

stop:
	docker-compose -p $(PROJECT) -f docker-compose.yml down

build:
	docker-compose -p $(PROJECT) -f docker-compose.yml build --no-cache

test:
	-docker-compose -p $(PROJECT) -f docker-compose.yml -f docker-compose.build.yml run --rm build buffalo test

css:
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css

cli:
	docker-compose -p $(PROJECT) -f docker-compose.yml exec dev bash
