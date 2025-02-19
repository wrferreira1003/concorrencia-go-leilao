package auctionusecase

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
	bidusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/bid_usecase"
)

func (u *AuctionUseCase) FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDto, error) {

	auction, err := u.auctionRepository.FindAuctionByID(ctx, id)
	if err != nil {
		return nil, err
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

func (u *AuctionUseCase) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category string, productName string) ([]AuctionOutputDto, error) {

	auctions, err := u.auctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
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

func (u *AuctionUseCase) FindWinnerBidByAuctionId(ctx context.Context, auctionID string) (*WinnerInfoOutputDto, error) {

	//Busca o leilão no banco de dados
	auction, err := u.auctionRepository.FindAuctionByID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	//Busca o lance vencedor do leilão
	bidWinner, err := u.bidRepository.FindWinnerBidByAuctionID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	//Converte o leilão para o DTO
	auctionOutputDto := AuctionOutputDto{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	if err != nil {
		//Se não houver lance vencedor, retorna o leilão sem o lance vencedor
		return &WinnerInfoOutputDto{
			AuctionID: auctionOutputDto,
			Bid:       nil,
		}, nil
	}

	//Converte o lance vencedor para o DTO caso exista
	bidOutputDto := &bidusecase.BidOutputDto{
		ID:        bidWinner.ID,
		UserID:    bidWinner.UserID,
		AuctionID: bidWinner.AuctionID,
		Amount:    bidWinner.Amount,
		Timestamp: bidWinner.Timestamp,
	}

	//Retorna o leilão com o lance vencedor
	return &WinnerInfoOutputDto{
		AuctionID: auctionOutputDto,
		Bid:       bidOutputDto,
	}, nil
}
