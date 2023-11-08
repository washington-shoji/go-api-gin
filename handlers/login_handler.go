package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (handler *LoginHandler) LoginRenderForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login", nil)
}

func (handler *LoginHandler) LoginRenderAuth(ctx *gin.Context) {

	if err := ctx.Request.Form; err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{
			Status: http.StatusBadRequest,
			Error:  []string{"Invalid input"},
		})
		return
	}

	username := ctx.Request.FormValue("username")
	password := ctx.Request.FormValue("password")

	result, err := handler.LoginService.Login(username, password)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusForbidden, Error: []string{"Not authorized"}})
		return
	}

	ctx.SetCookie("jwt", result.Token, 3600, "/", "localhost", false, true)

	ctx.Header("HX-Redirect", "/api/html/base")
}

func (handler *LoginHandler) LogOutAuth(ctx *gin.Context) {

	ctx.SetCookie("jwt", "", -3600, "/", "localhost", false, true)

	ctx.Header("HX-Redirect", "/api/auth")
}
