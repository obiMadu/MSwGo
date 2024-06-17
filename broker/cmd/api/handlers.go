package main

import (
	"bytes"
	"encoding/json"
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

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleSubmission(c *gin.Context) {
	var requestPayload RequestPayload

	err := c.ShouldBindBodyWithJSON(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Request format is bad.",
		})
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(c, requestPayload.Auth)
	default:
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Unknown action.",
		})
	}
}

func (app *Config) authenticate(c *gin.Context, a AuthPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://localhost:8082/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Error occured while contacting auth service.",
		})
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Error occured while contacting auth service.",
		})
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		c.JSON(http.StatusUnauthorized, jsonResponse{
			Error:   true,
			Message: "Invalid credendials.",
		})
		return
	} else if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Error calling auth service.",
		})
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Error occured while processing response from auth service.",
		})
		return
	}

	if jsonFromService.Error {
		c.JSON(http.StatusUnauthorized, jsonResponse{
			Error:   true,
			Message: "Invalid credendials.",
		})
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	c.JSON(http.StatusOK, payload)
}
