package controllers

import (
	"gorepair-rest-api/lib/database"
	"gorepair-rest-api/models"
	"gorepair-rest-api/models/tables"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func WorkshopRegisterCtrl(c echo.Context) error {
	var workshop models.SignUp
	c.Bind(&workshop)
	if workshop.Name == "" || workshop.Email == "" || workshop.Password == "" || workshop.Phone == "" {
		return c.JSON(http.StatusBadRequest, models.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := database.WorkshopRegister(workshop)
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

func UpdateWorkshopAddressCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}

	address := tables.WorkshopAddress{}
	c.Bind(&address)
	workshop := database.UpdateWorkshopAddress(c.Param("id"), address)
	if workshop == nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while updating data",
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Address successfully updated",
		Data:    workshop,
	})
}

func WorkshopLoginCtrl(c echo.Context) error {
	login := models.Login{}
	c.Bind(&login)
	workshop := database.WorkshopLogin(login)
	if workshop == nil {
		return c.JSON(http.StatusForbidden, models.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Email or Password is wrong",
			Data:    login,
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Login success",
		Data:    workshop,
	})
}

func GetWorkshopsCtrl(c echo.Context) error {
	workshop, e := database.GetWorkshops()
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
		Data:    workshop,
	})
}

func WorkshopDetailsCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}

	workshop, e := database.WorkshopDetails(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusNotFound, models.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Workshop is not exist",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    workshop,
	})
}

func UpdateWorkshopDescriptionCtrl(c echo.Context) error {
	_, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.ApiResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Id is not valid",
		})
	}
	description := tables.Description{}
	c.Bind(&description)
	result := database.UpdateWorkshopDescription(c.Param("id"), description)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while inputing data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Description successfully Updated",
		Data:    result,
	})
}

func FindWorkshopCtrl(c echo.Context) error {
	result := database.FindWorkshop(c.QueryParam("name"))
	if result == nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error while retrieving data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.ApiResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

