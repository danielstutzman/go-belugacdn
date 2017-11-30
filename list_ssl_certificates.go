package belugacdn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ListSslCertificatesOutput struct {
	Count        int              `json:"count"`
	Certificates []SslCertificate `json:"certificates"`
}

type SslCertificate struct {
	Id          string   `json:"id"`
	CommonName  string   `json:"common_name"`
	Enabled     bool     `json:"enabled"`
	CName       string   `json:"cname"`
	Names       []string `json:"names"`
	Protocols   string   `json:"protocols"`
	Ciphers     string   `json:"ciphers"`
	OcspExpires string   `json:"ocsp_expires"` // e.g. "2016-09-21 14:52:40-04"
	Status      string   `json:"status"`
}

func (config *Config) ListSslCertificates() (*ListSslCertificatesOutput, error) {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet,
		"https://api.belugacdn.com/api/cdn/v2/ssl-certificates", nil)
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

	if response.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("Error from ReadAll: %s", err)
		}
		return nil, fmt.Errorf("Non-OK from API: %s", body)
	}

	output := ListSslCertificatesOutput{}
	err = json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return nil, fmt.Errorf("Error from Decode: %s", err)
	}

	return &output, nil
}
