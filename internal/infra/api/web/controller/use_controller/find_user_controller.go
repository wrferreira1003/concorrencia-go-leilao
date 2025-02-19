package use_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/rest_err.go"
	userusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/user_usecase"
)

type UseController struct {
	userUseCase userusecase.UserUsecaseInterface
}

func NewUseController(userUseCase userusecase.UserUsecaseInterface) *UseController {
	return &UseController{userUseCase: userUseCase}
}

func (u *UseController) FindUserByID(c *gin.Context) {
	userID := c.Param("userId")

	// Validate user id
	if err := uuid.Validate(userID); err != nil {
		errRest := rest_err.NewBadRequestError("invalid user id", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid user id",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	user, err := u.userUseCase.FindUserByID(context.Background(), userID)
	if err != nil {
		restErr := rest_err.ConvertToRestErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
