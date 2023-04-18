package main

import (
	"jobsPortal/entities"
	"jobsPortal/routers"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbPort   = "5432"
	dbName   = "dans"
	db       *gorm.DB
	err      error
)

func init() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(entities.User{})
}

func main() {
	routers.StartService(db).Run(":4000")
}
