package handler

import (
	"gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"log"
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
// @Param Client body user.ClientRequest true "Client data"
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

	req.ClientType = "client"

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

	log.Println(limit, page)

	filter.Limit = int32(limit)
	filter.Page = int32(page)

	filter.CompanyId = c.MustGet("company_id").(string)
	filter.ClientType = "client"

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

	res, err := h.UserClient.UpdateClient(c, &user.ClientUpdateRequest{Id: id, FullName: req.FullName, Address: req.Address, Phone: req.Phone, CompanyId: companyID})
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

// GetStreetClientList godoc
// @Summary List all street clients
// @Description Retrieve a list of street clients with optional filters
// @Tags Clients
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query entity.StreetClientFilter false "Filter parameters"
// @Success 200 {object} user.ClientListResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /clients/street [get]
func (h *Handler) GetStreetClientList(c *gin.Context) {
	var filter user.FilterClientRequest

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing FilterClientRequest", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter.CompanyId = c.MustGet("company_id").(string)

	filter.ClientType = "street"

	res, err := h.UserClient.GetListClient(c, &filter)
	if err != nil {
		h.log.Error("Error retrieving client list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
