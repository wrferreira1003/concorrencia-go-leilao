package use_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/user_entity"
	userusecase "github.com/wrferreira1003/concorrencia-go-leilao/internal/usecase/user_usecase"
)

type UseController struct {
	userUseCase userusecase.UserUsecaseInterface
}

func NewUseController(userUseCase userusecase.UserUsecaseInterface) *UseController {
	return &UseController{userUseCase: userUseCase}
}

func (u *UseController) CreateUser(c *gin.Context) {
	var user userusecase.UserOutputDto

	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}

	userPtr, err := u.userUseCase.CreateUser(context.Background(), &user_entity.User{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, userPtr)
}

func (u *UseController) FindUserByID(c *gin.Context) {
	userID := c.Param("userId")

	// Validate user id
	if err := uuid.Validate(userID); err != nil {
		return
	}

	user, err := u.userUseCase.FindUserByID(context.Background(), userID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}
