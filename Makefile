SQLC_IMAGE := kjconroy/sqlc:1.16.0

generate:
	docker pull $(SQLC_IMAGE)
	docker run --rm -v $(PWD):/src -w /src/sqlc $(SQLC_IMAGE) generate

build: generate
	docker compose build

up:
	docker compose up

down:
	docker compose down

clean:
	docker image rm -f pismo pismo-db pismo-init
	docker volume rm -f pismo_pismo-data
