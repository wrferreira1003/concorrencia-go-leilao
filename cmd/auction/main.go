package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/database/mongodb"
)

func main() {

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// NewMongoDBConnection creates a new connection to the MongoDB database using the provided context.
	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error trying to connect with mongodb", err)
		return
	}

	fmt.Println("Connected to MongoDB")

}
