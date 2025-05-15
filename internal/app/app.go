package app

import (
	"api-gateway/config"

	"api-gateway/internal/client"
	"api-gateway/internal/controller"
	"api-gateway/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg        *config.Config
	router     *gin.Engine
	authClient *client.AuthClient
	userClient *client.UserClient
	gptClient  *client.GptClient
}

func NewServer(cfg *config.Config) (*Server, error) {
	authClient, err := client.NewAuthClient(cfg.Auth.Address)
	if err != nil {
		return nil, err
	}

	userClient, err := client.NewUserClient(cfg.User.Address)
	if err != nil {
		return nil, err
	}

	gptClient, err := client.NewGptClient(cfg.Gpt.Address)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	return &Server{
		cfg:        cfg,
		router:     router,
		authClient: authClient,
		userClient: userClient,
		gptClient:  gptClient,
	}, nil
}

func (s *Server) setupRoutes() {
	authController := controller.NewAuthController(s.authClient)
	userController := controller.NewUserController(s.userClient)
	gptController := controller.NewGptController(s.gptClient, s.userClient)

	// Public routes
	public := s.router.Group("/api/v1")
	{
		public.POST("/login", authController.Login)
	}

	// Protected routes
	protected := s.router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(s.authClient))
	{
		// User routes
		protected.POST("/user/products", userController.AddProduct)
		protected.DELETE("/user/products", userController.RemoveProduct)
		protected.PUT("/user/preferences", userController.UpdatePreferences)
		protected.GET("/user/products", userController.GetUserProducts)

		// GPT routes
		protected.POST("/gpt/generate", gptController.GenerateResponse)
	}
}

func (s *Server) Run(cfg *config.Config, devMode bool) error {
	defer s.authClient.Close()
	defer s.userClient.Close()
	defer s.gptClient.Close()

	s.setupRoutes()

	log.Printf("Server is running on port %s", s.cfg.Server.Port)
	return s.router.Run(":" + s.cfg.Server.Port)
}