help:
	    @echo ""
	    @echo "Makefile commands:"
	    @echo ""
	    @echo "DOCKER"
	    @echo ""
	    @echo "start           - Start the docker containers                     - ex: make start"
	    @echo "stop            - Stop the docker containers                      - ex: make stop"
	    @echo ""
	    @echo "CLIs"
	    @echo ""
	    @echo "server          - Access go container                             - ex: make server"
	    @echo "database        - Access database container                       - ex: make database"
	    @echo "database-cli    - Access database (psql) CLI                      - ex: make database-cli"
	    @echo ""
	    @echo "DATABASE"
	    @echo ""
	    @echo "migrate-up      - Run up migrations                               - ex: make migrate-up"
	    @echo "migrate-down    - Run down migrations                             - ex: make migrate-down"
	    @echo "migrate-create  - Creates up and down migration files             - ex: migrate-create name=migration_name_here"
	    @echo ""
	    @echo "JAVASCRIPT"
	    @echo ""
	    @echo "js-bundle       - Bundle JavaScript assets and watch for changes  - ex: make js-bundle"
	    @echo "js-install      - Install NPM package                             - ex: make js-install pkg=react"
	    @echo ""

start:
	docker-compose -f docker-compose.yml up -d

stop:
	docker-compose -f docker-compose.yml down

server:
	docker exec -it realestate_server /bin/sh

database:
	docker exec -it realestate_database /bin/bash

database-cli:
	docker-compose -f docker-compose.yml exec database psql -Urealestate -drealestate_db

migrate-create:
	 migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -database postgres://realestate:realestate_pass@localhost:5432/realestate_db?sslmode=disable -path db/migrations up

migrate-down:
	migrate -database postgres://realestate:realestate_pass@localhost:5432/realestate_db?sslmode=disable -path db/migrations down

js-bundle:
	docker exec -it realestate_server /bin/sh -c "npm run bundle"

js-install:
	docker exec -it realestate_server /bin/sh -c "npm install $(pkg)"
