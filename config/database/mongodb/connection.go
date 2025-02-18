package mongodb

import (
	"context"
	"os"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoDBURL := os.Getenv(MONGODB_URL)
	mongoDBDatabase := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoDBURL),
	)
	if err != nil {
		logger.Error("error trying to connect with mongodb", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("error trying to ping mongodb", err)
		return nil, err
	}

	return client.Database(mongoDBDatabase), nil
}
