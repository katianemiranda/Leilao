package user_usecase

import (
	"context"

	"github.com/katianemiranda/leilao/internal/entity/user_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
)

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   userEntity.Id,
		Name: userEntity.Nome,
	}, nil
}
