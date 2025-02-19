package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (b *BidController) FindBidAuctionByID(c *gin.Context) {
	auctionID := c.Param("auctionId")

	// Validate user id
	if err := uuid.Validate(auctionID); err != nil {
		return
	}

	bidoutputList, err := b.bidUseCase.FindBidByIDAuctionId(context.Background(), auctionID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, bidoutputList)
}
