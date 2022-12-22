# Transaction Service

This is a simple credit card transaction service written in Go.

## Requirements
- Linux
- Docker

This project is written with a Linux server in mind. It may need some modification to work on Windows. If Windows runtime is desired, please email me, and I can make the change.

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

### Stop and Clean Up
Use the following command to stop the services and clean up all containers and data.

```bash
make clean
```

## Example
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

The transaction service has the following components.

- REST API Server
- Database
- DB Init Container

### REST API Server

The REST API server is responsible for handling the HTTP requests.
### Database

A Postgres database is used to store the account and transaction data.

### DB Init Container

The init container is run once during the service start-up to create the database tables and initializing them with necessary data. It can be extended in the future to handle database migration.

## Development
The project was developed in Visual Studio Code on a Linux machine. Go extension (golang.go) is highly recommended.

## TODO
- Clean up docker compose environment variables.
- Create a package for processing transactions. Don't call db adapter directly from the HTTP handler.
- Add necessary indexes.
- Write tests.
