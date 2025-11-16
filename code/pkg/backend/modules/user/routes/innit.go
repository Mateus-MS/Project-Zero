package user_routes

import (
	user_service "MODULE_PATH/backend/modules/user/service"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, userService *user_service.IService) {
	router.GET("/users/register", registerRoute(*userService))
}
