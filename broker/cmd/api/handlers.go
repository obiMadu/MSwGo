package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Config) Broker(c *gin.Context) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}

	c.JSON(http.StatusOK, payload)
}
