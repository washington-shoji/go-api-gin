package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/washington-shoji/gin-api/config"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type LoginHandler struct {
	LoginService services.LoginService
}

func NewLoginHandler(service services.LoginService) *LoginHandler {
	return &LoginHandler{
		LoginService: service,
	}
}

func (handler *LoginHandler) Login(ctx *gin.Context) {
	loginReq := models.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{
			Status: http.StatusBadRequest,
			Error:  []string{"Invalid input"},
		})
		return
	}

	result, err := handler.LoginService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusForbidden, Error: []string{"Not authorized"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *LoginHandler) RefreshTokenEndpoint(ctx *gin.Context) {
	refreshTokenString := helpers.ExtractToken(ctx)

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("access denied")
		}
		return []byte(config.EnvConfig("JWT_SECRET")), nil
	})
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusForbidden, Error: []string{"Not authorized"}})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := fmt.Sprint(claims["username"])
		newAccessToken, err := helpers.GenerateToken(username)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: newAccessToken})
		} else {
			helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusForbidden, Error: []string{"Not authorized"}})
			return
		}
	}
}
