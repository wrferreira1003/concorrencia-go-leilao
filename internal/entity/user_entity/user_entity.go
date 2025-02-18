package user_entity

import (
	"context"

	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
)

type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type UserRepositoryInterface interface {
	FindUserByID(ctx context.Context, userID string) (*User, *internal_error.InternalError)
}
