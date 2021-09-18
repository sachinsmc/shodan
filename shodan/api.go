package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UsageLimits struct {
	ScanCredits  int `json:"scan_credits"`
	QueryCredits int `json:"query_credits"`
	MonitoredIps int `json:"monitored_ips"`
}

type APIInfo struct {
	ScanCredits  int         `json:"scan_credits"`
	UsageLimits  UsageLimits `json:"usage_limits"`
	Plan         string      `json:"plan"`
	HTTPS        bool        `json:"https"`
	Unlocked     bool        `json:"unlocked"`
	QueryCredits int         `json:"query_credits"`
	MonitoredIps int         `json:"monitored_ips"`
	UnlockedLeft int         `json:"unlocked_left"`
	Telnet       bool        `json:"telnet"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	var apiInfo APIInfo
	if err := json.NewDecoder(resp.Body).Decode(&apiInfo); err != nil {
		fmt.Println("🚀 ~ file: api.go ~ line 33 ~ iferr:=json.NewDecoder ~ err : ", err)
		return nil, err
	}
	return &apiInfo, nil
}

func (s *Client) HostSearch(query string) (*HostSearch, error) {
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, s.apiKey, query))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	var hostSearch HostSearch
	if err := json.NewDecoder(resp.Body).Decode(&hostSearch); err != nil {
		fmt.Println("🚀 ~ file: api.go ~ line 53 ~ iferr:=json.NewDecoder ~ err : ", err)
		return nil, err
	}
	return &hostSearch, nil

}
