package main

import (
	"log"

	"auth/internals/db"
	"auth/internals/env"
)

const webPort string = ":8082"

type Config struct{}

func main() {
	app := Config{}

	// load environment variables
	env.LoadConfig()

	// connect to the db
	db := db.ConnectToDB()
	if db == nil {
		log.Panic("Can't connect to Postgress")
	}

	// configure routes
	mux := app.routes()

	// start the server
	log.Printf("Starting auth service on port %s\n", webPort)
	mux.Run(webPort)
}
