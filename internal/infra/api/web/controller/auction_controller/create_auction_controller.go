package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/rest_err.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/infra/api/web/validation"
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
		restErr := validation.ValidateErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := a.auctionUseCase.CreateAuction(context.Background(), &auction)
	if err != nil {
		restErr := rest_err.ConvertToRestErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
