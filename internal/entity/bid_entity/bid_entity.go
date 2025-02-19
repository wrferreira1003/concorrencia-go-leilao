package bid_entity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
)

type Bid struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	AuctionID string    `json:"auction_id" bson:"auction_id"`
	Amount    float64   `json:"amount" bson:"amount"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bids []Bid) *internal_error.InternalError
	FindBidByID(ctx context.Context, auctionID string) ([]Bid, *internal_error.InternalError)
	FindWinnerBidByAuctionID(ctx context.Context, auctionID string) (*Bid, *internal_error.InternalError)
}

func NewBidEntity(userID string, auctionID string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := Bid{
		ID:        uuid.New().String(),
		UserID:    userID,
		AuctionID: auctionID,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return &bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	if _, err := uuid.Parse(b.UserID); err != nil {
		return internal_error.NewBadRequestError("user_id is not a valid uuid")
	}

	if _, err := uuid.Parse(b.AuctionID); err != nil {
		return internal_error.NewBadRequestError("auction_id is not a valid uuid")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("amount must be greater than 0")
	}

	return nil
}
