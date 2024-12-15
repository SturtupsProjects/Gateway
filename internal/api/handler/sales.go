package handler

import (
	"gateway/pkg/generated/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CalculateTotalSales godoc
// @Summary Calculate total sales
// @Description Calculate the total sales based on the sale request
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Sale body products.SaleRequest true "Sale data"
// @Success 200 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/calculate [post]
func (h *Handler) CalculateTotalSales(c *gin.Context) {
	var req products.SaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CalculateTotalSales request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.CalculateTotalSales(c, &req)
	if err != nil {
		h.log.Error("Error calculating total sales", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateSales godoc
// @Summary Create a new sale
// @Description Create a new sale with the provided details
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Sale body products.SaleRequest true "Sale data"
// @Success 201 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales [post]
func (h *Handler) CreateSales(c *gin.Context) {
	var req products.SaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateSales request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.CreateSales(c, &req)
	if err != nil {
		h.log.Error("Error creating sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateSales godoc
// @Summary Update an existing sale
// @Description Update the details of an existing sale by ID
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Sale ID"
// @Param Sale body products.SaleUpdate true "Updated sale data"
// @Success 200 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/{id} [put]
func (h *Handler) UpdateSales(c *gin.Context) {
	id := c.Param("id")
	var req products.SaleUpdate
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing UpdateSales request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.UpdateSales(c, &req)
	if err != nil {
		h.log.Error("Error updating sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetSales godoc
// @Summary Get a sale
// @Description Retrieve a sale by ID
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Sale ID"
// @Success 200 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/{id} [get]
func (h *Handler) GetSales(c *gin.Context) {
	id := c.Param("id")
	req := &products.SaleID{Id: id}

	res, err := h.ProductClient.GetSales(c, req)
	if err != nil {
		h.log.Error("Error fetching sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListSales godoc
// @Summary List all sales
// @Description Retrieve a list of sales with optional filters
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query products.SaleFilter false "Filter parameters"
// @Success 200 {object} products.SaleList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales [get]
func (h *Handler) GetListSales(c *gin.Context) {
	var filter products.SaleFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing SaleFilter", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.GetListSales(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving sales list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteSales godoc
// @Summary Delete a sale
// @Description Delete a sale by ID
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Sale ID"
// @Success 200 {object} products.Message
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/{id} [delete]
func (h *Handler) DeleteSales(c *gin.Context) {
	id := c.Param("id")
	req := &products.SaleID{Id: id}

	res, err := h.ProductClient.DeleteSales(c, req)
	if err != nil {
		h.log.Error("Error deleting sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
