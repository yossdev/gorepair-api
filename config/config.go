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

type Env struct {
	DBUsername   string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBName       string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// Database
var DB *gorm.DB

func InitDB() {
	// load env
	config, e := LoadConfig(".")
	if e != nil {
		log.Fatalln("Cannot load config", e)
	}

	// connect to DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
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
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
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
