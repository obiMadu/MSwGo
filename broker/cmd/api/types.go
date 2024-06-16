package main

type jsonResponse struct {
	Error   bool   `json:"error" binding:"required"`
	Message string `json:"message" binding:"required"`
	Data    any    `json:"data,omitempty"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
