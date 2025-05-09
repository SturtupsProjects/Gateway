package handler

import (
	"gateway/internal/entity"
	"gateway/internal/generated/debts"
	"gateway/internal/generated/products"
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CalculateTotalSales godoc
// @Summary Calculate total sales
// @Description Calculate the total sales based on the sale request
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Sale body entity.Sale true "Sale data"
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
	req.SoldBy = c.MustGet("id").(string)
	req.CompanyId = c.MustGet("company_id").(string)
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
// @Param branch_id header string true "Branch ID"
// @Param Sale body entity.Sale true "Sale data"
// @Success 201 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales [post]
func (h *Handler) CreateSales(c *gin.Context) {
	var req products.SaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateSales request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req.SoldBy, _ = c.MustGet("id").(string)
	req.CompanyId, _ = c.MustGet("company_id").(string)

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		h.log.Error("Missing branch_id in header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}
	req.BranchId = branchId

	if len(req.ClientId) < 16 {
		if req.ClientName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID or Client Name is required"})
			return
		}

		clientReq := user.ClientRequest{
			FullName:   req.ClientName,
			Address:    "No address",
			Phone:      req.ClientPhone,
			Type:       "client",
			ClientType: "street",
			CompanyId:  req.CompanyId,
		}

		client, err := h.UserClient.CreateClient(c, &clientReq)
		if err != nil {
			h.log.Error("Error creating client for sale", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		req.ClientId = client.Id
		if req.ClientId == "" {
			h.log.Error("Created client has no ID")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	res, err := h.ProductClient.CreateSales(c, &req)
	if err != nil {
		h.log.Error("Error creating sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.IsForDebt {
		debReq := debts.DebtsRequest{
			SaleId:       res.Id,
			CompanyId:    req.CompanyId,
			ClientId:     req.ClientId,
			TotalAmount:  res.TotalSalePrice,
			CurrencyCode: req.PaymentMethod,
			DebtType:     "debtor",
		}

		debtRes, err := h.DebtClient.CreateDebts(c, &debReq)
		if err != nil {
			h.log.Error("Error creating debt", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if debtRes.Id == "" {
			h.log.Error("Created debt has no ID")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if req.PaidAmount > 0 {
			reqPay := debts.PayDebtsReq{
				CompanyId:  req.CompanyId,
				DebtId:     debtRes.Id,
				PaidAmount: req.PaidAmount,
			}

			if _, err = h.DebtClient.PayDebts(c, &reqPay); err != nil {
				h.log.Error("Error processing debt payment", "error", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
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
// @Param Sale body entity.SaleUpdate true "Updated sale data"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/{id} [put]
func (h *Handler) UpdateSales(c *gin.Context) {
	var req products.SaleUpdate
	id := c.Param("id")
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing UpdateSales request body", "error", err.Error())
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
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/{id} [get]
func (h *Handler) GetSales(c *gin.Context) {
	id := c.Param("id")

	// Получаем branch_id из заголовков
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.SaleID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchId,
	}

	res, err := h.ProductClient.GetSales(c, req)
	if err != nil {
		h.log.Error("Error fetching sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListSales godoc
// @Summary Get list of sales
// @Description Retrieve a paginated list of sales with optional filters
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_name query string false "Filter by product_name"
// @Param limit query int false "Number of items per page (default: 10)"
// @Param page query int false "Page number (default: 1)"
// @Param start_date query string false "Start date for filtering (format: YYYY-MM-DD)"
// @Param end_date query string false "End date for filtering (format: YYYY-MM-DD)"
// @Param client_id query string false "Client ID to filter sales"
// @Param sold_by query string false "Sold by user ID to filter sales"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.SaleList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales [get]
func (h *Handler) GetListSales(c *gin.Context) {

	productName := c.Query("product_name")
	limitStr := c.Query("limit") // Значение по умолчанию - 10
	pageStr := c.Query("page")   // Значение по умолчанию - 1
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	clientId := c.Query("client_id")
	soldBy := c.Query("sold_by")
	branchId := c.GetHeader("branch_id") // Получаем из заголовков

	// Преобразуем limit и page в int64
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		h.log.Error("Error parsing limit", "error", err.Error())
		limit = 0
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		h.log.Error("Error parsing page", "error", err.Error())
		page = 0
	}

	// Логируем переданные параметры для отладки
	log.Println("Limit:", limit, "Page:", page, "StartDate:", startDate, "EndDate:", endDate, "ProductName", productName)

	// Проверяем, если branchId пустой, то возвращаем ошибку
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	// Получаем company_id из контекста
	companyId := c.MustGet("company_id").(string)

	// Создаем фильтр для передачи в запрос
	filter := products.SaleFilter{
		ProductName: productName,
		BranchId:    branchId,
		Limit:       limit,
		Page:        page,
		CompanyId:   companyId,
		SoldBy:      soldBy,
		ClientId:    clientId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	// Получаем список продаж с учетом фильтрации
	res, err := h.ProductClient.GetListSales(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving sales list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Дополнительная информация о клиентах, продавцах и товарах
	for i, sale := range res.Sales {
		clientRes, err := h.UserClient.GetClient(c, &user.UserIDRequest{Id: sale.ClientId, CompanyId: companyId})
		if err == nil {
			res.Sales[i].ClientName = clientRes.FullName
			res.Sales[i].ClientPhoneNumber = clientRes.Phone
		} else {
			h.log.Error("Error fetching customer details", "customer_id", sale.ClientId, "error", err.Error())
		}

		supplier, err := h.UserClient.GetUser(c, &user.UserIDRequest{Id: sale.SoldBy, CompanyId: companyId})
		if err == nil {
			res.Sales[i].SoldByName = supplier.FirstName
		} else {
			h.log.Error("Error fetching customer details", "customer_id", sale.SoldBy, "error", err.Error())
		}
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
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.Message
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales/{id} [delete]
func (h *Handler) DeleteSales(c *gin.Context) {
	id := c.Param("id")

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req := &products.SaleID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
		BranchId:  branchId,
	}

	res, err := h.ProductClient.DeleteSales(c, req)
	if err != nil {
		h.log.Error("Error deleting sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Payments godoc
// @Summary Bought products for cash or debt
// @Description Retrieve a list of payments
// @Tags Payments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Sale body entity.PaymentSale true "Sale data"
// @Success 200 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/payments [post]
func (h *Handler) Payments(c *gin.Context) {
	var req entity.PaymentSale

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CalculateTotalSales request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.PaymentMethod == "cash" {
		res, err := h.ProductClient.CalculateTotalSales(c, &products.SaleRequest{
			ClientId:      req.ClientId,
			SoldBy:        c.MustGet("id").(string),
			CompanyId:     c.MustGet("company_id").(string),
			PaymentMethod: req.PaymentMethod,
			BranchId:      req.BranchId,
			SoldProducts:  req.SoldProducts,
		})
		if err != nil {
			h.log.Error("Error calculating total sales", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}
	if req.PaymentMethod == "debt" {

		res, err := h.ProductClient.CalculateTotalSales(c, &products.SaleRequest{
			ClientId:      req.ClientId,
			SoldBy:        c.MustGet("id").(string),
			CompanyId:     c.MustGet("company_id").(string),
			PaymentMethod: req.PaymentMethod,
			BranchId:      req.BranchId,
			SoldProducts:  req.SoldProducts,
		})
		if err != nil {
			h.log.Error("Error calculating total sales", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res1, err := h.DebtClient.CreateDebts(c, &debts.DebtsRequest{
			ClientId:     req.ClientId,
			TotalAmount:  res.TotalSalePrice,
			CurrencyCode: req.CurrencyCode,
			CompanyId:    c.MustGet("company_id").(string),
		})
		if err != nil {
			h.log.Error("Error creating debt", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !req.IsFullyDebt {
			res2, err := h.DebtClient.PayDebts(c, &debts.PayDebtsReq{
				DebtId:     res1.Id,
				PaidAmount: req.PaidAmount,
				CompanyId:  c.MustGet("company_id").(string),
			})
			if err != nil {
				h.log.Error("Error processing payment", "error", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"sale":    res,
				"debt":    res1,
				"payment": res2,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"sale": res,
			"debt": res1,
		})
		return
	}

}
