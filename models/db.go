package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"social-service-sync/app/config"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DbHost, config.DbUser, config.DbName, config.DbPassword)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func DbConn() *gorm.DB {
	return db
}
