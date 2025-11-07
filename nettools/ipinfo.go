package nettools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// curl http://ip-api.com/json/179.218.123.197
// {"status":"success","country":"Brazil","countryCode":"BR","region":"RJ","regionName":"Rio de Janeiro","city":"Rio de Janeiro","zip":"22790","lat":-22.9072,"lon":-43.1883,"timezone":"America/Sao_Paulo","isp":"Claro NXT Telecomunicacoes Ltda","org":"Claro NXT Telecomunicacoes Ltda","as":"AS28573 Claro NXT Telecomunicacoes Ltda","query":"179.218.123.197"}
type IpLocation struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	State       string  `json:"regionName"`
	City        string  `json:"city"`
	Timezone    string  `json:"timezone"`
	Latitude    float32 `json:"lat"`
	Longitude   float32 `json:"lon"`
}

func GetIpLocation(ip string) (*IpLocation, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo IpLocation
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return nil, err
	}
	return &ipInfo, nil
}
