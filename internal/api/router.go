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
	"log/slog"
)

// @title API Gateway
// @version 1.0
// @description This is a sample server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @scheme http
func NewRouter(enf *casbin.Enforcer, cfg *config.Config, log *slog.Logger) *gin.Engine {
	// Initialize the Gin router
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	swagger := router.Group("/swagger", gin.BasicAuth(gin.Accounts{
		"smart-admin": "admin_846", // Логин и пароль
	}))
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Initialize the handler with config
	h := handler.NewHandlerRepo(cfg, log)

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
		user.POST("/get/access-token", h.GetAccessToken)
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
		products.GET("/dashboard/:currency", h.GetProductsDashboard)
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

	// Client routes group
	client := router.Group("/clients")
	{
		client.POST("", h.CreateClient)
		client.GET("", h.GetClientList)
		client.GET("/:id", h.GetClient)
		client.PUT("/:id", h.UpdateClient)
		client.DELETE("/:id", h.DeleteClient)

	}

	// Supplier routes group
	supplier := router.Group("/supplier")
	{
		supplier.POST("", h.CreateSupplier)
		supplier.GET("", h.GetSupplierList)
		supplier.GET("/:id", h.GetSupplier)
		supplier.PUT("/:id", h.UpdateSupplier)
		supplier.DELETE("/:id", h.DeleteSupplier)
	}

	// Company routes group
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

	// Branch routes group
	branch := router.Group("/branches")
	{
		branch.POST("/create", h.CreateBranch)
		branch.GET("/:branch_id", h.GetBranch)
		branch.PUT("/:branch_id", h.UpdateBranch)
		branch.DELETE("/:branch_id", h.DeleteBranch)
		branch.GET("/list", h.ListBranches)
	}

	// Statistics routes group
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

		statics.GET("/sale-statistics", h.GetSaleStatistics)
		statics.GET("/branch-income", h.GetBranchIncome)
		statics.GET("client-dashboard/:client_id", h.GetClientDashboard)
	}

	// CashFlow group
	cash := router.Group("/cash-flow")
	{
		cash.GET("", h.GetCashFlow)
		cash.POST("/income", h.CreateIncome)
		cash.POST("/expense", h.CreateExpense)
	}

	// Debts routes group
	debt := router.Group("/debts")
	{
		debt.POST("", h.CreateDebt)
		debt.GET("/:id", h.GetDebt)
		debt.GET("", h.GetListDebts)
		debt.GET("/client/:client_id", h.GetClientDebts)

		debt.GET("/excel/:currency", h.GetDebtsInExcel)

		debt.POST("/pay", h.PayDebt)
		debt.GET("/payments/:debt_id", h.GetPaymentsByDebtId)
		debt.GET("/payment/:id", h.GetPayment)
		debt.POST("/payments", h.Payments)
		debt.GET("/debts/payments/:user_id", h.GetUserPayments)

		debt.GET("/total-sum", h.GetTotalDebtSum)
		debt.GET("/total-sum/:user_id", h.GetUserTotalDebt)

	}

	creditor := router.Group("/creditor")
	{
		creditor.POST("", h.CreateCreditor)
		creditor.GET("/:id", h.GetCreditors)
		creditor.GET("", h.GetListCreditors)
		creditor.GET("/client/:supplier_id", h.GetCreditsFromSupplier)

		creditor.POST("/pay", h.PayCredit)
		creditor.GET("/payments/:creditor_id", h.GetPaymentsByCreditId)
		creditor.GET("/payment/:id", h.GetCreditPayment)
		creditor.GET("/creditors/payments/:supplier_id", h.GetPaymentsToSupplier)

		creditor.GET("/total-sum", h.GetTotalCreditSum)
		creditor.GET("/total-sum/:user_id", h.GetTotalCreditFromSupplier)

	}

	// Transfers routes group
	transfers := router.Group("/transfers")
	{
		transfers.POST("", h.CreateTransfers)
		transfers.GET("/:id", h.GetTransfers)
		transfers.GET("", h.GetTransferList)
	}

	//balance := router.Group("/company-balance")
	//{
	//	balance.POST("", h.CreateCompanyBalance)
	//	balance.GET("", h.GetCompanyBalance)
	//	balance.PUT("", h.UpdateCompanyBalance)
	//	balance.DELETE("/:company_id", h.DeleteCompanyBalance)
	//	balance.GET("/list", h.GetUsersBalanceList)
	//}

	// Salary routes group
	salary := router.Group("/salary")
	{
		salary.POST("", h.CreateSalary)
		salary.PUT("/:salary_id", h.UpdateSalary)
		salary.GET("/:salary_id", h.GetSalaryByID)
		salary.GET("", h.ListSalaries)
		salary.GET("/worker/:worker_id", h.GetWorkerAllInfo)
	}

	// Adjustment routes group
	adjustment := router.Group("/adjustment")
	{
		adjustment.POST("", h.CreateAdjustment)
		adjustment.PUT("/:adjustment_id", h.UpdateAdjustment)
		adjustment.PUT("/:adjustment_id/close", h.CloseAdjustment)
		adjustment.GET("/:adjustment_id", h.GetAdjustmentByID)
		adjustment.GET("", h.ListAdjustments)
	}

	return router
}
