# sqlc is used in generating the idiomatic Go code for the database adapter.
# https://docs.sqlc.dev/en/stable/index.html
SQLC_IMAGE := kjconroy/sqlc:1.16.0

# Generate code.
.PHONY: generate
generate:
	docker pull $(SQLC_IMAGE)
	docker run --rm -v $(PWD):/src -w /src/sqlc $(SQLC_IMAGE) generate

# Build the docker images.
.PHONY: build
build: generate
	docker compose build

# Bring up all services.
.PHONY: up
up:
	docker compose up

# Stop and remove all services.
.PHONY: down
down:
	docker compose down

# Clean up all images, volumns, etc.
.PHONY: clean
clean: down
	docker image rm -f pismo pismo-db pismo-dbinit
	docker volume rm -f pismo_pismo-data

# Clean up and rebuild the services from scratch.
.PHONY: reset
reset: clean up
