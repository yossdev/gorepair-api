package controllers

import (
	"encoding/json"
	"gorepair-rest-api/libs/database"
	"gorepair-rest-api/middlewares"
	"gorepair-rest-api/models"
	"gorepair-rest-api/models/responses"
	"gorepair-rest-api/models/tables"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserRegisterCtrl(c echo.Context) error {
	var user models.SignUp
	c.Bind(&user)
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Phone == "" {
		return c.JSON(http.StatusBadRequest, models.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := database.UserRegister(user)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while inputing data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, models.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Account created",
		Data:    result,
	})
}

func UpdateUserAddressCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}

	address := tables.UserAddress{}
	c.Bind(&address)
	user := database.UpdateUserAddress(c.Param("id"), address)
	if user == nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while updating data",
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
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
		return c.JSON(http.StatusForbidden, models.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Email or Password is wrong",
			Data:    login,
		})
	}

	byteData, _ := json.Marshal(user)
	var u tables.User
	json.Unmarshal([]byte(byteData), &u)
	
	jwtToken, e := middlewares.CreateToken(int(u.ID), u.Name)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
	}
	
	userResponse := responses.UserResponse{
		ID:        	u.ID,
		Email:     	u.Email,
		Name:		u.Name,
		Token:     	jwtToken,
		CreatedAt: 	u.CreatedAt,
		UpdatedAt: 	u.UpdatedAt,
	}
	
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Login success",
		Data:    userResponse,
	})
}

func GetUsersCtrl(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while retrieving data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}

func UserDetailsCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}

	users, e := database.UserDetails(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusNotFound, models.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "User is not exist",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}