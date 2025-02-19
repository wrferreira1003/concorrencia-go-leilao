package userusecase

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/user_entity"
)

func (u *UserUseCase) CreateUser(ctx context.Context, user *user_entity.User) (*UserOutputDto, error) {
	user, err := u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &UserOutputDto{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
