package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (handler *DashboardHandler) RenderDashboard(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dashboard", gin.H{})
}
