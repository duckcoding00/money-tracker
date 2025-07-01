include .env
export


.PHONY: migration-create
migration-create:
	@goose create  $(filter-out $@,$(MAKECMDGOALS)) sql

.PHONY: migrate-up
migrate-up:
	@goose up

.PHONY: migrate-up-one
migrate-up-one:
	@goose up-by-one

.PHONY: migrate-down
migrate-down:
	@goose down

.PHONY: migrate-down-one
migrate-down-one:
	@goose down-by-one

.PHONY: migrate-down-to
migrate-down-to:
	@goose down-to $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-status
migrate-status:
	@goose status

.PHONY: migrate-version
migrate-version:
	@goose version

.PHONY: migrate-reset
migrate-reset:
	@goose reset

.PHONY: migrate-redo
migrate-redo:
	@goose redo

.PHONY: go-tidy
go-tidy:
	@go mod tidy

# Helper untuk mencegah error saat passing arguments
%:
	@:

# Helper untuk mencegah error saat passing arguments
%:
	@: