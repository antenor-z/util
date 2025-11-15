package middle

import (
	"fmt"
	"time"
	"util/internal"
	"util/nettools"
)

var digCache, whoisCache, ipInfoCache, ipLocation ExpirableCache

func Dig(recordHost, recordType string) (string, error) {
	cacheResult, ok := digCache.GetString(fmt.Sprintf("%s:%s", recordHost, recordType))
	if ok {
		return cacheResult, nil
	}
	internalResult, err := internal.Dig(recordHost, recordType)
	if err != nil {
		return "", err
	}
	digCache.Set(fmt.Sprintf("%s:%s", recordHost, recordType), internalResult, time.Minute)
	return internalResult, nil

}

func Whois(recordHost string) (string, error) {
	cacheResult, ok := whoisCache.GetString(recordHost)
	if ok {
		return cacheResult, nil
	}
	internalResult, err := internal.Whois(recordHost)
	if err != nil {
		return "", err
	}
	whoisCache.Set(recordHost, internalResult, time.Hour*3)
	return internalResult, nil
}

func GetIpInfo(ip string) (*nettools.IpInfo, error) {
	cacheResultAny, ok := ipInfoCache.Get(ip)
	if ok {
		cacheResult := cacheResultAny.(*nettools.IpInfo)
		return cacheResult, nil
	}
	internalResult, err := nettools.GetIpInfo(ip)
	if err != nil {
		return &nettools.IpInfo{}, err
	}
	ipInfoCache.Set(ip, internalResult, time.Hour*12)
	return internalResult, nil
}

func GetIpLocation(ip string) (*nettools.IpLocation, error) {
	cacheResultAny, ok := ipLocation.Get(ip)
	if ok {
		cacheResult := cacheResultAny.(*nettools.IpLocation)
		return cacheResult, nil
	}
	internalResult, err := nettools.GetIpLocation(ip)
	if err != nil {
		return &nettools.IpLocation{}, err
	}
	ipLocation.Set(ip, internalResult, time.Hour*12)
	return internalResult, nil
}
