package user_repository_mongo

import (
	database_mongo "MODULE_PATH/backend/modules/database/mongo"
	user_model "MODULE_PATH/backend/modules/user/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	Collection *mongo.Collection
}

func New(client *database_mongo.Client) *repository {
	return &repository{
		Collection: client.DB.Collection("users"),
	}
}

func (repo *repository) Create(ctx context.Context, user user_model.UserEntity) error {
	_, err := repo.Collection.InsertOne(ctx, user)
	return err
}

func (repo *repository) Read(ctx context.Context, id string) (*user_model.UserEntity, error) {
	var user user_model.UserEntity
	err := repo.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return &user, err
}
