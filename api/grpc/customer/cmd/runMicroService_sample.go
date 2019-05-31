package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
	customerapi "github.com/sea350/ustart_tutorial/api/grpc/customer"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config customerapi.Config
	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	//Generating api object
	customerService, err := customerapi.New(&config)
	if err != nil {
		panic(err)
	}

	customerService.Run()
}
