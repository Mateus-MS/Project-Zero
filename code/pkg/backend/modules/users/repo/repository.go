package user_repository

import (
	"context"

	generic_repository "PLACEHOLDERPATH/backend/modules/common/repo"
	user_model "PLACEHOLDERPATH/backend/modules/users/model"
)

type IRepository interface {
	ReadByName(context.Context, string) (user_model.UserEntity, error)

	generic_repository.IGenericRepository[*user_model.UserEntity]
}
