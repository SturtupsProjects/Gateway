package router

import (
	"gateway/config"
	_ "gateway/internal/api/docs"
	"gateway/internal/api/handler"
	"gateway/internal/api/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description This is a sample server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @scheme http
func NewRouter(enf *casbin.Enforcer, cfg *config.Config) *gin.Engine {
	// Initialize the Gin router
	router := gin.Default()

	// Apply middleware for CORS and permission checks
	router.Use(middleware.CORSMiddleware())

	// Swagger Documentation Route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Use(middleware.PermissionMiddleware(enf))
	// Initialize the handler with config
	h := handler.NewHandlerRepo(cfg)

	// User routes group
	user := router.Group("/user")
	{
		user.POST("/admin/register", h.RegisterAdmin) // Register Admin
		user.POST("/register", h.CreateUser)          // Register User
		user.POST("/login", h.Login)                  // Login
		user.GET("/get/:id", h.GetUser)               // Get User by ID
		user.GET("/list", h.ListUser)                 // List all Users
		user.PUT("/update/:id", h.UpdateUser)         // Update User by ID
		user.DELETE("/delete/:id", h.DeleteUser)      // Delete User by ID
	}

	// Return the configured router
	return router
}
