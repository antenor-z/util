package main

import (
	"util/api"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")
	r.GET("/", home)
	r.GET("/conn", connection)
	r.GET("/dig", dig)
	r.GET("/whois", whois)
	r.GET("/codedecode", codedecode)
	apiGroup := r.Group("/api")
	apiGroup.GET("/dig", api.Dig)
	apiGroup.GET("/whois", api.Whois)
	apiGroup.GET("/ip", api.GetIPInfo)

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
	c.HTML(200, "connection.html",
		gin.H{
			"IP":        ip,
			"UserAgent": c.Request.UserAgent(),
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
