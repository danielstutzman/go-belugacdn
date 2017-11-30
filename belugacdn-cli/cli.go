package main

import (
	"flag"
	"log"

	"github.com/danielstutzman/go-belugacdn"
)

func main() {
	config := belugacdn.Config{}
	flag.StringVar(&config.Username, "username", "", "Username for BelugaCDN")
	flag.StringVar(&config.Password, "password", "", "Password for BelugaCDN")
	flag.Parse()

	if config.Username == "" {
		log.Fatalf("Missing -username")
	}
	if config.Password == "" {
		log.Fatalf("Missing -password")
	}

	if true {
		sites, err := config.ListSites()
		if err != nil {
			log.Fatalf("Error from ListSites: %s", err)
		}
		log.Printf("Sites: %+v", sites)
	}

	if false {
		site, err := config.CreateSite("images.yoursite.com",
			belugacdn.SiteConfiguration{
				Origin:    "origin-images.yoursite.com",
				Hostnames: []string{"images.yoursite.com"},
			})
		if err != nil {
			log.Fatalf("Error from CreateSite: %s", err)
		}
		log.Printf("Site: %+v", site)
	}

	if true {
		site, err := config.UpdateSite("example.danstutzman.com",
			belugacdn.SiteConfiguration{
				Origin:    "origin-images.yoursite.com",
				Hostnames: []string{"images.yoursite.com"},
			})
		if err != nil {
			log.Fatalf("Error from UpdateSite: %s", err)
		}
		log.Printf("Site: %+v", site)
	}

	if false {
		err := config.DeleteSite("images.yoursite.com")
		if err != nil {
			log.Fatalf("Error from DeleteSite: %s", err)
		}
	}
}
