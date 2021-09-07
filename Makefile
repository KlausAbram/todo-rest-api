.SILENT:

go-home:
	cd /home/klaus-abram/github.com/todo-rest-api/

migrate-up:
	migrate -path ./schema -database postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable up

migrate-down:
	migrate -path ./schema -database postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable down

down:
	docker stop 95e286827948 && docker stop 536c12783476

build:
	docker-compose build todo-api

run:
	docker-compose up todo-api

test:
	go test -v ./...

run-table:
	psql -U postgres

run-db:
	docker exec -it 536c12783476 /bin/bash


