package auctionusecase

import (
	"context"
	"time"

	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
)

type AuctionInputDto struct {
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Condition   string `json:"condition"`
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

type ProductCondition int
type AuctionStatus int

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInputDto *AuctionInputDto) *internal_error.InternalError
}

type AuctionUseCase struct {
	auctionRepository auction_entity.AuctionRepositoryInterface
}

func NewAuctionUseCase(
	auctionRepository auction_entity.AuctionRepositoryInterface,
) *AuctionUseCase {
	return &AuctionUseCase{
		auctionRepository: auctionRepository,
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
