package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/weileenchen/ustart_tutorial/api/rest/customer"
	//"github.com/lib/pq"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config customer.Config

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
	restAPI, err := customer.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/", nil)
	http.HandleFunc("/toggleAvailable", restAPI.ToggleAvailabke)
	http.HandleFunc("/register", restAPI.Register)
	http.HandleFunc("/search", restAPI.Search) //Not yet implemented for REST

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
