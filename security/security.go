package security

import (
	"os"
	"regexp"
	"slices"
	"strings"
)

func GetHost() string {
	contentB, err := os.ReadFile("host.txt")
	if err != nil {
		return "null"
	}
	return strings.TrimSpace(string(contentB))
}

func GetIP() string {
	contentB, err := os.ReadFile("ip.txt")
	if err != nil {
		return "null"
	}
	print(string(contentB))
	return strings.TrimSpace(string(contentB))
}

func IsRecordTypeValid(recordType string) bool {
	validTypes := []string{"A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "MX", "NS", "PTR", "SOA", "SRV", "TLSA", "TSIG", "TXT"}
	return slices.Contains(validTypes, recordType)
}

func IsURLValid(content string) bool {
	hostnameRegex := `^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`
	valid, err := regexp.MatchString(hostnameRegex, content)
	return valid && len(content) <= 253 && err == nil
}

func RemoveMyIP(whoisOutput string) string {
	sanitizedOutput := ""
	ip := GetIP()
	for _, line := range strings.Split(whoisOutput, "\n") {
		if !strings.Contains(line, ip) {
			sanitizedOutput += line + "\n"
		}
	}
	return sanitizedOutput
}
