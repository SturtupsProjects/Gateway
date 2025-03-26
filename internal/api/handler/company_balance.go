package handler

//import (
//	"net/http"
//	"strconv"
//
//	"gateway/internal/generated/company"
//	"github.com/gin-gonic/gin"
//)
//
//// ==================================================================
//// Create Company Balance
//// ==================================================================
//
//// @Summary Create Company Balance
//// @Description Create a new company balance
//// @Tags Company Balance
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param body body company.CompanyBalanceRequest true "Company balance request"
//// @Success 200 {object} company.CompanyBalanceResponse
//// @Failure 400 {object} products.Error
//// @Router /company-balance [post]
//func (h *Handler) CreateCompanyBalance(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	var req company.CompanyBalanceRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
//		return
//	}
//
//	companyID, exists := c.Get("company_id")
//	if !exists {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is missing in context"})
//		return
//	}
//	idStr, ok := companyID.(string)
//	if !ok {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID has invalid format"})
//		return
//	}
//	req.CompanyId = idStr
//
//	res, err := h.CompanyClient.CreateCompanyBalance(ctx, &req)
//	if err != nil {
//		h.log.Error("CreateCompanyBalance error: ", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}
//
//// ==================================================================
//// Get Company Balance
//// ==================================================================
//
//// @Summary Get Company Balance
//// @Description Get company balance by ID
//// @Tags Company Balance
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param company_id path string true "Company ID"
//// @Success 200 {object} company.CompanyBalanceResponse
//// @Failure 400 {object} products.Error
//// @Router /company-balance/{company_id} [get]
//func (h *Handler) GetCompanyBalance(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	companyID := c.Param("company_id")
//	if companyID == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
//		return
//	}
//
//	res, err := h.CompanyClient.GetCompanyBalance(ctx, &company.Id{Id: companyID})
//	if err != nil {
//		h.log.Error("GetCompanyBalance error: ", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}
//
//// ==================================================================
//// Update Company Balance
//// ==================================================================
//
//// @Summary Update Company Balance
//// @Description Update an existing company balance
//// @Tags Company Balance
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param body body company.CompanyBalanceRequest true "Company balance update request"
//// @Success 200 {object} company.CompanyBalanceResponse
//// @Failure 400 {object} products.Error
//// @Router /company-balance [put]
//func (h *Handler) UpdateCompanyBalance(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	var req company.CompanyBalanceRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
//		return
//	}
//
//	companyID, exists := c.Get("company_id")
//	if !exists {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is missing in context"})
//		return
//	}
//	idStr, ok := companyID.(string)
//	if !ok {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID has invalid format"})
//		return
//	}
//	req.CompanyId = idStr
//
//	res, err := h.CompanyClient.UpdateCompanyBalance(ctx, &req)
//	if err != nil {
//		h.log.Error("UpdateCompanyBalance error: ", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}
//
//// ==================================================================
//// Get Company Balance List
//// ==================================================================
//
//// @Summary List Company Balances
//// @Description Get list of company balances with pagination
//// @Tags Company Balance
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param page query int true "Page number"
//// @Param limit query int true "Limit per page"
//// @Success 200 {object} company.CompanyBalanceListResponse
//// @Failure 400 {object} products.Error
//// @Router /company-balance/list [get]
//func (h *Handler) GetUsersBalanceList(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	pageStr := c.Query("page")
//	limitStr := c.Query("limit")
//
//	pageInt, err := strconv.Atoi(pageStr)
//	if err != nil || pageInt <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
//		return
//	}
//	limitInt, err := strconv.Atoi(limitStr)
//	if err != nil || limitInt <= 0 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
//		return
//	}
//
//	req := &company.FilterCompanyBalanceRequest{
//		Page:  int32(pageInt),
//		Limit: int32(limitInt),
//	}
//
//	res, err := h.CompanyClient.GetUsersBalanceList(ctx, req)
//	if err != nil {
//		h.log.Error("GetUsersBalanceList error: ", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}
//
//// ==================================================================
//// Delete Company Balance
//// ==================================================================
//
//// @Summary Delete Company Balance
//// @Description Soft delete a company balance by ID
//// @Tags Company Balance
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param company_id path string true "Company ID"
//// @Success 200 {object} company.Message
//// @Failure 400 {object} products.Error
//// @Router /company-balance/{company_id} [delete]
//func (h *Handler) DeleteCompanyBalance(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	companyID := c.Param("company_id")
//	if companyID == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
//		return
//	}
//
//	res, err := h.CompanyClient.DeleteCompanyBalance(ctx, &company.Id{Id: companyID})
//	if err != nil {
//		h.log.Error("DeleteCompanyBalance error: ", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}
//
//// ==================================================================
//// Send SMS
//// ==================================================================
//
//// @Summary Send SMS
//// @Description Send an SMS notification
//// @Tags Company Balance
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Param body body company.SmsRequest true "SMS request"
//// @Success 200 {object} company.Message
//// @Failure 400 {object} products.Error
//// @Router /company-balance/sms [post]
//func (h *Handler) SendSMS(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	var smsReq company.SmsRequest
//	if err := c.ShouldBindJSON(&smsReq); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
//		return
//	}
//
//	companyID, exists := c.Get("company_id")
//	if !exists {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is missing in context"})
//		return
//	}
//	idStr, ok := companyID.(string)
//	if !ok {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID has invalid format"})
//		return
//	}
//	smsReq.CompanyId = idStr
//
//	if smsReq.Phone == "" || smsReq.Message == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number and message are required"})
//		return
//	}
//
//	res, err := h.CompanyClient.SendSMS(ctx, &smsReq)
//	if err != nil {
//		h.log.Error("SendSMS error: ", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, res)
//}
