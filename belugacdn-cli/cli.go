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

	if false {
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

	if false {
		cert, err := config.CreateSslCertificate(
			belugacdn.CreateSslCertificateInput{
				Certificate: "cert",
				Chain:       "chain",
				Key:         "key",
				Site:        "example.com",
			})
		if err != nil {
			log.Fatalf("Error from CreateSslCertificate: %s", err)
		}
		log.Printf("Created cert: %+v", cert)
	}

	if false {
		err := config.DeleteSslCertificate("monitoring.danstutzman.com")
		if err != nil {
			log.Fatalf("Error from DeleteSslCertificate: %s", err)
		}
	}

	if false {
		certs, err := config.ListSslCertificates()
		if err != nil {
			log.Fatalf("Error from ListSslCertificates: %s", err)
		}
		log.Printf("Certs: %+v", certs)
	}
}
