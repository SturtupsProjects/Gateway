package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	AccessSecretKey  string
	RefreshSecretKey string
	ExpiredAccess    int
)

func GenerateAccessToken(in *Claims) (string, error) {
	claims := Claims{
		Id:          in.Id,
		FirstName:   in.FirstName,
		PhoneNumber: in.PhoneNumber,
		CompanyId:   in.CompanyId,
		Role:        in.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(ExpiredAccess)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv(AccessSecretKey)))
}

func ExtractToken(tokenStr string, isAccessToken bool) (*Claims, error) {
	var secretKey string
	if isAccessToken {
		secretKey = AccessSecretKey
	} else {
		secretKey = RefreshSecretKey
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {

		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

func GetExpires() int {
	return ExpiredAccess
}
