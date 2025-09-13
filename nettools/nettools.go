package nettools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RDAPResponse struct {
	Country  string   `json:"country"`
	Entities []Entity `json:"entities"`
	Name     string   `json:"name"`
}

type Entity struct {
	VCardArray []any `json:"vcardArray"`
}

func parseVCard(vcardArray []any) (orgName string, country string) {
	if len(vcardArray) != 2 {
		return "", ""
	}
	entries, ok := vcardArray[1].([]any)
	if !ok {
		return "", ""
	}
	for _, e := range entries {
		entry, ok := e.([]any)
		if !ok || len(entry) < 4 {
			continue
		}
		key := entry[0].(string)
		switch key {
		case "fn":
			if val, ok := entry[3].(string); ok {
				orgName = val
			}
		case "adr":
			if obj, ok := entry[1].(map[string]any); ok {
				if label, ok := obj["label"].(string); ok {
					parts := strings.Split(strings.TrimSpace(label), "\n")
					if len(parts) > 0 {
						country = strings.TrimSpace(parts[len(parts)-1])
					}
				}
			}
		}
	}
	return orgName, country
}

func fetchRDAP(ip string) (*RDAPResponse, error) {
	url := fmt.Sprintf("https://rdap.arin.net/registry/ip/%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rdap RDAPResponse
	if err := json.Unmarshal(body, &rdap); err != nil {
		return nil, err
	}
	return &rdap, nil
}

type IpInfo struct {
	Organization string `json:"organization"`
	Country      string `json:"country"`
}

func GetIpInfo(ip string) (*IpInfo, error) {
	rdap, err := fetchRDAP(ip)
	if err != nil {
		return nil, err
	}
	ipOrganization, err := getIPOrganization(rdap, ip)
	if err != nil {
		return nil, err
	}
	ipCountry, err := getIPCountry(rdap, ip)
	if err != nil {
		return nil, err
	}
	return &IpInfo{
		Organization: ipOrganization,
		Country:      ipCountry,
	}, nil
}

func getIPOrganization(rdap *RDAPResponse, ip string) (string, error) {
	if len(rdap.Entities) > 0 {
		org, _ := parseVCard(rdap.Entities[0].VCardArray)
		if org != "" {
			return org, nil
		}
	}

	if rdap.Name != "" {
		return rdap.Name, nil
	}
	return "", fmt.Errorf("organization not found for IP %s", ip)
}

func getIPCountry(rdap *RDAPResponse, ip string) (string, error) {
	if rdap.Country != "" {
		return rdap.Country, nil
	}

	if len(rdap.Entities) > 0 {
		_, country := parseVCard(rdap.Entities[0].VCardArray)
		if country != "" {
			return country, nil
		}
	}

	return "", fmt.Errorf("country not found for IP %s", ip)
}
