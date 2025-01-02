package handler

import (
	"bytes"
	"gateway/internal/entity"
	"gateway/internal/generated/products"
	"gateway/internal/minio"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"

	"io"
	"log"
	"net/http"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details, including an optional file upload for the product image
// @Security ApiKeyAuth
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload product image (optional)"
// @Param category_id formData string true "ID of the product category"
// @Param name formData string true "Name of the product"
// @Param bill_format formData string true "Billing format of the product"
// @Param incoming_price formData float64 true "Incoming price of the product"
// @Param standard_price formData float64 true "Standard price of the product"
// @Success 201 {object} products.Product "Product successfully created"
// @Failure 400 {object} products.Error "Invalid input or bad request"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {

	var req entity.CreateProductRequest

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

	res, err := h.ProductClient.CreateProduct(c, &products.CreateProductRequest{CreatedBy: c.MustGet("id").(string),
		CategoryId:    req.CategoryID,
		Name:          req.Name,
		BillFormat:    req.BillFormat,
		IncomingPrice: req.IncomingPrice, StandardPrice: req.StandardPrice, ImageUrl: url, CompanyId: c.MustGet("company_id").(string)})
	if err != nil {
		h.log.Error("Error creating product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update the details of an existing product by ID, with optional media upload
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Param file formData file false "Upload product image (optional)"
// @Param name formData string true "Name of the product"
// @Param category_id formData string true "Category ID"
// @Param bill_format formData string false "Billing format"
// @Param incoming_price formData int64 true "Incoming price"
// @Param standard_price formData int64 true "Standard price"
// @Success 200 {object} products.Product
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var form UpdateProductForm
	if err := c.ShouldBind(&form); err != nil {
		h.log.Error("Error binding form data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var url string
	file, err := c.FormFile("file")
	if err == nil {
		// Upload the media file if provided
		url, err = minio.UploadMedia(file)
		if err != nil {
			h.log.Error("Error occurred while uploading file", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		h.log.Info("No file uploaded, continuing without an image")
	}

	// Build the update request
	req := products.UpdateProductRequest{
		Id:            id,
		CompanyId:     c.MustGet("company_id").(string),
		Name:          form.Name,
		CategoryId:    form.CategoryId,
		BillFormat:    form.BillFormat,
		IncomingPrice: form.IncomingPrice,
		StandardPrice: form.StandardPrice,
		ImageUrl:      url,
	}

	// Call the product service to update the product
	res, err := h.ProductClient.UpdateProduct(c, &req)
	if err != nil {
		h.log.Error("Error updating product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateProductForm defines the structure for form data binding
type UpdateProductForm struct {
	Name          string  `form:"name" binding:"required"`                   // Name of the product
	CategoryId    string  `form:"category_id" binding:"required"`            // ID of the product category
	BillFormat    string  `form:"bill_format"`                               // Optional billing format
	IncomingPrice float64 `form:"incoming_price" binding:"required,numeric"` // Incoming price of the product
	StandardPrice float64 `form:"standard_price" binding:"required,numeric"` // Standard price of the product
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
	req := &products.GetProductRequest{Id: id, CompanyId: c.MustGet("company_id").(string)}

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
	req := &products.GetProductRequest{Id: id, CompanyId: c.MustGet("company_id").(string)}

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
// @Param filter query entity.ProductFilter false "Filter parameters"
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

	filter.CompanyId = c.MustGet("company_id").(string)

	log.Println(filter.CompanyId)

	res, err := h.ProductClient.GetProductList(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving product list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UploadAndProcessExcel godoc
// @Summary Upload an Excel file and create products
// @Description Upload an Excel file containing product data, process it, and create products in bulk
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "Excel file containing products data"
// @Success 200 {object} entity.Error
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/excel-upload [post]
func (h *Handler) UploadAndProcessExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		h.log.Error("Error retrieving file", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Open the file
	fileContent, err := file.Open()
	if err != nil {
		h.log.Error("Error opening file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
		return
	}
	defer fileContent.Close()

	// Read the Excel file
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, fileContent); err != nil {
		h.log.Error("Error reading file content", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
		return
	}
	excelFile, err := excelize.OpenReader(buffer)
	if err != nil {
		h.log.Error("Error parsing Excel file", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Read the "Sheet1"
	rows, err := excelFile.GetRows("Оценка текущих товарных запасов")
	if err != nil {
		h.log.Error("Error reading sheet", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Process rows to create products
	var createdProducts []products.Product
	for i, row := range rows {
		// Skip header row
		if i == 0 {
			continue
		}
		if len(row) < 8 { // Ensure the row has all required columns
			h.log.Warn("Incomplete data in row", "row", i+1)
			continue
		}

		req := &products.CreateProductRequest{
			CategoryId:    "fec5dd49-9e44-4ee4-8499-4e332fac26a7", // Код (Column A)
			Name:          row[2],                                 // Наименование (Column B)
			BillFormat:    row[4],                                 // Ед. изм (Column D)
			IncomingPrice: parseToFloat64(row[7]),                 // Сумма себестоим. (Column G)
			StandardPrice: parseToFloat64(row[7]) * 1.1,           // Цена (Column H)
			ImageUrl:      "no image",                             // Default value
			CompanyId:     c.MustGet("company_id").(string),
			CreatedBy:     c.MustGet("id").(string),
		}

		// Call the product service
		product, err := h.ProductClient.CreateProduct(c, req)
		if err != nil {
			h.log.Error("Error creating product from row", "row", i+1, "error", err)
			continue
		}
		createdProducts = append(createdProducts, *product)
	}

	// Respond with created products
	c.JSON(http.StatusOK, gin.H{"message": "Products created successfully", "products": createdProducts})
}

// Helper function to parse string to float64
func parseToFloat64(value string) float64 {
	value = strings.ReplaceAll(value, ",", "") // Remove commas
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return parsedValue
}
