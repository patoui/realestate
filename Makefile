help:
	    @echo "Makefile commands:"
	    @echo "database - Access database (psql) CLI"
	    @echo "migrate-up - Run up migrations"
	    @echo "migrate-down - Run down migrations"
	    @echo "migrate-create name=migration_name_here - Creates up and down migration files"
	    @echo "cli-server - Access go container cli"
	    @echo "cli-database - Access database container cli"
	    @echo "bundle - Bundle frontend assets (via esbuild)"

start:
	docker-compose -f docker-compose.yml up -d

stop:
	docker-compose -f docker-compose.yml down

database:
	docker-compose -f docker-compose.yml exec database psql -Urealestate -drealestate_db

migrate-up:
	migrate -database postgres://realestate:realestate_pass@localhost:5432/realestate_db?sslmode=disable -path db/migrations up

migrate-down:
	migrate -database postgres://realestate:realestate_pass@localhost:5432/realestate_db?sslmode=disable -path db/migrations down

migrate-create:
	 migrate create -ext sql -dir db/migrations -seq $(name)

cli-server:
	docker exec -it realestate_server /bin/sh

cli-database:
	docker exec -it realestate_database /bin/bash

bundle:
	docker exec -it realestate_server /bin/sh -c "npm run bundle"
