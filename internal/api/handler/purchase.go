package handler

import (
	"gateway/pkg/generated/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreatePurchase godoc
// @Summary Create a new purchase
// @Description Create a new purchase with the provided details
// @Tags Purchases
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Purchase body entity.Purchase true "Purchase data"
// @Success 201 {object} products.PurchaseResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases [post]
func (h *Handler) CreatePurchase(c *gin.Context) {
	var req products.PurchaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreatePurchase request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.PurchasedBy = c.MustGet("id").(string)

	res, err := h.ProductClient.CreatePurchase(c, &req)
	if err != nil {
		h.log.Error("Error creating purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetPurchase godoc
// @Summary Get a purchase
// @Description Retrieve a purchase by ID
// @Tags Purchases
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Purchase ID"
// @Success 200 {object} products.PurchaseResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases/{id} [get]
func (h *Handler) GetPurchase(c *gin.Context) {
	id := c.Param("id")
	req := &products.PurchaseID{Id: id}

	res, err := h.ProductClient.GetPurchase(c, req)
	if err != nil {
		h.log.Error("Error fetching purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListPurchase godoc
// @Summary List all purchases
// @Description Retrieve a list of purchases with optional filters
// @Tags Purchases
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query products.FilterPurchase false "Filter parameters"
// @Success 200 {object} products.PurchaseList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases [get]
func (h *Handler) GetListPurchase(c *gin.Context) {
	var filter products.FilterPurchase

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing FilterPurchase", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.GetListPurchase(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving purchase list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdatePurchase godoc
// @Summary Update an existing purchase
// @Description Update the details of an existing purchase by ID
// @Tags Purchases
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Purchase ID"
// @Param Purchase body products.PurchaseUpdate true "Updated purchase data"
// @Success 200 {object} products.PurchaseResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases/{id} [put]
func (h *Handler) UpdatePurchase(c *gin.Context) {
	id := c.Param("id")
	var req products.PurchaseUpdate
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing UpdatePurchase request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.UpdatePurchase(c, &req)
	if err != nil {
		h.log.Error("Error updating purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeletePurchase godoc
// @Summary Delete a purchase
// @Description Delete a purchase by ID
// @Tags Purchases
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Purchase ID"
// @Success 200 {object} products.Message
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases/{id} [delete]
func (h *Handler) DeletePurchase(c *gin.Context) {
	id := c.Param("id")
	req := &products.PurchaseID{Id: id}

	res, err := h.ProductClient.DeletePurchase(c, req)
	if err != nil {
		h.log.Error("Error deleting purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
