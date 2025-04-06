package handler

import (
	"context"
	"gateway/internal/entity"
	"gateway/internal/generated/products"
	pbu "gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// TotalPriceOfProducts godoc
// @Summary Calculate the total price of products
// @Description Calculate the total price of all products for a specific company
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/products/total-price [get]
func (h *Handler) TotalPriceOfProducts(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02" // Человечный формат
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	// Call the gRPC method
	res, err := h.ProductClient.TotalPriceOfProducts(c, req)
	if err != nil {
		h.log.Error("Error calculating total price of products", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// TotalSoldProducts godoc
// @Summary Calculate the total quantity of sold products
// @Description Calculate the total quantity of sold products for a specific company
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/products/total-sold [get]
func (h *Handler) TotalSoldProducts(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02" // Человечный формат
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	// Call the gRPC method
	res, err := h.ProductClient.TotalSoldProducts(c, req)
	if err != nil {
		h.log.Error("Error calculating total sold products", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// TotalPurchaseProducts godoc
// @Summary Calculate the total purchase amount of products
// @Description Calculate the total purchase amount of products for a specific company
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/products/total-purchased [get]
func (h *Handler) TotalPurchaseProducts(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02" // Человечный формат
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	// Call the gRPC method
	res, err := h.ProductClient.TotalPurchaseProducts(c, req)
	if err != nil {
		h.log.Error("Error calculating total purchase products", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetMostSoldProductsByDay godoc
// @Summary Get the most sold products by day
// @Description Get a list of the most sold products between a given date range
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.MostSoldProductsResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/products/most-sold [get]
func (h *Handler) GetMostSoldProductsByDay(c *gin.Context) {

	log.Println("Mana keldi")

	companyId := c.MustGet("company_id").(string)
	branchId := c.GetHeader("branch_id") // Extract branch_id from header

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	req := &products.MostSoldProductsRequest{
		CompanyId: companyId,
		BranchId:  branchId, // Pass branch_id from header
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
	}

	h.log.Info("GetMostSoldProductsByDay", "req", req.CompanyId)

	res, err := h.ProductClient.GetMostSoldProductsByDay(c, req)
	if err != nil {
		h.log.Error("Error getting most sold products by day", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTopClients godoc
// @Summary Get top clients by value of purchases
// @Description Get the top clients for a company based on their purchase value in a given date range
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.GetTopEntitiesResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/top-clients [get]
func (h *Handler) GetTopClients(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02" // Человечный формат
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.GetTopEntitiesRequest{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339), // Переводим в RFC3339 для передачи
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	res, err := h.ProductClient.GetTopClients(c, req)
	if err != nil {
		h.log.Error("Error getting top clients", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var listCients entity.TopClientList

	for _, clientID := range res.Entities {
		var topClient entity.TopClient

		cl, err := h.UserClient.GetClient(context.Background(), &pbu.UserIDRequest{Id: clientID.SupplierId, CompanyId: companyId})
		if err == nil {
			topClient.ID = cl.Id
			topClient.Name = cl.FullName
			topClient.Phone = cl.Phone
			topClient.TotalSum = clientID.TotalValue
		} else {
			h.log.Error("Error getting client id", "error", err.Error())
			topClient.ID = clientID.SupplierId
			topClient.TotalSum = clientID.TotalValue
		}

		listCients.Clients = append(listCients.Clients, topClient)
	}

	c.JSON(http.StatusOK, listCients)
}

// GetTopSuppliers godoc
// @Summary Get top suppliers by value of products supplied
// @Description Get the top suppliers for a company based on the value of products supplied in a given date range
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.GetTopEntitiesResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/top-suppliers [get]
func (h *Handler) GetTopSuppliers(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02" // Человечный формат
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.GetTopEntitiesRequest{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	res, err := h.ProductClient.GetTopSuppliers(c, req)
	if err != nil {
		h.log.Error("Error getting top suppliers", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var listSuppliers entity.TopClientList

	for _, supplier := range res.Entities {
		var topSupplier entity.TopClient

		cl, err := h.UserClient.GetClient(context.Background(), &pbu.UserIDRequest{Id: supplier.SupplierId, CompanyId: companyId})
		if err == nil {
			topSupplier.ID = cl.Id
			topSupplier.Name = cl.FullName
			topSupplier.Phone = cl.Phone
			topSupplier.TotalSum = supplier.TotalValue
		} else {
			h.log.Error("Error getting supplier id", "error", err.Error())
			topSupplier.ID = supplier.SupplierId
			topSupplier.TotalSum = supplier.TotalValue
		}

		listSuppliers.Clients = append(listSuppliers.Clients, topSupplier)
	}

	c.JSON(http.StatusOK, listSuppliers)
}

// GetTotalIncome godoc
// @Summary Calculate the total income
// @Description Calculate the total income for a specific company
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/cash/total-income [get]
func (h *Handler) GetTotalIncome(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	// Call the repository method
	res, err := h.ProductClient.GetTotalIncome(context.Background(), req)
	if err != nil {
		h.log.Error("Error calculating total income", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTotalExpense godoc
// @Summary Calculate the total expense
// @Description Calculate the total expense for a specific company
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/cash/total-expense [get]
func (h *Handler) GetTotalExpense(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	// Call the repository method
	res, err := h.ProductClient.GetTotalExpense(context.Background(), req)
	if err != nil {
		h.log.Error("Error calculating total expense", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetNetProfit godoc
// @Summary Calculate the net profit
// @Description Calculate the net profit for a specific company
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /statistics/cash/net-profit [get]
func (h *Handler) GetNetProfit(c *gin.Context) {
	companyId := c.MustGet("company_id").(string)

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
		BranchId:  branchId,
	}

	// Call the repository method
	res, err := h.ProductClient.GetNetProfit(context.Background(), req)
	if err != nil {
		h.log.Error("Error calculating net profit", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetCashFlow godoc
// @Summary Get cash flow details for a company within a given date range
// @Description Get a list of cash flow transactions for a company based on a given date range
// @Tags Cash Flow
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Param limit query integer false "Limit of records per page (default: 10)"
// @Param page query integer false "Page number for pagination (default: 1)"
// @Param description query string false "Filter transactions by description"
// @Param transaction_type query string false "Transaction Type (income | expense)"
// @Param payment_method query string false "Payment Method (uzs | usd | card)"
// @Param branch_id header string true "Branch ID of the company"
// @Success 200 {object} products.ListCashFlow "List of cash flow transactions with pagination details"
// @Failure 400 {object} products.Error "Invalid input parameters or missing required values"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /cash-flow [get]
func (h *Handler) GetCashFlow(c *gin.Context) {

	companyId := c.MustGet("company_id").(string)
	description := c.Query("description")
	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")
	transactionType := c.Query("transaction_type")
	paymentMethod := c.Query("payment_method")
	limit := c.Query("limit")
	page := c.Query("page")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 // Default limit
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1 // Default page
	}

	if startDate == "" || endDate == "" {
		h.log.Error("Missing required query parameters: start_date or end_date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	layout := "2006-01-02"
	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		h.log.Error("Invalid start_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, expected YYYY-MM-DD"})
		return
	}

	parsedEndDate, err := time.Parse(layout, endDate)
	if err != nil {
		h.log.Error("Invalid end_date format", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, expected YYYY-MM-DD"})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.CashFlowReq{
		CompanyId:       companyId,
		BranchId:        branchId,
		StartDate:       parsedStartDate.Format(time.RFC3339),
		EndDate:         parsedEndDate.Format(time.RFC3339),
		TransactionType: transactionType,
		PaymentMethod:   paymentMethod,
		Description:     description,
		Limit:           int64(limitInt),
		Page:            int64(pageInt),
	}

	res, err := h.ProductClient.GetCashFlow(c, req)
	if err != nil {
		h.log.Error("Error getting cash flow", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateIncome godoc
// @Summary Create an income cash flow transaction for a company
// @Description Create an income record in the cash flow system
// @Tags Cash Flow
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body products.CashFlowRequest true "Income Cash Flow Data"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.CashFlow
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /cash-flow/income [post]
func (h *Handler) CreateIncome(c *gin.Context) {

	var request products.CashFlowRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Получаем branch_id из заголовка
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	// Устанавливаем branch_id в запрос
	request.BranchId = branchId

	// Получаем компанию и пользователя
	request.CompanyId = c.MustGet("company_id").(string)
	request.UserId = c.MustGet("id").(string)

	// Создание дохода
	res, err := h.ProductClient.CreateIncome(c, &request)
	if err != nil {
		h.log.Error("Error creating income", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateExpense godoc
// @Summary Create an expense cash flow transaction for a company
// @Description Create an expense record in the cash flow system
// @Tags Cash Flow
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body products.CashFlowRequest true "Expense Cash Flow Data"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.CashFlow
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /cash-flow/expense [post]
func (h *Handler) CreateExpense(c *gin.Context) {

	var request products.CashFlowRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Получаем branch_id из заголовка
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	// Устанавливаем branch_id в запрос
	request.BranchId = branchId

	// Получаем компанию и пользователя
	request.UserId = c.MustGet("id").(string)
	request.CompanyId = c.MustGet("company_id").(string)

	// Создание расхода
	res, err := h.ProductClient.CreateExpense(c, &request)
	if err != nil {
		h.log.Error("Error creating expense", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetSaleStatistics godoc
// @Summary Get sales statistics
// @Description Retrieve sales statistics based on a given time period
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Param period query string false "Period (e.g., daily, weekly, monthly)"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.SaleStatistics
// @Failure 400 {object} products.Error "Branch ID is required in the header"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /statistics/sale-statistics [get]
func (h *Handler) GetSaleStatistics(c *gin.Context) {

	var req products.SaleStatisticsReq

	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")
	req.Period = c.Query("period")

	req.CompanyId = c.MustGet("company_id").(string)

	req.BranchId = c.GetHeader("branch_id")
	if req.BranchId == "" {
		h.log.Error("Branch ID is required in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	res, err := h.ProductClient.GetSaleStatistics(c, &req)
	if err != nil {
		h.log.Error("Error getting sale statistics", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetBranchIncome godoc
// @Summary Get branch income
// @Description Retrieve total income for a branch within a specified date range
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} products.BranchIncomeRes
// @Failure 500 {object} products.Error "Internal server error"
// @Router /statistics/branch-income [get]
func (h *Handler) GetBranchIncome(c *gin.Context) {

	var req products.BranchIncomeReq

	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")
	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.ProductClient.GetBranchIncome(c, &req)
	if err != nil {
		h.log.Error("Error getting branch income", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetClientDashboard godoc
// @Summary Get client dashboard data
// @Description Retrieve dashboard data for a client
// @Tags Statistics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param client_id path string true "Client ID"
// @Param branch_id header string true "branch ID"
// @Success 200 {object} products.GetClientDashboardResponse
// @Failure 500 {object} products.Error "Internal server error"
// @Router /statistics/client-dashboard [get]
func (h *Handler) GetClientDashboard(c *gin.Context) {
	req := products.GetClientDashboardRequest{
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  c.GetHeader("branch_id"),
		ClientId:  c.Param("client_id"),
	}
	if req.BranchId == "" {
		h.log.Error("Branch ID is required in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	res, err := h.ProductClient.GetClientDashboard(c, &req)
	if err != nil {
		h.log.Error("Error getting client dashboard", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
