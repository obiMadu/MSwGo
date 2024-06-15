package main

import (
	"log"
	"os"
	"time"


	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const webPort string = ":80"

var counts int

type Config struct{}

func main() {
	app := Config{}

	// define http server

	// connect to the db
	db := connectToDB()
	if db == nil {
		log.Panic("Can't connect to Postgress")
	}

	// create http server
	router := gin.Default()

	app.routes(router)

	// start the server
	log.Printf("Starting auth service on port %s\n", webPort)
	router.Run(webPort)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// return *sql.DB from db(*gorm.DB) to enable Ping()
	gormDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// ping database
	err = gormDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *gorm.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for three seconds....")
		time.Sleep(3 * time.Second)
		continue
	}
}
