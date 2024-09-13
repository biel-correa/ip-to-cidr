package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	var requestIp = c.ClientIP()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"ip": requestIp,
	})
}
