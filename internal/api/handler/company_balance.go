package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"gateway/internal/generated/company"
	"github.com/gin-gonic/gin"
)

// @Summary Create Company Balance
// @Description Create a new company balance
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param balance body int true "Balance"
// @Success 200 {object} company.CompanyBalanceResponse
// @Failure 400 {object} string
// @Router /company-balance [post]
func (h *Handler) CreateCompanyBalance(c *gin.Context) {
	var request struct {
		Balance int64 `json:"balance"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid balance value" + err.Error()})
		return
	}
	res, err := h.CompanyClient.CreateCompanyBalance(c, &company.CompanyBalanceRequest{
		CompanyId: c.MustGet("company_id").(string),
		Balance:   request.Balance,
	})
	if err != nil {
		h.log.Error(fmt.Sprintf("CreateCompanyBalance request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Get Company Balance
// @Description Get company balance by ID
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Success 200 {object} company.CompanyBalanceResponse
// @Failure 400 {object} string
// @Router /company-balance/{company_id} [get]
func (h *Handler) GetCompanyBalance(c *gin.Context) {
	companyID := c.Param("company_id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}
	res, err := h.CompanyClient.GetCompanyBalance(c, &company.Id{Id: companyID})
	if err != nil {
		h.log.Error(fmt.Sprintf("GetCompanyBalance request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update Company Balance
// @Description Update a company balance
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param balance body int true "Balance"
// @Success 200 {object} company.CompanyBalanceResponse
// @Failure 400 {object} string
// @Router /company-balance [put]
func (h *Handler) UpdateCompanyBalance(c *gin.Context) {
	var request struct {
		Balance int64 `json:"balance"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid balance value"})
		return
	}
	res, err := h.CompanyClient.UpdateCompanyBalance(c, &company.CompanyBalanceRequest{
		CompanyId: c.MustGet("company_id").(string),
		Balance:   request.Balance,
	})
	if err != nil {
		h.log.Error(fmt.Sprintf("UpdateCompanyBalance request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Get Users Balance List
// @Description Get list of user balances filtered by company and user
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int true "Page"
// @Param limit query int true "Limit"
// @Success 200 {object} company.CompanyBalanceListResponse
// @Failure 400 {object} string
// @Router /company-balance/list [get]
func (h *Handler) GetUsersBalanceList(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}
	res, err := h.CompanyClient.GetUsersBalanceList(c, &company.FilterCompanyBalanceRequest{
		Limit: int32(limitInt),
		Page:  int32(pageInt),
	})
	if err != nil {
		h.log.Error(fmt.Sprintf("GetUsersBalanceList request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Delete Company Balance
// @Description Delete a company balance by ID
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Success 200 {object} company.Message
// @Failure 400 {object} string
// @Router /company-balance/{company_id} [delete]
func (h *Handler) DeleteCompanyBalance(c *gin.Context) {
	companyID := c.Param("company_id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}
	res, err := h.CompanyClient.DeleteCompanyBalance(c, &company.Id{Id: companyID})
	if err != nil {
		h.log.Error(fmt.Sprintf("DeleteCompanyBalance request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Send SMS
// @Description Send an SMS notification
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param phone_number body string true "Phone Number"
// @Param message body string true "Message content"
// @Success 200 {object} company.Message
// @Failure 400 {object} string
// @Router /company-balance/sms [post]
func (h *Handler) SendSMS(c *gin.Context) {
	smsRequest := company.SmsRequest{}
	// Bind incoming JSON data to the struct
	if err := c.ShouldBindJSON(&smsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	smsRequest.CompanyId = c.MustGet("company_id").(string)
	// Validate phone number and message
	if smsRequest.Phone == "" || smsRequest.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number and message are required"})
		return
	}

	// Call external SMS service or internal method to send the message
	res, err := h.CompanyClient.SendSMS(c, &smsRequest)
	if err != nil {
		h.log.Error(fmt.Sprintf("SendSMS request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, res)
}
