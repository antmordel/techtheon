# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# RUNNING DEV
# ==================================================================================== #

## run-local: run the application locally
.PHONY: run-local
run-local:
	@echo "🚀 Running application locally (remember to run docker-compose)"
	@go run app/services/rss/main.go


# ==================================================================================== #
# DATABASE
# ==================================================================================== #

## migrate-local: migrate local database
migrate-local:
	@echo "🚀 Migrating local database..."
	@DATABASE_URL=postgres://postgres:postgres@localhost:5432/feeds?sslmode=disable \
		dbmate migrate