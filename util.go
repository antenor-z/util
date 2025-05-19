package main

import (
	"os"
	"strings"
)

func GetHost() string {
	contentB, err := os.ReadFile("host.txt")
	if err != nil {
		return "null"
	}
	return strings.TrimSpace(string(contentB))
}
