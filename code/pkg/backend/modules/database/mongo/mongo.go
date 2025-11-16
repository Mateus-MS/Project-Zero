package database_mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB *mongo.Database
}

func New() *Client {
	return &Client{}
}
