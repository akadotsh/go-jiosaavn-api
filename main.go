package main

import (
	"fmt"
	"log"

	"github.com/akadotsh/go-jiosaavn-client/api"
)

func main() {

	port := ":8080"

	server := api.NewServer(port)

	fmt.Println("server running on port", port)
	log.Fatal(server.Start())

}
