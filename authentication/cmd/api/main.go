package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

const webPort string = ":80"

type Config struct{}

func main() {
	app := Config{}

	// define http server
	router := gin.Default()

	app.routes(router)

	// start the server
	log.Printf("Starting auth service on port %s\n", webPort)
	router.Run(webPort)
}
