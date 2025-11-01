package main

import (
	"util/api"
	"util/middle"

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
	r.GET("/beautify", beautify)
	r.GET("/wakelock", wakelock)
	r.GET("/note", postNote)
	r.GET("/note/:noteId", getNote)
	apiGroup := r.Group("/api")
	apiGroup.GET("/dig", api.Dig)
	apiGroup.GET("/whois", api.Whois)
	apiGroup.GET("/ip", api.GetIPInfo)
	apiGroup.POST("/note", api.PostNote)
	apiGroup.GET("/note/:noteId", api.GetNote)

	r.Run(":5200")
}

func home(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	ipInfo, _ := middle.GetIpInfo(ip)

	c.HTML(200, "main.html",
		gin.H{
			"IP":           ip,
			"UserAgent":    c.Request.UserAgent(),
			"Organization": ipInfo.Organization,
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

func beautify(c *gin.Context) {
	c.HTML(200, "beautify.html", gin.H{})
}

func wakelock(c *gin.Context) {
	c.HTML(200, "wakelock.html", gin.H{})
}

func postNote(c *gin.Context) {
	c.HTML(200, "note.html", gin.H{})
}

func getNote(c *gin.Context) {
	c.HTML(200, "note_view.html", gin.H{})
}
