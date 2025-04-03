package handler

import (
	"bytes"
	"context"
	"gateway/internal/generated/debts"
	pbu "gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"
	"strings"
)

// CreateDebt godoc
// @Summary Create debtor
// @Description Create a new debtor record.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body entity.DebtsRequest true "Debtor details"
// @Success 201 {object} debts.Debts "Created debtor record"
// @Failure 400 {object} products.Error "Invalid input"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts [post]
func (h *Handler) CreateDebt(c *gin.Context) {
	var req debts.DebtsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.DebtType = "debtor"

	res, err := h.DebtClient.CreateDebts(c, &req)
	if err != nil {
		h.log.Error("Error creating debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetDebt godoc
// @Summary Get debtor by ID
// @Description Retrieve debtor details using its ID.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Debtor ID"
// @Success 200 {object} debts.Debts "Debtor details"
// @Failure 400 {object} products.Error "Invalid ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/{id} [get]
func (h *Handler) GetDebt(c *gin.Context) {
	id := c.Param("id")

	req := &debts.DebtsID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
	}

	res, err := h.DebtClient.GetDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListDebts godoc
// @Summary List debtor records
// @Description Retrieve debtor records with optional filters.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param is_fully_pay query bool false "Filter by fully paid status"
// @Param currency_code query string false "Filter by currency code"
// @Param description query string false "Filter by description"
// @Param no_paid_debts query bool false "Filter by unpaid debts"
// @Param limit query int false "Maximum results"
// @Param page query int false "Page number for pagination"
// @Success 200 {object} debts.DebtsList "List of debtor records"
// @Failure 400 {object} products.Error "Invalid filter value"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts [get]
func (h *Handler) GetListDebts(c *gin.Context) {
	var filter debts.FilterDebts

	filter.IsFullyPay = c.Query("is_fully_pay")
	filter.CurrencyCode = c.Query("currency_code")
	filter.Description = c.Query("description")

	noPaidDebts := c.Query("no_paid_debts")
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	// Validate is_fully_pay value
	if filter.IsFullyPay != "true" && filter.IsFullyPay != "false" && filter.IsFullyPay != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for is_fully_pay"})
		return
	}

	if noPaidDebts == "true" {
		filter.NoPaidDebt = true
	}

	if limitStr != "" {
		if limit, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			filter.Limit = int32(limit)
		} else {
			h.log.Error("Invalid limit parameter", "value", limitStr, "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if pageStr != "" {
		if page, err := strconv.ParseInt(pageStr, 10, 64); err == nil {
			filter.Page = int32(page)
		} else {
			h.log.Error("Invalid page parameter", "value", pageStr, "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	filter.CompanyId = c.MustGet("company_id").(string)
	filter.DebtType = "debtor"

	res, err := h.DebtClient.GetListDebts(c, &filter)
	if err != nil {
		h.log.Error("Error fetching debt list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enrich each record with client details
	for i, debt := range res.Installments {
		client, err := h.UserClient.GetClient(context.Background(), &pbu.UserIDRequest{
			Id:        debt.ClientId,
			CompanyId: debt.CompanyId,
		})
		if err == nil {
			res.Installments[i].ClientName = client.FullName
			res.Installments[i].ClientPhone = client.Phone
		} else {
			h.log.Error("Error fetching client info", "client_id", debt.ClientId, "error", err.Error())
			res.Installments[i].ClientName = "Unknown"
			res.Installments[i].ClientPhone = "Unknown"
		}
	}

	c.JSON(http.StatusOK, res)
}

// GetClientDebts godoc
// @Summary Get client debtor records
// @Description Retrieve debtor records for a specific client.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param client_id path string true "Client ID"
// @Success 200 {object} debts.DebtsList "Debtor records for the client"
// @Failure 400 {object} products.Error "Invalid client ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/client/{client_id} [get]
func (h *Handler) GetClientDebts(c *gin.Context) {
	clientID := c.Param("client_id")
	req := &debts.ClientID{
		Id:        clientID,
		CompanyId: c.MustGet("company_id").(string),
		DebtType:  "debtor",
	}

	res, err := h.DebtClient.GetClientDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching client debts", "client_id", clientID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// PayDebt godoc
// @Summary Process debtor payment
// @Description Make a payment toward a debtor record.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body entity.PayDebtReq true "Payment details"
// @Success 200 {object} debts.Debts "Updated debtor record"
// @Failure 400 {object} products.Error "Invalid input"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/pay [post]
func (h *Handler) PayDebt(c *gin.Context) {
	var req debts.PayDebtsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.PayType = "in"

	res, err := h.DebtClient.PayDebts(c, &req)
	if err != nil {
		h.log.Error("Error processing payment", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPaymentsByDebtId godoc
// @Summary List payments by debtor ID
// @Description Retrieve all payments for a specific debtor record.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param debt_id path string true "Debtor ID"
// @Success 200 {object} debts.PaymentList "List of payments"
// @Failure 400 {object} products.Error "Invalid debtor ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/payments/{debt_id} [get]
func (h *Handler) GetPaymentsByDebtId(c *gin.Context) {
	debtId := c.Param("debt_id")
	req := &debts.PayDebtsID{
		Id:      debtId,
		PayType: "in",
	}

	res, err := h.DebtClient.GetPaymentsByDebtsId(c, req)
	if err != nil {
		h.log.Error("Error fetching payments for debt", "debt_id", debtId, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPayment godoc
// @Summary Get payment details
// @Description Retrieve details for a specific payment.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} debts.Payment "Payment details"
// @Failure 400 {object} products.Error "Invalid payment ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/payment/{id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	req := &debts.PaymentID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
	}

	res, err := h.DebtClient.GetPayment(c, req)
	if err != nil {
		h.log.Error("Error fetching payment", "payment_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTotalDebtSum godoc
// @Summary Get total debtor sum
// @Description Retrieve the total amount of debtor records for the company.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} debts.SumMoney "Total debtor sum"
// @Failure 400 {object} products.Error "Bad request"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/total-sum [get]
func (h *Handler) GetTotalDebtSum(c *gin.Context) {
	companyID := c.MustGet("company_id").(string)

	req := debts.CompanyID{
		Id:       companyID,
		DebtType: "debtor",
	}

	res, err := h.DebtClient.GetTotalDebtSum(c, &req)
	if err != nil {
		h.log.Error("Error fetching total debt sum", "company_id", companyID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserTotalDebt godoc
// @Summary Get user's total debtor sum
// @Description Retrieve the total debtor amount for a specific user.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} debts.SumMoney "User's total debtor sum"
// @Failure 400 {object} products.Error "Invalid user ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /debts/total-sum/{user_id} [get]
func (h *Handler) GetUserTotalDebt(c *gin.Context) {
	userID := c.Param("user_id")
	companyID := c.MustGet("company_id").(string)

	req := debts.ClientID{
		Id:        userID,
		CompanyId: companyID,
		DebtType:  "debtor",
	}

	res, err := h.DebtClient.GetUserTotalDebtSum(c, &req)
	if err != nil {
		h.log.Error("Error fetching user total debt sum", "user_id", userID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserPayments godoc
// @Summary List user payments
// @Description Retrieve all payments made by a specific user.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} debts.UserPaymentsRes "User payment records"
// @Failure 400 {object} products.Error "Invalid user ID"
// @Failure 500 {object} products.Error "Server error"
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
		DebtType:  "debtor",
	}

	res, err := h.DebtClient.GetUserPayments(c, req)
	if err != nil {
		h.log.Error("Error fetching user payments", "user_id", userID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetDebtsInExcel godoc
// @Summary Export debts to Excel
// @Description Export debtor records as an Excel file filtered by currency.
// @Tags Debts
// @Security ApiKeyAuth
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param currency path string true "Currency code"
// @Success 200 {file} file "Excel file with debtor records"
// @Failure 500 {object} entity.Error "Server error"
// @Router /debts/excel/{currency} [get]
func (h *Handler) GetDebtsInExcel(c *gin.Context) {
	var req debts.FilterExelDebt

	req.Currency = c.Param("currency")
	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.DebtClient.GetDebtsForExel(c, &req)
	if err != nil {
		h.log.Error("Error fetching debts for excel", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Client Name")
	f.SetCellValue("Sheet1", "B1", "Client Phone")
	f.SetCellValue("Sheet1", "C1", "Total Debt")
	f.SetCellValue("Sheet1", "D1", "Amount Paid")
	f.SetCellValue("Sheet1", "E1", "Remaining Debt")
	f.SetCellValue("Sheet1", "F1", "Currency")
	f.SetCellValue("Sheet1", "G1", "Last Payment Date")

	for i := 0; i < len(res.Debts); i++ {
		row := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", "A"+row, res.Debts[i].ClientFullName)
		f.SetCellValue("Sheet1", "B"+row, res.Debts[i].ClientPhone)
		f.SetCellValue("Sheet1", "C"+row, res.Debts[i].TotalAmount)
		f.SetCellValue("Sheet1", "D"+row, res.Debts[i].AmountPaid)
		f.SetCellValue("Sheet1", "E"+row, res.Debts[i].DebtsBalance)
		f.SetCellValue("Sheet1", "F"+row, strings.ToUpper(req.Currency))
		f.SetCellValue("Sheet1", "G"+row, res.Debts[i].LastPaymentDate)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		h.log.Error("Error writing Excel file", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=report.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}

// CreateCreditor godoc
// @Summary Create creditor
// @Description Create a new creditor record.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body entity.DebtsRequest true "Creditor details"
// @Success 201 {object} debts.Debts "Created creditor record"
// @Failure 400 {object} products.Error "Invalid input"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor [post]
func (h *Handler) CreateCreditor(c *gin.Context) {
	var req debts.DebtsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.DebtType = "creditor"

	res, err := h.DebtClient.CreateDebts(c, &req)
	if err != nil {
		h.log.Error("Error creating creditor", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetCreditors godoc
// @Summary Get creditor by ID
// @Description Retrieve creditor details using its ID.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Creditor ID"
// @Success 200 {object} debts.Debts "Creditor details"
// @Failure 400 {object} products.Error "Invalid ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/{id} [get]
func (h *Handler) GetCreditors(c *gin.Context) {
	id := c.Param("id")

	req := &debts.DebtsID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
	}

	res, err := h.DebtClient.GetDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching creditor", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListCreditors godoc
// @Summary List creditor records
// @Description Retrieve a list of creditor records with optional filters.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param is_fully_pay query bool false "Filter by fully paid status"
// @Param currency_code query string false "Filter by currency code"
// @Param description query string false "Filter by description"
// @Param no_paid_credits query bool false "Filter by unpaid credits"
// @Param limit query int false "Maximum results"
// @Param page query int false "Page number for pagination"
// @Success 200 {object} debts.DebtsList "List of creditor records"
// @Failure 400 {object} products.Error "Invalid filter value"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor [get]
func (h *Handler) GetListCreditors(c *gin.Context) {
	var filter debts.FilterDebts

	filter.IsFullyPay = c.Query("is_fully_pay")
	filter.CurrencyCode = c.Query("currency_code")
	filter.Description = c.Query("description")

	noPaidCredits := c.Query("no_paid_credits")
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	if filter.IsFullyPay != "true" && filter.IsFullyPay != "false" && filter.IsFullyPay != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for is_fully_pay"})
		return
	}

	if noPaidCredits == "true" {
		filter.NoPaidDebt = true
	}

	if limitStr != "" {
		if limit, err := strconv.ParseInt(limitStr, 10, 64); err == nil {
			filter.Limit = int32(limit)
		} else {
			h.log.Error("Invalid limit parameter", "value", limitStr, "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if pageStr != "" {
		if page, err := strconv.ParseInt(pageStr, 10, 64); err == nil {
			filter.Page = int32(page)
		} else {
			h.log.Error("Invalid page parameter", "value", pageStr, "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	filter.CompanyId = c.MustGet("company_id").(string)
	filter.DebtType = "creditor"

	res, err := h.DebtClient.GetListDebts(c, &filter)
	if err != nil {
		h.log.Error("Error fetching creditor list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enrich each record with supplier details
	for i, debt := range res.Installments {
		client, err := h.UserClient.GetClient(context.Background(), &pbu.UserIDRequest{
			Id:        debt.ClientId,
			CompanyId: debt.CompanyId,
		})
		if err == nil {
			res.Installments[i].ClientName = client.FullName
			res.Installments[i].ClientPhone = client.Phone
		} else {
			h.log.Error("Error fetching supplier info", "client_id", debt.ClientId, "error", err.Error())
			res.Installments[i].ClientName = "Unknown"
			res.Installments[i].ClientPhone = "Unknown"
		}
	}

	c.JSON(http.StatusOK, res)
}

// GetCreditsFromSupplier godoc
// @Summary Get credits records for supplier
// @Description Retrieve creditor records for a specific supplier.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param supplier_id path string true "Supplier ID"
// @Success 200 {object} debts.DebtsList "Creditor records for the supplier"
// @Failure 400 {object} products.Error "Invalid supplier ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/supplier/{supplier_id} [get]
func (h *Handler) GetCreditsFromSupplier(c *gin.Context) {
	supplierID := c.Param("supplier_id")
	req := &debts.ClientID{
		Id:       supplierID,
		DebtType: "creditor", // исправлено: кредитор, а не должник
	}

	res, err := h.DebtClient.GetClientDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching supplier creditor records", "supplier_id", supplierID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// PayCredit godoc
// @Summary Process creditor payment
// @Description Make a payment toward a creditor record.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body debts.PayDebtsReq true "Payment details"
// @Success 200 {object} debts.Debts "Updated creditor record"
// @Failure 400 {object} products.Error "Invalid input"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/pay [post]
func (h *Handler) PayCredit(c *gin.Context) {
	var req debts.PayDebtsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Invalid request data", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.PayType = "out" // для кредитора платежи исходящие

	res, err := h.DebtClient.PayDebts(c, &req)
	if err != nil {
		h.log.Error("Error processing creditor payment", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPaymentsByCreditId godoc
// @Summary List payments by creditor ID
// @Description Retrieve all payments for a specific creditor record.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param credit_id path string true "Creditor ID"
// @Success 200 {object} debts.PaymentList "List of payments"
// @Failure 400 {object} products.Error "Invalid creditor ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/payments/{credit_id} [get]
func (h *Handler) GetPaymentsByCreditId(c *gin.Context) {
	creditId := c.Param("credit_id")
	req := &debts.PayDebtsID{
		Id:      creditId,
		PayType: "in",
	}

	res, err := h.DebtClient.GetPaymentsByDebtsId(c, req)
	if err != nil {
		h.log.Error("Error fetching payments for creditor", "credit_id", creditId, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetCreditPayment godoc
// @Summary Get creditor payment details
// @Description Retrieve details for a specific creditor payment.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} debts.Payment "Payment details"
// @Failure 400 {object} products.Error "Invalid payment ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/payment/{id} [get]
func (h *Handler) GetCreditPayment(c *gin.Context) {
	id := c.Param("id")

	req := &debts.PaymentID{
		Id:        id,
		CompanyId: c.MustGet("company_id").(string),
	}

	res, err := h.DebtClient.GetPayment(c, req)
	if err != nil {
		h.log.Error("Error fetching creditor payment", "payment_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTotalCreditSum godoc
// @Summary Get total creditor sum
// @Description Retrieve the total amount of creditor records for the company.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} debts.SumMoney "Total creditor sum"
// @Failure 400 {object} products.Error "Bad request"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/total-sum [get]
func (h *Handler) GetTotalCreditSum(c *gin.Context) {
	companyID := c.MustGet("company_id").(string)

	req := debts.CompanyID{
		Id:       companyID,
		DebtType: "creditor",
	}

	res, err := h.DebtClient.GetTotalDebtSum(c, &req)
	if err != nil {
		h.log.Error("Error fetching total creditor sum", "company_id", companyID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTotalCreditFromSupplier godoc
// @Summary Get supplier's total creditor sum
// @Description Retrieve the total creditor amount for a specific supplier.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param supplier_id path string true "Supplier ID"
// @Success 200 {object} debts.SumMoney "Supplier's total creditor sum"
// @Failure 400 {object} products.Error "Invalid supplier ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/total-sum/{supplier_id} [get]
func (h *Handler) GetTotalCreditFromSupplier(c *gin.Context) {
	supplierID := c.Param("supplier_id")
	companyID := c.MustGet("company_id").(string)

	req := debts.ClientID{
		Id:        supplierID,
		CompanyId: companyID,
		DebtType:  "creditor", // исправлено на creditor
	}

	res, err := h.DebtClient.GetUserTotalDebtSum(c, &req)
	if err != nil {
		h.log.Error("Error fetching supplier total creditor sum", "supplier_id", supplierID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPaymentsToSupplier godoc
// @Summary List payments to supplier
// @Description Retrieve all payments made to a specific supplier for creditor records.
// @Tags Creditor
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param supplier_id path string true "Supplier ID"
// @Success 200 {object} debts.UserPaymentsRes "Supplier payment records"
// @Failure 400 {object} products.Error "Invalid supplier ID"
// @Failure 500 {object} products.Error "Server error"
// @Router /creditor/payments/{supplier_id} [get]
func (h *Handler) GetPaymentsToSupplier(c *gin.Context) {
	supplierID := c.Param("supplier_id")
	if supplierID == "" {
		h.log.Error("supplier_id not provided in URL")
		c.JSON(http.StatusBadRequest, gin.H{"error": "supplier_id is required"})
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
		Id:        supplierID,
		CompanyId: companyID,
		DebtType:  "creditor", // исправлено на creditor
	}

	res, err := h.DebtClient.GetUserPayments(c, req)
	if err != nil {
		h.log.Error("Error fetching supplier payments", "supplier_id", supplierID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, res)
}
