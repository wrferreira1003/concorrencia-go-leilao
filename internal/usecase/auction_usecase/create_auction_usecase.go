package auctionusecase

import (
	"context"
	"time"

	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/bid_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
	bidusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/bid_usecase"
)

type AuctionInputDto struct {
	ProductName string `json:"product_name" binding:"required, min=1"`
	Category    string `json:"category" binding:"required, min=2"`
	Description string `json:"description" binding:"required, min=10"`
	Condition   string `json:"condition" binding:"required, min=1"`
}

type AuctionOutputDto struct {
	ID          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02T15:04:05Z"`
}

type WinnerInfoOutputDto struct {
	AuctionID AuctionOutputDto         `json:"auction"`
	Bid       *bidusecase.BidOutputDto `json:"bid,omitempty"`
}

type ProductCondition int
type AuctionStatus int

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInputDto *AuctionInputDto) *internal_error.InternalError
	FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDto, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category string, productName string) ([]AuctionOutputDto, *internal_error.InternalError)
	FindWinnerBidByAuctionId(ctx context.Context, auctionID string) (*WinnerInfoOutputDto, *internal_error.InternalError)
}

type AuctionUseCase struct {
	auctionRepository auction_entity.AuctionRepositoryInterface
	bidRepository     bid_entity.BidRepositoryInterface
}

func NewAuctionUseCase(
	auctionRepository auction_entity.AuctionRepositoryInterface,
	bidRepository bid_entity.BidRepositoryInterface,
) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepository: auctionRepository,
		bidRepository:     bidRepository,
	}
}

func (u *AuctionUseCase) CreateAuction(ctx context.Context, auctionInputDto *AuctionInputDto) *internal_error.InternalError {

	conditions, err := auction_entity.StringToProductCondition(auctionInputDto.Condition)
	if err != nil {
		return internal_error.NewBadRequestError(err.Error())
	}

	auction, err := auction_entity.NewAuctionRepository(
		auctionInputDto.ProductName,
		auctionInputDto.Category,
		auctionInputDto.Description,
		conditions,
	)
	if err != nil {
		return internal_error.NewBadRequestError(err.Error())
	}

	// TODO: Implement the auction creation
	err = u.auctionRepository.CreateAuction(ctx, auction)
	if err != nil {
		return internal_error.NewBadRequestError(err.Error())
	}

	return nil
}
