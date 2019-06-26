package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/weileenchen/ustart_tutorial/api/rest/car"
	"github.com/weileenchen/ustart_tutorial/car"
	//"github.com/lib/pq"
	//prof "github.com/weileenchen/ustart_tutorial/car"
	//"github.com/weileenchen/ustart_tutorial/car/storage"
	//elasticstore "github.com/weileenchen/ustart_tutorial/car/storage/elastic"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config car.Config

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
	restAPI, err := car.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/", nil)
	http.HandleFunc("/toggleAvailable", restAPI.ToggleAvailable)
	http.HandleFunc("/register", restAPI.Register)
	http.HandleFunc("/search", nil) //Not yet implemented for REST

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
