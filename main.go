package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Server
	http.HandleFunc("/v1/users", getUser)
	http.HandleFunc("/v1/workshops", getWorkshop)
	fmt.Println("Starting REST API web server at http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}

// GoRepair REST API
type User struct {
	Id       int
	Email    string
	Password string
	Address  string
}

type Workshop struct {
	Id       int
	Email    string
	Password string
	Address  string
}

var user = []User{
	{1, "axe@gmail.com", "axe123", geolocation(42.36399, -71.05493)},
	{2, "kunkka@gmail.com", "kunkka123", geolocation(37.4224764, -122.0842499)},
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

var workshop = []Workshop{
	{1, "bengkelaxe@gmail.com", "bengkelaxe123", geolocation(42.36399, -71.05493)},
	{2, "bengkelkunkka@gmail.com", "bengkelkunkka123", geolocation(37.4224764, -122.0842499)},
}

func getWorkshop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(workshop)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// HERE Geocoding and Search API
const API_KEY = "ydHGJC-tW01mfH8P68eLNt9HHdA4wG_YTYtnr0qT-1M" //"HERE_API_KEY"

type output struct {
    Items []struct {
        Title string `json:"title"`
    } `json:"items"`
}

func geolocation(lat, lng float64) string {
	var address string
	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey=" + API_KEY + "&at=" + fmt.Sprint(lat) + "," + fmt.Sprint(lng)
    res, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err)
    }
    // fmt.Println(string(body))
	var data output
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	for _, add := range data.Items {
		address = add.Title
	}
	fmt.Println(address)
	return address
}