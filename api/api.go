package api

import (
	"strings"
	gotoURL "util/gotoUrl"
	"util/middle"
	"util/note"
	"util/qr"
	"util/security"

	"github.com/gin-gonic/gin"
)

func Whois(c *gin.Context) {
	recordHost := c.Query("host")
	if !security.IsURLValid(recordHost) {
		c.String(400, "invalid hostname")
		return
	}
	recordHost = security.GetHostname(recordHost)
	response, err := middle.Whois(recordHost)
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
	recordHost = security.GetHostname(recordHost)
	response, err := middle.Dig(recordHost, recordType)
	if err != nil {
		c.String(400, "unknown error")
		return
	}

	c.String(200, response)
}

func IsGraphicalBrowser(c *gin.Context) bool {
	return strings.Contains(c.Request.UserAgent(), "Mozilla")
}

func Ip(c *gin.Context) string {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	return ip
}

func GetIPInfo(c *gin.Context) {
	ip := Ip(c)
	ipInfo, err := middle.GetIpInfo(ip)
	if err != nil {
		c.String(400, "unknown error")
		return
	}
	c.JSON(200, ipInfo)
}

func PostNote(c *gin.Context) {
	var newNote note.NoteDto
	err := c.ShouldBindBodyWithJSON(&newNote)
	if err != nil {
		c.String(400, "invalid parameters")
		return
	}
	err = note.CreateNote(newNote)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, "ok")
}

func GetNote(c *gin.Context) {
	noteId := c.Param("noteId")
	noteFromCache, err := note.GetNote(noteId)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	c.JSON(200, noteFromCache)
}

func GetQRCode(c *gin.Context) {
	text := c.Query("text")
	png, err := qr.QrGen(text)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.Data(200, "image/png", png)
}

func URLShortener(c *gin.Context) {
	var newGoto gotoURL.GotoUrlDto
	err := c.ShouldBindBodyWithJSON(&newGoto)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	err = gotoURL.Set(newGoto)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(200, "ok")
}
