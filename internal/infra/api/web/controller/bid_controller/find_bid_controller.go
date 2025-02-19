package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/rest_err.go"
)

func (b *BidController) FindBidAuctionByID(c *gin.Context) {
	auctionID := c.Param("auctionId")

	// Validate user id
	if err := uuid.Validate(auctionID); err != nil {
		errRest := rest_err.NewBadRequestError("invalid auction id", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid auction id",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	bidoutputList, err := b.bidUseCase.FindBidByIDAuctionId(context.Background(), auctionID)
	if err != nil {
		restErr := rest_err.ConvertToRestErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, bidoutputList)
}
