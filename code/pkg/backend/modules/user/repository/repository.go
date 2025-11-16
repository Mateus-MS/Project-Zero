package user_repository

import (
	user_model "MODULE_PATH/backend/modules/user/model"
	"context"
)

type IRepository interface {
	Create(context.Context, user_model.UserEntity) error
	Read(context.Context, string) (user_model.UserEntity, error)
}
