package database

import (
	"fmt"
	"log"
	"os"

	"github.com/vaibhavvvvv/go-trivia-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //options for gorm
	}) //postgres.Open(dsn) is gorm_dialector. dsn->data source name string

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2) //to exit program with error code 2
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{}) //to create tables that we need from our gorm models

	DB = Dbinstance{
		Db: db, //assigning value to Global DB variable with the db we set up
	}
}
