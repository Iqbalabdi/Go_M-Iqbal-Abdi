package middleware

import (
	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type jwtMiddleware struct {
	key string
}

type JWTService interface {
	GenerateToken() (string, error)
	JwtMiddleware() echo.MiddlewareFunc
}

// enkapsulasi jwt
func NewJwtService(secretKey string) JWTService {
	return &jwtMiddleware{
		key: secretKey,
	}
}

// fungsi middleware echo untuk dipasang pada route
func (h *jwtMiddleware) JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(h.key),
	})
}

// fungsi generate token ketika login user
func (h *jwtMiddleware) GenerateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	key := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return key.SignedString([]byte(h.key))
}
