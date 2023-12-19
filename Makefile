# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## cmd/api: run the cmd/api application
cmd/api:
	go run ./cmd/api

## cmd/api-log: run the cmd/api with wrap log application
cmd/api-log:
	go run ./cmd/api | go run cmd/logfmt/main.go

.PHONY: help confirm run/api run/api-log run/collector run/admin db/psql db/migrations/new db/migrations/up audit vendor test/coverage swagger profil