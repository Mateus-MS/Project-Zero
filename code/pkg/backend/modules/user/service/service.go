package user_service

import (
	user_model "MODULE_PATH/backend/modules/user/model"
	user_repository "MODULE_PATH/backend/modules/user/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userEntity = user_model.UserEntity

type IService interface {
	Register(context.Context, string, string) error

	Create(context.Context, userEntity) error
	ReadByName(context.Context, string) (userEntity, error)
	DeleteByID(context.Context, primitive.ObjectID) error
}

type service struct {
	repository *user_repository.IRepository
}

func New(repo *user_repository.IRepository) *service {
	return &service{
		repository: repo,
	}
}
