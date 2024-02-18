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
# RUN INSIDE KIND
# ==================================================================================== #

POSTGRES     := postgres:15-alpine
TELEPRESENCE := docker.io/datawire/tel2:2.10.4

KIND_CLUSTER := techtheon-cluster

dev-kind-up:
	kind create cluster \
		--image kindest/node:v1.25.3@sha256:f52781bc0d7a19fb6c405c2af83abfeb311f130707a0e219175677e366cc45d1 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml
	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-up: dev-kind-up
	kind load docker-image $(TELEPRESENCE) --name $(KIND_CLUSTER)
	telepresence --context=kind-$(KIND_CLUSTER) helm install
	telepresence --context=kind-$(KIND_CLUSTER) connect


# ==================================================================================== #
# DATABASE
# ==================================================================================== #

## migrate-local: migrate local database
migrate-local:
	@echo "🚀 Migrating local database..."
	@DATABASE_URL=postgres://postgres:postgres@localhost:5432/feeds?sslmode=disable \
		dbmate migrate