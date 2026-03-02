package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/katianemiranda/leilao/configuration/logger"
	"github.com/katianemiranda/leilao/internal/entity/user_entity"
	"github.com/katianemiranda/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {

	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserById(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {
	fillter := bson.M{"_id": id}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, fillter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found this id = %s", id), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("User not found this id = %s", id))
		}

		logger.Error("Error trying to find user by userId", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by userId")
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Nome: userEntityMongo.Name,
	}

	return userEntity, nil
}
