package handler

import (
	"gateway/internal/entity"
	"gateway/internal/generated/products"
	"gateway/internal/minio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateCategory godoc
// @Summary Create Product Category
// @Description Create a new product category by specifying its name and optionally uploading an image
// @Tags Category
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file false "Upload category image (optional)"
// @Param name formData string true "Name of the category"
// @Success 201 {object} products.Category "Category successfully created"
// @Failure 400 {object} entity.Error "Invalid input or bad request"
// @Failure 500 {object} entity.Error "Internal server error"
// @Router /products/category [post]
func (h *Handler) CreateCategory(c *gin.Context) {

	req := entity.Names{}

	if err := c.ShouldBind(&req); err != nil {
		h.log.Error("bind json err", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	var url string
	file, err := c.FormFile("file")
	if err == nil {
		url, err = minio.UploadMedia(file)
		if err != nil {
			log.Println("Error occurred while uploading file")
			h.log.Error("Error occurred while uploading file:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		url = "no image"
		log.Println("No file uploaded, continuing without an image")
	}

	res, err := h.ProductClient.CreateCategory(c, &products.CreateCategoryRequest{Name: req.Name, CreatedBy: c.MustGet("id").(string), ImageUrl: url, CompanyId: c.MustGet("company_id").(string)})
	if err != nil {
		h.log.Error("Error creating category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateCategory godoc
// @Summary Update Product Category
// @Description Update a product category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Category ID"
// @Param file formData file false "Upload category image (optional)"
// @Param name formData string true "Name of the category"// @Success 200 {object} products.Category
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/category/{id} [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	req := &products.UpdateCategoryRequest{Id: id, CompanyId: c.MustGet("company_id").(string)}

	if err := c.ShouldBindJSON(req); err != nil {
		h.log.Error("Error parsing UpdateCategory request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var url string
	file, err := c.FormFile("file")
	if err == nil {
		url, err = minio.UploadMedia(file)
		if err != nil {
			log.Println("Error occurred while uploading file")
			h.log.Error("Error occurred while uploading file:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		log.Println("No file uploaded, continuing without an image")
	}
	req.ImageUrl = url

	res, err := h.ProductClient.UpdateCategory(c, req)
	if err != nil {
		h.log.Error("Error updating category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
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
	req := &products.GetCategoryRequest{Id: id, CompanyId: c.MustGet("company_id").(string)}

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
	req.CompanyId = c.MustGet("company_id").(string)
	//req.CreatedBy = c.MustGet("id").(string)
	//req.CreatedBy = uuid.New().String()
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
	req := &products.GetCategoryRequest{Id: id, CompanyId: c.MustGet("company_id").(string)}

	res, err := h.ProductClient.DeleteCategory(c, req)
	if err != nil {
		h.log.Error("Error deleting category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
