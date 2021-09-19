package main

import (
	"encoding/json"
	"fmt"
	"gorepair-rest-api/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	InitDB() 
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
	Code    int			`json:"code"`
	Message string		`json:"message"`
	Data    interface{}	`json:"data"`
}

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SignUp struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone	 string	`json:"phone" form:"phone"`
}

type User struct {
	ID         	uint64         	 `gorm:"primaryKey; autoIncrement" json:"id"`
	Email      	string         	 `gorm:"size:255; unique; not null" json:"email" form:"email"`
	Password   	string         	 `gorm:"size:255; not null" json:"password" form:"password"`
	Name       	string         	 `gorm:"size:125; not null" json:"name" form:"name"`
	Gender     	string         	 `gorm:"size:1" json:"gender" form:"gender"`
	DOB        	datatypes.Date	 `json:"dob" form:"dob"`
	Phone		string			 `gorm:"size:13; not null" json:"phone" form:"phone"`
	Address 	UserAddress      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
	Orders		[]Order			 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`	
	Ratings		[]UserRating 	 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ratings"`	
	CreatedAt 	time.Time		 `json:"createdAt"`
	UpdatedAt 	time.Time		 `json:"updatedAt"`
}

type UserAddress struct {
	ID 				uint64 			`gorm:"primaryKey; autoIncrement" json:"id"`
	UserID 			uint64			`json:"userId"`
	BuildingNumber 	uint16			`json:"buildingNumber" form:"buildingNumber"`
	Street 			string			`gorm:"size:255" json:"street" form:"street"`
	City 			string			`gorm:"size:50" json:"city" form:"city"`
	CountryCode 	string			`gorm:"size:5" json:"countryCode" form:"countryCode"`
	PostalCode 		string			`gorm:"size:10" json:"postalCode" form:"postalCode"`
	Province		string			`gorm:"size:50" json:"province" form:"province"`
	CreatedAt 		time.Time		`json:"createdAt"`
	UpdatedAt 		time.Time		`json:"updatedAt"`
	DeletedAt 		gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}

type Workshop struct {
	ID         			uint64         	`gorm:"primaryKey; autoIncrement" json:"id"`
	Email      			string         	`gorm:"size:255; unique; not null" json:"email" form:"email"`
	Password   			string         	`gorm:"size:255; not null" json:"password" form:"password"`
	Name       			string         	`gorm:"size:125; not null" json:"name" form:"name"`
	Phone				string			`gorm:"size:15; not null" json:"phone" form:"phone"`
	OperationalStart 	string			`gorm:"size:6; not null" json:"operationalStart" form:"operationalStart"`
	OperationalEnd 		string			`gorm:"size:6; not null" json:"operationalEnd" form:"operationalEnd"`
	Description 		Description		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"description"`
	Address 			WorkshopAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
	Services			[]Service		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"services"`	
	Orders				[]Order			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`	
	Ratings				[]WorkshopRating`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ratings"`
	CreatedAt 			time.Time		`json:"createdAt"`
	UpdatedAt 			time.Time		`json:"updatedAt"`
}

type WorkshopAddress struct {
	ID 				uint64 			`gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID 		uint64			`json:"workshopId"`
	BuildingNumber 	uint16			`json:"buildingNumber" form:"buildingNumber"`
	Street 			string			`gorm:"size:255" json:"street" form:"street"`
	City 			string			`gorm:"size:50" json:"city" form:"city"`
	CountryCode 	string			`gorm:"size:5" json:"countryCode" form:"countryCode"`
	PostalCode 		string			`gorm:"size:10" json:"postalCode" form:"postalCode"`
	Province		string			`gorm:"size:50" json:"province" form:"province"`
	CreatedAt 		time.Time		`json:"createdAt"`
	UpdatedAt 		time.Time		`json:"updatedAt"`
	DeletedAt 		gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}

type Description struct {
	ID 			uint64 			`gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID 	uint64			`json:"workshopId"`
	Description string			`json:"description" form:"description"`
	CreatedAt 	time.Time		`json:"createdAt"`
	UpdatedAt 	time.Time		`json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}

type Service struct {
	ID 					uint64			`gorm:"primaryKey; autoIncrement" json:"id"`
	WorkshopID 			uint64			`json:"workshopId"`
	Vehicle 			string			`gorm:"size:125" json:"vehicle" form:"vehicle"`
	Type 				string			`gorm:"size:45" json:"type" form:"type"`
	Fullservice 		bool			`json:"fullservice"`
	FsPrice 			uint64			`json:"fsPrice" form:"fsPrice"`
	RoutineMaintenance 	bool			`json:"routineMaintenance"`
	RmPrice 			uint64			`json:"rmPrice" form:"rmPrice"`
	MachineRepair 		bool			`json:"machineRepair"`
	MrPrice 			uint64			`json:"mrPrice" form:"mrPrice"`
	BodyRepair 			bool			`json:"bodyRepair"`
	BrPrice 			uint64			`json:"brPrice" form:"brPrice"`
	ElectricalRepair 	bool			`json:"electricalRepair"`
	ErPrice 			uint64			`json:"erPrice" form:"erPrice"`
	Orders				[]Order			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`	
	CreatedAt 			time.Time		`json:"createdAt"`
	UpdatedAt 			time.Time		`json:"updatedAt"`
	DeletedAt 			gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}

