package handler

import (
	"gateway/internal/generated/products"
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
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
	req.CompanyId = c.MustGet("company_id").(string)

	log.Println(req.CompanyId)

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
	req := &products.PurchaseID{Id: id, CompanyId: c.MustGet("company_id").(string)}

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
// @Param filter query entity.FilterPurchase false "Filter parameters"
// @Success 200 {object} products.PurchaseList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases [get]
func (h *Handler) GetListPurchase(c *gin.Context) {
	var filter products.FilterPurchase

	// Parse query parameters into filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing FilterPurchase", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter.CompanyId = c.MustGet("company_id").(string)

	// Fetch the list of purchases
	res, err := h.ProductClient.GetListPurchase(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving purchase list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enhance the response with additional details
	for i, purchase := range res.Purchases {
		// Fetch client (supplier) details
		clientRes, err := h.UserClient.GetClient(c, &user.UserIDRequest{Id: purchase.SupplierId, CompanyId: c.MustGet("company_id").(string)})
		if err == nil {
			res.Purchases[i].SupplierName = clientRes.FullName
		} else {
			h.log.Error("Error fetching supplier details", "supplier_id", purchase.SupplierId, "error", err.Error())
		}

		// Fetch purchaser details for phone number
		purchaserRes, err := h.UserClient.GetClient(c, &user.UserIDRequest{Id: purchase.PurchasedBy, CompanyId: c.MustGet("company_id").(string)})
		if err == nil {
			res.Purchases[i].PurchaserPhoneNumber = purchaserRes.Phone
		} else {
			h.log.Error("Error fetching purchaser details", "purchased_by", purchase.PurchasedBy, "error", err.Error())
		}

		// Fetch product names for each item in the purchase
		for j, item := range purchase.Items {
			productRes, err := h.ProductClient.GetProduct(c, &products.GetProductRequest{
				Id:        item.ProductId,
				CompanyId: filter.CompanyId,
			})
			if err == nil {
				res.Purchases[i].Items[j].ProductName = productRes.Name
				res.Purchases[i].Items[j].ProductImage = productRes.ImageUrl
			} else {
				h.log.Error("Error fetching product details", "product_id", item.ProductId, "error", err.Error())
			}
		}
	}

	// Return the enhanced purchase list response
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
// @Param Purchase body entity.PurchaseUpdate true "Updated purchase data"
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
	req.CompanyId = c.MustGet("company_id").(string)

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
	req := &products.PurchaseID{Id: id, CompanyId: c.MustGet("company_id").(string)}

	res, err := h.ProductClient.DeletePurchase(c, req)
	if err != nil {
		h.log.Error("Error deleting purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
