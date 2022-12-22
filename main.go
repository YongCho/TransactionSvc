package main

import (
	"log"
	"os"

	"pismo.io/db"
	"pismo.io/rest"
	"pismo.io/transaction"
	"pismo.io/util"

	_ "github.com/lib/pq"
)

func main() {
	// Create a database connection.
	dbConn, err := db.GetDBHandle()
	if err != nil {
		log.Printf("Could not get DB handle: %s", err)
		os.Exit(1)
	}

	// Create a database adapter object.
	dbAdapter, err := db.NewDBAdapter(dbConn)
	if err != nil {
		log.Printf("Could not create DB adapter: %s", err)
		os.Exit(1)
	}

	// Create the transaction manager.
	txnMgr, err := transaction.NewTransactionMgr(dbAdapter)
	if err != nil {
		log.Printf("Could not create transaction manager: %s", err)
		os.Exit(1)
	}

	// Create a REST API handler object.
	handler, err := rest.NewHandler(txnMgr)
	if err != nil {
		log.Printf("Could not create API handler: %s", err)
		os.Exit(1)
	}

	// Create and start the HTTP server.
	server := rest.NewServer(handler)
	server.Run(util.Env.GetListenPort())
}
