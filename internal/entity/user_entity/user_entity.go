package user_entity

import (
	"context"

	"github.com/katianemiranda/leilao/internal/internal_error"
)

type User struct {
	Id   string
	Nome string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
