PROJECT=schutzstreifen

install:
	buffalo db create -a

start:
	docker-compose -p $(PROJECT) up -d

stop:
	docker-compose -p $(PROJECT) down

test:
	-docker-compose -p $(PROJECT) -f docker-compose.yml -f docker-compose.build.yml run --rm build buffalo test
	docker-compose -p $(PROJECT) down

css:
	sassc -t compressed public/assets/scss/application.scss public/assets/application.css
