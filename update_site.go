package belugacdn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UpdateSiteInput struct {
	Configuration SiteConfiguration `json:"configuration"`
}

func (config *Config) UpdateSite(siteName string, newConfig SiteConfiguration) (*CreateSiteOutput, error) {
	client := &http.Client{}

	input := UpdateSiteInput{Configuration: newConfig}

	inputJson, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("Error from json.Marshal: %s", err)
	}

	request, err := http.NewRequest(http.MethodPut,
		"https://api.belugacdn.com/api/cdn/v2/sites/"+siteName,
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

	return &output, nil
}
