package api

import (
	"os/exec"
	"strings"
	"util/security"

	"github.com/gin-gonic/gin"
)

func Whois(c *gin.Context) {
	recordHost := c.Query("host")
	if !security.IsURLValid(recordHost) {
		c.String(400, "invalid hostname")
		return
	}
	cmd := exec.Command("whois", recordHost)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.String(200, security.RemoveMyIP(string(output)))
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
	cmd := exec.Command("dig", "@1.1.1.1", recordHost, recordType)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.String(200, extractAnswerSection(string(output)))
}

func extractAnswerSection(s string) string {
	answerStarted := false
	answerSection := ""
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "ANSWER SECTION") {
			answerStarted = true
			continue
		}
		if answerStarted {
			if strings.Contains(line, ";;") {
				return answerSection
			}
			answerSection += line + "\n"
		}
	}
	return answerSection
}

func Ip(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	c.String(200, "%s\n", ip)
}
