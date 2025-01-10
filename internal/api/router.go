package router

import (
	"gateway/config"
	_ "gateway/internal/api/docs"
	"gateway/internal/api/handler"
	"gateway/internal/api/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
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

	router.Use(middleware.CORSMiddleware())

	// Apply middleware for CORS and permission checks
	router.Use(cors.Default())

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
		pcategory.PUT("/:id", h.UpdateCategory)
		pcategory.DELETE("/:id", h.DeleteCategory)
	}

	// Product routes group
	products := router.Group("/products")
	{
		products.POST("", h.CreateProduct)
		products.POST("/bulk/:category_id", h.CreateBulkProducts)
		products.GET("", h.GetProductList)
		products.GET("/:id", h.GetProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
		products.POST("/excel-upload/:category_id", h.UploadAndProcessExcel)
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

	//Client routes group
	client := router.Group("/clients")
	{
		client.POST("", h.CreateClient)
		client.GET("", h.GetClientList)
		client.GET("/:id", h.GetClient)
		client.PUT("/:id", h.UpdateClient)
		client.DELETE("/:id", h.DeleteClient)

		client.GET("/street", h.GetStreetClientList)
	}
	company := router.Group("/companies")
	{
		// admin
		company.POST("/admin", h.CreateCompanyA)
		company.GET("/admin/:company_id", h.GetCompanyA)
		company.PUT("/admin/:company_id", h.UpdateCompanyA)
		company.DELETE("/admin/:company_id", h.DeleteCompanyA)
		company.GET("/admin/all", h.GetAllCompaniesA)
		company.GET("/admin/:company_id/users", h.ListCompanyUsersA)
		company.POST("/admin/:company_id/users", h.CreateCompanyUserA)

		// user
		company.GET("", h.GetCompany)
		company.PUT("", h.UpdateCompany)
		company.GET("/users", h.ListCompanyUsers)
		company.POST("/users", h.CreateCompanyUser)

	}
	branch := router.Group("/branches")
	{
		branch.POST("/admin", h.CreateBranch)
		branch.GET("/:branch_id", h.GetBranch)
		branch.PUT("/:branch_id", h.UpdateBranch)
		branch.DELETE("/:branch_id", h.DeleteBranch)
		branch.GET("/company/:company_id", h.ListBranches)
	}

	statics := router.Group("/statistics")
	{
		statics.GET("/products/total-price", h.TotalPriceOfProducts)
		statics.GET("/products/total-sold", h.TotalSoldProducts)
		statics.GET("/products/total-purchased", h.TotalPurchaseProducts)
		statics.GET("/products/get-most-sold", h.GetMostSoldProductsByDay)

		statics.GET("/top-clients", h.GetTopClients)
		statics.GET("/top-suppliers", h.GetTopSuppliers)

		statics.GET("/cash/total-income", h.GetTotalIncome)
		statics.GET("/cash/total-expense", h.GetTotalExpense)
		statics.GET("/cash/net-profit", h.GetNetProfit)
	}

	cash := router.Group("/cash-flow")
	{
		cash.GET("", h.GetCashFlow)
		cash.POST("/income", h.CreateIncome)
		cash.POST("/expense", h.CreateExpense)
	}

	debt := router.Group("/debts")
	{
		debt.POST("", h.CreateDebt)
		debt.GET("/:id", h.GetDebt)
		debt.GET("", h.GetListDebts)
		debt.GET("/client/:client_id", h.GetClientDebts)

		debt.POST("/pay", h.PayDebt)
		debt.GET("/payments/:debt_id", h.GetPaymentsByDebtId)
		debt.GET("/payment/:id", h.GetPayment)
		//debt.GET("/payment", h.GetPayments)
	}

	// Return the configured router
	return router
}
