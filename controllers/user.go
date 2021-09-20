package controllers

import (
	"gorepair-rest-api/lib/database"
	"gorepair-rest-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserRegisterCtrl(c echo.Context) error {
	var user models.SignUp
	c.Bind(&user)
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Phone == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := database.UserRegister(user)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while inputing data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Account created",
		Data:    result,
	})
}

func UpdateUserAddressCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}

	address := models.UserAddress{}
	c.Bind(&address)
	user := database.UpdateUserAddress(c.Param("id"), address)
	if user == nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while updating data",
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Address successfully updated",
		Data:    user,
	})
}

func UserLoginCtrl(c echo.Context) error {
	login := models.Login{}
	c.Bind(&login)
	user := database.UserLogin(login)
	if user == nil {
		return c.JSON(http.StatusForbidden, models.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "Email or Password is wrong",
			Data:    login,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login success",
		Data:    user,
	})
}

func GetUsersCtrl(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while retrieving data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}

func UserDetailsCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}

	users, e := database.UserDetails(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "User is not exist",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}