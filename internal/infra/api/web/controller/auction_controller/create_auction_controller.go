package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	auctionusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/auction_usecase"
)

type AuctionController struct {
	auctionUseCase auctionusecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auctionusecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{auctionUseCase: auctionUseCase}
}

func (a *AuctionController) CreateAuction(c *gin.Context) {

	var auction auctionusecase.AuctionInputDto
	if err := c.ShouldBindJSON(&auction); err != nil {
		return
	}

	err := a.auctionUseCase.CreateAuction(context.Background(), &auction)
	if err != nil {
		return
	}

	c.Status(http.StatusCreated)
}
