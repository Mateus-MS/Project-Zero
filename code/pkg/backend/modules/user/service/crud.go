package user_service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *service) ReadByName(context.Context, string) (userEntity, error) {
	return userEntity{}, nil
}

func (svc *service) DeleteByID(context.Context, primitive.ObjectID) error {
	return nil
}
