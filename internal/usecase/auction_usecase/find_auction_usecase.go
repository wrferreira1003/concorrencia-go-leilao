package auctionusecase

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
)

func (u *AuctionUseCase) FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDto, *internal_error.InternalError) {

	auction, err := u.auctionRepository.FindAuctionByID(ctx, id)
	if err != nil {
		return nil, internal_error.NewBadRequestError(err.Error())
	}

	auctionOutputDto := &AuctionOutputDto{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	return auctionOutputDto, nil
}

func (u *AuctionUseCase) UpdateAuction(ctx context.Context, status AuctionStatus, category string, productName string) ([]AuctionOutputDto, *internal_error.InternalError) {

	auctions, err := u.auctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, internal_error.NewBadRequestError(err.Error())
	}

	var auctionOutputDtos []AuctionOutputDto
	for _, value := range auctions {
		auctionOutputDtos = append(auctionOutputDtos, AuctionOutputDto{
			ID:          value.ID,
			ProductName: value.ProductName,
			Category:    value.Category,
			Description: value.Description,
			Condition:   ProductCondition(value.Condition),
			Status:      AuctionStatus(value.Status),
			Timestamp:   value.Timestamp,
		})
	}

	return auctionOutputDtos, nil
}
