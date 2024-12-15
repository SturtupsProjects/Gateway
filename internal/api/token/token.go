package token

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// ExtractClaims validates and extracts claims from the token
func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	// Determine which secret key to use
	var secretKey string

	// Parse the token with claims
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure token uses HMAC signing
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract and return claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract token claims")
	}

	return claims, nil
}

// ValidateToken checks the validity of the token
func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaims(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}
