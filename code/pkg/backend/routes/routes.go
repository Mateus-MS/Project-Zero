package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "alive")
	})
}
