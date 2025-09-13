package main

import (
	"util/api"
	"util/nettools"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")
	r.GET("/", home)
	r.GET("/conn", connection)
	r.GET("/api/dig", api.Dig)
	r.GET("/api/whois", api.Whois)
	r.GET("/whois", whois)
	r.GET("/dig", dig)
	r.GET("/codedecode", codedecode)

	r.Run(":5200")
}

func home(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")

	c.HTML(200, "main.html",
		gin.H{
			"IP":        ip,
			"UserAgent": c.Request.UserAgent(),
		})
}

func connection(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	ipOrganization, err := nettools.GetIPOrganization(ip)
	if err != nil {
		ipOrganization = ""
	}
	ipCountry, err := nettools.GetIPCountry(ip)
	if err != nil {
		ipCountry = ""
	}
	c.HTML(200, "connection.html",
		gin.H{
			"IP":             ip,
			"UserAgent":      c.Request.UserAgent(),
			"IPOrganization": ipOrganization,
			"IPCountry":      ipCountry,
		})
}

func whois(c *gin.Context) {
	c.HTML(200, "whois.html", gin.H{})
}

func dig(c *gin.Context) {
	c.HTML(200, "dig.html", gin.H{})
}

func codedecode(c *gin.Context) {
	c.HTML(200, "codedecode.html", gin.H{})
}
