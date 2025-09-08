package main

import (
	"util/api"
	"util/security"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", home)
	r.GET("/api/dig", api.Dig)
	r.GET("/api/whois", api.Whois)
	r.GET("/whois", whois)
	r.GET("/dig", dig)
	r.GET("/codedecode", codedecode)

	r.Run(":5200")
}

func home(c *gin.Context) {
	c.HTML(200, "main.html", gin.H{"IP": security.GetIP(), "UserAgent": c.Request.UserAgent()})
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
