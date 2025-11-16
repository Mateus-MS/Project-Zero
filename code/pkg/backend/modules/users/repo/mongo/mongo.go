package user_repository_mongo

import (
	generic_repository_mongo "PLACEHOLDERPATH/backend/modules/common/repo/mongo"
	user_model "PLACEHOLDERPATH/backend/modules/users/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*generic_repository_mongo.GenericRepository[*user_model.UserEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository_mongo.New[*user_model.UserEntity](coll),
	}
}
