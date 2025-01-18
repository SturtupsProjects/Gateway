package handler

import (
	"gateway/internal/entity"
	"gateway/internal/generated/products"
	"gateway/internal/minio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CreateCategory godoc
// @Summary Create Product Category
// @Description Create a new product category by specifying its name and optionally uploading an image
// @Tags Category
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id header string true "Branch ID"
// @Param file formData file false "Upload category image (optional)"
// @Param name formData string true "Name of the category"
// @Success 201 {object} products.Category "Category successfully created"
// @Failure 400 {object} entity.Error "Invalid input or bad request"
// @Failure 500 {object} entity.Error "Internal server error"
// @Router /products/category [post]
func (h *Handler) CreateCategory(c *gin.Context) {

	branchID := c.GetHeader("branch_id")
	if branchID == "" {
		h.log.Error("Branch ID is missing in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

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

	res, err := h.ProductClient.CreateCategory(c, &products.CreateCategoryRequest{
		Name:      req.Name,
		CreatedBy: c.MustGet("id").(string),
		ImageUrl:  url,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchID,
	})
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
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id header string true "Branch ID"
// @Param id path string true "Category ID"
// @Param file formData file false "Upload category image (optional)"
// @Param name formData string true "Name of the category"
// @Success 200 {object} products.Category "Category successfully updated"
// @Failure 400 {object} products.Error "Invalid input or bad request"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /products/category/{id} [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	branchID := c.GetHeader("branch_id")
	if branchID == "" {
		h.log.Error("Branch ID is missing in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	id := c.Param("id")
	req := &products.UpdateCategoryRequest{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchID,
	}

	name := entity.Names{}

	if err := c.ShouldBind(&name); err != nil {
		h.log.Error("Error parsing UpdateCategory request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(name.Name)

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
	req.Name = name.Name

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
// @Param branch_id header string true "Branch ID"
// @Param id path string true "Category ID"
// @Success 200 {object} products.Category "Category details"
// @Failure 400 {object} products.Error "Invalid input or bad request"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /products/category/{id} [get]
func (h *Handler) GetCategory(c *gin.Context) {
	branchID := c.GetHeader("branch_id")
	if branchID == "" {
		h.log.Error("Branch ID is missing in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	id := c.Param("id")
	req := &products.GetCategoryRequest{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchID,
	}

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
// @Param branch_id header string true "Branch ID"
// @Param name query string false "Filter by category name"
// @Param limit query string false "Limit"
// @Param offset query string false "Page"
// @Success 200 {object} products.CategoryList "List of categories"
// @Failure 400 {object} products.Error "Bad request due to invalid query parameters"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /products/category [get]
func (h *Handler) GetListCategory(c *gin.Context) {
	branchID := c.GetHeader("branch_id")
	if branchID == "" {
		h.log.Error("Branch ID is missing in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	var req products.CategoryName
	req.Name = c.Query("name")
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 0
	}
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}
	req.Limit = limit
	req.Page = page
	req.CompanyId = c.MustGet("company_id").(string)
	req.BranchId = branchID

	res, err := h.ProductClient.GetListCategory(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("Failed to retrieve category list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
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
// @Security ApiKeyAuth
// @Param branch_id header string true "Branch ID"
// @Param id path string true "Category ID"
// @Success 200 {object} products.Message "Category successfully deleted"
// @Failure 400 {object} products.Error "Invalid input or bad request"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /products/category/{id} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	branchID := c.GetHeader("branch_id")
	if branchID == "" {
		h.log.Error("Branch ID is missing in the header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	id := c.Param("id")
	req := &products.GetCategoryRequest{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchID,
	}

	res, err := h.ProductClient.DeleteCategory(c, req)
	if err != nil {
		h.log.Error("Error deleting category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
