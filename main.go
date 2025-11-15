package main

import (
	"util/api"
	"util/nettools"

	"github.com/gin-gonic/gin"
)

func main() {
	nettools.GetIP()
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
	r.GET("/g/:alias", goTo)
	r.GET("/goto", URLShortenerPage)
	r.GET("qr", getQR)
	apiGroup := r.Group("/api")
	apiGroup.GET("/dig", api.Dig)
	apiGroup.GET("/whois", api.Whois)
	apiGroup.POST("/note", api.PostNote)
	apiGroup.GET("/note/:noteId", api.GetNote)
	apiGroup.GET("/qr", api.GetQRCode)
	apiGroup.POST("/goto", api.URLShortener)

	r.Run(":5200")
}
