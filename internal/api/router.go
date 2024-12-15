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

	// Initialize the handler with config
	h := handler.NewHandlerRepo(cfg)

	// User routes group
	user := router.Group("/user")
	{
		user.POST("/admin/register", h.RegisterAdmin)
		user.POST("/register", h.CreateUser)
		user.POST("/login", h.Login)
		user.GET("/get/:id", h.GetUser)
		user.GET("/list", h.ListUser)
		user.PUT("/update/:id", h.UpdateUser)
		user.DELETE("/delete/:id", h.DeleteUser)
	}
	router.Use(middleware.PermissionMiddleware(enf))
	// Product Category routes group
	pcategory := router.Group("/products/category")
	{
		pcategory.POST("", h.CreateCategory)
		pcategory.GET("", h.GetListCategory)
		pcategory.GET("/:id", h.GetCategory)
		//pcategory.PUT("/update/:id", h.)
		pcategory.DELETE("/:id", h.DeleteCategory)
	}

	// Product routes group
	products := router.Group("/products")
	{
		products.POST("", h.CreateProduct)
		products.GET("", h.GetProductList)
		products.GET("/:id", h.GetProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
	}

	// Purchase routes group
	purchase := router.Group("/purchases")
	{
		purchase.POST("", h.CreatePurchase)
		purchase.GET("", h.GetListPurchase)
		purchase.GET("/:id", h.GetPurchase)
		purchase.PUT("/:id", h.UpdatePurchase)
		purchase.DELETE("/:id", h.DeletePurchase)
	}

	// Sales routes group
	sales := router.Group("/sales")
	{
		sales.POST("", h.CreateSales)
		sales.GET("", h.GetListSales)
		sales.GET("/:id", h.GetSales)
		sales.PUT("/:id", h.UpdateSales)
		sales.DELETE("/:id", h.DeleteSales)
		sales.POST("/calculate", h.CalculateTotalSales)
	}

	// Return the configured router
	return router
}
