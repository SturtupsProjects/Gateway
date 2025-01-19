package handler

import (
	"fmt"
	"gateway/internal/entity"
	"gateway/internal/generated/company"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Company
// @Description Create a new company
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body entity.CreateCompanyRequest true "Company details"
// @Success 200 {object} company.CompanyResponse
// @Failure 400 {object} string
// @Router /companies/admin [post]
func (h *Handler) CreateCompanyA(c *gin.Context) {
	var req entity.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.CompanyClient.CreateCompany(c, &company.CreateCompanyRequest{
		Name:    req.Name,
		Website: req.Website,
		UserId:  c.MustGet("id").(string),
	})
	if err != nil {
		h.log.Error(fmt.Sprintf("CreateCompany request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Company
// @Description Get company details by ID
// @Tags companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} company.CompanyResponse
// @Failure 400 {object} string
// @Router /companies [get]
func (h *Handler) GetCompany(c *gin.Context) {
	req := &company.GetCompanyRequest{CompanyId: c.MustGet("company_id").(string)}
	res, err := h.CompanyClient.GetCompany(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("GetCompany request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Company
// @Description Get company details by ID
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Success 200 {object} company.CompanyResponse
// @Failure 400 {object} string
// @Router /companies/admin/{company_id} [get]
func (h *Handler) GetCompanyA(c *gin.Context) {
	req := &company.GetCompanyRequest{CompanyId: c.Param("company_id")}
	res, err := h.CompanyClient.GetCompany(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("GetCompany request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Company
// @Description Update company details
// @Tags companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body entity.CreateCompanyRequest true "Updated company details"
// @Success 200 {object} company.CompanyResponse
// @Failure 400 {object} string
// @Router /companies [put]
func (h *Handler) UpdateCompany(c *gin.Context) {
	var req entity.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.CompanyClient.UpdateCompany(c, &company.UpdateCompanyRequest{
		CompanyId: c.MustGet("company_id").(string),
		Name:      req.Name,
		Website:   req.Website,
		Logo:      req.Logo,
	})
	if err != nil {
		h.log.Error(fmt.Sprintf("UpdateCompany request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Company
// @Description Update company details admin
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Param input body entity.UpdateCompanyRequest true "Updated company details"
// @Success 200 {object} company.CompanyResponse
// @Failure 400 {object} string
// @Router /companies/admin/{company_id} [put]
func (h *Handler) UpdateCompanyA(c *gin.Context) {
	var req company.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CompanyId = c.Param("company_id")

	res, err := h.CompanyClient.UpdateCompany(c, &req)
	if err != nil {
		h.log.Error(fmt.Sprintf("UpdateCompany request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

//// @Summary Delete Company
//// @Description Delete a company by ID
//// @Tags companies
//// @Accept json
//// @Produce json
//// @Security ApiKeyAuth
//// @Success 200 {object} company.Message
//// @Failure 400 {object} entity.Error
//// @Router /companies [delete]
//func (h *Handler) DeleteCompany(c *gin.Context) {
//	companyID := c.MustGet("company_id").(string)
//	if companyID == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})
//		return
//	}
//	req := &company.DeleteCompanyRequest{CompanyId: companyID}
//	res, err := h.CompanyClient.DeleteCompany(c, req)
//	if err != nil {
//		h.log.Error(fmt.Sprintf("DeleteCompany request error: %v", err))
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, res)
//}

// @Summary Delete Company
// @Description Delete a company by ID
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Success 200 {object} company.Message
// @Failure 400 {object} entity.Error
// @Router /companies/admin/{company_id} [delete]
func (h *Handler) DeleteCompanyA(c *gin.Context) {
	companyID := c.Param("company_id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company_id is required"})
		return
	}
	req := &company.DeleteCompanyRequest{CompanyId: companyID}
	res, err := h.CompanyClient.DeleteCompany(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("DeleteCompany request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get All Companies
// @Description Get all companies
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query string false "Limit"
// @Param page query string false "Page"
// @Success 200 {object} []company.CompanyResponse
// @Failure 400 {object} string
// @Router /companies/admin/all [get]
func (h *Handler) GetAllCompaniesA(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	// Convert limit and page to integers and calculate offset
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}
	offset := (pageInt - 1) * limitInt

	req := &company.ListCompaniesRequest{
		Limit: int32(limitInt),
		Page:  int32(offset),
	}
	res, err := h.CompanyClient.ListCompanies(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("GetAllCompanies request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary ListCompanyUsers
// @Description Get all users for a company
// @Tags companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query string false "Limit"
// @Param page query string false "Page"
// @Success 200 {array} company.ListCompanyUsersResponse
// @Failure 400 {object} string
// @Router /companies/users [get]
func (h *Handler) ListCompanyUsers(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	name := c.Query("name")
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}
	req := &company.ListCompanyUsersRequest{CompanyId: c.MustGet("company_id").(string), Limit: int32(limitInt), Page: int32(pageInt), Name: name}
	res, err := h.CompanyClient.ListCompanyUsers(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("ListCompanyUsers request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary ListCompanyUsersA
// @Description Get all users for a company
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Param limit query string false "Limit"
// @Param page query string false "Page"
// @Success 200 {array} company.ListCompanyUsersResponse
// @Failure 400 {object} string
// @Router /companies/admin/{company_id}/users [get]
func (h *Handler) ListCompanyUsersA(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}
	req := &company.ListCompanyUsersRequest{CompanyId: c.Param("company_id"), Limit: int32(limitInt), Page: int32(pageInt)}
	res, err := h.CompanyClient.ListCompanyUsers(c, req)
	if err != nil {
		h.log.Error(fmt.Sprintf("ListCompanyUsers request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Create Company User
// @Description Create a new user for a company
// @Tags companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body entity.CreateUserToCompanyRequest true "Company user details"
// @Success 200 {object} company.Message
// @Failure 400 {object} string
// @Router /companies/users [post]
func (h *Handler) CreateCompanyUser(c *gin.Context) {
	var req company.CreateUserToCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.CompanyClient.CreateUserToCompany(c, &company.CreateUserToCompanyRequest{CompanyId: c.MustGet("company_id").(string), FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Role: req.Role, Username: req.Username, Password: req.Password, PhoneNumber: req.PhoneNumber})
	if err != nil {
		h.log.Error(fmt.Sprintf("CreateCompanyUser request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Create Company User
// @Description Create a new user for a company
// @Tags Admin Companies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company_id path string true "Company ID"
// @Param input body entity.CreateUserToCompanyRequest true "Company user details"
// @Success 200 {object} company.Message
// @Failure 400 {object} string
// @Router /companies/admin/{company_id}/users [post]
func (h *Handler) CreateCompanyUserA(c *gin.Context) {
	var req entity.CreateUserToCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.CompanyClient.CreateUserToCompany(c, &company.CreateUserToCompanyRequest{CompanyId: c.Param("company_id"), FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Role: req.Role, Username: req.Username, Password: req.Password, PhoneNumber: req.PhoneNumber})
	if err != nil {
		h.log.Error(fmt.Sprintf("CreateCompanyUser request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
