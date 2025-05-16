package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", home)
	r.GET("/ip", getIp)
	r.GET("/ua", getUA)
	r.GET("/lang", getLang)
	r.GET("/encoding", getEnc)
	r.GET("/mime", getMime)
	r.GET("/charset", getCharset)
	r.GET("/forwarded", getFwd)
	r.GET("/all", getAll)
	r.GET("/all.json", getAllJSON)

	r.Run(":5200")
}

type a4Info struct {
	IpAddr     string `json:"ip_addr"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
	Port       string `json:"port"`
	Language   string `json:"language"`
	Referer    string `json:"referer"`
	Connection string `json:"connection"`
	KeepAlive  string `json:"keep_alive"`
	Method     string `json:"method"`
	Encoding   string `json:"encoding"`
	Mime       string `json:"mime"`
	Charset    string `json:"charset"`
	Via        string `json:"via"`
	Forwarded  string `json:"forwarded"`
}

type allInfo struct {
	Info    a4Info
	All     string
	AllJson string
	Root    string
}

func serializePlain(i a4Info) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ip_addr: %s\n", i.IpAddr))
	sb.WriteString(fmt.Sprintf("remote_host: %s\n", i.RemoteHost))
	sb.WriteString(fmt.Sprintf("user_agent: %s\n", i.UserAgent))
	sb.WriteString(fmt.Sprintf("port: %s\n", i.Port))
	sb.WriteString(fmt.Sprintf("language: %s\n", i.Language))
	sb.WriteString(fmt.Sprintf("referer: %s\n", i.Referer))
	sb.WriteString(fmt.Sprintf("connection: %s\n", i.Connection))
	sb.WriteString(fmt.Sprintf("keep_alive: %s\n", i.KeepAlive))
	sb.WriteString(fmt.Sprintf("method: %s\n", i.Method))
	sb.WriteString(fmt.Sprintf("encoding: %s\n", i.Encoding))
	sb.WriteString(fmt.Sprintf("mime: %s\n", i.Mime))
	sb.WriteString(fmt.Sprintf("charset: %s\n", i.Charset))
	sb.WriteString(fmt.Sprintf("via: %s\n", i.Via))
	sb.WriteString(fmt.Sprintf("forwarded: %s\n", i.Forwarded))
	return sb.String()
}

func home(c *gin.Context) {
	if !strings.Contains(c.Request.UserAgent(), "Mozilla") {
		ip := c.RemoteIP()
		c.String(200, ip)
		return
	}
	host := c.Request.Host
	info := getInfo(c)
	var infoJsonStr string
	infoJsonBytes, err := json.Marshal(info)
	if err != nil {
		infoJsonStr = ""
	} else {
		infoJsonStr = string(infoJsonBytes)
	}

	a := allInfo{
		Info:    info,
		Root:    host,
		All:     serializePlain(info),
		AllJson: infoJsonStr,
	}

	c.HTML(200, "main.html", a)
}

func getIp(c *gin.Context) {
	ip := c.Request.Header.Get("CF-Connecting-IP")
	c.String(200, "%s\n", ip)
}

func getUA(c *gin.Context) {
	ip := c.Request.UserAgent()
	c.String(200, "%s\n", ip)
}

func getLang(c *gin.Context) {
	ip := c.Request.Header.Get("Accept-Language")
	c.String(200, "%s\n", ip)
}

func getEnc(c *gin.Context) {
	ip := c.Request.Header.Get("Accept-Encoding")
	c.String(200, "%s\n", ip)
}

func getMime(c *gin.Context) {
	ip := c.Request.Header.Get("Accept")
	c.String(200, "%s\n", ip)
}

func getCharset(c *gin.Context) {
	ip := c.Request.Header.Get("Accept-Charset")
	c.String(200, "%s\n", ip)
}

func getFwd(c *gin.Context) {
	ip := c.Request.Header.Get("forwarded")
	c.String(200, "%s\n", ip)
}

func getInfo(c *gin.Context) a4Info {
	vals := slices.Collect(maps.Values(c.Request.Header))
	fmt.Println(vals)
	return a4Info{
		IpAddr:     c.Request.Header.Get("CF-Connecting-IP"),
		RemoteHost: c.Request.RemoteAddr,
		UserAgent:  c.Request.UserAgent(),
		Port:       c.Request.URL.Port(),
		Language:   c.Request.Header.Get("Accept-Language"),
		Referer:    c.Request.Header.Get("Referer"),
		Connection: c.Request.Header.Get("Connection"),
		KeepAlive:  c.Request.Header.Get("Keep-Alive"),
		Method:     c.Request.Method,
		Encoding:   c.Request.Header.Get("Accept-Encoding"),
		Mime:       c.Request.Header.Get("Accept"),
		Charset:    c.Request.Header.Get("Accept-Charset"),
		Via:        c.Request.Header.Get("via"),
		Forwarded:  c.Request.Header.Get("X-Forwarded-For"),
	}
}

func getAll(c *gin.Context) {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(200, serializePlain(getInfo(c)))
}

func getAllJSON(c *gin.Context) {
	c.JSON(200, getInfo(c))
}
