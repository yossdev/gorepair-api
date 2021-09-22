package config

import (
	"context"
	"fmt"
	"gorepair-rest-api/models/tables"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadConfig() {
	viper.SetConfigFile(`app.env`)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Cannot load config", err)
	}
}

// Database
var DB *gorm.DB

func InitDB() {
	// connect to DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString(`DB_USER`),
		viper.GetString(`DB_PASSWORD`),
		viper.GetString(`DB_HOST`),
		viper.GetString(`DB_PORT`),
		viper.GetString(`DB_NAME`),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection DB Failed")
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(
		&tables.User{},
		&tables.UserAddress{},
		&tables.Workshop{},
		&tables.WorkshopAddress{},
		&tables.Description{},
		&tables.Service{},
		&tables.Order{},
		&tables.OrderStatus{},
		&tables.Rating{},
	)
}

// MongoDB for saving data log
var Client *mongo.Client

func InitMongo() {
	// Set client options
	clientOptions := options.Client().ApplyURI(viper.GetString(`MongoDB`))
	
	// Connect to MongoDB
	var e error
	Client, e = mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
        log.Fatalln(e)
    }
	
	// Check the connection
	e = Client.Ping(context.TODO(), nil)
	if e != nil {
        log.Fatalln(e)
    }
}
