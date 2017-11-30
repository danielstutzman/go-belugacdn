package belugacdn

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (config *Config) DeleteSslCertificate(commonName string) error {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodDelete,
		"https://api.belugacdn.com/api/cdn/v2/ssl-certificates/"+commonName,
		nil)
	if err != nil {
		return fmt.Errorf("Error from NewRequest: %s", err)
	}

	request.SetBasicAuth(config.Username, config.Password)
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error from client.Do: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("Error from ReadAll: %s", err)
		}
		return fmt.Errorf("Non-OK from API: %s", body)
	}

	return nil
}
