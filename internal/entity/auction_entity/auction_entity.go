package auction_entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
)

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	// Assume that the auction status is a number between 1 and 2
	Active    AuctionStatus = iota // 0
	Completed                      // 1
)

const (
	// Assume that the product condition is a number between 1 and 3
	New         ProductCondition = iota // 0
	Used                                // 1
	Refurbished                         // 2
)

// AuctionRepositoryInterface is the interface for the auction repository
type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internal_error.InternalError
	FindAuctionByID(ctx context.Context, auctionID string) (*Auction, *internal_error.InternalError)
	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category string,
		productName string,
	) ([]*Auction, *internal_error.InternalError)
}

// NewAuctionRepository creates a new auction repository
func NewAuctionRepository(
	productName string,
	category string,
	description string,
	condition ProductCondition,
) (*Auction, error) {

	auction := &Auction{
		ID:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	err := auction.Validate()
	if err != nil {
		return nil, errors.New("invalid auction")
	}

	return auction, nil
}

func (a *Auction) Validate() error {
	if len(a.ProductName) <= 1 || len(a.Category) <= 1 || len(a.Description) <= 1 {
		return errors.New("product name, category and description must be at least 1 character")
	}

	if a.Condition < New || a.Condition > Refurbished {
		return errors.New("condition must be between 0 and 2")
	}

	return nil
}

func StringToProductCondition(condition string) (ProductCondition, error) {
	switch condition {
	case "new":
		return New, nil
	case "used":
		return Used, nil
	case "refurbished":
		return Refurbished, nil
	default:
		return -1, fmt.Errorf("invalid product condition: %s", condition)
	}
}
