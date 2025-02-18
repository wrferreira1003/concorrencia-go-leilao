package userusecase

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/user_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
)

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUsecaseInterface interface {
	FindUserByID(ctx context.Context, id string) (*UserOutputDto, *internal_error.InternalError)
}

func (u *UserUseCase) FindUserByID(ctx context.Context, id string) (*UserOutputDto, *internal_error.InternalError) {
	user, err := u.UserRepository.FindUserByID(ctx, id)
	if err != nil {
		logger.Error("error finding user", err)
		return nil, internal_error.NewInternalServerError("error finding user")
	}

	return &UserOutputDto{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
