package auction

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
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
	database *mongo.Database,
) *AuctionRepositoryMongo {
	return &AuctionRepositoryMongo{
		Collection: database.Collection("auctions"),
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

	return nil
}
