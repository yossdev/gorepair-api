package web

import (
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/infrastructures/local_db"

	"github.com/gofiber/fiber/v2"
)

type RouterStruct struct {
	Web       *fiber.App
	MysqlDB   db.MysqlDB
	MongoDB	  db.MongoDB
	ScribleDB local_db.ScribleDB
}