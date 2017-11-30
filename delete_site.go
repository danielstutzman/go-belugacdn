package belugacdn

import (
	"fmt"
	"net/http"
)

func (config *Config) DeleteSite(siteName string) error {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodDelete,
		"https://api.belugacdn.com/api/cdn/v2/sites/"+siteName,
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

	return nil
}
