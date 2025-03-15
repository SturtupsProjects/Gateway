package handler

import (
	"context"
	"gateway/internal/generated/debts"
	pbu "gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CreateDebt godoc
// @Summary Create a new debt
// @Description Add a new debt for a client with the specified details
// @Security ApiKeyAuth
// @Tags Debts
// @Accept json
// @Produce json
// @Param data body debts.DebtsRequest true "Debt creation details"
// @Success 201 {object} debts.Debts "Debt successfully created"
// @Failure 400 {object} products.Error "Invalid input or bad request"
// @Failure 500 {object} products.Error "Internal server error"
// @Router /debts [post]
func (h *Handler) CreateDebt(c *gin.Context) {
	var req debts.DebtsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.DebtClient.CreateDebts(c, &req)
	if err != nil {
		h.log.Error("Error creating debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetDebt godoc
// @Summary Get a debt by ID
// @Description Retrieve details of a specific debt by ID
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Debt ID"
// @Success 200 {object} debts.Debts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/{id} [get]
func (h *Handler) GetDebt(c *gin.Context) {
	id := c.Param("id")
	req := &debts.DebtsID{Id: id, CompanyId: c.MustGet("company_id").(string)}

	res, err := h.DebtClient.GetDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListDebts godoc
// @Summary Get list of debts
// @Description Retrieve a list of debts with optional filters
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param description query string false "Filter by description"
// @Param currencyCode query string false "Filter by currency code"
// @Param total_amount_min query number false "Filter by minimum total amount"
// @Param total_amount_max query number false "Filter by maximum total amount"
// @Param limit query int false "Number of results to return"
// @Param page query int false "Page number for pagination"
// @Param is_fully_pay query bool false "filter by fully pay"
// @Success 200 {object} debts.DebtsList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts [get]
func (h *Handler) GetListDebts(c *gin.Context) {
	var filter debts.FilterDebts

	// Extract query parameters
	filter.Description = c.Query("description")
	filter.CurrencyCode = c.Query("currencyCode")
	filter.IsFullyPay = c.Query("is_fully_pay")

	log.Println(filter.IsFullyPay)

	if filter.IsFullyPay == "true" || filter.IsFullyPay == "false" || filter.IsFullyPay == "" {
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for is_fully_pay"})
		return
	}

	// Parse numeric filters
	if totalAmountMin := c.Query("total_amount_min"); totalAmountMin != "" {
		if min, err := strconv.ParseFloat(totalAmountMin, 64); err == nil {
			filter.TotalAmountMin = min
		} else {
			h.log.Error("Invalid totalAmountMin parameter", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid totalAmountMin parameter"})
			return
		}
	}

	if totalAmountMax := c.Query("total_amount_max"); totalAmountMax != "" {
		if max, err := strconv.ParseFloat(totalAmountMax, 64); err == nil {
			filter.TotalAmountMax = max
		} else {
			h.log.Error("Invalid totalAmountMax parameter", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid totalAmountMax parameter"})
			return
		}
	}

	// Parse limit and page for pagination
	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	if limitStr != "" {
		if limit, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			filter.Limit = int32(limit)
		} else {
			h.log.Error("Invalid limit parameter", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if pageStr != "" {
		if page, err := strconv.ParseInt(pageStr, 10, 64); err == nil {
			filter.Page = int32(page)
		} else {
			h.log.Error("Invalid page parameter", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	// Extract company ID from context (must be set by middleware)
	companyID, exists := c.Get("company_id")
	if !exists {
		h.log.Error("Missing company_id in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	filter.CompanyId = companyID.(string)

	// Fetch debts from DebtClient
	res, err := h.DebtClient.GetListDebts(c, &filter)
	if err != nil {
		h.log.Error("Error fetching debt list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch client details for each debt
	for i, debt := range res.Installments {
		client, err := h.UserClient.GetClient(context.Background(), &pbu.UserIDRequest{
			Id:        debt.ClientId,
			CompanyId: debt.CompanyId,
		})
		if err == nil {
			res.Installments[i].ClientName = client.FullName
			res.Installments[i].ClientPhone = client.Phone
		} else {
			h.log.Error("Error fetching client info", "error", err.Error())
			res.Installments[i].ClientName = "Unknown"
			res.Installments[i].ClientPhone = "Unknown"
		}
	}

	// Respond with the results
	c.JSON(http.StatusOK, res)
}

// GetClientDebts godoc
// @Summary Get debts for a client
// @Description Retrieve a list of debts associated with a specific client
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param client_id path string true "Client ID"
// @Success 200 {object} debts.DebtsList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/client/{client_id} [get]
func (h *Handler) GetClientDebts(c *gin.Context) {
	clientID := c.Param("client_id")
	req := &debts.ClientID{Id: clientID}

	res, err := h.DebtClient.GetClientDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching client debts", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// PayDebt godoc
// @Summary Pay a debt
// @Description Make a payment towards a specific debt
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body debts.PayDebtsReq true "Payment details"
// @Success 200 {object} debts.Debts
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/pay [post]
func (h *Handler) PayDebt(c *gin.Context) {
	var req debts.PayDebtsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.DebtClient.PayDebts(c, &req)
	if err != nil {
		h.log.Error("Error processing payment", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPaymentsByDebtId godoc
// @Summary Get payments by debt ID
// @Description Retrieve all payments associated with a specific debt
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param debt_id path string true "Debt ID"
// @Success 200 {object} debts.PaymentList
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/payments/{debt_id} [get]
func (h *Handler) GetPaymentsByDebtId(c *gin.Context) {
	debtId := c.Param("debt_id")
	req := &debts.DebtsID{Id: debtId}

	res, err := h.DebtClient.GetPaymentsByDebtsId(c, req)
	if err != nil {
		h.log.Error("Error fetching payments for debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

//// GetPayments godoc
//// @Summary Get a list of payments
//// @Description Retrieve a list of payments with optional filters
//// @Tags Debts
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param filter query debts.FilterPayment false "Filter parameters"
//// @Success 200 {object} debts.PaymentList
//// @Failure 400 {object} products.Error
//// @Failure 500 {object} products.Error
//// @Router /debts/payment [get]
//func (h *Handler) GetPayments(c *gin.Context) {
//	var filter debts.FilterPayment
//
//	if err := c.ShouldBindQuery(&filter); err != nil {
//		h.log.Error("Error parsing filter", "error", err.Error())
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	filter.CompanyId = c.MustGet("company_id").(string)
//
//	res, err := h.DebtClient.GetPayments(c, &filter)
//	if err != nil {
//		h.log.Error("Error fetching payments list", "error", err.Error())
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}

// GetPayment godoc
// @Summary Get a payment by ID
// @Description Retrieve details of a specific payment by ID
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} debts.Payment
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/payment/{id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("id")
	req := &debts.PaymentID{Id: id, CompanyId: c.MustGet("company_id").(string)}

	res, err := h.DebtClient.GetPayment(c, req)
	if err != nil {
		h.log.Error("Error fetching payment", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTotalDebtSum godoc
// @Summary Get Total Sum Of Debts
// @Description Get total sum of debts from company
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} debts.SumMoney
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/total-sum [get]
func (h *Handler) GetTotalDebtSum(c *gin.Context) {

	companyID := c.MustGet("company_id").(string)

	req := debts.CompanyID{
		Id: companyID,
	}

	res, err := h.DebtClient.GetTotalDebtSum(c, &req)
	if err != nil {
		h.log.Error("Error fetching total debt sum", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserTotalDebt godoc
// @Summary Get Total Debts Sum Of User
// @Description Get total debts sum for a specific user from company
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} debts.SumMoney
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/total-sum/{user_id} [get]
func (h *Handler) GetUserTotalDebt(c *gin.Context) {
	// Извлекаем user_id из URL-параметра
	userID := c.Param("user_id")
	if userID == "" {
		h.log.Error("user_id not provided in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	// Извлекаем company_id из контекста
	companyID := c.MustGet("company_id").(string)

	req := debts.ClientID{
		Id:        userID,
		CompanyId: companyID,
	}

	res, err := h.DebtClient.GetUserTotalDebtSum(c, &req)
	if err != nil {
		h.log.Error("Error fetching user total debt sum", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserPayments godoc
// @Summary Get All Payments for a Client
// @Description Retrieve all payments made by the client over time
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} debts.UserPaymentsRes
// @Failure 400 {object} products.Error
// @Failure 500 {object} products.Error
// @Router /debts/payments/{user_id} [get]
func (h *Handler) GetUserPayments(c *gin.Context) {

	userID := c.Param("user_id")
	if userID == "" {
		h.log.Error("user_id not provided in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	companyVal, exists := c.Get("company_id")
	if !exists {
		h.log.Error("company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})
		return
	}
	companyID, ok := companyVal.(string)
	if !ok || companyID == "" {
		h.log.Error("company_id is not a valid string")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company_id"})
		return
	}

	req := &debts.ClientID{
		Id:        userID,
		CompanyId: companyID,
	}

	res, err := h.DebtClient.GetUserPayments(c, req)
	if err != nil {
		h.log.Error("Error fetching user payments", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, res)
}
