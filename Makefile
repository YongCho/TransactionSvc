# sqlc is used in generating the idiomatic Go code for the database adapter.
# https://docs.sqlc.dev/en/stable/index.html
# It is only used in build time, and is not part of the runtime stack.
SQLC_IMAGE := kjconroy/sqlc:1.16.0

# Runtime images.
API_SERVER_IMAGE := pismo-api
DB_INIT_IMAGE := pismo-dbinit
POSTGRES_IMAGE := postgres:15.1

# Credentials. Override them using environment variables as necessary.
DB_USER ?= pismo
DB_PASSWORD ?= pismo

# Generate code.
.PHONY: generate
generate:
	docker pull $(SQLC_IMAGE)
	docker run --rm -v $(PWD):/src -w /src/sqlc $(SQLC_IMAGE) generate

# Generate .env file.
.PHONY: dotenv
dotenv:
	@> .env && \
		echo "API_SERVER_IMAGE=$(API_SERVER_IMAGE)" >> .env && \
		echo "DB_INIT_IMAGE=$(DB_INIT_IMAGE)" >> .env && \
		echo "POSTGRES_IMAGE=$(POSTGRES_IMAGE)" >> .env && \
		echo "DB_USER=$(DB_USER)" >> .env && \
		echo "DB_PASSWORD=$(DB_PASSWORD)" >> .env

# Build the docker images.
.PHONY: build
build: dotenv generate
	docker compose build

# Bring up all services.
.PHONY: up
up: dotenv build
	docker compose up

# Stop and remove all services.
.PHONY: down
down: dotenv
	docker compose down

# Clean up all images, volumns, etc.
.PHONY: clean
clean: down
	docker image rm -f $(API_SERVER_IMAGE) $(DB_INIT_IMAGE) $(POSTGRES_IMAGE)
	docker volume rm -f pismo_pismo-data
	rm -f .env

# Clean up and rebuild the services from scratch.
.PHONY: reset
reset: clean up
