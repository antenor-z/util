package middle

import (
	"fmt"
	"time"
	"util/internal"
	"util/nettools"
)

var expirableCache ExpirableCache

func Dig(recordHost, recordType string) (string, error) {
	cacheResult, ok := expirableCache.GetString(fmt.Sprintf("DIG:%s:%s", recordHost, recordType))
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
	cacheResult, ok := expirableCache.GetString(fmt.Sprintf("WHOIS:%s", recordHost))
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

func GetIpInfo(ip string) (*nettools.IpInfo, error) {
	cacheResultAny, ok := expirableCache.Get(fmt.Sprintf("IPINFO:%s", ip))
	if ok {
		cacheResult := cacheResultAny.(*nettools.IpInfo)
		return cacheResult, nil
	}
	internalResult, err := nettools.GetIpInfo(ip)
	if err != nil {
		return &nettools.IpInfo{}, err
	}
	expirableCache.Set(fmt.Sprintf("IPINFO:%s", ip), internalResult, time.Hour*12)
	return internalResult, nil
}

func GetIpLocation(ip string) (*nettools.IpLocation, error) {
	cacheResultAny, ok := expirableCache.Get(fmt.Sprintf("IPLOCATION:%s", ip))
	if ok {
		cacheResult := cacheResultAny.(*nettools.IpLocation)
		return cacheResult, nil
	}
	internalResult, err := nettools.GetIpLocation(ip)
	if err != nil {
		return &nettools.IpLocation{}, err
	}
	expirableCache.Set(fmt.Sprintf("IPLOCATION:%s", ip), internalResult, time.Hour*12)
	return internalResult, nil
}
