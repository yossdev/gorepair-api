package middlewares

import (
	"gorepair-rest-api/config"
	"gorepair-rest-api/models/tables"

	"github.com/labstack/echo/v4"
)

func BasicAuth(email, password string, c echo.Context) (bool, error) {
	var user tables.User
	result := config.DB.Where("email = ? AND password = ?", email, password).Preload("Address").Preload("Orders").Preload("Ratings").Find(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return false, result.Error
	}
	return true, nil
}