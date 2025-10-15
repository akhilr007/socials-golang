# Load environment variables (recommended to use direnv or export manually)
# This Makefile assumes DB_MIGRATION_ADDR is set in the environment

MIGRATIONS_PATH = ./cmd/migrate/migrations

# Tell Make to not treat these as files
.PHONY: migrate-create migrate-up migrate-down

## Create a new migration file
migrate-create:
ifndef NAME
	$(error NAME is required. Usage: make migrate-create NAME=create_users_table)
endif
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(NAME)

## Run all up migrations
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DB_ADDR)" up

## Roll back N migrations (e.g. make migrate-down STEPS=1)
migrate-down:
ifndef STEPS
	$(error STEPS is required. Usage: make migrate-down STEPS=1)
endif
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DB_ADDR)" down $(STEPS)

## Clear dirty state manually (Usage: make migrate-fix VERSION=4)
migrate-fix:
ifndef VERSION
	$(error VERSION is required. Usage: make migrate-fix VERSION=4)
endif
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DB_ADDR)" force $(VERSION)
