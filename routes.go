package main

import (
	"util/api"
	gotoURL "util/gotoUrl"
	"util/middle"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	if !api.IsGraphicalBrowser(c) {
		ip := api.Ip(c)
		c.String(200, ip)
		return
	}
	ip := api.Ip(c)

	c.HTML(200, "main.html",
		gin.H{
			"IP":        ip,
			"UserAgent": c.Request.UserAgent(),
		})
}

func connection(c *gin.Context) {
	ip := api.Ip(c)
	ipLocation, _ := middle.GetIpLocation(ip)
	ipInfo, _ := middle.GetIpInfo(ip)
	c.HTML(200, "connection.html",
		gin.H{
			"IP":           ip,
			"UserAgent":    c.Request.UserAgent(),
			"RemoteHost":   c.Request.RemoteAddr,
			"Port":         c.Request.URL.Port(),
			"Language":     c.Request.Header.Get("Accept-Language"),
			"Referer":      c.Request.Header.Get("Referer"),
			"Connection":   c.Request.Header.Get("Connection"),
			"KeepAlive":    c.Request.Header.Get("Keep-Alive"),
			"Method":       c.Request.Method,
			"Encoding":     c.Request.Header.Get("Accept-Encoding"),
			"Mime":         c.Request.Header.Get("Accept"),
			"Charset":      c.Request.Header.Get("Accept-Charset"),
			"Via":          c.Request.Header.Get("via"),
			"Forwarded":    c.Request.Header.Get("X-Forwarded-For"),
			"Organization": ipInfo.Organization,
			"Country":      ipInfo.Country,
			"City":         ipLocation.City,
			"State":        ipLocation.State,
			"Timezone":     ipLocation.Timezone,
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

func getQR(c *gin.Context) {
	c.HTML(200, "qr.html", gin.H{})
}

func URLShortenerPage(c *gin.Context) {
	c.HTML(200, "goto.html", gin.H{})
}

func goTo(c *gin.Context) {
	alias := c.Param("alias")
	url, err := gotoURL.Get(alias)
	if err != nil {
		c.Redirect(302, "/goto")
	}
	c.Redirect(302, url)
}
