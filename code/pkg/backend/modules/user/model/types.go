package user_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	// TODO: remove this mongo dependency
	ID       primitive.ObjectID `json:"ID" bson:"_id"`
	Name     string             `json:"name"         binding:"required" bson:"name"`
	Password string             `json:"password"     binding:"required" bson:"password"`
	IsAdmin  bool               `json:"isAdmin"                         bson:"isAdmin"`
}
