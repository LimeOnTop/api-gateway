package controller

import (
	// "net/http"

	"api-gateway/internal/client"
	// "github.com/gin-gonic/gin"
)

type UserController struct {
	userClient *client.UserClient
}

func NewUserController(userClient *client.UserClient) *UserController {
	return &UserController{
		userClient: userClient,
	}
}