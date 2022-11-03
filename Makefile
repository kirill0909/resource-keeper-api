.PHONY:
.SILENT:

build:
	sudo docker-compose up --build  app

up:
	sudo docker-compose up -d

stop:
	docker-compose stop
	
migrate:
	./migrate.sh

ping:
	curl -k -X GET https://localhost:8000/auth/ping
