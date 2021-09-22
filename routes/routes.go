package routes

import (
	"gorepair-rest-api/controllers"
	"gorepair-rest-api/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func New() *echo.Echo {
	// New Instance
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.BodyDump(middlewares.BodyDumpLog))

	v1 := e.Group("v1/api/")

	// Restricted group
	r := e.Group("v1/api/restricted/")
	r.Use(middleware.JWT([]byte(viper.GetString(`SECRET_JWT`))))
	r.GET("users/:id", controllers.UserDetailsCtrl)

	// User routes
	v1.GET("users", controllers.GetUsersCtrl)
	v1.POST("sign-up/user", controllers.UserRegisterCtrl)
	v1.POST("sign-in/user", controllers.UserLoginCtrl)
	v1.PUT("users/:id/address", controllers.UpdateUserAddressCtrl)

	// Workshop routes
	v1.GET("workshops", controllers.GetWorkshopsCtrl)
	v1.GET("workshops/:id", controllers.WorkshopDetailsCtrl)
	v1.GET("workshops/find", controllers.FindWorkshopCtrl)
	v1.POST("sign-up/workshop", controllers.WorkshopRegisterCtrl)
	v1.POST("sign-in/workshop", controllers.WorkshopLoginCtrl)
	v1.PUT("workshops/:id/description", controllers.UpdateWorkshopDescriptionCtrl)
	v1.PUT("workshops/:id/address", controllers.UpdateWorkshopAddressCtrl)

	return e
}