type Order struct {
	ID 			uint64 			`gorm:"primaryKey; autoIncrement" json:"id"`
	UserID 		uint64			`json:"userId"`
	WorkshopID 	uint64			`json:"workshopId"`
	ServiceID	uint64			`json:"serviceId"`
	OnSite 		bool			`json:"onSite"`
	PickUp 		bool			`json:"pickUp"`
	Note 		string			`json:"note" form:"note"`
	TotalPrice 	uint64			`json:"totalPrice" form:"totalPrice"`
	OrderStatus OrderStatus		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orderStatus"`	
	CreatedAt 	time.Time		`json:"createdAt"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
}

type OrderStatus struct {
	OrderID 	uint64		`json:"orderId"`
	Pending 	bool		`json:"pending"`
	OnProcess 	bool		`json:"onProcess"`
	ReadyToTake bool		`json:"readyToTake"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}

type UserRating struct {
	UserID 		uint64		`json:"userId"`
	WorkshopID 	uint64		`json:"workshopId"`
	Rating 		string		`gorm:"size:1; not null" json:"rating"`
	Description string		`json:"description" form:"description"`
	CreatedAt 	time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}

type WorkshopRating struct {
	WorkshopID 	uint64		`json:"workshopId"`
	UserID 		uint64		`json:"userId"`
	Rating 		string		`gorm:"size:1; not null" json:"rating"`
	Description string		`json:"description" form:"description"`
	CreatedAt 	time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`	
}

// hardcoded data
// var user = []User{
// 	{1, "axe@gmail.com", "axe123", geolocation(42.36399, -71.05493)},
// 	{2, "kunkka@gmail.com", "kunkka123", geolocation(37.4224764, -122.0842499)},
// }

func UserRegister(c echo.Context) error {
	var user SignUp
	c.Bind(&user)
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Phone == "" {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Request is not valid",
			Data:    nil,
		})
	}

	var userDB User
	userDB.Name = user.Name
	userDB.Email = user.Email
	userDB.Password = user.Password
	userDB.Phone = user.Phone
	userDB.Address = UserAddress{CountryCode: "ID"} // TODO: Later replace with geolocation API function (Hardcoded for now)

	result := db.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while input user data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, BaseResponse{
		Code:    http.StatusCreated,
		Message: "Account created",
		Data:    userDB,
	})
}

func UserLogin(c echo.Context) error {
	login := Login{}
	c.Bind(&login)
	var userDB User
	result := db.Where("email = ? AND password = ?", login.Email, login.Password).Find(&userDB)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "User is not registered",
			Data:    nil,
		})
	}
	var addressDB UserAddress
	db.Where("user_id = ?", userDB.ID).Find(&addressDB)
	userDB.Address = addressDB
	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Login success",
		Data:    userDB,
	})
}

func GetUsers(c echo.Context) error {
	var usersDB []User
	result := db.Preload("Address").Preload("Orders").Preload("Ratings").Find(&usersDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while retrieving user data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    usersDB,
	})
}

func UserDetails(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Id is not valid",
		})
	}

	var userDB User
	db.Preload("Address").Preload("Orders").Preload("Ratings").First(&userDB, "id = ?", c.Param("id"))
	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userDB,
	})
}

//============================================================================================
// hardcoded data
// var workshop = []Workshop{
// 	{1, "bengkelaxe@gmail.com", "bengkelaxe123", geolocation(42.36399, -71.05493)},
// 	{2, "bengkelkunkka@gmail.com", "bengkelkunkka123", geolocation(37.4224764, -122.0842499)},
// }

func GetWorkshops(c echo.Context) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Success",
		Data: Workshop{},
	})
}

func WorkshopDetails(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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
		Data: Workshop{ID: id}, // will check to database for now it will be empty
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

// Database
var db *gorm.DB

func InitDB() {
	// load env
	config, cErr := config.LoadConfig(".")
	if cErr != nil {
		log.Fatalln("Cannot load config", cErr)
	}
	// connect to DB
	var err error
	dsn := config.DBUsername + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection DB Failed")
	}
	InitMigration()
}

func InitMigration() {
	db.AutoMigrate(
		&User{},
		&UserAddress{},
		&Workshop{},
		&WorkshopAddress{},
		&Description{},
		&Service{},
		&Order{},
		&OrderStatus{},
		&UserRating{},
		&WorkshopRating{},
	)
}