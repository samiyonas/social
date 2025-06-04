-include .envrc
MIGRATION_PATH = "./cmd/migrate/migration"

.PHONY: migrate-create migrate-up migrate-down migrate-version

migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up

migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

migrate-version:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) version

%:
	@:
