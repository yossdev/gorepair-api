package routes

import (
	"gorepair-rest-api/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// New Instance
	e := echo.New()
	v1 := e.Group("v1/api/")

	// User routes
	v1.GET("users", controllers.GetUsersCtrl)
	v1.GET("users/:id", controllers.UserDetailsCtrl)
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
