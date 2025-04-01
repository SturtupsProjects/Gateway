package handler

import (
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateSalary godoc
// @Summary Create Salary
// @Description Create a new salary record for a user.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param salary body entity.SalaryRequest true "Salary Request"
// @Success 201 {object} user.SalaryResponse "Salary created successfully"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /salary [post]
func (h *Handler) CreateSalary(c *gin.Context) {
	var req user.SalaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("CreateSalary: error parsing request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("CreateSalary: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.CreateSalary(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("CreateSalary: error creating salary", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// UpdateSalary godoc
// @Summary Update Salary
// @Description Update an existing salary record.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param salary body entity.SalaryUpdate true "Salary Update Request"
// @Param salary_id path string true "Salary ID"
// @Success 200 {object} user.SalaryUpdate "Salary updated successfully"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /salary/{salary_id} [put]
func (h *Handler) UpdateSalary(c *gin.Context) {
	var req user.SalaryUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("UpdateSalary: error parsing request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.SalaryId = c.Param("salary_id")
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("UpdateSalary: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.UpdateSalary(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("UpdateSalary: error updating salary", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetSalaryByID godoc
// @Summary Get Salary by ID
// @Description Retrieve salary record by salary ID.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param salary_id path string true "Salary ID"
// @Success 200 {object} user.SalaryResponse "Salary data returned"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /salary/{salary_id} [get]
func (h *Handler) GetSalaryByID(c *gin.Context) {
	var req user.ID
	req.Id = c.Param("salary_id")
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("GetSalaryByID: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}

	req.CompanyId = companyID.(string)

	res, err := h.UserClient.GetSalaryByID(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("GetSalaryByID: error getting salary", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ListSalaries godoc
// @Summary List Salaries
// @Description List salary records with pagination, sorting and filtering.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit (set 0 to get all records)"
// @Param page query int false "Page (set 0 to get all records)"
// @Param sort_field query string false "Sort field (e.g., salary_date, created_at, salary_amount)"
// @Param sort_order query string false "Sort order (ASC or DESC)"
// @Success 200 {object} user.GetSalaryList "List of salaries"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /salary [get]
func (h *Handler) ListSalaries(c *gin.Context) {
	var req user.GetSalaryRequest
	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	req.SortField = c.Query("sort_field")
	req.Order = c.Query("sort_order")

	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("ListSalaries: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	if limitStr == "" {
		req.Limit = 0
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.log.Error("ListSalaries: error parsing limit", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		req.Limit = int64(limit)
	}

	if pageStr == "" {
		req.Page = 0
	} else {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error("ListSalaries: error parsing page", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
			return
		}
		req.Page = int64(page)
	}

	res, err := h.UserClient.ListSalaries(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("ListSalaries: error getting salaries", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CreateAdjustment godoc
// @Summary Create Adjustment
// @Description Create a new salary adjustment (bonus/penalty) for a user.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param adjustment body entity.AdjustmentRequest true "Adjustment Request"
// @Success 201 {object} user.AdjustmentResponse "Adjustment created successfully"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /adjustment [post]
func (h *Handler) CreateAdjustment(c *gin.Context) {
	var req user.AdjustmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("CreateAdjustment: error parsing request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("CreateAdjustment: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.CreateAdjustment(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("CreateAdjustment: error creating adjustment", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// UpdateAdjustment godoc
// @Summary Update Adjustment
// @Description Update an existing salary adjustment record.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param adjustment body entity.AdjustmentUpdate true "Adjustment Update Request"
// @Param adjustment_id path string true "Adjustment ID"
// @Success 200 {object} user.AdjustmentResponse "Adjustment updated successfully"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /adjustment/{adjustment_id} [put]
func (h *Handler) UpdateAdjustment(c *gin.Context) {
	var req user.AdjustmentUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("UpdateAdjustment: error parsing request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.AdjustmentId = c.Param("adjustment_id")
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("UpdateAdjustment: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.UpdateAdjustment(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("UpdateAdjustment: error updating adjustment", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CloseAdjustment godoc
// @Summary Close Adjustment
// @Description Mark an adjustment as inactive (close it).
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param adjustment_id path string true "Adjustment ID"
// @Success 200 {object} user.AdjustmentResponse "Adjustment closed successfully"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /adjustment/{adjustment_id}/close [put]
func (h *Handler) CloseAdjustment(c *gin.Context) {
	var req user.ID
	req.Id = c.Param("adjustment_id")
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("CloseAdjustment: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.CloseAdjustment(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("CloseAdjustment: error closing adjustment", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAdjustmentByID godoc
// @Summary Get Adjustment by ID
// @Description Retrieve a salary adjustment record by its ID.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param adjustment_id path string true "Adjustment ID"
// @Success 200 {object} user.AdjustmentResponse "Adjustment data returned"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /adjustment/{adjustment_id} [get]
func (h *Handler) GetAdjustmentByID(c *gin.Context) {
	var req user.ID
	req.Id = c.Param("adjustment_id")
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("GetAdjustmentByID: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.GetAdjustmentByID(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("GetAdjustmentByID: error getting adjustment", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ListAdjustments godoc
// @Summary List Adjustments
// @Description List salary adjustments with pagination, filtering and sorting.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit (set 0 for no limit)"
// @Param page query int false "Page (set 0 for no pagination)"
// @Param sort_field query string false "Sort field (e.g., adjustment_date, created_at, amount)"
// @Param sort_order query string false "Sort order (ASC or DESC)"
// @Param adjustment_type query string false "Filter by adjustment type (e.g., bonus, penalty)"
// @Param is_active query string false "Filter by active status (true or false)"
// @Success 200 {object} user.AdjustmentList "List of adjustments"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /adjustment [get]
func (h *Handler) ListAdjustments(c *gin.Context) {
	var req user.GetAdjustmentRequest
	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	req.SortField = c.Query("sort_field")
	req.Order = c.Query("sort_order")
	req.AdjustmentType = c.Query("adjustment_type")
	req.IsActive = c.Query("is_active")

	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("ListAdjustments: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	if limitStr == "" {
		req.Limit = 0
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.log.Error("ListAdjustments: error parsing limit", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		req.Limit = int64(limit)
	}

	if pageStr == "" {
		req.Page = 0
	} else {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error("ListAdjustments: error parsing page", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
			return
		}
		req.Page = int64(page)
	}

	res, err := h.UserClient.ListAdjustments(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("ListAdjustments: error getting adjustments", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetWorkerAllInfo godoc
// @Summary Get Worker All Info
// @Description Retrieve full information about a worker including user data, current salary, and all adjustments.
// @Tags Adjustment-Salary-User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param worker_id path string true "Worker ID"
// @Success 200 {object} user.WorkerAllInfo "Worker full info returned"
// @Failure 400 {object} entity.Error "Bad request"
// @Router /salary/worker/{worker_id} [get]
func (h *Handler) GetWorkerAllInfo(c *gin.Context) {
	var req user.ID
	req.Id = c.Param("worker_id")
	companyID, ok := c.Get("company_id")
	if !ok {
		h.log.Error("GetWorkerAllInfo: company_id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id not provided"})
		return
	}
	req.CompanyId = companyID.(string)

	res, err := h.UserClient.GetWorkerAllInfo(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("GetWorkerAllInfo: error getting worker info", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
