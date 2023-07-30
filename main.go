package main

import (
	"log"

	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/database"
	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()

	if err != nil {
		log.Fatalf("failed to initialise Database Client: %s", err)
	}

	srv := server.NewEchoServer(db)

	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
