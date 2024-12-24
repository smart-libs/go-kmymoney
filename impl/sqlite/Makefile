
GO=go
GO_TESTS_INT=./test/integration/...
SQLITE3_DB_DIR=$(PWD)/test/integration/db
#SQLITE3_DB=file:$(SQLITE3_DB_DIR)/test.db?mode=memory&cache=shared
SQLITE3_DB="$(SQLITE3_DB_DIR)/test.db?_fk=true"
SQL_MIGRATIONS_DIR=$(PWD)/build/db/migrations

#init.repo: Initialize the repository
init.repo:
	$(GO) install github.com/pressly/goose/v3/cmd/goose@latest

#migrate.new: @ creates a new migration file
migrate.new:
ifndef NAME
	@echo
	@echo ">>>>>>>>>>>> Please call make with NAME=<data-migration-name>"
	@echo "In order to create a new SQL file that performs the database change"
	@echo
	@exit 1
endif
	mkdir -p $(SQLITE3_DB_DIR)
	goose -dir $(SQL_MIGRATIONS_DIR) sqlite3 $(SQLITE3_DB) create $(NAME) sql

#migrate.up: @ run all migrations in the specified database (see README for details)
migrate.up:
	@echo ====[Migrating UP]=======================================================================
	mkdir -p $(SQLITE3_DB_DIR)
	goose -dir $(SQL_MIGRATIONS_DIR) sqlite3 $(SQLITE3_DB) up

#migrate.down: @ downgrade to last change (see README for database migration details)
migrate.down:
	goose -dir $(SQL_MIGRATIONS_DIR) sqlite3 $(SQLITE3_DB) down

tests.integration:
	@echo "Run migrations"
	$(MAKE) migrate.up
	@echo "Run tests"
	export SQLITE3_DB="$(SQLITE3_DB)" && \
	$(GO) test -p 1 -v $(GO_TESTS_INT)