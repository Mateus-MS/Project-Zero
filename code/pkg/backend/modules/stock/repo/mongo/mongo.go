package stock_repository_mongo

import (
	generic_repository_mongo "PLACEHOLDERPATH/backend/modules/common/repo/mongo"
	stock_model "PLACEHOLDERPATH/backend/modules/stock/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*generic_repository_mongo.GenericRepository[*stock_model.StockEntity]
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{
		GenericRepository: generic_repository_mongo.New[*stock_model.StockEntity](coll),
	}
}
