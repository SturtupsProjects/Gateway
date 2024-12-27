package handler

import (
	"gateway/internal/generated/products"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/total-price [get]
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

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
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
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/total-sold [get]
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

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
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
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/total-purchased [get]
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

	req := &products.StatisticReq{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
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
// @Success 200 {object} products.MostSoldProductsResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/most-sold [get]
func (h *Handler) GetMostSoldProductsByDay(c *gin.Context) {

	log.Println("Mana keldi")

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

	req := &products.MostSoldProductsRequest{
		CompanyId: companyId,
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
// @Success 200 {object} products.GetTopEntitiesResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/top-clients [get]
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

	req := &products.GetTopEntitiesRequest{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339), // Переводим в RFC3339 для передачи
		EndDate:   parsedEndDate.Format(time.RFC3339),
	}

	res, err := h.ProductClient.GetTopClients(c, req)
	if err != nil {
		h.log.Error("Error getting top clients", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
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
// @Success 200 {object} products.GetTopEntitiesResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/top-suppliers [get]
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

	req := &products.GetTopEntitiesRequest{
		CompanyId: companyId,
		StartDate: parsedStartDate.Format(time.RFC3339),
		EndDate:   parsedEndDate.Format(time.RFC3339),
	}

	res, err := h.ProductClient.GetTopSuppliers(c, req)
	if err != nil {
		h.log.Error("Error getting top suppliers", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
