package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/katianemiranda/leilao/configuration/database/mongodb"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")

		return
	}

	databaseClient, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
		return
	}

	fmt.Println("Successfully connected to MongoDB:", databaseClient.Name())

}
