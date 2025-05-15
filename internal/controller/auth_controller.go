package controller

import (
	// "net/http"

	"api-gateway/internal/client"
	// "github.com/gin-gonic/gin"
)

type AuthController struct {
	authClient *client.AuthClient
}

func NewAuthController(authClient *client.AuthClient) *AuthController {
	return &AuthController{
		authClient: authClient,
	}
}
