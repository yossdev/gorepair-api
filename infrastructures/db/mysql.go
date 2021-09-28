package db

import (
	"fmt"
	"gorepair-rest-api/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB interface {
	DB() *gorm.DB
}

type mysqlDB struct {
	db *gorm.DB
}

func NewMysqlClient() MysqlDB {
	// log.Println("Initialize Database connection")
	var db *gorm.DB
	var err error
	dbHost := config.Get().DbHost
	dbPort := config.Get().DbPort
	dbName := config.Get().DbName
	dbUser := config.Get().DbUsername
	dbPassword := config.Get().DbPassword

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, 
		dbPassword, 
		dbHost, 
		dbPort, 
		dbName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// defer db.Close()

	if err != nil {
		log.Println(fmt.Sprintf("Error to loading Database %s", err))
		return nil
	}

	log.Println("Connected to MySql!")

	return &mysqlDB{
		db: db,
	}
}

func (c mysqlDB) DB() *gorm.DB {
	return c.db
}
