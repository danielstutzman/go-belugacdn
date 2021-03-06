package belugacdn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreateSiteInput struct {
	Name          string            `json:"name"`
	Configuration SiteConfiguration `json:"configuration"`
}

type SiteConfiguration struct {
	Origin     string     `json:"origin"`
	OriginPort int        `json:"originPort"`
	Hostnames  []string   `json:"hostnames"`
	Rules      []SiteRule `json:"rules"`
}

type SiteRule struct {
	Paths   []string     `json:"paths"`
	Actions []SiteAction `json:"actions"`
}

type SiteAction struct {
	Action string `json:"action"`
	Scheme string `json:"scheme"`
}

type CreateSiteOutput struct {
	// If success:
	Name                string            `json:"name"`
	ConfigurationSource string            `json:"configuration-source"` // e.g. "internal-json"
	CreatedAt           string            `json:"created"`              // e.g. "2017-11-30 14:45:20.943849"
	DomainId            int               `json:"domain_id"`            // e.g. 16967
	SyncStatus          string            `json:"sync-status"`          // e.g. "pending"
	ReloadStatus        string            `json:"reload-status"`        // e.g. "pending"
	Status              string            `json:"status"`               // e.g. pending
	Configuration       SiteConfiguration `json:"configuration"`
	CName               string            `json:"cname"`
	// TODO: statistics, throughput, requests

	// If failure:
	Result  string `json:"result"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (config *Config) CreateSite(siteName string, newConfig SiteConfiguration) (*CreateSiteOutput, error) {
	client := &http.Client{}

	input := CreateSiteInput{Name: siteName, Configuration: newConfig}
	inputJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("Error from json.Marshal: %s", err)
	}

	request, err := http.NewRequest(http.MethodPost,
		"https://api.belugacdn.com/api/cdn/v2/sites",
		bytes.NewBuffer(inputJson))
	if err != nil {
		return nil, fmt.Errorf("Error from NewRequest: %s", err)
	}

	request.SetBasicAuth(config.Username, config.Password)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Error from client.Do: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("Error from ReadAll: %s", err)
		}
		return nil, fmt.Errorf("Non-OK from API: %s", body)
	}

	output := CreateSiteOutput{}
	err = json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return nil, fmt.Errorf("Error from Decode: %s", err)
	}

	if output.Result == "failure" {
		return nil, fmt.Errorf("Got failure with message: %s", output.Message)
	}

	return &output, nil
}
