package handler

import (
	"net/http"

	"gateway/internal/generated/debts"
	"github.com/gin-gonic/gin"
)

// GetPayment godoc
// @Summary Get a payment by ID
// @Description Retrieve a payment by its unique ID
// @Tags Payments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} debts.Payment
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /payments/{id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("id")
	req := &debts.PaymentID{Id: id}

	res, err := h.PaymentClient.GetPayment(c, req)
	if err != nil {
		h.log.Error("Error fetching payment", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPaymentsByDebtId godoc
// @Summary Get payments by debt ID
// @Description Retrieve all payments associated with a specific debt
// @Tags Payments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param debt_id path string true "Debt ID"
// @Success 200 {object} debts.PaymentList
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /payments/debt/{debt_id} [get]
func (h *Handler) GetPaymentsByDebtId(c *gin.Context) {
	debtId := c.Param("debt_id")
	req := &debts.DebtID{Id: debtId}

	res, err := h.PaymentClient.GetPaymentsByDebtId(c, req)
	if err != nil {
		h.log.Error("Error fetching payments for debt", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPayments godoc
// @Summary Get a list of payments
// @Description Retrieve a list of payments with optional filters
// @Tags Payments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param filter query debts.FilterPayment false "Filter parameters"
// @Success 200 {object} debts.PaymentList
// @Failure 400 {object} debts.Error
// @Failure 500 {object} debts.Error
// @Router /payments [get]
func (h *Handler) GetPayments(c *gin.Context) {
	var filter debts.FilterPayment

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.log.Error("Error parsing filter", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.PaymentClient.GetPayments(c, &filter)
	if err != nil {
		h.log.Error("Error fetching payments list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
