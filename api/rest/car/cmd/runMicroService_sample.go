package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/weileenchen/ustart_tutorial/api/rest/car"
	//prof "github.com/sea350/ustart_tutorial/customer"
	//"github.com/sea350/ustart_tutorial/customer/storage"
	//elasticstore "github.com/sea350/ustart_tutorial/customer/storage/elastic"
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
<<<<<<< HEAD
	http.HandleFunc("/toggleAvailable", restAPI.toggleAvailabke)
=======
	http.HandleFunc("/toggleAvailable", restAPI.ToggleAvailable)
>>>>>>> 373e2289441a2399564408571f86c1eb70feb9d3
	http.HandleFunc("/register", restAPI.Register)
	http.HandleFunc("/search", nil) //Not yet implemented for REST

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
