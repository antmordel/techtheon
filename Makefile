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
	@echo "🚀 Running application locally"
	@go run app/services/rss/main.go