package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (config *Config) list_sites() {
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet,
		"https://api.belugacdn.com/api/cdn/v2/sites", nil)
	if err != nil {
		log.Fatalf("Error from NewRequest: %s", err)
	}

	request.SetBasicAuth(config.Username, config.Password)
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error from http.Get: %s", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error from ReadAll: %s", err)
	}

	log.Printf("Got body: %s", body)
}
