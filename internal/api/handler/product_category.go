package handler

import (
	"gateway/pkg/generated/products"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
)

// CreateCategory godoc
// @Summary Create Product Category
// @Description Create a new product category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body products.CreateCategoryRequest true "Category data"
// @Success 201 {object} products.Category
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/category [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	var req products.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ProductClient.CreateCategory(c, &req)
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
// @Success 200 {array} products.CategoryList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/category [get]
func (h *Handler) GetListCategory(c *gin.Context) {
	res, err := h.ProductClient.GetListCategory(c.Request.Context(), &empty.Empty{})
	if err != nil {
		h.log.Error("Error retrieving category list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteCategory godoc
// @Summary Delete Product Category
// @Description Delete a product category by ID
// @Tags Category
// @Accept json
// @Produce json
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
