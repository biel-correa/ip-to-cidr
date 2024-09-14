package handler

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	var requestIp = c.ClientIP()
	var ip = c.DefaultQuery("ip", "")

	if ip == "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip": requestIp,
		})
		return
	}

	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":    requestIp,
			"error": "Invalid IP address",
		})
		return
	}

	if parsedIp.To4() != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":    requestIp,
			"error": "IPv4 is not supported",
		})
		return
	}

	cidrFormat := fmt.Sprintf("%s/64", ip)
	_, cidr, _ := net.ParseCIDR(cidrFormat)

	if cidr == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":    requestIp,
			"error": "Could not parse to CIDR/64",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"ip":   requestIp,
		"cidr": cidr,
	})
}
