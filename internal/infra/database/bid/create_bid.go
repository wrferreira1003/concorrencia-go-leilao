package bid_repository

import (
	"context"
	"sync"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/bid_entity"
	auction_repository "github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/database/auction"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	ID        string  `bson:"_id"`
	UserID    string  `bson:"user_id"`
	AuctionID string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepositoryMongo struct {
	collection        *mongo.Collection
	AuctionRepository *auction_repository.AuctionRepositoryMongo
}

func NewBidRepositoryMongo(
	collection *mongo.Database,
	auctionRepository *auction_repository.AuctionRepositoryMongo,
) *BidRepositoryMongo {
	return &BidRepositoryMongo{
		collection:        collection.Collection("bids"),
		AuctionRepository: auctionRepository,
	}
}

func (r *BidRepositoryMongo) CreateBid(ctx context.Context, bids []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bids {
		wg.Add(1)
		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auction, err := r.AuctionRepository.FindAuctionByID(ctx, bidValue.AuctionID)
			if err != nil {
				logger.Error("error finding auction", err)
				return
			}

			// Check if the auction is active
			if auction.Status != auction_entity.Active {
				logger.Error("auction is not active", internal_error.NewNotFoundError("auction is not active"))
				return
			}

			bidMongo := BidEntityMongo{
				ID:        bidValue.ID,
				UserID:    bidValue.UserID,
				AuctionID: bidValue.AuctionID,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			if _, err := r.collection.InsertOne(ctx, bidMongo); err != nil {
				logger.Error("error creating bid", err)
				return
			}
		}(bid)
	}

	wg.Wait()

	return nil
}
