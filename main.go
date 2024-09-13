package main

import (
	config "app/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	config.CreateRoutes(router)
	router.Run(":8080")
}
