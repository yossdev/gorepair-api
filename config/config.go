package config

import (
	"fmt"
	"gorepair-rest-api/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Env struct {
	HERE_API_KEY string `mapstructure:"HERE_API_KEY"`
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
		&models.User{},
		&models.UserAddress{},
		&models.Workshop{},
		&models.WorkshopAddress{},
		&models.Description{},
		&models.Service{},
		&models.Order{},
		&models.OrderStatus{},
		&models.Rating{},
	)
}