package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error" binding:"required"`
	Message string `json:"message" binding:"required"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(c *gin.Context) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker!",
	}

	c.JSON(http.StatusOK, payload)
}
