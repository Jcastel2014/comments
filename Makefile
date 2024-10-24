include .envrc

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@eco 'Running application...'
	@go run ./cmd/api -port=4000 -env=development -db-dsn=${COMMENTS_DB_DSN}

## db/psql : connect to the database using psql (terminal)

.PHONY: db/psql
db/psql:
	psql ${COMMENTS_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## .PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${COMMENTS_DB_DSN} up