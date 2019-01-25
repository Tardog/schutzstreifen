PROJECT=schutzstreifen

install:
	buffalo db create -a

start:
	docker-compose -p $(PROJECT) up -d

stop:
	docker-compose -p $(PROJECT) down

build:
	docker-compose -p $(PROJECT) build

test:
	buffalo db drop -e test
	buffalo test

css:
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css
