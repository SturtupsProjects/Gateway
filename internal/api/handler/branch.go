package handler

import (
	"fmt"
	"gateway/internal/generated/company"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Branch
// @Description Create a new branch
// @Tags Admin Branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body company.CreateBranchRequest true "Branch details"
// @Success 200 {object} company.BranchResponse
// @Failure 400 {object} string
// @Router /branches/admin [post]
func (h *Handler) CreateBranch(c *gin.Context) {
	var req company.CreateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
// @Tags branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} company.BranchResponse
// @Failure 400 {object} string
// @Router /branches/{branch_id} [get]
func (h *Handler) GetBranch(c *gin.Context) {
	req := &company.GetBranchRequest{BranchId: c.Param("branch_id")}
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
// @Tags branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body company.UpdateBranchRequest true "Updated branch details"
// @Success 200 {object} company.BranchResponse
// @Failure 400 {object} string
// @Router /branches/{branch_id} [put]
func (h *Handler) UpdateBranch(c *gin.Context) {
	var req company.UpdateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.BranchId = c.Param("branch_id")
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
// @Tags branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id path string true "Branch ID"
// @Success 200 {object} company.Message
// @Failure 400 {object} entity.Error
// @Router /branches/{branch_id} [delete]
func (h *Handler) DeleteBranch(c *gin.Context) {
	branchID := c.Param("branch_id")
	if branchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_id is required"})
		return
	}
	req := &company.DeleteBranchRequest{BranchId: branchID}
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
// @Tags branches
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Success 200 {object} company.ListBranchesResponse
// @Failure 400 {object} string
// @Router /branches/company/{company_id} [get]
func (h *Handler) ListBranches(c *gin.Context) {
	companyID := c.Param("company_id")
	req := &company.ListBranchesRequest{CompanyId: companyID}
	res, err := h.CompanyClient.ListBranches(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("ListBranches request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
