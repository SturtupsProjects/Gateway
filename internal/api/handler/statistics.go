package handler

import (
	"gateway/internal/generated/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TotalPriceOfProducts godoc
// @Summary Calculate the total price of products
// @Description Calculate the total price of all products for a specific company
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Company ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/total-price [post]
func (h *Handler) TotalPriceOfProducts(c *gin.Context) {

	// Call the gRPC method
	res, err := h.ProductClient.TotalPriceOfProducts(c, &products.CompanyID{Id: c.MustGet("company_id").(string)})
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
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param CompanyID body products.CompanyID true "Company ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/total-sold [post]
func (h *Handler) TotalSoldProducts(c *gin.Context) {
	// Call the gRPC method
	res, err := h.ProductClient.TotalSoldProducts(c, &products.CompanyID{Id: c.MustGet("company_id").(string)})
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
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param CompanyID body products.CompanyID true "Company ID"
// @Success 200 {object} products.PriceProducts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/total-purchased [post]
func (h *Handler) TotalPurchaseProducts(c *gin.Context) {

	// Call the gRPC method
	res, err := h.ProductClient.TotalPurchaseProducts(c, &products.CompanyID{Id: c.MustGet("company_id").(string)})
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
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param MostSoldProductsRequest body entity.MostSoldProductsRequest true "Date Range and Company ID"
// @Success 200 {object} products.MostSoldProductsResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/most-sold-by-day [post]
func (h *Handler) GetMostSoldProductsByDay(c *gin.Context) {
	var req products.MostSoldProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing GetMostSoldProductsByDay request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CompanyId = c.MustGet("company_id").(string)

	// Call the gRPC method
	res, err := h.ProductClient.GetMostSoldProductsByDay(c, &req)
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
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param GetTopEntitiesRequest body entity.GetTopEntitiesRequest true "Date Range and Company ID"
// @Success 200 {object} products.GetTopEntitiesResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/top-clients [post]
func (h *Handler) GetTopClients(c *gin.Context) {
	var req products.GetTopEntitiesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing GetTopClients request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CompanyId = c.MustGet("company_id").(string)

	// Call the gRPC method
	res, err := h.ProductClient.GetTopClients(c, &req)
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
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param GetTopEntitiesRequest body entity.GetTopEntitiesRequest true "Date Range and Company ID"
// @Success 200 {object} products.GetTopEntitiesResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/top-suppliers [post]
func (h *Handler) GetTopSuppliers(c *gin.Context) {
	var req products.GetTopEntitiesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing GetTopSuppliers request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CompanyId = c.MustGet("company_id").(string)

	// Call the gRPC method
	res, err := h.ProductClient.GetTopSuppliers(c, &req)
	if err != nil {
		h.log.Error("Error getting top suppliers", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
