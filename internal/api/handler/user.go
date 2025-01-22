package handler

import (
	"gateway/internal/api/token"
	"gateway/internal/entity"
	user "gateway/internal/generated/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAdmin godoc
// @Summary Register an Admin
// @Description Register a new admin account
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} user.MessageResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/admin/register [post]
func (a *Handler) RegisterAdmin(c *gin.Context) {
	res, err := a.UserClient.RegisterAdmin(c.Request.Context(), &user.MessageResponse{Message: "Register Admin!"})

	if err != nil {
		a.log.Error("Error registering admin", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Admin Login
// @Description Admin login
// @Tags User
// @Accept json
// @Produce json
// @Param Login body user.LogInRequest true "Admin login"
// @Success 200 {object} user.TokenResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/login [post]
func (a *Handler) Login(c *gin.Context) {
	var req user.LogInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.log.Error("Error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.UserClient.LogIn(c.Request.Context(), &user.LogInRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})

	if err != nil {
		a.log.Error("Error logging in", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateUser godoc
// @Summary Create User
// @Description Register a new user account
// @Tags User
// @Accept json
// @Produce json
// @Param CreateUser body entity.UserUpdateRequest true "Create user"
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/register [post]
func (a *Handler) CreateUser(c *gin.Context) {
	var req entity.UserUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.log.Error("Error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.UserClient.AddUser(c.Request.Context(), &user.UserRequest{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        req.Role,
		Password:    req.Password,
	})

	if err != nil {
		a.log.Error("Error creating user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUser godoc
// @Summary Update User
// @Description Update user details
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param UpdateUser body entity.UserUpdateRequest true "Update user"
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/update/{id} [put]
func (a *Handler) UpdateUser(c *gin.Context) {
	var req entity.UserUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.log.Error("Error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Using the 'id' from URL param and mapping fields correctly
	res, err := a.UserClient.UpdateUser(c.Request.Context(), &user.UserRequest{
		UserId:      c.Param("id"),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Role:        req.Role,
		CompanyId:   req.CompanyId,
	})

	if err != nil {
		a.log.Error("Error updating user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete a user account
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.MessageResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/delete/{id} [delete]
func (a *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	res, err := a.UserClient.DeleteUser(c.Request.Context(), &user.UserIDRequest{
		Id: id, // Correct usage of 'Id' as per Proto definition
	})

	if err != nil {
		a.log.Error("Error deleting user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUser godoc
// @Summary Get User
// @Description Retrieve user details by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/get/{id} [get]
func (a *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	res, err := a.UserClient.GetUser(c.Request.Context(), &user.UserIDRequest{
		Id: id, // Correct usage of 'Id' as per Proto definition
	})

	if err != nil {
		a.log.Error("Error retrieving user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListUser godoc
// @Summary List Users
// @Description Retrieve a list of users with optional filters
// @Tags User
// @Accept json
// @Produce json
// @Param FilterUser query user.FilterUserRequest false "User filter parameters"
// @Success 200 {array} user.UserListResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/list [get]
func (a *Handler) ListUser(c *gin.Context) {
	var req user.FilterUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		a.log.Error("Error parsing query parameters", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.UserClient.GetUserList(c.Request.Context(), &user.FilterUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	})

	if err != nil {
		a.log.Error("Error retrieving user list", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAccessToken godoc
// @Summary Access Token
// @Description Get Access token with refresh token
// @Tags User
// @Accept json
// @Produce json
// @Param RefreshToken body Token true "Refresh token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /user/get/access-token [post]
func (a *Handler) GetAccessToken(c *gin.Context) {
	var tokenRequest Token

	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		a.log.Error("Error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	claims, err := token.ExtractToken(tokenRequest.Token, false)
	if err != nil {
		a.log.Error("Error extracting token", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	accessToken, err := token.GenerateAccessToken(claims)
	if err != nil {
		a.log.Error("Error generating access token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

type Token struct {
	Token string `json:"token"`
}
