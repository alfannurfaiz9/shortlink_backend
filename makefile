include ./.env

DB_URL=postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable

MIGRATE_PATH=db/migrations

migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq create_$(NAME)_table

migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATE_PATH) up

migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATE_PATH) down $(N)

db-reset:
	make migrate-down migrate-up