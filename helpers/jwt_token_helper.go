package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/washington-shoji/gin-api/config"
)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.EnvConfig("JWT_SECRET")))
}

func GenerateRefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString([]byte(config.EnvConfig("JWT_SECRET")))
}

func TokenValid(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.EnvConfig("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenUsername(ctx *gin.Context) (string, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.EnvConfig("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := fmt.Sprint(claims["username"])
		return username, nil
	}

	return "", nil
}

func TokenValidCookie(ctx *gin.Context) error {
	tokenString, err := ctx.Cookie("jwt")
	if err != nil {
		return err
	}
	if _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.EnvConfig("JWT_SECRET")), nil
	}); err != nil {
		return err
	}

	return nil
}
