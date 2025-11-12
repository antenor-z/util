package nettools

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var externalIp string

func GetIP() string {
	if len(externalIp) >= 6 {
		return externalIp
	}
	req, _ := http.NewRequest("GET", "https://ifconfig.me", nil)
	req.Header.Set("User-Agent", "curl/7.64.1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	externalIp = strings.TrimSpace(string(body))

	fmt.Printf("External IP: %s", externalIp)

	return externalIp
}
