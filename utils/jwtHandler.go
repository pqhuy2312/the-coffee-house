package utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pqhuy2312/the-coffee-house/exceptions"
	"github.com/pqhuy2312/the-coffee-house/models"
)

type JwtCustomClaims struct {
	Id   int       `json:"id"`
	Role string    `json:"role"`
	Version int `json:"version"`
	jwt.StandardClaims
}

func GenerateToken(u *models.User, ExpiresAt int64) string {

	claims := &JwtCustomClaims{
		Id: u.Id,
		Role: u.Role,
		Version: u.TokenVersion,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return result
}

func ParseToken(tokenString string, c *fiber.Ctx) (*JwtCustomClaims, error) {

	token, _ := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, exceptions.InValidTokenException(c)
	}

}