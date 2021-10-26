help:
	    @echo "Makefile commands:"
	    @echo "database - Access database (psql) CLI"
	    @echo "migrate-up - Run up migrations"
	    @echo "migrate-down - Run down migrations"
	    @echo "cli-server - Access go container cli"
	    @echo "cli-database - Access database container cli"

database:
	docker-compose -f docker-compose.yml exec database psql -Urealestate -drealestate_db

migrate-up:
	migrate -database postgres://realestate:realestate_pass@localhost:5432/realestate_db?sslmode=disable -path db/migrations up

migrate-down:
	migrate -database postgres://realestate:realestate_pass@localhost:5432/realestate_db?sslmode=disable -path db/migrations down

cli-server:
	docker exec -it realestate_server_1 /bin/sh

cli-database:
	docker exec -it realestate_database_1 /bin/bash
