package handler

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	var requestIp = c.ClientIP()
	var ip = c.DefaultQuery("ip", "")
	var subnetMask = c.DefaultQuery("subnetMask", "32")

	if ip == "" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip": requestIp,
		})
		return
	}

	cidr, error := convertToCidr64(ip, subnetMask)
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

func convertToCidr64(ip string, subnetMask string) (*net.IPNet, string) {
	parsedSubnetMask, err := strconv.Atoi(subnetMask)
	if err != nil || parsedSubnetMask < 0 || parsedSubnetMask > 128 {
		return nil, "Invalid subnet mask"
	}

	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return nil, "Invalid IP address"
	}

	if parsedIp.To4() != nil && parsedSubnetMask > 32 {
		return nil, "Invalid subnet mask for IPv4. Max 32"
	}

	cidrFormat := fmt.Sprintf("%s/%s", ip, subnetMask)
	_, cidr, _ := net.ParseCIDR(cidrFormat)

	if cidr == nil {
		return nil, fmt.Sprintf("Could not parse to CIDR/%s", subnetMask)
	}

	return cidr, ""
}
