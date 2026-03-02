package user_usecase

import (
	"context"

	"github.com/katianemiranda/leilao/internal/entity/user_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
)

type UseruseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCaseInterface interface {
}

func (u *UseruseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   userEntity.Id,
		Name: userEntity.Nome,
	}, nil
}
