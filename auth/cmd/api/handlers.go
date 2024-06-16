package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Authenticate(c *gin.Context) {

	requestPayload := jsonRequest{}

	err := c.ShouldBindBodyWithJSON(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Invalid credentials.",
		})
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Invalid credentials.",
		})
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Invalid credentials.",
		})
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	c.JSON(http.StatusAccepted, payload)
}
