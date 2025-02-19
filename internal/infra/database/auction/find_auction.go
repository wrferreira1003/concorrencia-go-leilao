package auction_repository

import (
	"context"
	"errors"
	"time"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *AuctionRepositoryMongo) FindAuctionByID(ctx context.Context, auctionID string) (*auction_entity.Auction, error) {

	// Define the filter to find the auction by ID
	filter := bson.M{"_id": auctionID}

	var auction AuctionEntityMongo

	err := r.Collection.FindOne(ctx, filter).Decode(&auction)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("auction not found", err)
			return nil, internal_error.NewNotFoundError("auction not found")
		}
		logger.Error("error finding auction", err)
		return nil, internal_error.NewInternalServerError("error finding auction")
	}

	return &auction_entity.Auction{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   time.Unix(auction.Timestamp, 0),
	}, nil
}

func (r *AuctionRepositoryMongo) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category string,
	productName string,
) ([]*auction_entity.Auction, error) {

	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{Pattern: productName, Options: "i"}
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("error finding auctions", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var auctions []AuctionEntityMongo
	if err := cursor.All(ctx, &auctions); err != nil {
		logger.Error("error decoding auctions", err)
		return nil, err
	}

	var auctionEntities []*auction_entity.Auction
	for _, auction := range auctions {
		auctionEntities = append(auctionEntities, &auction_entity.Auction{
			ID:          auction.ID,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   auction.Condition,
			Status:      auction.Status,
			Timestamp:   time.Unix(auction.Timestamp, 0),
		})
	}

	return auctionEntities, nil
}
