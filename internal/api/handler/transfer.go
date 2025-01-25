package handler

import (
	"gateway/internal/generated/products"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CreateTransfers godoc
// @Summary Create a new transfer
// @Description Create a new transfer between branches
// @Tags Transfers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id header string true "Branch ID"
// @Param Transfer body entity.TransferReq true "Transfer data"
// @Success 201 {object} products.Transfer
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /transfers [post]
func (h *Handler) CreateTransfers(c *gin.Context) {
	var req products.TransferReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateTransfers request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.TransferredBy = c.MustGet("id").(string)
	req.FromBranchId = branchId

	res, err := h.ProductClient.CreateTransfers(c, &req)
	if err != nil {
		h.log.Error("Error creating transfer", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetTransfers godoc
// @Summary Get a transfer by ID
// @Description Retrieve a transfer by its ID
// @Tags Transfers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Transfer ID"
// @Success 200 {object} products.Transfer
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /transfers/{id} [get]
func (h *Handler) GetTransfers(c *gin.Context) {

	id := c.Param("id")
	companyId := c.MustGet("company_id").(string)

	req := &products.TransferID{Id: id, CompanyId: companyId}

	res, err := h.ProductClient.GetTransfers(c, req)
	if err != nil {
		h.log.Error("Error fetching transfer", "transfer_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetTransferList godoc
// @Summary Get a list of transfers
// @Description Retrieve a paginated list of transfers with optional filters
// @Tags Transfers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param branch_id header string true "Branch ID"
// @Param limit query int false "Number of products "
// @Param page query int false "Offset for pagination "
// @Param product_name query string false "filter by product_name"
// @Success 200 {object} products.TransferList
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /transfers [get]
func (h *Handler) GetTransferList(c *gin.Context) {
	var filter products.TransferFilter

	limitSt := c.Query("limit")
	pageSt := c.Query("page")
	filter.ProductName = c.Query("product_name")

	limit, err := strconv.Atoi(limitSt)
	if err != nil {
		h.log.Error("Error parsing limit", "error", err.Error())
		limit = 0
	}
	page, err := strconv.Atoi(pageSt)
	if err != nil {
		h.log.Error("Error parsing page", "error", err.Error())
		page = 0
	}
	branchId := c.GetHeader("branch_id")
	if branchId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required in the header"})
		return
	}

	log.Println(limit, page)

	filter.Limit = int64(limit)
	filter.Page = int64(page)
	filter.BranchId = branchId
	filter.CompanyId = c.MustGet("company_id").(string)

	res, err := h.ProductClient.GetTransferList(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving transfer list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
