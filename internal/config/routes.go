package config

import (
	"app/api/handler"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine) {
	router.GET("/", handler.IndexHandler)
}
