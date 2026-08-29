package main

import (
	"github.com/akadotsh/go-jiosaavn-client/api"
	"charm.land/log/v2"
)

func main() {
	port := ":8080"

	server := api.NewServer(port)

	log.Info("server running on port", port)
	log.Fatal(server.Start())
}
