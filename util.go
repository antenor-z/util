package main

import "os"

func GetHost() string {
	contentB, err := os.ReadFile("host.txt")
	if err != nil {
		return "null"
	}
	return string(contentB)
}
