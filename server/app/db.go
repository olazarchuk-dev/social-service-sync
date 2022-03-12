package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/olazarchuk-dev/go-social-service/server/app/config"
)

var db *sql.DB

func init() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbUri := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbPort, config.DbName)
	fmt.Println(dbUri)

	conn, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
}

func DbConn() *sql.DB {
	return db
}
