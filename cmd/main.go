package main

import (
	"github.com/akadotsh/go-jiosaavn-client/api"
	"github.com/charmbracelet/log"
)

func main() {
	port := ":8080"

	server := api.NewServer(port)

	log.Info("server running on port", port)
	log.Fatal(server.Start())

}
