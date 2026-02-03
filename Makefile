.DEFAULT_GOAL := restart

up: docker-up
down: docker-down
restart: down up

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-build:
	docker-compose build

create-env:
	cp .env.example .env