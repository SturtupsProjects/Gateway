package handler

import (
	"gateway/internal/generated/debts"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDebt godoc
// @Summary Create a new debt
// @Description Add a new debt for a client with the specified details
// @Security ApiKeyAuth
// @Tags Debts
// @Accept json
// @Produce json
// @Param data body debts.DebtRequest true "Debt creation details"
// @Success 201 {object} debts.Debt "Debt successfully created"
// @Failure 400 {object} debts.Error "Invalid input or bad request"
// @Failure 500 {object} debts.Error "Internal server error"
// @Router /debts [post]
func (h *Handler) CreateDebt(c *gin.Context) {
	var req debts.DebtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error binding JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	res, err := h.DebtClient.CreateDebt(c, &req)
	if err != nil {
		h.log.Error("Error creating debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetDebt godoc
// @Summary Get a debt by ID
// @Description Retrieve details of a specific debt by ID
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Debt ID"
// @Success 200 {object} debts.Debt
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /debts/{id} [get]
func (h *Handler) GetDebt(c *gin.Context) {
	id := c.Param("id")
	req := &debts.DebtID{Id: id}

	res, err := h.DebtClient.GetDebt(c, req)
	if err != nil {
		h.log.Error("Error fetching debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// PayDebt godoc
// @Summary Pay a debt
// @Description Make a payment towards a specific debt
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body debts.PayDebtReq true "Payment details"
// @Success 200 {object} debts.Debt
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /debts/pay [post]
func (h *Handler) PayDebt(c *gin.Context) {
	var req debts.PayDebtReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error binding JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	res, err := h.DebtClient.PayDebt(c, &req)
	if err != nil {
		h.log.Error("Error processing payment", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListDebts godoc
// @Summary Get list of debts
// @Description Retrieve a list of debts with optional filters
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query debts.FilterDebt false "Filter parameters"
// @Success 200 {object} debts.DebtsList
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /debts [get]
func (h *Handler) GetListDebts(c *gin.Context) {
	var filter debts.FilterDebt
	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing filter", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.DebtClient.GetListDebts(c, &filter)
	if err != nil {
		h.log.Error("Error fetching debt list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetClientDebts godoc
// @Summary Get debts for a client
// @Description Retrieve a list of debts associated with a specific client
// @Tags Debts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param client_id path string true "Client ID"
// @Success 200 {object} debts.DebtsList
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /debts/client/{client_id} [get]
func (h *Handler) GetClientDebts(c *gin.Context) {
	clientID := c.Param("client_id")
	req := &debts.ClientID{Id: clientID}

	res, err := h.DebtClient.GetClientDebts(c, req)
	if err != nil {
		h.log.Error("Error fetching client debts", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
