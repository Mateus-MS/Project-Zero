package user_routes

import (
	user_service "MODULE_PATH/backend/modules/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoute(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := userService.Register(c, username, password)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "User registered into DB successfully")
	}
}
