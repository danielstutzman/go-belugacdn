package main

import (
	"flag"
	"log"
)

func main() {
	config := Config{}
	flag.StringVar(&config.Username, "username", "", "Username for BelugaCDN")
	flag.StringVar(&config.Password, "password", "", "Password for BelugaCDN")
	flag.Parse()

	if config.Username == "" {
		log.Fatalf("Missing -username")
	}
	if config.Password == "" {
		log.Fatalf("Missing -password")
	}

	config.list_sites()
}
