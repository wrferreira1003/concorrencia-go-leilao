package bidusecase

import (
	"context"
)

func (u *BidUseCase) FindBidByIDAuctionId(ctx context.Context, auctionID string) ([]BidOutputDto, error) {
	bids, err := u.bidRepository.FindBidByID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	var bidsOutputDto []BidOutputDto
	for _, bid := range bids {
		bidsOutputDto = append(bidsOutputDto, BidOutputDto{
			ID:        bid.ID,
			UserID:    bid.UserID,
			AuctionID: bid.AuctionID,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidsOutputDto, nil
}

func (u *BidUseCase) FindWinnerBidByAuctionId(ctx context.Context, auctionID string) (*BidOutputDto, error) {
	winnerBid, err := u.bidRepository.FindWinnerBidByAuctionID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	return &BidOutputDto{
		ID:        winnerBid.ID,
		UserID:    winnerBid.UserID,
		AuctionID: winnerBid.AuctionID,
		Amount:    winnerBid.Amount,
		Timestamp: winnerBid.Timestamp,
	}, nil
}
