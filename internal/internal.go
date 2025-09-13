package internal

import (
	"os/exec"
	"strings"
	"util/security"
)

func Dig(recordHost, recordType string) (string, error) {
	cmd := exec.Command("dig", "@1.1.1.1", recordHost, recordType)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return extractAnswerSection(string(output)), nil
}

func Whois(recordHost string) (string, error) {
	cmd := exec.Command("whois", recordHost)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return security.RemoveMyIP(string(output)), nil
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
