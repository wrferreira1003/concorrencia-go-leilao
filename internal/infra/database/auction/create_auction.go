package auction_repository

import (
	"context"
	"os"
	"time"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	ID          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}

type AuctionRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewAuctionRepositoryMongo(
	collection *mongo.Database,
) *AuctionRepositoryMongo {
	return &AuctionRepositoryMongo{
		Collection: collection.Collection("auctions"),
	}
}

func (r *AuctionRepositoryMongo) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {

	// Convert the auction entity to the auction entity mongo
	auctionMongo := &AuctionEntityMongo{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	// Insert the auction entity mongo into the database
	_, err := r.Collection.InsertOne(ctx, auctionMongo)
	if err != nil {
		logger.Error("error creating auction", err)
		return internal_error.NewInternalServerError("error creating auction")
	}

	// Fica aguardando o intervalo de tempo e atualiza o status para completed
	go func() {
		select {
		case <-time.After(getAuctionInterval()):
			update := bson.M{
				"$set": bson.M{
					"status": auction_entity.Completed,
				},
			}
			filter := bson.M{
				"_id": auction.ID,
			}
			_, err := r.Collection.UpdateOne(ctx, filter, update)
			if err != nil {
				logger.Error("error updating auction", err)
				return
			}
		}
	}()

	return nil
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	if duration, err := time.ParseDuration(auctionInterval); err == nil {
		return duration
	}
	return time.Minute * 5
}
