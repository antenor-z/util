package gotoURL

import (
	"errors"
	"strings"
	"time"
	"util/middle"
	"util/security"
)

type GotoUrlDto struct {
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

var gotoCache middle.ExpirableCache

func Set(gotoUrlDto GotoUrlDto) error {
	_, ok := gotoCache.GetString(gotoUrlDto.Alias)
	if ok {
		return errors.New("an alias already exists")
	}
	if !security.IsURLValid(gotoUrlDto.Url) {
		return errors.New("invalid URL")
	}
	if !strings.HasPrefix(gotoUrlDto.Url, "https://") || !strings.HasPrefix(gotoUrlDto.Url, "http://") {
		gotoUrlDto.Url = "https://" + gotoUrlDto.Url
	}
	gotoCache.Set(gotoUrlDto.Alias, gotoUrlDto.Url, time.Hour*12)
	return nil
}

func Get(alias string) (string, error) {
	str, ok := gotoCache.GetString(alias)
	if !ok {
		return "", errors.New("not found")
	}
	return str, nil
}
