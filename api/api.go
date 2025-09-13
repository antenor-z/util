package api

import (
	"util/internal"
	"util/nettools"
	"util/security"

	"github.com/gin-gonic/gin"
)

func Whois(c *gin.Context) {
	recordHost := c.Query("host")
	if !security.IsURLValid(recordHost) {
		c.String(400, "invalid hostname")
		return
	}
	response, err := internal.Whois(recordHost)
	if err != nil {
		c.String(400, "unknown error")
		return
	}

	c.String(200, response)
}

func Dig(c *gin.Context) {
	recordType := c.Query("type")
	recordHost := c.Query("host")
	if !security.IsRecordTypeValid(recordType) {
		c.String(400, "invalid DNS record")
		return
	}
	if !security.IsURLValid(recordHost) {
		c.String(400, "invalid hostname")
		return
	}
	response, err := internal.Dig(recordHost, recordType)
	if err != nil {
		c.String(400, "unknown error")
		return
	}

	c.String(200, response)
}

func Ip(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	c.String(200, "%s\n", ip)
}

func GetIPOrganization(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	organization, err := nettools.GetIPOrganization(ip)
	if err != nil {
		c.String(400, "unknown error")
		return
	}
	c.String(200, organization)
}

func GetIPCountry(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	country, err := nettools.GetIPOrganization(ip)
	if err != nil {
		c.String(400, "unknown error")
		return
	}
	c.String(200, country)
}
