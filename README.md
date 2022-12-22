# Transaction Service

This is a simple credit card transaction service written in Go.

## Requirements
- Linux
- Docker
- Ports 8080 and 5432 available

This project was written with a Linux server in mind. It may need some modification to work on Windows. If Windows runtime is desired, please email me, and I can make the change.

Also, the project was developed and tested using Docker Engine version `20.10`.

The following ports must be available for binding:
- `8080`: Used for serving the REST API.
- `5432`: Used for connecting to the database, mainly for development purpose.

## Build and Run

### Build
Use the following command to build the project.

```bash
make build
```

### Run
Use the following command to run all the services.

```bash
make up
```

### View Logs
Use the following command to tail the container logs.

```bash
make logs
```

### Run Test
Use the following command to run the tests.

```
make test
```

### Stop and Clean Up
Use the following command to stop the services and clean up all containers and data.

```bash
make clean
```

## API Request Example
Create a new account.

```bash
curl -X POST http://localhost:8080/accounts -H "Content-Type: application/json" --data '{"document_number": "F432"}'

# {"ID":2,"document_number":"F432"}
```

Fetch an existing account.

```bash
curl -X GET http://localhost:8080/accounts/2

# {"ID":2,"document_number":"F432"}
```

Create a new transaction.

```bash
curl -X POST http://localhost:8080/transactions -H "Content-Type: application/json" --data '{"account_id": 2, "operation_type_id": 4, "amount": 59.99}'

# {"ID":2,"account_id":2,"operation_type_id":4,"Amount":59.99}
```

## Components

The transaction service has the following containers.

- `pismo-api`: a REST API server for handling the HTTP requests.
- `pismo-db`: a Postgres database used for storing the account and transaction data.
- `pismo-dbinit`: a ephemeral container used for initializing the database.

## Development
The project was developed in Visual Studio Code on a Linux machine. Installing Go extension (golang.go) is highly recommended.

There are some notable external libraries used in the code to facilitate the development.

### SQLC
[sqlc](https://docs.sqlc.dev/en/stable/index.html#) is used for generating Go code for interacting with the database.

### Gin-Gonic
[gin-gonic](https://gin-gonic.com/docs/introduction/) is a HTTP web framework used for implementing the REST server.

## TODO
- Increase test coverage.
- Make sure concurrent transactions are handled correctly.
