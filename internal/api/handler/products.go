package handler

import (
	"gateway/internal/entity"
	"gateway/internal/generated/products"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Product body entity.CreateProductRequest true "Product data"
// @Success 201 {object} products.Product
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var req entity.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateProduct request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.CreateProduct(c, &products.CreateProductRequest{CreatedBy: c.MustGet("id").(string), CategoryId: req.CategoryID, Name: req.Name,
		BillFormat: req.BillFormat, IncomingPrice: req.IncomingPrice, StandardPrice: req.StandardPrice})
	if err != nil {
		h.log.Error("Error creating product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update the details of an existing product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Param Product body products.UpdateProductRequest true "Updated product data"
// @Success 200 {object} products.Product
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req products.UpdateProductRequest
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing UpdateProduct request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.UpdateProduct(c, &req)
	if err != nil {
		h.log.Error("Error updating product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} products.Message
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	req := &products.GetProductRequest{Id: id}

	res, err := h.ProductClient.DeleteProduct(c, req)
	if err != nil {
		h.log.Error("Error deleting product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProduct godoc
// @Summary Get a product
// @Description Retrieve a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} products.Product
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	req := &products.GetProductRequest{Id: id}

	res, err := h.ProductClient.GetProduct(c, req)
	if err != nil {
		h.log.Error("Error fetching product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProductList godoc
// @Summary List all products
// @Description Retrieve a list of products with optional filters
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query products.ProductFilter false "Filter parameters"
// @Success 200 {object} products.ProductList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products [get]
func (h *Handler) GetProductList(c *gin.Context) {
	var filter products.ProductFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing ProductFilter", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.GetProductList(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving product list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
