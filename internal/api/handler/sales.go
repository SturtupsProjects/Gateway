package handler

import (
	"gateway/internal/entity"
	"gateway/internal/generated/debts"
	"gateway/internal/generated/products"
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
// @Param Sale body entity.Sale true "Sale data"
// @Param branch_id header string true "Branch ID"
// @Success 201 {object} products.SaleResponse
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales [post]
func (h *Handler) CreateSales(c *gin.Context) {
	var req products.SaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateSales request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID := c.MustGet("company_id").(string)

	if len(req.ClientId) < 16 {
		if req.ClientName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing client id or client_name"})
			return
		}
		if req.ClientPhone == "" {
			req.ClientPhone = "no phone"
		}

		clientReq := user.ClientRequest{
			FullName:   req.ClientName,
			Address:    "no address",
			Phone:      req.ClientPhone,
			Type:       "client",
			ClientType: "street",
			CompanyId:  companyID,
		}

		client, err := h.UserClient.CreateClient(c, &clientReq)
		if err != nil {
			h.log.Error("Error creating sales", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Println(client.Id)

		req.ClientId = client.Id
	}

	req.SoldBy = c.MustGet("id").(string)
	req.CompanyId = companyID

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}
	req.BranchId = branchId

	res, err := h.ProductClient.CreateSales(c, &req)
	if err != nil {
		h.log.Error("Error creating sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
// @Summary List all sales
// @Description Retrieve a list of sales with optional filters
// @Tags Sales
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query entity.SaleFilter false "Filter parameters"
// @Param branch_id header string true "Branch ID"
// @Success 200 {object} products.SaleList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /sales [get]
func (h *Handler) GetListSales(c *gin.Context) {
	var filter products.SaleFilter

	// Преобразуем параметры фильтра из запроса
	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing SaleFilter", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем branch_id из заголовков
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	filter.CompanyId = c.MustGet("company_id").(string)
	filter.BranchId = branchId

	// Получаем список продаж с учетом фильтрации
	res, err := h.ProductClient.GetListSales(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving sales list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Дополнительная информация о клиентах, продавцах и товарах
	for i, sale := range res.Sales {
		clientRes, err := h.UserClient.GetClient(c, &user.UserIDRequest{Id: sale.ClientId, CompanyId: c.MustGet("company_id").(string)})
		if err == nil {
			res.Sales[i].ClientName = clientRes.FullName
			res.Sales[i].ClientPhoneNumber = clientRes.Phone
		} else {
			h.log.Error("Error fetching customer details", "customer_id", sale.ClientId, "error", err.Error())
		}

		supplier, err := h.UserClient.GetUser(c, &user.UserIDRequest{Id: sale.SoldBy, CompanyId: c.MustGet("company_id").(string)})
		if err == nil {
			res.Sales[i].SoldByName = supplier.FirstName
		} else {
			h.log.Error("Error fetching customer details", "customer_id", sale.SoldBy, "error", err.Error())
		}

		// Информация о проданных товарах
		for j, item := range sale.SoldProducts {
			productRes, err := h.ProductClient.GetProduct(c, &products.GetProductRequest{
				Id:        item.ProductId,
				CompanyId: filter.CompanyId,
			})
			if err == nil {
				res.Sales[i].SoldProducts[j].ProductName = productRes.Name
				res.Sales[i].SoldProducts[j].ProductImage = productRes.ImageUrl
			} else {
				h.log.Error("Error fetching customer details", "customer_id", filter.CompanyId, "error", err.Error())
			}
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
			ClientName:    req.ClientName,
			ClientPhone:   req.ClientPhone,
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
			ClientName:    req.ClientName,
			ClientPhone:   req.ClientPhone,
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
