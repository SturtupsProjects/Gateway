package handler

import (
	"fmt"
	"net/http"

	"gateway/internal/generated/company"
	"github.com/gin-gonic/gin"
)

// @Summary Create Company Balance
// @Description Create a new company balance
// @Tags Company Balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body company.CompanyBalanceRequest true "Company Balance details"
// @Success 200 {object} company.CompanyBalanceResponse
// @Failure 400 {object} string
// @Router /company-balance [post]
func (h *Handler) CreateCompanyBalance(c *gin.Context) {
	var req company.CompanyBalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.CompanyClient.CreateCompanyBalance(c, &company.CompanyBalanceRequest{
		CompanyId: req.CompanyId,
		Balance:   req.Balance,
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
// @Param input body company.CompanyBalanceRequest true "Updated balance details"
// @Success 200 {object} company.CompanyBalanceResponse
// @Failure 400 {object} string
// @Router /company-balance [put]
func (h *Handler) UpdateCompanyBalance(c *gin.Context) {
	var req company.CompanyBalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.CompanyClient.UpdateCompanyBalance(c, &company.CompanyBalanceRequest{
		CompanyId: req.CompanyId,
		Balance:   req.Balance,
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
// @Param input body company.FilterCompanyBalanceRequest true "Filter details"
// @Success 200 {object} company.CompanyBalanceListResponse
// @Failure 400 {object} string
// @Router /company-balance/list [post]
func (h *Handler) GetUsersBalanceList(c *gin.Context) {
	var req company.FilterCompanyBalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.CompanyClient.GetUsersBalanceList(c, &company.FilterCompanyBalanceRequest{
		Limit: req.Limit,
		Page:  req.Page,
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
	res, err := h.CompanyClient.DeleteCompanyBalance(c, &company.Id{Id: companyID})
	if err != nil {
		h.log.Error(fmt.Sprintf("DeleteCompanyBalance request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
