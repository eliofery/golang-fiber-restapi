package database

import (
	"fmt"
	"github.com/eliofery/golang-fiber-restapi/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func Connect() error {
	op := "database.Connect"

	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	dbConnection, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("%v: %w", op, err)
	}

	DB = dbConnection

	return nil
}

func Migrations() error {
	op := "database.Migrations"

	err := DB.Debug().AutoMigrate(
		&model.Cashier{},
		&model.Category{},
		&model.Discount{},
		&model.Order{},
		&model.Payment{},
		&model.PaymentType{},
		&model.Product{},
	)
	if err != nil {
		return fmt.Errorf("%v: %w", op, err)
	}

	return nil
}
