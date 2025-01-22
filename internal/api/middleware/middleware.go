package middleware

import (
	"errors"
	"fmt"
	"gateway/internal/api/token"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CasbinPermission struct {
	enforcer *casbin.Enforcer
}

// GetRole extracts and validates the user role from the authorization token.
func (c *CasbinPermission) GetRole(ctx *gin.Context) (string, error) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		return "", errors.New("missing authorization token")
	}

	claims, err := token.ExtractToken(tokenStr, true)
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// Set claims in context for later use
	ctx.Set("id", claims.Id)
	ctx.Set("role", claims.Role)
	ctx.Set("company_id", claims.CompanyId)
	ctx.Set("claims", claims)

	return claims.Role, nil
}

// CheckPermission verifies if the user has permission for the requested action.
func (c *CasbinPermission) CheckPermission(ctx *gin.Context) (bool, error) {
	role, err := c.GetRole(ctx)
	if err != nil {
		return false, err
	}

	action := ctx.Request.Method
	object := ctx.Request.URL.Path

	allowed, err := c.enforcer.Enforce(role, object, action)
	if err != nil {
		return false, fmt.Errorf("error during permission check: %w", err)
	}

	return allowed, nil
}

// PermissionMiddleware creates a Gin middleware for Casbin permission checks.
func PermissionMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	permissionHandler := &CasbinPermission{
		enforcer: enforcer,
	}

	return func(ctx *gin.Context) {
		allowed, err := permissionHandler.CheckPermission(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})
			return
		}

		ctx.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Branch-Id, branch_id")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
