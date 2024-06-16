package main

import (
	"database/sql"
	"log"

	"auth/internals/db"
	"auth/internals/env"
	"auth/internals/models"
)

const webPort string = ":8082"

type Config struct {
	DB     *sql.DB
	Models models.Models
}

func main() {
	app := Config{}

	// load environment variables
	env.LoadConfig()

	// configure the db
	gormDB, err := db.NewDB()
	if err != nil {
		log.Panicf("Unable to configure DB %s\n", err.Error())
	}

	rawDB := db.RawDB(gormDB)

	app = Config{
		DB:     rawDB,
		Models: models.New(rawDB),
	}

	// configure routes
	mux := app.routes()

	// start the server
	log.Printf("Starting auth service on port %s\n", webPort)
	mux.Run(webPort)
}
