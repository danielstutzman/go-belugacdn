package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListSitesOutput struct {
	Sites []ListSitesOutputSite `json:"sites"`
}

type ListSitesOutputSite struct {
	Name          string            `json:"name"`
	Created       string            `json:"created"`   // e.g. 2016-06-29 16:01:54.836804
	DomainId      int               `json:"domain_id"` // e.g. 1234
	Configuration SiteConfiguration `json:"configuration"`
}

func (config *Config) ListSites() (*ListSitesOutput, error) {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet,
		"https://api.belugacdn.com/api/cdn/v2/sites", nil)
	if err != nil {
		return nil, fmt.Errorf("Error from NewRequest: %s", err)
	}

	request.SetBasicAuth(config.Username, config.Password)
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Error from client.Do: %s", err)
	}
	defer response.Body.Close()

	output := ListSitesOutput{}
	err = json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return nil, fmt.Errorf("Error from Decode: %s", err)
	}

	return &output, nil
}
