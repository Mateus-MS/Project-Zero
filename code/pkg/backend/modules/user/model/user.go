package user_model

import (
	"time"

	// TODO: remove this mongo dependency
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(username, password string) *UserEntity {
	user := UserEntity{
		Name:     username,
		Password: password,
	}

	user.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	return &user
}
