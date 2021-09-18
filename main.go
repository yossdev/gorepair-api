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
	v1 := e.Group("v1/api/")
	
	// User routes
	v1.GET("users", GetUsers)
	v1.GET("user/:id", UserDetails)
	v1.POST("sign-up/user", UserRegister)
	v1.POST("sign-in/user", UserLogin)

	// Workshop routes
	v1.GET("workshops", GetWorkshops)
	v1.GET("workshop/:id", WorkshopDetails)
	v1.GET("workshops/find", FindWorkshop)
	v1.POST("sign-up/workshop", WorkshopRegister)
	v1.POST("sign-in/workshop", WorkshopLogin)

	e.Start(":8000")
}

// GoRepair REST API
type BaseResponse struct {
	Code 	int
	Message string
	Data 	interface{}
}

type User struct {
	Id       int	`json:"id" form:"id"`
	Email    string	`json:"email" form:"email"`
	Password string	`json:"password" form:"password"`
	Address  string	`json:"address" form:"address"`
}

type Workshop struct {
	Id       int	`json:"id" form:"id"`
	Email    string	`json:"email" form:"email"`
	Password string	`json:"password" form:"password"`
	Address  string	`json:"address" form:"address"`
}

type Login struct {
	Email    string	`json:"email" form:"email"`
	Password string	`json:"password" form:"password"`
}

type SignUp struct {
	Name 	 string `json:"name" form:"name"`
	Email    string	`json:"email" form:"email"`
	Password string	`json:"password" form:"password"`
}

// hardcoded data
var user = []User{
	{1, "axe@gmail.com", "axe123", geolocation(42.36399, -71.05493)},
	{2, "kunkka@gmail.com", "kunkka123", geolocation(37.4224764, -122.0842499)},
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: user,
	})
}

func UserDetails(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: "Id is not valid",
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: User{Id: id}, // will check to database for now it will be empty
	})
}

func UserRegister(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user SignUp
	user.Name = name
	user.Email = email
	user.Password = password

	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: user,
	})
}

func UserLogin(c echo.Context) error {
	login := Login{}
	c.Bind(&login)
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: login,
	})
}

// hardcoded data
var workshop = []Workshop{
	{1, "bengkelaxe@gmail.com", "bengkelaxe123", geolocation(42.36399, -71.05493)},
	{2, "bengkelkunkka@gmail.com", "bengkelkunkka123", geolocation(37.4224764, -122.0842499)},
}

func GetWorkshops(c echo.Context) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: workshop,
	})
}

func WorkshopDetails(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: "Id is not valid",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: Workshop{Id: id}, // will check to database for now it will be empty
	})
}

func FindWorkshop(c echo.Context) error {
	match := c.QueryParam("name")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name": match,
		"result": []string{"jaya bengkel", "honda", "suzuki"}, // hardcoded data for now
	})
}

func WorkshopRegister(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var workshop SignUp
	workshop.Name = name
	workshop.Email = email
	workshop.Password = password

	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: workshop,
	})
}

func WorkshopLogin(c echo.Context) error {
	login := Login{}
	c.Bind(&login)
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: login,
	})
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