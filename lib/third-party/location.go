package controllers

import (
	"encoding/json"
	"fmt"
	"gorepair-rest-api/config"
	"gorepair-rest-api/models"
	"io/ioutil"
	"log"
	"net/http"
)

// HERE RevGeocoding and Search API
func GetRevGeocode(lat, lng float64) string {
	var address string
	// load env
	config, cErr := config.LoadConfig(".")
	if cErr != nil {
		log.Fatalln("Cannot load config", cErr)
	}

	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey=" + config.HERE_API_KEY + "&at=" + fmt.Sprint(lat) + "," + fmt.Sprint(lng)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(body))
	var data models.RevOutput
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	for _, add := range data.Items {
		address = add.Title
	}
	// fmt.Println(address)
	return address
}