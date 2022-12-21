package main

import (
	"log"
	"os"

	"pismo.io/db"
	"pismo.io/rest"
	"pismo.io/util"

	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.GetDBHandle()
	if err != nil {
		log.Printf("Could not get DB handle: %s", err)
		os.Exit(1)
	}
	dbAdapter, err := db.NewDBAdapter(dbConn)
	if err != nil {
		log.Printf("Could not create DB adapter: %s", err)
		os.Exit(1)
	}
	handler, err := rest.NewHandler(dbAdapter)
	if err != nil {
		log.Printf("Could not create API handler: %s", err)
		os.Exit(1)
	}

	server := rest.NewServer()
	server.AddHandlerFunc("/accounts", handler.CreateAccount)
	server.ListenAndServe(util.Env.GetListenPort())
}
