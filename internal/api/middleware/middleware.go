package middleware

import (
	"errors"
	"fmt"
	"gateway/internal/api/token"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func (c *casbinPermission) GetRole(ctx *gin.Context) (string, int) {
	Token := ctx.GetHeader("Authorization")
	if Token == "" {
		fmt.Println("Authorization token is missing")
		return "1.unauthorized", http.StatusUnauthorized
	}

	claims, err := token.ExtractClaims(Token)
	if err != nil {
		return "", http.StatusUnauthorized // Token xatolik yuz berdi
	}
	role, ok := claims["role"].(string)
	if !ok || role == "" {
		fmt.Println("Role is missing or invalid")
		return "3.unauthorized", http.StatusUnauthorized
	}

	userID, ok := claims["id"].(string)
	if !ok {
		fmt.Println(err)
		return "4.unauthorized", http.StatusUnauthorized
	}
	compId, ok := claims["company_id"].(string)
	if !ok {
		fmt.Println(err)
		return "4.unauthorized", http.StatusUnauthorized
	}

	ctx.Set("id", userID)
	ctx.Set("role", role)
	ctx.Set("company_id", compId)

	return role, 0
}

func (c *casbinPermission) CheckPermission(ctx *gin.Context) (bool, error) {
	subject, status := c.GetRole(ctx)
	if status != 0 {
		return false, errors.New("error while getting a role: " + subject)
	}
	action := ctx.Request.Method
	object := ctx.Request.URL.Path
	fmt.Println("subject: ", subject, "action: ", action, "object: ", object)

	allow, err := c.enforcer.Enforce(subject, object, action)
	if err != nil {
		return false, fmt.Errorf("Casbin enforce error: %w", err)
	}
	return allow, nil
}

func PermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(ctx *gin.Context) {
		res, err := casbHandler.CheckPermission(ctx)

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		if !res {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "You don't have permission"})
			return
		}
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": "authorization header is required"})
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid token: %s", err)})
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid token claims: %s", err)})
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin, Accept, X-Requested-With,X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
