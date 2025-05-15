package controller

import (
	// "net/http"

	"api-gateway/internal/client"
	// "github.com/gin-gonic/gin"
)

type UserController struct {
	userClient *client.UserClient
}

func NewUserClient(userClient *client.UserClient) *UserController {
	return &UserController{
		userClient: userClient,
	}
}