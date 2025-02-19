package bid_repository

import (
	"context"
	"time"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/bid_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *BidRepositoryMongo) FindBidByID(ctx context.Context, auctionID string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		logger.Error("error finding bid", err)
		return nil, internal_error.NewNotFoundError("bid not found")
	}

	var bids []BidEntityMongo
	if err := cursor.All(ctx, &bids); err != nil {
		logger.Error("error finding bid", err)
		return nil, internal_error.NewNotFoundError("bid not found")
	}

	var bidsEntity []bid_entity.Bid
	for _, bid := range bids {
		bidsEntity = append(bidsEntity, bid_entity.Bid{
			ID:        bid.ID,
			UserID:    bid.UserID,
			AuctionID: bid.AuctionID,
			Amount:    bid.Amount,
			Timestamp: time.Unix(bid.Timestamp, 0),
		})
	}
	return bidsEntity, nil
}

func (r *BidRepositoryMongo) FindWinnerBidByAuctionID(ctx context.Context, auctionID string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionID}

	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	var bid BidEntityMongo
	if err := r.collection.FindOne(ctx, filter, opts).Decode(&bid); err != nil {
		logger.Error("error finding bid", err)
		return nil, internal_error.NewNotFoundError("bid not found")
	}

	return &bid_entity.Bid{
		ID:        bid.ID,
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Amount:    bid.Amount,
		Timestamp: time.Unix(bid.Timestamp, 0),
	}, nil
}
