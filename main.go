package main

import (
	"log"

	"github.com/hareem7bilal/go-microservice/internal/database"
	"github.com/hareem7bilal/go-microservice/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialise database client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
