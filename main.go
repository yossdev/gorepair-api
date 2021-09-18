package main

import (
	"encoding/json"
	"fmt"
	"gorepair-rest-api/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	// Server
	e := echo.New()
	
	// User routes
	e.GET("/v1/users", GetUser)
	e.GET("/v1/users/:id", GetUser)

	// Workshop routes
	e.GET("/v1/workshops", GetWorkshop)
	e.GET("/v1/workshops/:id", GetWorkshop)

	fmt.Println("Starting REST API web server at http://localhost:8000/")
	e.Start(":8000")
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

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range user {
		if user[i].Id == id {
			return c.JSON(http.StatusOK, user[i])
		}
	}
	return c.JSON(http.StatusOK, "User not found")
}

var workshop = []Workshop{
	{1, "bengkelaxe@gmail.com", "bengkelaxe123", geolocation(42.36399, -71.05493)},
	{2, "bengkelkunkka@gmail.com", "bengkelkunkka123", geolocation(37.4224764, -122.0842499)},
}

func GetWorkshop(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range workshop {
		if workshop[i].Id == id {
			return c.JSON(http.StatusOK, workshop[i])
		}
	}
	return c.JSON(http.StatusOK, "Workshop not found")
}

// HERE Geocoding and Search API
type output struct {
    Items []struct {
        Title string `json:"title"`
    } `json:"items"`
}

func geolocation(lat, lng float64) string {
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
	var data output
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	for _, add := range data.Items {
		address = add.Title
	}
	// fmt.Println(address)
	return address
}