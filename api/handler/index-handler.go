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

	cidr, error := convertToCidr64(ip)
	if error != "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":    requestIp,
			"error": error,
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"ip":   requestIp,
		"cidr": cidr,
	})
}

func convertToCidr64(ip string) (*net.IPNet, string) {
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return nil, "Invalid IP address"
	}

	if parsedIp.To4() != nil {
		return nil, "IPv4 is not supported"
	}

	cidrFormat := fmt.Sprintf("%s/64", ip)
	_, cidr, _ := net.ParseCIDR(cidrFormat)

	if cidr == nil {
		return nil, "Could not parse to CIDR/64"
	}

	return cidr, ""
}
