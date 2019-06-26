package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sea350/ustart_tutorial/api/rest/customer"
	//prof "github.com/sea350/ustart_tutorial/customer"
	//"github.com/sea350/ustart_tutorial/customer/storage"
	//elasticstore "github.com/sea350/ustart_tutorial/customer/storage/elastic"
)

func main() {
	//prints into terminal but gives additional options
	//set log flags to the file
	//lshortfile -- prints logname as well as number
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

	//there for posterity
	http.HandleFunc("/", nil)	//route endpoint -- instead of returning an error response it returns nil


	http.HandleFunc("/pull", restAPI.Pull)
	http.HandleFunc("/register", restAPI.Register)
	http.HandleFunc("/search", nil) //Not yet implemented for REST

	//Hear and handle
	//runs the restapi on the config port
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
