package main

import (
	"log"
	"os"

	"github.com/rlaskowski/myip"
	"github.com/rlaskowski/myip/config"
)

func main() {
	myip.Flags()

	config, err := config.ConfigFile(myip.ConfigPath)
	if err != nil {
		log.Fatalf("Unexpected error: %s", err.Error())
		os.Exit(1)
	}

	srv := myip.NewService(config)
	srv.Run()
}
