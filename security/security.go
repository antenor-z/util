package security

import (
	"net/url"
	"slices"
	"strings"
	"util/nettools"
)

func IsRecordTypeValid(recordType string) bool {
	validTypes := []string{"A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "MX", "NS", "PTR", "SOA", "SRV", "TLSA", "TSIG", "TXT"}
	return slices.Contains(validTypes, recordType)
}

func IsURLValid(rawURL string) bool {
	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}

	if len(rawURL) > 2083 {
		return false
	}

	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if parsed.Host == "" {
		return false
	}

	return true
}

func GetHostname(rawURL string) string {
	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	return parsed.Hostname()
}
func RemoveMyIP(whoisOutput string) string {
	sanitizedOutput := ""
	ip := nettools.GetIP()
	for _, line := range strings.Split(whoisOutput, "\n") {
		if !strings.Contains(line, ip) {
			sanitizedOutput += line + "\n"
		}
	}
	return sanitizedOutput
}

func EscapeHTML(s string) string {
	s = strings.ReplaceAll(s, "<", "&lt")
	s = strings.ReplaceAll(s, ">", "&gt")
	return s
}
