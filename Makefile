.PHONY:
.SILENT:

build:
	docker-compose up --build  app

up:
	docker-compose up -d

stop:
	docker-compose stop
	
migrate:
	./migrate.sh

ping:
	curl -k -X GET https://localhost:8000/auth/ping
