package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func СonnectDB() *gorm.DB {
	dns := "host=localhost user=user password=pass dbname=notes port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}
	err = db.AutoMigrate(&Note{}, &User{})
	if err != nil {
		log.Fatal("Failed to migrate DB", err)
	}
	log.Println("Successfully connected to DB")
	return db

}
