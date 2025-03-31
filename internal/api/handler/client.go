package handler

import (
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client with the provided details
// @Tags Clients
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Client body entity.ClientRequest true "Client data"
// @Success 201 {object} user.ClientResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /clients [post]
func (h *Handler) CreateClient(c *gin.Context) {
	var req user.ClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateClient request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.ClientType = "client" // Если для клиента всегда "client", оставляем так.
	req.Type = "client"

	res, err := h.UserClient.CreateClient(c, &req)
	if err != nil {
		h.log.Error("Error creating client", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetClient godoc
// @Summary Get a client
// @Description Retrieve a client by ID
// @Tags Clients
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Client ID"
// @Success 200 {object} user.ClientResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /clients/{id} [get]
func (h *Handler) GetClient(c *gin.Context) {
	id := c.Param("id")
	req := &user.UserIDRequest{Id: id}
	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.UserClient.GetClient(c, req)
	if err != nil {
		h.log.Error("Error fetching client", "client_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetClientList godoc
// @Summary List all clients
// @Description Retrieve a list of clients with optional filters
// @Tags Clients
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query entity.ClientFilter false "Filter parameters"
// @Success 200 {object} user.ClientListResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /clients [get]
func (h *Handler) GetClientList(c *gin.Context) {
	var filter user.FilterClientRequest

	filter.Phone = c.Query("phone")
	filter.Address = c.Query("address")
	filter.FullName = c.Query("full_name")
	filter.ClientType = c.Query("client_type")
	if filter.ClientType == "" {
		filter.ClientType = "client"
	}

	limitSt := c.Query("limit")
	pageSt := c.Query("page")
	limit, err := strconv.Atoi(limitSt)
	if err != nil {
		limit = 0
	}
	page, err := strconv.Atoi(pageSt)
	if err != nil {
		page = 0
	}
	filter.Limit = int32(limit)
	filter.Page = int32(page)
	filter.CompanyId = c.MustGet("company_id").(string)
	filter.Type = "client"

	res, err := h.UserClient.GetListClient(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving client list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateClient godoc
// @Summary Update an existing client
// @Description Update the details of an existing client by ID
// @Tags Clients
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Client ID"
// @Param Client body entity.Client true "Updated client data"
// @Success 200 {object} user.ClientResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /clients/{id} [put]
func (h *Handler) UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var req user.ClientUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing UpdateClient request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID := c.MustGet("company_id").(string)
	res, err := h.UserClient.UpdateClient(c, &user.ClientUpdateRequest{
		Id:        id,
		FullName:  req.FullName,
		Address:   req.Address,
		Phone:     req.Phone,
		CompanyId: companyID,
	})
	if err != nil {
		h.log.Error("Error updating client", "client_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteClient godoc
// @Summary Delete a client
// @Description Delete a client by ID
// @Tags Clients
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Client ID"
// @Success 200 {object} user.MessageResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /clients/{id} [delete]
func (h *Handler) DeleteClient(c *gin.Context) {
	id := c.Param("id")
	req := &user.UserIDRequest{Id: id}
	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.UserClient.DeleteClient(c, req)
	if err != nil {
		h.log.Error("Error deleting client", "client_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ----------------------------- Supplier Endpoints -----------------------------------------------------------------------------

// CreateSupplier godoc
// @Summary Create a new supplier
// @Description Create a new supplier with the provided details
// @Tags Supplier
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Supplier body entity.ClientRequest true "Supplier data"
// @Success 201 {object} user.ClientResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /supplier [post]
func (h *Handler) CreateSupplier(c *gin.Context) {
	var req user.ClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing CreateSupplier request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.CompanyId = c.MustGet("company_id").(string)
	req.Type = "supplier"

	res, err := h.UserClient.CreateClient(c, &req)
	if err != nil {
		h.log.Error("Error creating supplier", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetSupplier godoc
// @Summary Get a supplier
// @Description Retrieve a supplier by ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Supplier ID"
// @Success 200 {object} user.ClientResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /supplier/{id} [get]
func (h *Handler) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	req := &user.UserIDRequest{Id: id}
	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.UserClient.GetClient(c, req)
	if err != nil {
		h.log.Error("Error fetching supplier", "supplier_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetSupplierList godoc
// @Summary List all suppliers
// @Description Retrieve a list of suppliers with optional filters
// @Tags Supplier
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query entity.SupplierFilter false "Filter parameters"
// @Success 200 {object} user.ClientListResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /supplier [get]
func (h *Handler) GetSupplierList(c *gin.Context) {
	var filter user.FilterClientRequest

	filter.Phone = c.Query("phone")
	filter.Address = c.Query("address")
	filter.FullName = c.Query("full_name")

	limitSt := c.Query("limit")
	pageSt := c.Query("page")
	limit, err := strconv.Atoi(limitSt)
	if err != nil {
		limit = 0
	}
	page, err := strconv.Atoi(pageSt)
	if err != nil {
		page = 0
	}
	filter.Limit = int32(limit)
	filter.Page = int32(page)
	filter.CompanyId = c.MustGet("company_id").(string)

	filter.ClientType = "client"
	filter.Type = "supplier"

	res, err := h.UserClient.GetListClient(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving supplier list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateSupplier godoc
// @Summary Update an existing supplier
// @Description Update the details of an existing supplier by ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Supplier ID"
// @Param Supplier body entity.Client true "Updated supplier data"
// @Success 200 {object} user.ClientResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /supplier/{id} [put]
func (h *Handler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var req user.ClientUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error parsing UpdateSupplier request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID := c.MustGet("company_id").(string)
	res, err := h.UserClient.UpdateClient(c, &user.ClientUpdateRequest{
		Id:        id,
		FullName:  req.FullName,
		Address:   req.Address,
		Phone:     req.Phone,
		CompanyId: companyID,
	})
	if err != nil {
		h.log.Error("Error updating supplier", "supplier_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteSupplier godoc
// @Summary Delete a supplier
// @Description Delete a supplier by ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Supplier ID"
// @Success 200 {object} user.MessageResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /supplier/{id} [delete]
func (h *Handler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	req := &user.UserIDRequest{Id: id}
	req.CompanyId = c.MustGet("company_id").(string)

	res, err := h.UserClient.DeleteClient(c, req)
	if err != nil {
		h.log.Error("Error deleting supplier", "supplier_id", id, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
