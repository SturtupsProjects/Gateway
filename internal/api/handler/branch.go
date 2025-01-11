package handler

import (
	"fmt"
	"gateway/internal/generated/company"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Branch
// @Description Create a new branch
// @Tags Branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body company.CreateBranchRequest true "Branch details"
// @Success 200 {object} company.BranchResponse
// @Failure 400 {object} map[string]string
// @Router /branches/create [post]
func (h *Handler) CreateBranch(c *gin.Context) {
	var req company.CreateBranchRequest

	// Получение company_id из контекста
	companyID, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "company_id not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.CompanyId = companyID.(string)
	res, err := h.CompanyClient.CreateBranch(c, &req)
	if err != nil {
		h.log.Error(fmt.Sprintf("CreateBranch request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Get Branch
// @Description Get branch details by ID
// @Tags Branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id path string true "Branch ID"
// @Success 200 {object} company.BranchResponse
// @Failure 400 {object} map[string]string
// @Router /branches/{branch_id} [get]
func (h *Handler) GetBranch(c *gin.Context) {
	req := &company.GetBranchRequest{
		BranchId: c.Param("branch_id"),
	}

	companyID, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "company_id not found"})
		return
	}

	req.CompanyId = companyID.(string)
	res, err := h.CompanyClient.GetBranch(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("GetBranch request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update Branch
// @Description Update branch details
// @Tags Branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id path string true "Branch ID"
// @Param input body company.UpdateBranchRequest true "Updated branch details"
// @Success 200 {object} company.BranchResponse
// @Failure 400 {object} map[string]string
// @Router /branches/{branch_id} [put]
func (h *Handler) UpdateBranch(c *gin.Context) {
	var req company.UpdateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "company_id not found"})
		return
	}

	req.BranchId = c.Param("branch_id")
	req.CompanyId = companyID.(string)
	res, err := h.CompanyClient.UpdateBranch(c, &req)
	if err != nil {
		h.log.Error(fmt.Sprintf("UpdateBranch request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Delete Branch
// @Description Delete a branch by ID
// @Tags Branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id path string true "Branch ID"
// @Success 200 {object} company.Message
// @Failure 400 {object} map[string]string
// @Router /branches/{branch_id} [delete]
func (h *Handler) DeleteBranch(c *gin.Context) {
	branchID := c.Param("branch_id")
	if branchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_id is required"})
		return
	}

	companyID, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "company_id not found"})
		return
	}

	req := &company.DeleteBranchRequest{
		BranchId:  branchID,
		CompanyId: companyID.(string),
	}

	res, err := h.CompanyClient.DeleteBranch(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("DeleteBranch request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary List Branches
// @Description List all branches for a company
// @Tags Branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} company.ListBranchesResponse
// @Failure 400 {object} map[string]string
// @Router /branches/list [get]
func (h *Handler) ListBranches(c *gin.Context) {
	companyID, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "company_id not found"})
		return
	}

	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

	req := &company.ListBranchesRequest{
		CompanyId: companyID.(string),
		Limit:     stringToInt(limit),
		Page:      stringToInt(page),
	}

	res, err := h.CompanyClient.ListBranches(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("ListBranches request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// stringToInt is a helper function to convert strings to integers
func stringToInt(s string) int32 {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int32(value)
}
