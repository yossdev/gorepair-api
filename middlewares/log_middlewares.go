package middlewares

import (
	"context"
	"gorepair-rest-api/config"
	"gorepair-rest-api/models"
	"log"

	"github.com/labstack/echo/v4"
)

func BodyDumpLog(c echo.Context, reqBody, resBody []byte) {
	collection := config.Client.Database("play").Collection("apiLog")

	var apilog models.BodyDumpLog
	apilog.URI = c.Request().RequestURI
	apilog.Method = c.Request().Method
	apilog.Status = c.Response().Status

	_, e := collection.InsertOne(context.TODO(), apilog)
	if e != nil {
		log.Fatalln(e)
	}
}