package handler

import (
	"gateway/internal/generated/products"
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CreatePurchase godoc
// @Summary Create a new purchase
// @Description Create a new purchase with the provided details
// @Tags Purchases
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Purchase body entity.Purchase true "Purchase data"
// @Param branch_id header string true "Branch ID"
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

	// Получение branch_id из заголовка
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}
	req.BranchId = branchId

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
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PurchaseResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases/{id} [get]
func (h *Handler) GetPurchase(c *gin.Context) {
	id := c.Param("id")
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.PurchaseID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchId,
	}

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
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.PurchaseList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases [get]
func (h *Handler) GetListPurchase(c *gin.Context) {
	var filter products.FilterPurchase

	// Извлекаем параметры запроса индивидуально
	productName := c.Query("product_name")
	supplierId := c.Query("supplier_id")
	purchasedBy := c.Query("purchased_by")
	companyId := c.MustGet("company_id").(string)
	createdAt := c.Query("created_at")
	branchId := c.GetHeader("branch_id") // Получаем из заголовков
	description := c.Query("description")

	// Пагинация
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	// Преобразуем limit и page в int64
	var limit int64 = 0 // Значение по умолчанию
	var page int64 = 0  // Значение по умолчанию

	if limitStr != "" {
		var err error
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			h.log.Error("Error parsing limit", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}

	if pageStr != "" {
		var err error
		page, err = strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			h.log.Error("Error parsing page", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	// Логируем параметры запроса для отладки
	log.Println("ProductId:", productName, "SupplierId:", supplierId, "PurchasedBy:", purchasedBy, "CreatedAt:", createdAt, "BranchId:", branchId, "Limit:", limit, "Page:", page, "ProductName", productName)

	// Проверяем наличие branchId
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	// Создаем фильтр
	filter = products.FilterPurchase{
		ProductName: productName,
		SupplierId:  supplierId,
		PurchasedBy: purchasedBy,
		CompanyId:   companyId,
		CreatedAt:   createdAt,
		BranchId:    branchId,
		Limit:       limit,
		Page:        page,
		Description: description,
	}

	// Получаем список покупок с учетом фильтрации
	res, err := h.ProductClient.GetListPurchase(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving purchase list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Дополнительная информация (например, поставщик и покупатель)
	for i, purchase := range res.Purchases {
		// Получаем информацию о поставщике
		clientRes, err := h.UserClient.GetClient(c, &user.UserIDRequest{Id: purchase.SupplierId, CompanyId: filter.CompanyId})
		if err == nil {
			res.Purchases[i].SupplierName = clientRes.FullName
		} else {
			h.log.Error("Error fetching supplier details", "supplier_id", purchase.SupplierId, "error", err.Error())
		}

		// Получаем информацию о покупателе
		purchaserRes, err := h.UserClient.GetClient(c, &user.UserIDRequest{Id: purchase.PurchasedBy, CompanyId: filter.CompanyId})
		if err == nil {
			res.Purchases[i].PurchaserPhoneNumber = purchaserRes.Phone
		} else {
			h.log.Error("Error fetching purchaser details", "purchased_by", purchase.PurchasedBy, "error", err.Error())
		}
	}

	// Возвращаем ответ
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
// @Param branch_id header string true "Branch ID"
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

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}
	req.BranchId = branchId

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
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.Message
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /purchases/{id} [delete]
func (h *Handler) DeletePurchase(c *gin.Context) {
	id := c.Param("id")

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.PurchaseID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchId,
	}

	res, err := h.ProductClient.DeletePurchase(c, req)
	if err != nil {
		h.log.Error("Error deleting purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
