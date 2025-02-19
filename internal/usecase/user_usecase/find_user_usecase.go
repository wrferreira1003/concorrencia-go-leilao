package userusecase

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/user_entity"
)

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUsecaseInterface interface {
	FindUserByID(ctx context.Context, id string) (*UserOutputDto, error)
	CreateUser(ctx context.Context, user *user_entity.User) (*UserOutputDto, error)
}

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) FindUserByID(ctx context.Context, id string) (*UserOutputDto, error) {
	user, err := u.UserRepository.FindUserByID(ctx, id)
	if err != nil {
		logger.Error("error finding user", err)
		return nil, err
	}

	return &UserOutputDto{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
