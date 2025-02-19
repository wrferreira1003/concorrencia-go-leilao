package auction_controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/auction_entity"
)

func (a *AuctionController) FindAuctionByID(c *gin.Context) {
	auctionID := c.Param("auctionId")

	// Validate user id
	if err := uuid.Validate(auctionID); err != nil {
		return
	}

	auction, err := a.auctionUseCase.FindAuctionByID(context.Background(), auctionID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (a *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	product := c.Query("product")

	auctionStatus, errConvert := strconv.Atoi(status)
	if errConvert != nil {
		return
	}

	auctions, err := a.auctionUseCase.FindAuctions(context.Background(), auction_entity.AuctionStatus(auctionStatus), category, product)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (a *AuctionController) FindAuctionsWinningBid(c *gin.Context) {
	auctionID := c.Param("auctionId")

	// Validate user id
	if err := uuid.Validate(auctionID); err != nil {
		return
	}

	auction, err := a.auctionUseCase.FindWinnerBidByAuctionId(context.Background(), auctionID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, auction)
}
