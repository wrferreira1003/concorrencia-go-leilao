package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	bidusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/bid_usecase"
)

type BidController struct {
	bidUseCase bidusecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bidusecase.BidUseCaseInterface) *BidController {
	return &BidController{bidUseCase: bidUseCase}
}

func (b *BidController) CreateBid(c *gin.Context) {

	var bid bidusecase.BidInputDto
	if err := c.ShouldBindJSON(&bid); err != nil {
		return
	}

	err := b.bidUseCase.CreateBid(context.Background(), &bid)
	if err != nil {
		return
	}

	c.Status(http.StatusCreated)
}
