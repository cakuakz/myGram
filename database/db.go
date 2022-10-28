package database

import (
	
	"fmt"
	"final-project/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)

var (
	host 		= "localhost"
	user 		= "postgres"
	password 	= "123456"
	port		= 5432
	name		= "myGram"
	db			*gorm.DB
	err			error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, name, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB{
	return db
}