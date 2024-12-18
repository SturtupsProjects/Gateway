package handler

import (
	"fmt"
	"gateway/internal/entity"
	"gateway/pkg/generated/products"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateCategory godoc
// @Summary Create Product Category
// @Description Create a new product category by specifying its name
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Category body entity.Names true "Category data"
// @Success 201 {object} products.Category
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/category [post]
func (h *Handler) CreateCategory(c *gin.Context) {

	req := entity.Names{}

	// Bind the JSON request to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	res, err := h.ProductClient.CreateCategory(c, &products.CreateCategoryRequest{Name: req.Name, CreatedBy: c.MustGet("id").(string)})
	if err != nil {
		h.log.Error("Error creating category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetCategory godoc
// @Summary Get Product Category
// @Description Retrieve a product category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Category ID"
// @Success 200 {object} products.Category
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/category/{id} [get]
func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	req := &products.GetCategoryRequest{Id: id}

	res, err := h.ProductClient.GetCategory(c, req)
	if err != nil {
		h.log.Error("Error fetching category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListCategory godoc
// @Summary List Product Categories
// @Description Retrieve a list of product categories
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "Filter by category name"
// @Success 200 {object} products.CategoryList "List of categories"
// @Failure 400 {object} products.Error "Bad request due to invalid query parameters"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /products/category [get]
func (h *Handler) GetListCategory(c *gin.Context) {
	var req products.CategoryName
	req.Name = c.Query("name")
	//req.CreatedBy = c.MustGet("id").(string)
	req.CreatedBy = uuid.New().String()
	fmt.Println(req.CreatedBy)
	// Call the ProductClient to get the list of categories
	res, err := h.ProductClient.GetListCategory(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("Failed to retrieve category list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	// Return the successful response
	c.JSON(http.StatusOK, res)
}

// DeleteCategory godoc
// @Summary Delete Product Category
// @Description Delete a product category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Category ID"
// @Success 200 {object} products.Message
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/category/{id} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	req := &products.GetCategoryRequest{Id: id}

	res, err := h.ProductClient.DeleteCategory(c, req)
	if err != nil {
		h.log.Error("Error deleting category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
