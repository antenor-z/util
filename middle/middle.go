package middle

import (
	"fmt"
	"time"
	"util/internal"
)

var expirableCache ExpirableCache

func Dig(recordHost, recordType string) (string, error) {
	cacheResult, ok := expirableCache.Get(fmt.Sprintf("DIG:%s:%s", recordHost, recordType))
	if ok {
		return cacheResult, nil
	}
	internalResult, err := internal.Dig(recordHost, recordType)
	if err != nil {
		return "", err
	}
	expirableCache.Set(fmt.Sprintf("DIG:%s:%s", recordHost, recordType), internalResult, time.Minute)
	return internalResult, nil

}

func Whois(recordHost string) (string, error) {
	cacheResult, ok := expirableCache.Get(fmt.Sprintf("WHOIS:%s", recordHost))
	if ok {
		return cacheResult, nil
	}
	internalResult, err := internal.Whois(recordHost)
	if err != nil {
		return "", err
	}
	expirableCache.Set(fmt.Sprintf("WHOIS:%s", recordHost), internalResult, time.Hour*3)
	return internalResult, nil
}
