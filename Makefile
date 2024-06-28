ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=db_user password=db_password dbname=birthday_service host=localhost port=5432 sslmode=disable
endif

INTERNAL_PATH=$(CURDIR)/internal
MOCKGEN_TAG=1.6.0
MIGRATION_FOLDER=$(INTERNAL_PATH)/database/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down
