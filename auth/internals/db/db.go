package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var counts int

func OpenDB(dsn string) (*gorm.DB, error) {
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

func ConnectToDB() *gorm.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := OpenDB(dsn)
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
